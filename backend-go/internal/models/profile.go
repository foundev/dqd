package models

// ProfileJSON represents a Dremio query profile
type ProfileJSON struct {
	ID                           *ID                          `json:"id,omitempty"`
	Query                        string                       `json:"query,omitempty"`
	Plan                         string                       `json:"plan,omitempty"`
	JSONPlan                     string                       `json:"jsonPlan,omitempty"`
	User                         string                       `json:"user,omitempty"`
	State                        int                          `json:"state,omitempty"`
	Start                        int64                        `json:"start,omitempty"`
	End                          int64                        `json:"end,omitempty"`
	PlanningStart                int64                        `json:"planningStart,omitempty"`
	PlanningEnd                  int64                        `json:"planningEnd,omitempty"`
	TotalFragments               int64                        `json:"totalFragments,omitempty"`
	FinishedFragments            int64                        `json:"finishedFragments,omitempty"`
	FragmentProfile              []FragmentProfile            `json:"fragmentProfile,omitempty"`
	NodeProfile                  []NodeProfile                `json:"nodeProfile,omitempty"`
	Foreman                      *Foreman                     `json:"foreman,omitempty"`
	ClientInfo                   *ClientInfo                  `json:"clientInfo,omitempty"`
	DremioVersion                string                       `json:"dremioVersion,omitempty"`
	NonDefaultOptionsJSON        string                       `json:"nonDefaultOptionsJSON,omitempty"`
	ResourceSchedulingProfile    *ResourceSchedulingProfile   `json:"resourceSchedulingProfile,omitempty"`
	AccelerationProfile          *AccelerationProfile         `json:"accelerationProfile,omitempty"`
	DatasetProfile               []DatasetProfile             `json:"datasetProfile,omitempty"`
	PlanPhases                   []PlanPhases                 `json:"planPhases,omitempty"`
	StateList                    []ProfileState               `json:"stateList,omitempty"`
	CommandPoolWaitMillis        int64                        `json:"commandPoolWaitMillis,omitempty"`
	NumPlanCacheUsed             int64                        `json:"numPlanCacheUsed,omitempty"`
	SerializedPlan               string                       `json:"serializedPlan,omitempty"`
	OperatorTypeMetricsMap       *OperatorTypeMetricsMap      `json:"operatorTypeMetricsMap,omitempty"`
}

// ID represents a profile identifier
type ID struct {
	Part1 int64 `json:"part1,omitempty"`
	Part2 int64 `json:"part2,omitempty"`
}

// FragmentProfile represents a query fragment
type FragmentProfile struct {
	MajorFragmentID     int                   `json:"majorFragmentId,omitempty"`
	MinorFragmentProfile []MinorFragmentProfile `json:"minorFragmentProfile,omitempty"`
}

// MinorFragmentProfile represents a minor fragment execution
type MinorFragmentProfile struct {
	State                       int               `json:"state,omitempty"`
	MinorFragmentID             int               `json:"minorFragmentId,omitempty"`
	OperatorProfile             []OperatorProfile `json:"operatorProfile,omitempty"`
	StartTime                   int64             `json:"startTime,omitempty"`
	EndTime                     int64             `json:"endTime,omitempty"`
	MemoryUsed                  int64             `json:"memoryUsed,omitempty"`
	MaxMemoryUsed               int64             `json:"maxMemoryUsed,omitempty"`
	Endpoint                    *Endpoint         `json:"endpoint,omitempty"`
	LastUpdate                  int64             `json:"lastUpdate,omitempty"`
	LastProgress                int64             `json:"lastProgress,omitempty"`
	BlockedOnDownstreamDuration int64             `json:"blockedOnDownstreamDuration,omitempty"`
	BlockedOnUpstreamDuration   int64             `json:"blockedOnUpstreamDuration,omitempty"`
	SleepingDuration            int64             `json:"sleepingDuration,omitempty"`
}

// OperatorProfile represents an operator's performance metrics
type OperatorProfile struct {
	OperatorID               int64          `json:"operatorId,omitempty"`
	OperatorType             int            `json:"operatorType,omitempty"`
	OperatorSubtype          int64          `json:"operatorSubtype,omitempty"`
	SetupNanos               int64          `json:"setupNanos,omitempty"`
	ProcessNanos             int64          `json:"processNanos,omitempty"`
	WaitNanos                int64          `json:"waitNanos,omitempty"`
	PeakLocalMemoryAllocated int64          `json:"peakLocalMemoryAllocated,omitempty"`
	Metric                   []Metric       `json:"metric,omitempty"`
	InputProfile             []InputProfile `json:"inputProfile,omitempty"`
}

// Metric represents a performance metric
type Metric struct {
	MetricID   int   `json:"metricId,omitempty"`
	LongValue  int64 `json:"longValue,omitempty"`
	DoubleValue float64 `json:"doubleValue,omitempty"`
}

// InputProfile represents input data profile
type InputProfile struct {
	Records int64 `json:"records,omitempty"`
	Batches int64 `json:"batches,omitempty"`
	Schemas int64 `json:"schemas,omitempty"`
}

// NodeProfile represents a node's profile
type NodeProfile struct {
	Endpoint     *Endpoint `json:"endpoint,omitempty"`
	MaxMemoryUsed int64    `json:"maxMemoryUsed,omitempty"`
	TimeEnqueuedBeforeSubmitMs int64 `json:"timeEnqueuedBeforeSubmitMs,omitempty"`
}

// Endpoint represents a network endpoint
type Endpoint struct {
	Address        string  `json:"address,omitempty"`
	FabricPort     int     `json:"fabricPort,omitempty"`
	UserPort       int     `json:"userPort,omitempty"`
	Roles          *Roles  `json:"roles,omitempty"`
	AvailableCores int     `json:"availableCores,omitempty"`
	NodeTag        string  `json:"nodeTag,omitempty"`
}

// Roles represents node roles
type Roles struct {
	SQLQuery       bool `json:"sqlQuery,omitempty"`
	LogicalPlan    bool `json:"logicalPlan,omitempty"`
	PhysicalPlan   bool `json:"physicalPlan,omitempty"`
	JavaExecutor   bool `json:"javaExecutor,omitempty"`
	DistributedCache bool `json:"distributedCache,omitempty"`
	Master         bool `json:"master,omitempty"`
}

// Foreman represents the coordinator
type Foreman struct {
	Address    string `json:"address,omitempty"`
	FabricPort int    `json:"fabricPort,omitempty"`
	UserPort   int    `json:"userPort,omitempty"`
}

// ClientInfo represents client information
type ClientInfo struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Application string `json:"application,omitempty"`
}

// ResourceSchedulingProfile represents resource scheduling info
type ResourceSchedulingProfile struct {
	ResourceSchedulingStart int64 `json:"resourceSchedulingStart,omitempty"`
	ResourceSchedulingEnd   int64 `json:"resourceSchedulingEnd,omitempty"`
	QueueName               string `json:"queueName,omitempty"`
	QueueID                 string `json:"queueId,omitempty"`
	RuleName                string `json:"ruleName,omitempty"`
	RuleID                  string `json:"ruleId,omitempty"`
	RuleContent             string `json:"ruleContent,omitempty"`
}

// AccelerationProfile represents acceleration info
type AccelerationProfile struct {
	Accelerated      bool   `json:"accelerated,omitempty"`
	LayoutID         string `json:"layoutId,omitempty"`
	NumSubstitutions int    `json:"numSubstitutions,omitempty"`
}

// DatasetProfile represents dataset information
type DatasetProfile struct {
	DatasetPath    string `json:"datasetPath,omitempty"`
	DatasetType    string `json:"datasetType,omitempty"`
	SQL            string `json:"sql,omitempty"`
	BatchSchema    string `json:"batchSchema,omitempty"`
	AllowPartialMatch bool `json:"allowPartialMatch,omitempty"`
}

// PlanPhases represents query planning phases
type PlanPhases struct {
	PhaseName       string `json:"phaseName,omitempty"`
	DurationMillis  int64  `json:"durationMillis,omitempty"`
	PlannerType     string `json:"plannerType,omitempty"`
}

// ProfileState represents a state change
type ProfileState struct {
	State          int   `json:"state,omitempty"`
	StateStartTime int64 `json:"stateStartTime,omitempty"`
}

// OperatorTypeMetricsMap represents operator type metrics
type OperatorTypeMetricsMap struct {
	OperatorTypeMetrics []OperatorTypeMetrics `json:"operatorTypeMetrics,omitempty"`
}

// OperatorTypeMetrics represents metrics for an operator type
type OperatorTypeMetrics struct {
	OperatorType int      `json:"operatorType,omitempty"`
	Metric       []Metric `json:"metric,omitempty"`
}
