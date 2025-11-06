package profile

import (
	"fmt"
	"sort"

	"github.com/dremio/dqd/internal/models"
)

// DetailedAnalysis contains comprehensive profile analysis
type DetailedAnalysis struct {
	Profile           *models.ProfileJSON
	TotalDurationMs   float64
	PlanningDurationMs float64
	ExecutionDurationMs float64

	// Performance metrics
	BottleneckOperators []BottleneckOperator
	LongestOperators    []OperatorSummary
	MemoryIntensiveOps  []OperatorSummary

	// Resource usage
	TotalMemoryMB      float64
	PeakMemoryMB       float64
	TotalRecords       int64

	// Query phases
	PlanningPhases     []PlanningPhase

	// Fragment analysis
	FragmentStats      []FragmentStat

	// Operator breakdown
	OperatorsByType    map[string]OperatorTypeStats
}

// BottleneckOperator represents a performance bottleneck
type BottleneckOperator struct {
	FragmentID         string
	OperatorID         int64
	OperatorType       string
	ProcessTimeMs      float64
	WaitTimeMs         float64
	BlockedTimeMs      float64
	PercentOfTotal     float64
	Issue              string
}

// OperatorSummary represents an operator with key metrics
type OperatorSummary struct {
	FragmentID       string
	OperatorID       int64
	OperatorType     string
	ProcessTimeMs    float64
	WaitTimeMs       float64
	PeakMemoryMB     float64
	Records          int64
	PercentOfTotal   float64
}

// PlanningPhase represents a query planning phase
type PlanningPhase struct {
	Name         string
	DurationMs   int64
	PercentOfTotal float64
}

// FragmentStat represents fragment-level statistics
type FragmentStat struct {
	FragmentID      int
	ThreadCount     int
	TotalTimeMs     float64
	PeakMemoryMB    float64
	RecordsProcessed int64
}

// OperatorTypeStats represents aggregated stats by operator type
type OperatorTypeStats struct {
	Type            string
	Count           int
	TotalTimeMs     float64
	TotalMemoryMB   float64
	TotalRecords    int64
	AvgTimeMs       float64
}

// AnalyzeDetailed performs comprehensive profile analysis
func AnalyzeDetailed(profile *models.ProfileJSON) *DetailedAnalysis {
	analysis := &DetailedAnalysis{
		Profile:         profile,
		OperatorsByType: make(map[string]OperatorTypeStats),
	}

	// Calculate durations
	analysis.TotalDurationMs = float64(profile.End-profile.Start) / 1e6
	analysis.PlanningDurationMs = float64(profile.PlanningEnd-profile.PlanningStart) / 1e6
	analysis.ExecutionDurationMs = analysis.TotalDurationMs - analysis.PlanningDurationMs

	// Analyze operators
	operators := extractAllOperators(profile)
	analysis.analyzeOperators(operators)

	// Find bottlenecks
	analysis.findBottlenecks(operators)

	// Analyze fragments
	analysis.analyzeFragments(profile)

	// Analyze planning phases
	analysis.analyzePlanningPhases(profile)

	return analysis
}

func extractAllOperators(profile *models.ProfileJSON) []operatorWithContext {
	var operators []operatorWithContext

	for _, fragment := range profile.FragmentProfile {
		for _, minorFragment := range fragment.MinorFragmentProfile {
			fragmentID := fmt.Sprintf("%d-%d", fragment.MajorFragmentID, minorFragment.MinorFragmentID)

			for _, op := range minorFragment.OperatorProfile {
				operators = append(operators, operatorWithContext{
					FragmentID: fragmentID,
					Operator:   op,
				})
			}
		}
	}

	return operators
}

type operatorWithContext struct {
	FragmentID string
	Operator   models.OperatorProfile
}

func (a *DetailedAnalysis) analyzeOperators(operators []operatorWithContext) {
	for _, op := range operators {
		processTimeMs := float64(op.Operator.ProcessNanos) / 1e6
		memoryMB := float64(op.Operator.PeakLocalMemoryAllocated) / 1024 / 1024
		records := getTotalRecords(op.Operator.InputProfile)

		// Update totals
		a.TotalMemoryMB += memoryMB
		if memoryMB > a.PeakMemoryMB {
			a.PeakMemoryMB = memoryMB
		}
		a.TotalRecords += records

		// Track by type
		opType := getOperatorTypeName(op.Operator.OperatorType)
		stats := a.OperatorsByType[opType]
		stats.Type = opType
		stats.Count++
		stats.TotalTimeMs += processTimeMs
		stats.TotalMemoryMB += memoryMB
		stats.TotalRecords += records
		a.OperatorsByType[opType] = stats
	}

	// Calculate averages
	for opType, stats := range a.OperatorsByType {
		if stats.Count > 0 {
			stats.AvgTimeMs = stats.TotalTimeMs / float64(stats.Count)
			a.OperatorsByType[opType] = stats
		}
	}
}

func (a *DetailedAnalysis) findBottlenecks(operators []operatorWithContext) {
	// Find longest running operators
	longestOps := make([]OperatorSummary, 0, len(operators))

	for _, op := range operators {
		processTimeMs := float64(op.Operator.ProcessNanos) / 1e6
		waitTimeMs := float64(op.Operator.WaitNanos) / 1e6
		memoryMB := float64(op.Operator.PeakLocalMemoryAllocated) / 1024 / 1024
		records := getTotalRecords(op.Operator.InputProfile)

		longestOps = append(longestOps, OperatorSummary{
			FragmentID:     op.FragmentID,
			OperatorID:     op.Operator.OperatorID,
			OperatorType:   getOperatorTypeName(op.Operator.OperatorType),
			ProcessTimeMs:  processTimeMs,
			WaitTimeMs:     waitTimeMs,
			PeakMemoryMB:   memoryMB,
			Records:        records,
			PercentOfTotal: (processTimeMs / a.ExecutionDurationMs) * 100,
		})
	}

	// Sort by process time descending
	sort.Slice(longestOps, func(i, j int) bool {
		return longestOps[i].ProcessTimeMs > longestOps[j].ProcessTimeMs
	})

	// Keep top 10
	if len(longestOps) > 10 {
		a.LongestOperators = longestOps[:10]
	} else {
		a.LongestOperators = longestOps
	}

	// Find memory intensive operators
	memoryOps := make([]OperatorSummary, len(longestOps))
	copy(memoryOps, longestOps)

	sort.Slice(memoryOps, func(i, j int) bool {
		return memoryOps[i].PeakMemoryMB > memoryOps[j].PeakMemoryMB
	})

	if len(memoryOps) > 10 {
		a.MemoryIntensiveOps = memoryOps[:10]
	} else {
		a.MemoryIntensiveOps = memoryOps
	}

	// Identify bottlenecks (operators taking >5% of execution time)
	for _, op := range longestOps {
		if op.PercentOfTotal > 5.0 {
			issue := ""
			if op.WaitTimeMs > op.ProcessTimeMs {
				issue = "High wait time (I/O or blocking)"
			} else if op.PeakMemoryMB > 100 {
				issue = "High memory usage"
			} else {
				issue = "Long processing time"
			}

			a.BottleneckOperators = append(a.BottleneckOperators, BottleneckOperator{
				FragmentID:     op.FragmentID,
				OperatorID:     op.OperatorID,
				OperatorType:   op.OperatorType,
				ProcessTimeMs:  op.ProcessTimeMs,
				WaitTimeMs:     op.WaitTimeMs,
				BlockedTimeMs:  op.WaitTimeMs,
				PercentOfTotal: op.PercentOfTotal,
				Issue:          issue,
			})
		}
	}
}

func (a *DetailedAnalysis) analyzeFragments(profile *models.ProfileJSON) {
	for _, fragment := range profile.FragmentProfile {
		stat := FragmentStat{
			FragmentID:  fragment.MajorFragmentID,
			ThreadCount: len(fragment.MinorFragmentProfile),
		}

		for _, minorFragment := range fragment.MinorFragmentProfile {
			duration := float64(minorFragment.EndTime-minorFragment.StartTime) / 1e6
			memoryMB := float64(minorFragment.MaxMemoryUsed) / 1024 / 1024

			stat.TotalTimeMs += duration
			if memoryMB > stat.PeakMemoryMB {
				stat.PeakMemoryMB = memoryMB
			}

			for _, op := range minorFragment.OperatorProfile {
				stat.RecordsProcessed += getTotalRecords(op.InputProfile)
			}
		}

		a.FragmentStats = append(a.FragmentStats, stat)
	}
}

func (a *DetailedAnalysis) analyzePlanningPhases(profile *models.ProfileJSON) {
	var totalPlanningMs int64

	for _, phase := range profile.PlanPhases {
		totalPlanningMs += phase.DurationMillis
	}

	for _, phase := range profile.PlanPhases {
		pct := float64(0)
		if totalPlanningMs > 0 {
			pct = (float64(phase.DurationMillis) / float64(totalPlanningMs)) * 100
		}

		a.PlanningPhases = append(a.PlanningPhases, PlanningPhase{
			Name:           phase.PhaseName,
			DurationMs:     phase.DurationMillis,
			PercentOfTotal: pct,
		})
	}
}
