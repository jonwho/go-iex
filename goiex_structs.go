package goiex

import (
	"net/http"
)

type Batch struct {
	Quote Quote
	News  News
	Chart Chart
}

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

type NewsDTO struct {
	Datetime string `json:"datetime"`
	Headline string `json:"headline"`
	Source   string `json:"source"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	Related  string `json:"related"`
	Image    string `json:"image"`
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

type News struct {
	News []NewsDTO
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

type RefDataCorporateActions struct {
	CorporateActions []CorporateActionDTO
}

type CorporateActionDTO struct {
	RecordID                     string `json:"RecordID"`
	DailyListTimestamp           string `json:"DailyListTimestamp"`
	EffectiveDate                string `json:"EffectiveDate"`
	IssueEvent                   string `json:"IssueEvent"`
	CurrentSymbolinINETSymbology string `json:"CurrentSymbolinINETSymbology"`
	CurrentSymbolinCQSSymbology  string `json:"CurrentSymbolinCQSSymbology"`
	CurrentSymbolinCMSSymbology  string `json:"CurrentSymbolinCMSSymbology"`
	NewSymbolinINETSymbology     string `json:"NewSymbolinINETSymbology"`
	NewSymbolinCQSSymbology      string `json:"NewSymbolinCQSSymbology"`
	NewSymbolinCMSSymbology      string `json:"NewSymbolinCMSSymbology"`
	CurrentSecurityName          string `json:"CurrentSecurityName"`
	NewSecurityName              string `json:"NewSecurityName"`
	CurrentCompanyName           string `json:"CurrentCompanyName"`
	NewCompanyName               string `json:"NewCompanyName"`
	CurrentListingCenter         string `json:"CurrentListingCenter"`
	NewListingCenter             string `json:"NewListingCenter"`
	DelistingReason              string `json:"DelistingReason"`
	CurrentRoundLotSize          string `json:"CurrentRoundLotSize"`
	NewRoundLotSize              string `json:"NewRoundLotSize"`
	CurrentLULDTierIndicator     string `json:"CurrentLULDTierIndicator"`
	NewLULDTierIndicator         string `json:"NewLULDTierIndicator"`
	ExpirationDate               string `json:"ExpirationDate"`
	SeparationDate               string `json:"SeparationDate"`
	SettlementDate               string `json:"SettlementDate"`
	MaturityDate                 string `json:"MaturityDate"`
	RedemptionDate               string `json:"RedemptionDate"`
	CurrentFinancialStatus       string `json:"CurrentFinancialStatus"`
	NewFinancialStatus           string `json:"NewFinancialStatus"`
	WhenIssuedFlag               string `json:"WhenIssuedFlag"`
	WhenDistributedFlag          string `json:"WhenDistributedFlag"`
	IPOFlag                      string `json:"IPOFlag"`
	NotesforEachEntry            string `json:"NotesforEachEntry"`
	RecordUpdateTime             string `json:"RecordUpdateTime"`
}

type RefDataDividends struct {
	Dividends []DividendDTO
}

type DividendDTO struct {
	RecordID                 string `json:"RecordID"`
	DailyListTimestamp       string `json:"DailyListTimestamp"`
	EventType                string `json:"EventType"`
	SymbolinINETSymbology    string `json:"SymbolinINETSymbology"`
	SymbolinCQSSymbology     string `json:"SymbolinCQSSymbology"`
	SymbolinCMSSymbology     string `json:"SymbolinCMSSymbology"`
	SecurityName             string `json:"SecurityName"`
	CompanyName              string `json:"CompanyName"`
	DeclarationDate          string `json:"DeclarationDate"`
	AmountDescription        string `json:"AmountDescription"`
	PaymentFrequency         string `json:"PaymentFrequency"`
	ExDate                   string `json:"ExDate"`
	RecordDate               string `json:"RecordDate"`
	PaymentDate              string `json:"PaymentDate"`
	DividendTypeID           string `json:"DividendTypeID"`
	StockAdjustmentFactor    string `json:"StockAdjustmentFactor"`
	StockAmount              string `json:"StockAmount"`
	CashAmount               string `json:"CashAmount"`
	PostSplitShares          string `json:"PostSplitShares"`
	PreSplitShares           string `json:"PreSplitShares"`
	QualifiedDividend        string `json:"QualifiedDividend"`
	ExercisePriceAmount      string `json:"ExercisePriceAmount"`
	ElectionorExpirationDate string `json:"ElectionorExpirationDate"`
	GrossAmount              string `json:"GrossAmount"`
	NetAmount                string `json:"NetAmount"`
	BasisNotes               string `json:"BasisNotes"`
	NotesforEachEntry        string `json:"NotesforEachEntry"`
	RecordUpdateTime         string `json:"RecordUpdateTime"`
}

type RefDataNextDayExDates struct {
	NextDayExDates []NextDayExDateDTO
}

type NextDayExDateDTO struct {
	RecordID                 string `json:"RecordID"`
	DailyListTimestamp       string `json:"DailyListTimestamp"`
	ExDate                   string `json:"ExDate"`
	SymbolinINETSymbology    string `json:"SymbolinINETSymbology"`
	SymbolinCQSSymbology     string `json:"SymbolinCQSSymbology"`
	SymbolinCMSSymbology     string `json:"SymbolinCMSSymbology"`
	SecurityName             string `json:"SecurityName"`
	CompanyName              string `json:"CompanyName"`
	DividendTypeID           string `json:"DividendTypeID"`
	AmountDescription        string `json:"AmountDescription"`
	PaymentFrequency         string `json:"PaymentFrequency"`
	StockAdjustmentFactor    string `json:"StockAdjustmentFactor"`
	StockAmount              string `json:"StockAmount"`
	CashAmount               string `json:"CashAmount"`
	PostSplitShares          string `json:"PostSplitShares"`
	PreSplitShares           string `json:"PreSplitShares"`
	QualifiedDividend        string `json:"QualifiedDividend"`
	ExercisePriceAmount      string `json:"ExercisePriceAmount"`
	ElectionorExpirationDate string `json:"ElectionorExpirationDate"`
	GrossAmount              string `json:"GrossAmount"`
	NetAmount                string `json:"NetAmount"`
	BasisNotes               string `json:"BasisNotes"`
	NotesforEachEntry        string `json:"NotesforEachEntry"`
	RecordUpdateTime         string `json:"RecordUpdateTime"`
}

type RefDataSymbolDirectories struct {
	SymbolDirectories []SymbolDirectoryDTO
}

type SymbolDirectoryDTO struct {
	RecordID                             string `json:"RecordID"`
	DailyListTimestamp                   string `json:"DailyListTimestamp"`
	SymbolinINETSymbology                string `json:"SymbolinINETSymbology"`
	SymbolinCQSSymbology                 string `json:"SymbolinCQSSymbology"`
	SymbolinCMSSymbology                 string `json:"SymbolinCMSSymbology"`
	SecurityName                         string `json:"SecurityName"`
	CompanyName                          string `json:"CompanyName"`
	TestIssue                            string `json:"TestIssue"`
	IssueDescription                     string `json:"IssueDescription"`
	IssueType                            string `json:"IssueType"`
	IssueSubType                         string `json:"IssueSubType"`
	SICCode                              string `json:"SICCode"`
	TransferAgent                        string `json:"TransferAgent"`
	FinancialStatus                      string `json:"FinancialStatus"`
	RoundLotSize                         string `json:"RoundLotSize"`
	PreviousOfficialClosingPrice         string `json:"PreviousOfficialClosingPrice"`
	AdjustedPreviousOfficialClosingPrice string `json:"AdjustedPreviousOfficialClosingPrice"`
	WhenIssuedFlag                       string `json:"WhenIssuedFlag"`
	WhenDistributedFlag                  string `json:"WhenDistributedFlag"`
	IPOFlag                              string `json:"IPOFlag"`
	FirstDateListed                      string `json:"FirstDateListed"`
	LULDTierIndicator                    string `json:"LULDTierIndicator"`
	CountryofIncorporation               string `json:"CountryofIncorporation"`
	LeveragedETPFlag                     string `json:"LeveragedETPFlag"`
	LeveragedETPRatio                    string `json:"LeveragedETPRatio"`
	InverseETPFlag                       string `json:"InverseETPFlag"`
	RecordUpdateTime                     string `json:"RecordUpdateTime"`
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
