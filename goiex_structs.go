package goiex

type EarningsToday struct {
	Bto BTO
	Amc AMC
}

type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []earnings
}

type earnings struct {
	ActualEPS              float32 `json:"actualEPS"`
	ConcensusEPS           float32 `json:"consensusEPS"`
	EstimatedEPS           float32 `json:"estimatedEPS"`
	AnnounceTime           string  `json:"announceTime"`
	NumberOfEstimates      int32   `json:"numberOfEstimates"`
	EPSSurpriseDollar      float32 `json:"EPSSurpriseDollar"`
	EPSReportDate          string  `json:"EPSReportDate"` // make date type?
	FiscalPeriod           string  `json:"fiscalPeriod"`
	FiscalEndDate          string  `json:"fiscalEndDate"`
	YearAgo                float32 `json:"yearAgo"`
	YearAgoChangePercent   float32 `json:"yearAgoChangePercent"`
	EstimatedChangePercent float32 `json:"estimatedChangePercent"`
	SymbolId               int32   `json:"symbolId"`
}

type BTO struct {
	Earnings
}

type AMC struct {
	Earnings
}
