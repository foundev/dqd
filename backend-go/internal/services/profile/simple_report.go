package profile

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/dremio/dqd/internal/models"
)

// OperatorRow represents a row in the operator table
type OperatorRow struct {
	FragmentID    string
	OperatorID    int64
	OperatorType  string
	ProcessTimeMs float64
	WaitTimeMs    float64
	SetupTimeMs   float64
	PeakMemoryMB  float64
	Records       int64
}

// GenerateSimpleReport generates a simple HTML report from a profile
func GenerateSimpleReport(profile *models.ProfileJSON) (string, error) {
	// Extract operator data
	operators := extractOperators(profile)

	// Calculate summary stats
	totalDuration := float64(profile.End-profile.Start) / 1e6 // Convert to ms
	planningDuration := float64(profile.PlanningEnd-profile.PlanningStart) / 1e6

	data := struct {
		Profile           *models.ProfileJSON
		Operators         []OperatorRow
		TotalDurationMs   float64
		PlanningDurationMs float64
		OperatorCount     int
		QueryPreview      string
	}{
		Profile:           profile,
		Operators:         operators,
		TotalDurationMs:   totalDuration,
		PlanningDurationMs: planningDuration,
		OperatorCount:     len(operators),
		QueryPreview:      truncate(profile.Query, 200),
	}

	tmpl, err := template.New("simple").Funcs(template.FuncMap{
		"formatTime": formatTime,
	}).Parse(simpleReportTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func extractOperators(profile *models.ProfileJSON) []OperatorRow {
	var operators []OperatorRow

	for _, fragment := range profile.FragmentProfile {
		for _, minorFragment := range fragment.MinorFragmentProfile {
			fragmentID := fmt.Sprintf("%d-%d", fragment.MajorFragmentID, minorFragment.MinorFragmentID)

			for _, op := range minorFragment.OperatorProfile {
				row := OperatorRow{
					FragmentID:    fragmentID,
					OperatorID:    op.OperatorID,
					OperatorType:  getOperatorTypeName(op.OperatorType),
					ProcessTimeMs: float64(op.ProcessNanos) / 1e6,
					WaitTimeMs:    float64(op.WaitNanos) / 1e6,
					SetupTimeMs:   float64(op.SetupNanos) / 1e6,
					PeakMemoryMB:  float64(op.PeakLocalMemoryAllocated) / 1024 / 1024,
					Records:       getTotalRecords(op.InputProfile),
				}
				operators = append(operators, row)
			}
		}
	}

	return operators
}

func getTotalRecords(inputs []models.InputProfile) int64 {
	var total int64
	for _, input := range inputs {
		total += input.Records
	}
	return total
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func formatTime(millis int64) string {
	if millis == 0 {
		return "N/A"
	}
	t := time.Unix(millis/1000, (millis%1000)*1e6)
	return t.Format("2006-01-02 15:04:05 MST")
}

func getOperatorTypeName(opType int) string {
	// Map of operator types from Dremio
	operatorTypes := map[int]string{
		0:  "SINGLE_SENDER",
		1:  "BROADCAST_SENDER",
		2:  "FILTER",
		3:  "HASH_AGGREGATE",
		4:  "HASH_JOIN",
		5:  "MERGE_JOIN",
		6:  "HASH_PARTITION_SENDER",
		7:  "LIMIT",
		8:  "MERGING_RECEIVER",
		9:  "ORDERED_PARTITION_SENDER",
		10: "PROJECT",
		11: "UNORDERED_RECEIVER",
		12: "RANGE_SENDER",
		13: "SCREEN",
		14: "SELECTION_VECTOR_REMOVER",
		15: "STREAMING_AGGREGATE",
		16: "TOP_N_SORT",
		17: "EXTERNAL_SORT",
		18: "TRACE",
		19: "UNION",
		20: "OLD_SORT",
		21: "PARQUET_ROW_GROUP_SCAN",
		22: "HIVE_SUB_SCAN",
		23: "SYSTEM_TABLE_SCAN",
		24: "MOCK_SUB_SCAN",
		25: "PARQUET_WRITER",
		26: "DIRECT_SUB_SCAN",
		27: "TEXT_SUB_SCAN",
		28: "TEXT_WRITER",
		29: "JSON_SUB_SCAN",
		30: "INFO_SCHEMA_SUB_SCAN",
		31: "COMPLEX_TO_JSON",
		32: "PRODUCER_CONSUMER",
		33: "HBASE_SUB_SCAN",
		34: "WINDOW",
		35: "NESTED_LOOP_JOIN",
		36: "AVRO_SUB_SCAN",
	}

	if name, ok := operatorTypes[opType]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_%d", opType)
}

const simpleReportTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>DQD Simple Profile Analysis</title>
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
            padding: 20px;
        }
        .container {
            max-width: 1400px;
            margin: 0 auto;
            background: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        h1 {
            color: #006493;
            margin-bottom: 10px;
            font-size: 28px;
        }
        h2 {
            color: #006493;
            margin-top: 30px;
            margin-bottom: 15px;
            font-size: 20px;
            border-bottom: 2px solid #006493;
            padding-bottom: 5px;
        }
        .summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin: 20px 0;
        }
        .summary-card {
            background: #f8f9fa;
            padding: 15px;
            border-radius: 6px;
            border-left: 4px solid #006493;
        }
        .summary-card .label {
            font-size: 12px;
            color: #666;
            text-transform: uppercase;
            margin-bottom: 5px;
        }
        .summary-card .value {
            font-size: 24px;
            font-weight: bold;
            color: #006493;
        }
        .query-preview {
            background: #f8f9fa;
            padding: 15px;
            border-radius: 6px;
            margin: 20px 0;
            font-family: 'Courier New', monospace;
            font-size: 13px;
            white-space: pre-wrap;
            border: 1px solid #ddd;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
        }
        th {
            background: #006493;
            color: white;
            padding: 12px 8px;
            text-align: left;
            font-weight: 600;
            position: sticky;
            top: 0;
            z-index: 10;
        }
        td {
            padding: 10px 8px;
            border-bottom: 1px solid #e0e0e0;
        }
        tr:hover {
            background: #f5f5f5;
        }
        .number {
            text-align: right;
            font-family: 'Courier New', monospace;
        }
        .footer {
            margin-top: 30px;
            padding-top: 20px;
            border-top: 1px solid #e0e0e0;
            color: #666;
            font-size: 12px;
            text-align: center;
        }
        .metadata {
            display: grid;
            grid-template-columns: auto 1fr;
            gap: 10px;
            margin: 20px 0;
            font-size: 14px;
        }
        .metadata dt {
            font-weight: 600;
            color: #666;
        }
        .metadata dd {
            color: #333;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>DQD Simple Profile Analysis</h1>
        <p style="color: #666; margin-bottom: 20px;">Generated by Go Backend on {{ .Profile.DremioVersion }}</p>

        <div class="summary">
            <div class="summary-card">
                <div class="label">Total Duration</div>
                <div class="value">{{ printf "%.2f" .TotalDurationMs }} ms</div>
            </div>
            <div class="summary-card">
                <div class="label">Planning Time</div>
                <div class="value">{{ printf "%.2f" .PlanningDurationMs }} ms</div>
            </div>
            <div class="summary-card">
                <div class="label">Total Operators</div>
                <div class="value">{{ .OperatorCount }}</div>
            </div>
            <div class="summary-card">
                <div class="label">Fragments</div>
                <div class="value">{{ .Profile.TotalFragments }}</div>
            </div>
        </div>

        <h2>Query</h2>
        <div class="query-preview">{{ .QueryPreview }}</div>

        <dl class="metadata">
            <dt>User:</dt>
            <dd>{{ .Profile.User }}</dd>
            <dt>Query ID:</dt>
            <dd>{{ if .Profile.ID }}{{ printf "%x-%x" .Profile.ID.Part1 .Profile.ID.Part2 }}{{ else }}N/A{{ end }}</dd>
            <dt>Start Time:</dt>
            <dd>{{ formatTime .Profile.Start }}</dd>
            <dt>End Time:</dt>
            <dd>{{ formatTime .Profile.End }}</dd>
        </dl>

        <h2>Operators ({{ .OperatorCount }} total)</h2>
        <div style="overflow-x: auto;">
            <table>
                <thead>
                    <tr>
                        <th>Fragment</th>
                        <th>Operator ID</th>
                        <th>Type</th>
                        <th class="number">Process Time (ms)</th>
                        <th class="number">Wait Time (ms)</th>
                        <th class="number">Setup Time (ms)</th>
                        <th class="number">Peak Memory (MB)</th>
                        <th class="number">Records</th>
                    </tr>
                </thead>
                <tbody>
                    {{ range .Operators }}
                    <tr>
                        <td>{{ .FragmentID }}</td>
                        <td>{{ .OperatorID }}</td>
                        <td>{{ .OperatorType }}</td>
                        <td class="number">{{ printf "%.2f" .ProcessTimeMs }}</td>
                        <td class="number">{{ printf "%.2f" .WaitTimeMs }}</td>
                        <td class="number">{{ printf "%.2f" .SetupTimeMs }}</td>
                        <td class="number">{{ printf "%.2f" .PeakMemoryMB }}</td>
                        <td class="number">{{ .Records }}</td>
                    </tr>
                    {{ end }}
                </tbody>
            </table>
        </div>

        <div class="footer">
            <p>Dremio Query Doctor (DQD) - Simple Profile Analysis Report</p>
            <p>Generated with Go backend • {{ .OperatorCount }} operators analyzed</p>
        </div>
    </div>
</body>
</html>`
