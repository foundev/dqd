package models

// AboutResponse represents the response from /about.json
type AboutResponse struct {
	Version string `json:"version"`
}

// ProfileAnalysisRequest represents a profile analysis request
type ProfileAnalysisRequest struct {
	ShowPlanDetails  bool `json:"showPlanDetails"`
	ShowConvertToRel bool `json:"showConvertToRel"`
}

// QueriesAnalysisRequest represents a queries.json analysis request
type QueriesAnalysisRequest struct {
	StartDate string `json:"startDate"`
	StartTime string `json:"startTime"`
	EndDate   string `json:"endDate"`
	EndTime   string `json:"endTime"`
	Limit     int    `json:"limit"`
	Window    int    `json:"window"`
}

// ReproductionRequest represents a schema generation request
type ReproductionRequest struct {
	Records            int    `json:"records"`
	Timeout            int    `json:"timeout"`
	DefaultCtasFormat  string `json:"defaultCtasFormat,omitempty"`
	NasPath            string `json:"nasPath,omitempty"`
	ColumnDefYaml      string `json:"columnDefYaml,omitempty"`
}
