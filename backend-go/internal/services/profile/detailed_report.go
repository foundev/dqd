package profile

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/dremio/dqd/internal/models"
)

// GenerateDetailedReport generates a comprehensive HTML report from a profile
func GenerateDetailedReport(profile *models.ProfileJSON) (string, error) {
	// Perform detailed analysis
	analysis := AnalyzeDetailed(profile)

	data := struct {
		Profile  *models.ProfileJSON
		Analysis *DetailedAnalysis
	}{
		Profile:  profile,
		Analysis: analysis,
	}

	tmpl, err := template.New("detailed").Funcs(template.FuncMap{
		"formatTime":   formatTime,
		"formatNumber": formatNumber,
		"formatBytes":  formatBytes,
		"formatPct":    formatPct,
		"divf":         divf,
		"mulf":         mulf,
		"float64":      toFloat64,
	}).Parse(detailedReportTemplate)

	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func formatNumber(n int64) string {
	// Format number with thousands separator
	if n < 0 {
		return "-" + formatNumber(-n)
	}
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}

	s := fmt.Sprintf("%d", n)
	var result []byte
	for i, digit := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(digit))
	}
	return string(result)
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatPct(val float64) string {
	return fmt.Sprintf("%.1f%%", val)
}

func divf(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func mulf(a, b float64) float64 {
	return a * b
}

func toFloat64(n int64) float64 {
	return float64(n)
}

const detailedReportTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DQD Detailed Profile Analysis</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
            background: #f5f5f5;
            color: #333;
        }
        .header {
            background: linear-gradient(135deg, #006493 0%, #004d73 100%);
            color: white;
            padding: 30px;
            box-shadow: 0 2px 8px rgba(0,0,0,0.15);
        }
        .header h1 {
            font-size: 32px;
            margin-bottom: 10px;
        }
        .header .subtitle {
            font-size: 16px;
            opacity: 0.9;
        }
        .container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 20px;
        }
        .nav {
            background: white;
            padding: 15px 30px;
            margin-bottom: 20px;
            border-radius: 8px;
            box-shadow: 0 1px 3px rgba(0,0,0,0.1);
            position: sticky;
            top: 0;
            z-index: 100;
        }
        .nav a {
            color: #006493;
            text-decoration: none;
            margin-right: 20px;
            font-weight: 500;
        }
        .nav a:hover {
            text-decoration: underline;
        }
        .section {
            background: white;
            margin-bottom: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .section-header {
            background: #f8f9fa;
            padding: 20px;
            border-bottom: 2px solid #006493;
        }
        .section-header h2 {
            color: #006493;
            font-size: 24px;
        }
        .section-content {
            padding: 20px;
        }
        .metrics-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 20px;
            margin-bottom: 20px;
        }
        .metric-card {
            background: linear-gradient(135deg, #006493 0%, #0077b3 100%);
            color: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 8px rgba(0,100,147,0.3);
        }
        .metric-card .label {
            font-size: 14px;
            opacity: 0.9;
            margin-bottom: 8px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        .metric-card .value {
            font-size: 32px;
            font-weight: bold;
        }
        .metric-card .subvalue {
            font-size: 14px;
            opacity: 0.8;
            margin-top: 5px;
        }
        .alert {
            padding: 15px 20px;
            margin-bottom: 15px;
            border-radius: 6px;
            border-left: 4px solid;
        }
        .alert-warning {
            background: #fff3cd;
            border-color: #ffc107;
            color: #856404;
        }
        .alert-info {
            background: #d1ecf1;
            border-color: #0dcaf0;
            color: #0c5460;
        }
        .alert-danger {
            background: #f8d7da;
            border-color: #dc3545;
            color: #721c24;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 15px;
        }
        th {
            background: #006493;
            color: white;
            padding: 12px;
            text-align: left;
            font-weight: 600;
            position: sticky;
            top: 60px;
            z-index: 10;
        }
        td {
            padding: 12px;
            border-bottom: 1px solid #e0e0e0;
        }
        tr:hover {
            background: #f5f5f5;
        }
        .number {
            text-align: right;
            font-family: 'Courier New', monospace;
        }
        .badge {
            display: inline-block;
            padding: 4px 12px;
            border-radius: 12px;
            font-size: 12px;
            font-weight: 600;
        }
        .badge-danger {
            background: #dc3545;
            color: white;
        }
        .badge-warning {
            background: #ffc107;
            color: #000;
        }
        .badge-info {
            background: #0dcaf0;
            color: #000;
        }
        .query-box {
            background: #f8f9fa;
            border: 1px solid #dee2e6;
            border-radius: 6px;
            padding: 15px;
            font-family: 'Courier New', monospace;
            font-size: 13px;
            white-space: pre-wrap;
            overflow-x: auto;
        }
        .footer {
            text-align: center;
            padding: 30px;
            color: #666;
            font-size: 14px;
        }
        .progress-bar {
            height: 24px;
            background: #e9ecef;
            border-radius: 4px;
            overflow: hidden;
            margin-top: 8px;
        }
        .progress-fill {
            height: 100%;
            background: linear-gradient(90deg, #006493 0%, #0077b3 100%);
            display: flex;
            align-items: center;
            padding: 0 10px;
            color: white;
            font-size: 12px;
            font-weight: 600;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>🔍 Detailed Profile Analysis</h1>
        <div class="subtitle">Comprehensive query performance diagnostics</div>
    </div>

    <div class="nav">
        <a href="#summary">Summary</a>
        <a href="#bottlenecks">Bottlenecks</a>
        <a href="#operators">Operators</a>
        <a href="#resources">Resources</a>
        <a href="#planning">Planning</a>
        <a href="#fragments">Fragments</a>
    </div>

    <div class="container">
        <!-- Executive Summary -->
        <div class="section" id="summary">
            <div class="section-header">
                <h2>📊 Executive Summary</h2>
            </div>
            <div class="section-content">
                <div class="metrics-grid">
                    <div class="metric-card">
                        <div class="label">Total Duration</div>
                        <div class="value">{{ printf "%.0f" .Analysis.TotalDurationMs }} ms</div>
                        <div class="subvalue">{{ printf "%.2f" (divf .Analysis.TotalDurationMs 1000) }} seconds</div>
                    </div>
                    <div class="metric-card">
                        <div class="label">Planning Time</div>
                        <div class="value">{{ printf "%.0f" .Analysis.PlanningDurationMs }} ms</div>
                        <div class="subvalue">{{ printf "%.1f%%" (mulf (divf .Analysis.PlanningDurationMs .Analysis.TotalDurationMs) 100) }} of total</div>
                    </div>
                    <div class="metric-card">
                        <div class="label">Execution Time</div>
                        <div class="value">{{ printf "%.0f" .Analysis.ExecutionDurationMs }} ms</div>
                        <div class="subvalue">{{ printf "%.1f%%" (mulf (divf .Analysis.ExecutionDurationMs .Analysis.TotalDurationMs) 100) }} of total</div>
                    </div>
                    <div class="metric-card">
                        <div class="label">Total Memory</div>
                        <div class="value">{{ printf "%.0f" .Analysis.PeakMemoryMB }} MB</div>
                        <div class="subvalue">Peak memory used</div>
                    </div>
                    <div class="metric-card">
                        <div class="label">Records Processed</div>
                        <div class="value">{{ formatNumber .Analysis.TotalRecords }}</div>
                        <div class="subvalue">{{ printf "%.0f" (divf (float64 .Analysis.TotalRecords) (divf .Analysis.ExecutionDurationMs 1000)) }} rec/sec</div>
                    </div>
                    <div class="metric-card">
                        <div class="label">Fragments</div>
                        <div class="value">{{ .Profile.TotalFragments }}</div>
                        <div class="subvalue">{{ .Profile.FinishedFragments }} finished</div>
                    </div>
                </div>

                <h3 style="margin: 20px 0 10px 0;">Query Information</h3>
                <table>
                    <tr>
                        <td style="width: 200px; font-weight: 600;">Query ID</td>
                        <td>{{ if .Profile.ID }}{{ printf "%x-%x" .Profile.ID.Part1 .Profile.ID.Part2 }}{{ else }}N/A{{ end }}</td>
                    </tr>
                    <tr>
                        <td style="font-weight: 600;">User</td>
                        <td>{{ .Profile.User }}</td>
                    </tr>
                    <tr>
                        <td style="font-weight: 600;">Dremio Version</td>
                        <td>{{ .Profile.DremioVersion }}</td>
                    </tr>
                    <tr>
                        <td style="font-weight: 600;">Start Time</td>
                        <td>{{ formatTime .Profile.Start }}</td>
                    </tr>
                    <tr>
                        <td style="font-weight: 600;">End Time</td>
                        <td>{{ formatTime .Profile.End }}</td>
                    </tr>
                </table>

                <h3 style="margin: 20px 0 10px 0;">Query</h3>
                <div class="query-box">{{ .Profile.Query }}</div>
            </div>
        </div>

        <!-- Performance Bottlenecks -->
        <div class="section" id="bottlenecks">
            <div class="section-header">
                <h2>⚠️ Performance Bottlenecks</h2>
            </div>
            <div class="section-content">
                {{ if .Analysis.BottleneckOperators }}
                <div class="alert alert-warning">
                    <strong>{{ len .Analysis.BottleneckOperators }} bottleneck(s) detected</strong> - These operators are consuming significant execution time
                </div>

                <table>
                    <thead>
                        <tr>
                            <th>Fragment</th>
                            <th>Operator</th>
                            <th>Type</th>
                            <th class="number">Process Time</th>
                            <th class="number">Wait Time</th>
                            <th class="number">% of Total</th>
                            <th>Issue</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Analysis.BottleneckOperators }}
                        <tr>
                            <td>{{ .FragmentID }}</td>
                            <td>{{ .OperatorID }}</td>
                            <td>{{ .OperatorType }}</td>
                            <td class="number">{{ printf "%.2f" .ProcessTimeMs }} ms</td>
                            <td class="number">{{ printf "%.2f" .WaitTimeMs }} ms</td>
                            <td class="number">
                                <span class="badge badge-danger">{{ printf "%.1f%%" .PercentOfTotal }}</span>
                            </td>
                            <td>{{ .Issue }}</td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
                {{ else }}
                <div class="alert alert-info">
                    <strong>No significant bottlenecks detected</strong> - Query execution appears balanced
                </div>
                {{ end }}
            </div>
        </div>

        <!-- Operator Analysis -->
        <div class="section" id="operators">
            <div class="section-header">
                <h2>🔧 Operator Analysis</h2>
            </div>
            <div class="section-content">
                <h3>Top 10 Longest Running Operators</h3>
                <table>
                    <thead>
                        <tr>
                            <th>Fragment</th>
                            <th>Operator</th>
                            <th>Type</th>
                            <th class="number">Process Time</th>
                            <th class="number">Wait Time</th>
                            <th class="number">Memory (MB)</th>
                            <th class="number">Records</th>
                            <th class="number">% of Total</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Analysis.LongestOperators }}
                        <tr>
                            <td>{{ .FragmentID }}</td>
                            <td>{{ .OperatorID }}</td>
                            <td>{{ .OperatorType }}</td>
                            <td class="number">{{ printf "%.2f" .ProcessTimeMs }} ms</td>
                            <td class="number">{{ printf "%.2f" .WaitTimeMs }} ms</td>
                            <td class="number">{{ printf "%.2f" .PeakMemoryMB }}</td>
                            <td class="number">{{ formatNumber .Records }}</td>
                            <td class="number">{{ printf "%.1f%%" .PercentOfTotal }}</td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>

                <h3 style="margin-top: 30px;">Operators by Type</h3>
                <table>
                    <thead>
                        <tr>
                            <th>Operator Type</th>
                            <th class="number">Count</th>
                            <th class="number">Total Time</th>
                            <th class="number">Avg Time</th>
                            <th class="number">Total Memory</th>
                            <th class="number">Records</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range $type, $stats := .Analysis.OperatorsByType }}
                        <tr>
                            <td>{{ $stats.Type }}</td>
                            <td class="number">{{ $stats.Count }}</td>
                            <td class="number">{{ printf "%.2f" $stats.TotalTimeMs }} ms</td>
                            <td class="number">{{ printf "%.2f" $stats.AvgTimeMs }} ms</td>
                            <td class="number">{{ printf "%.2f" $stats.TotalMemoryMB }} MB</td>
                            <td class="number">{{ formatNumber $stats.TotalRecords }}</td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
            </div>
        </div>

        <!-- Resource Usage -->
        <div class="section" id="resources">
            <div class="section-header">
                <h2>💾 Resource Usage</h2>
            </div>
            <div class="section-content">
                <h3>Top 10 Memory Intensive Operators</h3>
                <table>
                    <thead>
                        <tr>
                            <th>Fragment</th>
                            <th>Operator</th>
                            <th>Type</th>
                            <th class="number">Peak Memory (MB)</th>
                            <th class="number">Process Time</th>
                            <th class="number">Records</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Analysis.MemoryIntensiveOps }}
                        <tr>
                            <td>{{ .FragmentID }}</td>
                            <td>{{ .OperatorID }}</td>
                            <td>{{ .OperatorType }}</td>
                            <td class="number">
                                <strong>{{ printf "%.2f" .PeakMemoryMB }}</strong>
                            </td>
                            <td class="number">{{ printf "%.2f" .ProcessTimeMs }} ms</td>
                            <td class="number">{{ formatNumber .Records }}</td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
            </div>
        </div>

        <!-- Planning Phases -->
        <div class="section" id="planning">
            <div class="section-header">
                <h2>📋 Query Planning Phases</h2>
            </div>
            <div class="section-content">
                {{ if .Analysis.PlanningPhases }}
                <table>
                    <thead>
                        <tr>
                            <th>Phase</th>
                            <th class="number">Duration (ms)</th>
                            <th>% of Planning</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Analysis.PlanningPhases }}
                        <tr>
                            <td>{{ .Name }}</td>
                            <td class="number">{{ .DurationMs }}</td>
                            <td>
                                <div class="progress-bar">
                                    <div class="progress-fill" style="width: {{ printf "%.1f%%" .PercentOfTotal }}">
                                        {{ printf "%.1f%%" .PercentOfTotal }}
                                    </div>
                                </div>
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
                {{ else }}
                <div class="alert alert-info">No planning phase information available</div>
                {{ end }}
            </div>
        </div>

        <!-- Fragment Statistics -->
        <div class="section" id="fragments">
            <div class="section-header">
                <h2>🧩 Fragment Statistics</h2>
            </div>
            <div class="section-content">
                <table>
                    <thead>
                        <tr>
                            <th>Fragment ID</th>
                            <th class="number">Threads</th>
                            <th class="number">Total Time (ms)</th>
                            <th class="number">Peak Memory (MB)</th>
                            <th class="number">Records Processed</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .Analysis.FragmentStats }}
                        <tr>
                            <td>{{ .FragmentID }}</td>
                            <td class="number">{{ .ThreadCount }}</td>
                            <td class="number">{{ printf "%.2f" .TotalTimeMs }}</td>
                            <td class="number">{{ printf "%.2f" .PeakMemoryMB }}</td>
                            <td class="number">{{ formatNumber .RecordsProcessed }}</td>
                        </tr>
                        {{ end }}
                    </tbody>
                </table>
            </div>
        </div>
    </div>

    <div class="footer">
        <p><strong>Dremio Query Doctor (DQD)</strong> - Detailed Profile Analysis</p>
        <p>Generated by Go Backend • {{ len .Analysis.LongestOperators }} operators analyzed</p>
    </div>
</body>
</html>`
