package goiex

import (
	"net/http"
)

type Chart struct {
	Charts []chart
}

type Client struct {
	httpClient *http.Client
}

type Earnings struct {
	Symbol   string
	Earnings []earnings
}

type EarningsToday struct {
	Bto []earningsReport
	Amc []earningsReport
}

type Quote struct {
	Symbol                string
	CompanyName           string
	PrimaryExchange       string
	Sector                string
	CalculationPrice      string
	Open                  float32
	OpenTime              int64
	Close                 float32
	CloseTime             int64
	High                  float32
	Low                   float32
	LatestPrice           float32
	LatestSource          string
	LatestTime            string
	LatestUpdate          int64
	LatestVolume          int32
	IexRealtimePrice      float32
	IexRealtimeSize       int32
	IexLastUpdated        int64
	DelayedPrice          float32
	DelayedPriceTime      int64
	ExtendedPrice         float32
	ExtendedChange        float32
	ExtendedChangePercent float32
	ExtendedPriceTime     int64
	PreviousClose         float32
	Change                float32
	ChangePercent         float32
	IexMarketPercent      float32
	IexVolume             int32
	AvgTotalVolume        int32
	IexBidPrice           float32
	IexBidSize            int32
	IexAskPrice           float32
	IexAskSize            int32
	MarketCap             int64
	PeRatio               float32
	Week52High            float32
	Week52Low             float32
	YtdChange             float32
}

type chart struct {
	Date             string
	Open             float32
	High             float32
	Low              float32
	Close            float32
	Volume           int64
	UnadjustedVolume int64
	Change           float32
	ChangePercent    float32
	Vwap             float32
	Label            string
	ChangeOverTime   float32
}

type earningsReport struct {
	earnings
	Symbol   string
	Headline string
	Quote    Quote
}

type earnings struct {
	ActualEPS              float32
	ConcensusEPS           float32
	EstimatedEPS           float32
	AnnounceTime           string
	NumberOfEstimates      int32
	EPSSurpriseDollar      float32
	EPSReportDate          string
	FiscalPeriod           string
	FiscalEndDate          string
	YearAgo                float32
	YearAgoChangePercent   float32
	EstimatedChangePercent float32
	SymbolId               int32
}
