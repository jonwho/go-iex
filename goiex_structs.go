package goiex

import (
	"net/http"
)

type AskDTO struct {
	Price     float32
	Size      int32
	Timestamp float64
}

type BidDTO struct {
	Price     float32
	Size      int32
	Timestamp float64
}

type TradeDTO struct {
	Price                 float32
	Size                  int32
	TradeID               int32
	IsISO                 bool
	IsOddLot              bool
	IsOutsideRegularHours bool
	IsSinglePriceCross    bool
	IsTradeThroughExempt  bool
	Timestamp             float64
}

type SystemEvent struct {
	SystemEvent string
	Timestamp   float64
}

type Book struct {
	Quote       Quote
	Bids        []BidDTO
	Asks        []AskDTO
	Trades      []TradeDTO
	SystemEvent SystemEvent
}

type Chart struct {
	Charts []ChartDTO
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

type Earnings struct {
	Symbol   string
	Earnings []EarningsDTO
}

type EarningsToday struct {
	BTO []EarningsReportDTO
	AMC []EarningsReportDTO
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

type RefDataSymbols struct {
	Symbols []SymbolDTO
}

type SymbolDTO struct {
	Symbol    string
	Name      string
	Date      string
	IsEnabled bool
	Type      string
	// iex API returns iexId as string sometimes
	// UnmarshalJSON for SymbolDTO will use type switching to convert iexId to int
	IexId interface{}
}

type ChartDTO struct {
	Date                 string
	Minute               string
	Label                string
	High                 float32
	Low                  float32
	Average              float32
	Volume               int32
	Notional             float32
	NumberOfTrades       int32
	MarketHigh           float32
	MarketLow            float32
	MarketAverage        float32
	MarketVolume         int64
	MarketNotional       float64
	MarketNumberOfTrades int32
	Open                 float32
	Close                float32
	MarketOpen           float32
	MarketClose          float32
	ChangeOverTime       float32
	MarketChangeOverTime float32
}

type EarningsReportDTO struct {
	EarningsDTO
	Symbol   string
	Headline string
	Quote    Quote
}

type EarningsDTO struct {
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

type KeyStat struct {
	KeyStatDTO
}

type KeyStatDTO struct {
	CompanyName         string
	MarketCap           int64
	Beta                float32
	Week52High          float32
	Week52Low           float32
	Week52Change        float32
	ShortInterest       int32
	ShortDate           int64
	DividendRate        float32
	DividendYield       float32
	ExDividendDate      interface{}
	LatestEPS           float32
	LatestEPSDate       string
	SharesOutstanding   int64
	Float               int64
	ReturnOnEquity      float32
	ConsensusEPS        float32
	NumberOfEstimates   int32
	Symbol              string
	EBITDA              int64
	Revenue             int64
	GrossProfit         int64
	Cash                int64
	Debt                int64
	TtmEPS              float32
	RevenuePerShare     float32
	RevenuePerEmployee  float32
	PeRatioHigh         float32
	PeRatioLow          float32
	EPSSurpriseDollar   interface{}
	EPSSurprisePercent  float32
	ReturnOnAssets      float32
	ReturnOnCapital     interface{}
	ProfitMargin        float32
	PriceToSales        float32
	PriceToBook         float32
	Day200MovingAvg     float32
	Day50MovingAvg      float32
	InstitutionPercent  float32
	InsiderPercent      interface{}
	ShortRatio          interface{}
	Year5ChangePercent  float32
	Year2ChangePercent  float32
	Year1ChangePercent  float32
	YtdChangePercent    float32
	Month6ChangePercent float32
	Month3ChangePercent float32
	Month1ChangePercent float32
	Day5ChangePercent   float32
}
