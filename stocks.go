package goiex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// IndicatorName for TechnicalIndicator API
type IndicatorName int

const (
	// ABS Vector Absolute Value
	ABS IndicatorName = iota
	// ACOS Vector Arccosine
	ACOS
	// AD Accumulation/Distribution Line
	AD
	// ADD Vector Addition
	ADD
	// ADOSC Accumulation/Distribution Oscillator
	ADOSC
	// ADX Average Directional Movement Index
	ADX
	// ADXR Average Directional Movement Rating
	ADXR
	// AO Awesome Oscillator
	AO
	// APO Absolute Price Oscillator
	APO
	// AROON Aroon
	AROON
	// AROONOSC Aroon Oscillator
	AROONOSC
	// ASIN Vector Arcsine
	ASIN
	// ATAN Vector Arctangent
	ATAN
	// ATR Average True Range
	ATR
	// AVGPRICE Average Price
	AVGPRICE
	// BBANDS Bollinger Bands
	BBANDS
	// BOP Balance of Power
	BOP
	// CCI Commodity Channel Index
	CCI
	// CEIL Vector Ceiling
	CEIL
	// CMO Change Momentum Oscillator
	CMO
	// COS Vector Cosine
	COS
	// COSH Vector Hyperbolic Cosine
	COSH
	// CROSSANY Crossany
	CROSSANY
	// CROSSOVER Crossover
	CROSSOVER
	// CVI Chaikins Volatility
	CVI
	// DECAY Linear Decay
	DECAY
	// DEMA Double Exponential Moving Average
	DEMA
	// DI Directional Indicator
	DI
	// DIV Vector Division
	DIV
	// DM Directional Movement
	DM
	// DPO Detrended Price Oscillator
	DPO
	// DX Directional Movement Index
	DX
	// EDECAY Exponential Decay
	EDECAY
	// EMA Exponential Moving Average
	EMA
	// EMV Ease of Movement
	EMV
	// EXP Vector Exponential
	EXP
	// FISHER Fisher Transform
	FISHER
	// FLOOR Vector Floor
	FLOOR
	// FOSC Forecast Oscillator
	FOSC
	// HMA Hull Moving Average
	HMA
	// KAMA Kaufman Adaptive Moving Average
	KAMA
	// KVO Klinger Volume Oscillator
	KVO
	// LAG Lag
	LAG
	// LINREG Linear Regression
	LINREG
	// LINREGINTERCEPT Linear Regression Intercept
	LINREGINTERCEPT
	// LINREGSLOPE Linear Regression Slope
	LINREGSLOPE
	// LN Vector Natural Log
	LN
	// LOG10 Vector Base-10 Log
	LOG10
	// MACD Moving Average Convergence/Divergence
	MACD
	// MARKETFI Market Facilitation Index
	MARKETFI
	// MASS Mass Index
	MASS
	// MAX Maximum In Period
	MAX
	// MD Mean Deviation Over Period
	MD
	// MEDPRICE Median Price
	MEDPRICE
	// MFI Money Flow Index
	MFI
	// MIN Minimum In Period
	MIN
	// MOM Momentum
	MOM
	// MSW Mesa Sine Wave
	MSW
	// MUL Vector Multiplication
	MUL
	// NATR Normalized Average True Range
	NATR
	// NVI Negative Volume Index
	NVI
	// OBV On Balance Volume
	OBV
	// PPO Percentage Price Oscillator
	PPO
	// PSAR Parabolic SAR
	PSAR
	// PVI Positive Volume Index
	PVI
	// QSTICK Qstick
	QSTICK
	// ROC Rate of Change
	ROC
	// ROCR Rate of Change Ratio
	ROCR
	// ROUND Vector Round
	ROUND
	// RSI Relative Strength Index
	RSI
	// SIN Vector Sine
	SIN
	// SINH Vector Hyperbolic Sine
	SINH
	// SMA Simple Moving Average
	SMA
	// SQRT Vector Square Root
	SQRT
	// STDDEV Standard Deviation Over Period
	STDDEV
	// STOCH Stochastic Oscillator
	STOCH
	// STOCHRSI Stochastic RSI
	STOCHRSI
	// SUB Vector Subtraction
	SUB
	// SUM Sum Over Period
	SUM
	// TAN Vector Tangent
	TAN
	// TANH Vector Hyperbolic Tangent
	TANH
	// TEMA Triple Exponential Moving Average
	TEMA
	// TODEG Vector Degree Conversion
	TODEG
	// TORAD Vector Radian Conversion
	TORAD
	// TR True Range
	TR
	// TRIMA Triangular Moving Average
	TRIMA
	// TRIX Trix
	TRIX
	// TRUNC Vector Truncate
	TRUNC
	// TSF Time Series Forecast
	TSF
	// TYPPRICE Typical Price
	TYPPRICE
	// ULTOSC Ultimate Oscillator
	ULTOSC
	// VAR Variance Over Period
	VAR
	// VHF Vertical Horizontal Filter
	VHF
	// VIDYA Variable Index Dynamic Average
	VIDYA
	// VOLATILITY Annualized Historical Volatility
	VOLATILITY
	// VOSC Volume Oscillator
	VOSC
	// VWMA Volume Weighted Moving Average
	VWMA
	// WAD Williams Accumulation/Distribution
	WAD
	// WCPRICE Weight Close Price
	WCPRICE
	// WILDERS Wilders Smoothing
	WILDERS
	// WILLR Williams %R
	WILLR
	// WMA Weighted Moving Average
	WMA
	// ZLEMA Zero-Lag Exponential Moving Average
	ZLEMA
)

// ChartRange for Chart API
type ChartRange int

const (
	// ChartRangeMax chart range
	ChartRangeMax ChartRange = iota
	// ChartRangeFiveYear chart range
	ChartRangeFiveYear
	// ChartRangeTwoYear chart range
	ChartRangeTwoYear
	// ChartRangeOneYear chart range
	ChartRangeOneYear
	// ChartRangeYearToDate chart range
	ChartRangeYearToDate
	// ChartRangeSixMonth chart range
	ChartRangeSixMonth
	// ChartRangeThreeMonth chart range
	ChartRangeThreeMonth
	// ChartRangeOneMonth chart range
	ChartRangeOneMonth
	// ChartRangeOneDay chart range
	ChartRangeOneDay
)

// ChartQueryParams optional query parameters
type ChartQueryParams struct {
	ChartCloseOnly  bool       `url:"chartCloseOnly,omitempty"`
	ChartByDay      bool       `url:"chartByDay,omitempty"`
	ChartSimplify   bool       `url:"chartSimplify,omitempty"`
	ChartInterval   uint       `url:"chartInterval,omitempty"`
	ChangeFromClose bool       `url:"changeFromClose,omitempty"`
	ChartLast       uint       `url:"chartLast,omitempty"`
	Range           ChartRange `url:"range,omitempty"`
	// ExactDate date formatted as YYYYMMDD
	ExactDate string `url:"exactDate,omitempty"`
	// Sort can be `asc` or `desc`. Defaults to `desc`.
	Sort string `url:"sort,omitempty"`
	// IncludeToday appends current trading to data
	IncludeToday bool `url:"includeToday,omitempty"`
}

// DividendRange for Dividend API
type DividendRange int

const (
	// DividendRangeFiveYear dividend range
	DividendRangeFiveYear DividendRange = iota
	// DividendRangeTwoYear dividend range
	DividendRangeTwoYear
	// DividendRangeOneYear dividend range
	DividendRangeOneYear
	// DividendRangeYearToDate dividend range
	DividendRangeYearToDate
	// DividendRangeSixMonth dividend range
	DividendRangeSixMonth
	// DividendRangeThreeMonth dividend range
	DividendRangeThreeMonth
	// DividendRangeOneMonth dividend range
	DividendRangeOneMonth
	// DividendRangeNext dividend range
	DividendRangeNext
)

// SplitRange for Split API
type SplitRange int

const (
	// SplitRangeFiveYear split range
	SplitRangeFiveYear SplitRange = iota
	// SplitRangeTwoYear split range
	SplitRangeTwoYear
	// SplitRangeOneYear split range
	SplitRangeOneYear
	// SplitRangeYearToDate split range
	SplitRangeYearToDate
	// SplitRangeSixMonth split range
	SplitRangeSixMonth
	// SplitRangeThreeMonth split range
	SplitRangeThreeMonth
	// SplitRangeOneMonth split range
	SplitRangeOneMonth
	// SplitRangeNext split range
	SplitRangeNext
)

// PeriodQueryParameter accepted values for query parameter `period`
type PeriodQueryParameter int

const (
	// PeriodAnnual annual period
	PeriodAnnual PeriodQueryParameter = iota
	// PeriodQuarter quarter period
	PeriodQuarter
)

// Stock struct to interface with /stock endpoints
type Stock struct {
	iex
}

// AdvancedStat struct
type AdvancedStat struct {
	KeyStat
	TotalCash                int64       `json:"totalCash"`
	CurrentDebt              int64       `json:"currentDebt"`
	Revenue                  int64       `json:"revenue"`
	GrossProfit              int64       `json:"grossProfit"`
	TotalRevenue             int64       `json:"totalRevenue"`
	EBITDA                   int64       `json:"EBITDA"`
	RevenuePerShare          float64     `json:"revenuePerShare"`
	RevenuePerEmployee       float64     `json:"revenuePerEmployee"`
	DebtToEquity             float64     `json:"debtToEquity"`
	ProfitMargin             float64     `json:"profitMargin"`
	EnterpriseValue          int64       `json:"enterpriseValue"`
	EnterpriseValueToRevenue float64     `json:"enterpriseValueToRevenue"`
	PriceToSales             float64     `json:"priceToSales"`
	PriceToBook              float64     `json:"priceToBook"`
	ForwardPERatio           interface{} `json:"forwardPERatio"`
	PegRatio                 float64     `json:"pegRatio"`
}

// Asks struct
type Asks []struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Timestamp int64   `json:"timestamp"`
}

// BalanceSheetParams query parameters
type BalanceSheetParams struct {
	// Period specify either "annual" or "quarter" with PeriodQueryParameter
	Period PeriodQueryParameter `url:"period"`
	// Last with "quarter" period can specify up to 12 and up to 4 with "annual" period
	Last int `url:"last"`
}

// BalanceSheet struct
type BalanceSheet struct {
	Symbol       string `json:"symbol"`
	BalanceSheet []struct {
		ReportDate              string      `json:"reportDate"`
		CurrentCash             int64       `json:"currentCash"`
		ShortTermInvestments    int64       `json:"shortTermInvestments"`
		Receivables             int64       `json:"receivables"`
		Inventory               int64       `json:"inventory"`
		OtherCurrentAssets      int64       `json:"otherCurrentAssets"`
		CurrentAssets           int64       `json:"currentAssets"`
		LongTermInvestments     int64       `json:"longTermInvestments"`
		PropertyPlantEquipment  int64       `json:"propertyPlantEquipment"`
		Goodwill                interface{} `json:"goodwill"`
		IntangibleAssets        interface{} `json:"intangibleAssets"`
		OtherAssets             int64       `json:"otherAssets"`
		TotalAssets             int64       `json:"totalAssets"`
		AccountsPayable         int64       `json:"accountsPayable"`
		CurrentLongTermDebt     int64       `json:"currentLongTermDebt"`
		OtherCurrentLiabilities int64       `json:"otherCurrentLiabilities"`
		TotalCurrentLiabilities int64       `json:"totalCurrentLiabilities"`
		LongTermDebt            int64       `json:"longTermDebt"`
		OtherLiabilities        int64       `json:"otherLiabilities"`
		MinorityInterest        int         `json:"minorityInterest"`
		TotalLiabilities        int64       `json:"totalLiabilities"`
		CommonStock             int64       `json:"commonStock"`
		RetainedEarnings        int64       `json:"retainedEarnings"`
		TreasuryStock           interface{} `json:"treasuryStock"`
		CapitalSurplus          interface{} `json:"capitalSurplus"`
		ShareholderEquity       int64       `json:"shareholderEquity"`
		NetTangibleAssets       int64       `json:"netTangibleAssets"`
	} `json:"balancesheet"`
}

// Batch struct
type Batch struct {
	Quote Quote
	News  News
	Chart []Chart
}

// Bids struct
type Bids []struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Timestamp int64   `json:"timestamp"`
}

// Book struct
type Book struct {
	Asks        Asks
	Bids        Bids
	Quote       Quote
	Trades      Trades
	SystemEvent SystemEvent
}

// CashFlow struct
type CashFlow struct {
	Symbol   string `json:"symbol"`
	CashFlow []struct {
		ReportDate              string      `json:"reportDate"`
		NetIncome               int64       `json:"netIncome"`
		Depreciation            int64       `json:"depreciation"`
		ChangesInReceivables    int64       `json:"changesInReceivables"`
		ChangesInInventories    int         `json:"changesInInventories"`
		CashChange              int64       `json:"cashChange"`
		CashFlow                int64       `json:"cashFlow"`
		CapitalExpenditures     int64       `json:"capitalExpenditures"`
		Investments             int         `json:"investments"`
		InvestingActivityOther  int         `json:"investingActivityOther"`
		TotalInvestingCashFlows int64       `json:"totalInvestingCashFlows"`
		DividendsPaid           int64       `json:"dividendsPaid"`
		NetBorrowings           int         `json:"netBorrowings"`
		OtherFinancingCashFlows int         `json:"otherFinancingCashFlows"`
		CashFlowFinancing       int64       `json:"cashFlowFinancing"`
		ExchangeRateEffect      interface{} `json:"exchangeRateEffect"`
	} `json:"cashflow"`
}

// Chart struct
type Chart struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Close          float64 `json:"close"`
	Volume         int     `json:"volume"`
	UOpen          float64 `json:"uOpen"`
	UHigh          float64 `json:"uHigh"`
	ULow           float64 `json:"uLow"`
	UClose         float64 `json:"uClose"`
	UVolume        int     `json:"uVolume"`
	Change         float64 `json:"change"`
	ChangePercent  float64 `json:"changePercent"`
	Label          string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

// Collection struct
type Collection []struct {
	Quote
}

// Company struct
type Company struct {
	Symbol      string   `json:"symbol"`
	CompanyName string   `json:"companyName"`
	Employees   int      `json:"employees"`
	Exchange    string   `json:"exchange"`
	Industry    string   `json:"industry"`
	Website     string   `json:"website"`
	Description string   `json:"description"`
	CEO         string   `json:"CEO"`
	IssueType   string   `json:"issueType"`
	Sector      string   `json:"sector"`
	Tags        []string `json:"tags"`
}

// DelayedQuote struct
type DelayedQuote struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	DelayedSize      int     `json:"delayedSize"`
	DelayedPriceTime int64   `json:"delayedPriceTime"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	TotalVolume      int     `json:"totalVolume"`
	ProcessedTime    int64   `json:"processedTime"`
}

// Dividends struct {
type Dividends []struct {
	Symbol       string  `json:"symbol"`
	ExDate       string  `json:"exDate"`
	PaymentDate  string  `json:"paymentDate"`
	RecordDate   string  `json:"recordDate"`
	DeclaredDate string  `json:"declaredDate"`
	Amount       float64 `json:"amount,string"`
	Flag         string  `json:"flag"`
	Currency     string  `json:"currency"`
	Description  string  `json:"description"`
	Frequency    string  `json:"frequency"`
}

// Earnings struct
type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []struct {
		ActualEPS            float64 `json:"actualEPS"`
		ConsensusEPS         float64 `json:"consensusEPS"`
		AnnounceTime         string  `json:"announceTime"`
		NumberOfEstimates    int     `json:"numberOfEstimates"`
		EPSSurpriseDollar    float64 `json:"EPSSurpriseDollar"`
		EPSReportDate        string  `json:"EPSReportDate"`
		FiscalPeriod         string  `json:"fiscalPeriod"`
		FiscalEndDate        string  `json:"fiscalEndDate"`
		YearAgo              float64 `json:"yearAgo"`
		YearAgoChangePercent float64 `json:"yearAgoChangePercent"`
	} `json:"earnings"`
}

// EarningsToday struct
type EarningsToday struct {
	BTO   []EarningsTodayDTO `json:"bto"`
	AMC   []EarningsTodayDTO `json:"amc"`
	DMT   []EarningsTodayDTO `json:"dmt"`
	Other []EarningsTodayDTO `json:"other"`
}

// EarningsTodayDTO struct
type EarningsTodayDTO struct {
	ConsensusEPS      float64 `json:"consensusEPS"`
	AnnounceTime      string  `json:"announceTime"`
	NumberOfEstimates int     `json:"numberOfEstimates"`
	FiscalPeriod      string  `json:"fiscalPeriod"`
	FiscalEndDate     string  `json:"fiscalEndDate"`
	Symbol            string  `json:"symbol"`
	Quote             Quote   `json:"quote"`
}

// Estimates struct
type Estimates struct {
	Symbol    string `json:"symbol"`
	Estimates []struct {
		ConsensusEPS      float64 `json:"consensusEPS"`
		NumberOfEstimates int     `json:"numberOfEstimates"`
		FiscalPeriod      string  `json:"fiscalPeriod"`
		FiscalEndDate     string  `json:"fiscalEndDate"`
		ReportDate        string  `json:"reportDate"`
	} `json:"estimates"`
}

// Financials struct
type Financials struct {
	Symbol     string `json:"symbol"`
	Financials []struct {
		ReportDate             string      `json:"reportDate"`
		GrossProfit            int64       `json:"grossProfit"`
		CostOfRevenue          int64       `json:"costOfRevenue"`
		OperatingRevenue       int64       `json:"operatingRevenue"`
		TotalRevenue           int64       `json:"totalRevenue"`
		OperatingIncome        int64       `json:"operatingIncome"`
		NetIncome              int64       `json:"netIncome"`
		ResearchAndDevelopment int64       `json:"researchAndDevelopment"`
		OperatingExpense       int64       `json:"operatingExpense"`
		CurrentAssets          int64       `json:"currentAssets"`
		TotalAssets            int64       `json:"totalAssets"`
		TotalLiabilities       int64       `json:"totalLiabilities"`
		CurrentCash            int64       `json:"currentCash"`
		CurrentDebt            int64       `json:"currentDebt"`
		TotalCash              int64       `json:"totalCash"`
		TotalDebt              int64       `json:"totalDebt"`
		ShareholderEquity      int64       `json:"shareholderEquity"`
		CashChange             int         `json:"cashChange"`
		CashFlow               int64       `json:"cashFlow"`
		OperatingGainsLosses   interface{} `json:"operatingGainsLosses"`
	} `json:"financials"`
}

// FundOwnership struct
type FundOwnership []struct {
	AdjHolding       int    `json:"adjHolding"`
	AdjMv            int    `json:"adjMv"`
	EntityProperName string `json:"entityProperName"`
	ReportDate       int64  `json:"reportDate"`
	ReportedHolding  int    `json:"reportedHolding"`
	ReportedMv       int    `json:"reportedMv"`
}

// IncomeStatement struct
type IncomeStatement struct {
	Symbol string `json:"symbol"`
	Income []struct {
		ReportDate             string `json:"reportDate"`
		TotalRevenue           int64  `json:"totalRevenue"`
		CostOfRevenue          int64  `json:"costOfRevenue"`
		GrossProfit            int64  `json:"grossProfit"`
		ResearchAndDevelopment int64  `json:"researchAndDevelopment"`
		SellingGeneralAndAdmin int64  `json:"sellingGeneralAndAdmin"`
		OperatingExpense       int64  `json:"operatingExpense"`
		OperatingIncome        int64  `json:"operatingIncome"`
		OtherIncomeExpenseNet  int    `json:"otherIncomeExpenseNet"`
		Ebit                   int64  `json:"ebit"`
		InterestIncome         int    `json:"interestIncome"`
		PretaxIncome           int64  `json:"pretaxIncome"`
		IncomeTax              int64  `json:"incomeTax"`
		MinorityInterest       int    `json:"minorityInterest"`
		NetIncome              int64  `json:"netIncome"`
		NetIncomeBasic         int64  `json:"netIncomeBasic"`
	} `json:"income"`
}

// InsiderRoster struct
type InsiderRoster []struct {
	EntityName string `json:"entityName'"`
	Position   int    `json:"position"`
	ReportDate int64  `json:"reportDate"`
}

// InsiderSummary struct
type InsiderSummary []struct {
	FullName      string `json:"fullName"`
	NetTransacted int    `json:"netTransacted"`
	ReportedTitle string `json:"reportedTitle"`
	TotalBought   int    `json:"totalBought"`
	TotalSold     int    `json:"totalSold"`
}

// InsiderTransactions struct
type InsiderTransactions []struct {
	EffectiveDate int64   `json:"effectiveDate"`
	FullName      string  `json:"fullName"`
	ReportedTitle string  `json:"reportedTitle"`
	TranPrice     float64 `json:"tranPrice"`
	TranShares    float64 `json:"tranShares"`
	TranValue     float64 `json:"tranValue"`
}

// InstitutionalOwnership struct
type InstitutionalOwnership []struct {
	AdjHolding       int    `json:"adjHolding"`
	AdjMv            int    `json:"adjMv"`
	EntityProperName string `json:"entityProperName"`
	ReportDate       int64  `json:"reportDate"`
	ReportedHolding  int    `json:"reportedHolding"`
}

// IntradayPrices struct
type IntradayPrices []struct {
	Date                 string  `json:"date"`
	Minute               string  `json:"minute"`
	Label                string  `json:"label"`
	MarktOpen            float64 `json:"marktOpen"`
	MarketClose          float64 `json:"marketClose"`
	MarktHigh            float64 `json:"marktHigh"`
	MarketLow            float64 `json:"marketLow"`
	MarketAverage        float64 `json:"marketAverage"`
	MarketVolume         int     `json:"marketVolume"`
	MarketNotional       float64 `json:"marketNotional"`
	MarketNumberOfTrades int     `json:"marketNumberOfTrades"`
	MarketChangeOverTime float64 `json:"marketChangeOverTime"`
	High                 float64 `json:"high"`
	Low                  float64 `json:"low"`
	Open                 float64 `json:"open"`
	Close                float64 `json:"close"`
	Average              float64 `json:"average"`
	Volume               int     `json:"volume"`
	Notional             float64 `json:"notional"`
	NumberOfTrades       int     `json:"numberOfTrades"`
	ChangeOverTime       float64 `json:"changeOverTime"`
}

// IPOCalendar struct
type IPOCalendar struct {
	RawData []struct {
		Symbol                 string      `json:"symbol"`
		CompanyName            string      `json:"companyName"`
		ExpectedDate           string      `json:"expectedDate"`
		LeadUnderwriters       []string    `json:"leadUnderwriters"`
		Underwriters           []string    `json:"underwriters"`
		CompanyCounsel         []string    `json:"companyCounsel"`
		UnderwriterCounsel     []string    `json:"underwriterCounsel"`
		Auditor                string      `json:"auditor"`
		Market                 string      `json:"market"`
		Cik                    string      `json:"cik"`
		Address                string      `json:"address"`
		City                   string      `json:"city"`
		State                  string      `json:"state"`
		Zip                    string      `json:"zip"`
		Phone                  string      `json:"phone"`
		Ceo                    string      `json:"ceo"`
		Employees              int         `json:"employees"`
		URL                    string      `json:"url"`
		Status                 string      `json:"status"`
		SharesOffered          int         `json:"sharesOffered"`
		PriceLow               float64     `json:"priceLow"`
		PriceHigh              float64     `json:"priceHigh"`
		OfferAmount            interface{} `json:"offerAmount"`
		TotalExpenses          float64     `json:"totalExpenses"`
		SharesOverAlloted      int         `json:"sharesOverAlloted"`
		ShareholderShares      interface{} `json:"shareholderShares"`
		SharesOutstanding      int         `json:"sharesOutstanding"`
		LockupPeriodExpiration string      `json:"lockupPeriodExpiration"`
		QuietPeriodExpiration  string      `json:"quietPeriodExpiration"`
		Revenue                int         `json:"revenue"`
		NetIncome              int         `json:"netIncome"`
		TotalAssets            int         `json:"totalAssets"`
		TotalLiabilities       int         `json:"totalLiabilities"`
		StockholderEquity      int         `json:"stockholderEquity"`
		CompanyDescription     string      `json:"companyDescription"`
		BusinessDescription    string      `json:"businessDescription"`
		UseOfProceeds          string      `json:"useOfProceeds"`
		Competition            string      `json:"competition"`
		Amount                 int         `json:"amount"`
		PercentOffered         string      `json:"percentOffered"`
	} `json:"rawData"`
	ViewData []struct {
		Company  string `json:"Company"`
		Symbol   string `json:"Symbol"`
		Price    string `json:"Price"`
		Shares   string `json:"Shares"`
		Amount   string `json:"Amount"`
		Float    string `json:"Float"`
		Percent  string `json:"Percent"`
		Market   string `json:"Market"`
		Expected string `json:"Expected"`
	} `json:"viewData"`
}

// KeyStat struct
type KeyStat struct {
	CompanyName         string  `json:"companyName"`
	Marketcap           int64   `json:"marketcap"`
	Week52High          float64 `json:"week52high"`
	Week52Low           float64 `json:"week52low"`
	Week52Change        float64 `json:"week52change"`
	SharesOutstanding   int64   `json:"sharesOutstanding"`
	Float               float64 `json:"float"`
	Symbol              string  `json:"symbol"`
	Avg10Volume         float64 `json:"avg10Volume"`
	Avg30Volume         float64 `json:"avg30Volume"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	Employees           int     `json:"employees"`
	TtmEPS              float64 `json:"ttmEPS"`
	TtmDividendRate     float64 `json:"ttmDividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	NextDividendDate    string  `json:"nextDividendDate"`
	ExDividendDate      string  `json:"exDividendDate"`
	NextEarningsDate    string  `json:"nextEarningsDate"`
	PeRatio             float64 `json:"peRatio"`
	Beta                float64 `json:"beta"`
	MaxChangePercent    float64 `json:"maxChangePercent"`
	Year5ChangePercent  float64 `json:"year5ChangePercent"`
	Year2ChangePercent  float64 `json:"year2ChangePercent"`
	Year1ChangePercent  float64 `json:"year1ChangePercent"`
	YtdChangePercent    float64 `json:"ytdChangePercent"`
	Month6ChangePercent float64 `json:"month6ChangePercent"`
	Month3ChangePercent float64 `json:"month3ChangePercent"`
	Month1ChangePercent float64 `json:"month1ChangePercent"`
	Day30ChangePercent  float64 `json:"day30ChangePercent"`
	Day5ChangePercent   float64 `json:"day5ChangePercent"`
}

// LargestTrades struct
type LargestTrades []struct {
	Price     float64 `json:"price"`
	Size      int     `json:"size"`
	Time      int64   `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}

// Logo struct
type Logo struct {
	URL string `json:"url"`
}

// MarketVolume struct
type MarketVolume []struct {
	Mic           string  `json:"mic"`
	TapeID        string  `json:"tapeId"`
	VenueName     string  `json:"venueName"`
	Volume        int     `json:"volume"`
	TapeA         int     `json:"tapeA"`
	TapeB         int     `json:"tapeB"`
	TapeC         int     `json:"tapeC"`
	MarketPercent float64 `json:"marketPercent"`
	LastUpdated   int64   `json:"lastUpdated"`
}

// News struct
type News []struct {
	Datetime   int64  `json:"datetime"`
	Headline   string `json:"headline"`
	Source     string `json:"source"`
	URL        string `json:"url"`
	Summary    string `json:"summary"`
	Related    string `json:"related"`
	Image      string `json:"image"`
	Lang       string `json:"lang"`
	HasPaywall bool   `json:"hasPaywall"`
}

// OHLC struct
type OHLC struct {
	Open struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"open"`
	Close struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"close"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

// Option struct
type Option struct {
	Symbol         string  `json:"symbol"`
	ID             string  `json:"id"`
	ExpirationDate string  `json:"expirationDate"`
	ContractSize   int     `json:"contractSize"`
	StrikePrice    int     `json:"strikePrice"`
	ClosingPrice   float64 `json:"closingPrice"`
	Side           string  `json:"side"`
	Type           string  `json:"type"`
	Volume         int     `json:"volume"`
	OpenInterest   int     `json:"openInterest"`
	Bid            float64 `json:"bid"`
	Ask            float64 `json:"ask"`
	LastUpdated    string  `json:"lastUpdated"`
}

// PreviousDayPrice struct
type PreviousDayPrice struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Volume         float64 `json:"volume"`
	UOpen          float64 `json:"uOpen"`
	UClose         float64 `json:"uClose"`
	UHigh          float64 `json:"uHigh"`
	ULow           float64 `json:"uLow"`
	UVolume        float64 `json:"uVolume"`
	Change         float64 `json:"change"`
	ChangePercent  float64 `json:"changePercent"`
	ChangeOverTime float64 `json:"changeOverTime"`
	Symbol         string  `json:"symbol"`
}

// PriceTarget struct
type PriceTarget struct {
	Symbol             string  `json:"symbol"`
	UpdatedDate        string  `json:"updatedDate"`
	PriceTargetAverage float64 `json:"priceTargetAverage"`
	PriceTargetHigh    float64 `json:"priceTargetHigh"`
	PriceTargetLow     float64 `json:"priceTargetLow"`
	NumberOfAnalysts   int     `json:"numberOfAnalysts"`
}

// Quote struct
type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              int64   `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             int64   `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          int64   `json:"latestUpdate"`
	LatestVolume          int     `json:"latestVolume"`
	IexRealtimePrice      float64 `json:"iexRealtimePrice"`
	IexRealtimeSize       int     `json:"iexRealtimeSize"`
	IexLastUpdated        int64   `json:"iexLastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      int64   `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     int64   `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	IexMarketPercent      float64 `json:"iexMarketPercent"`
	IexVolume             int     `json:"iexVolume"`
	AvgTotalVolume        int     `json:"avgTotalVolume"`
	IexBidPrice           float64 `json:"iexBidPrice"`
	IexBidSize            int     `json:"iexBidSize"`
	IexAskPrice           float64 `json:"iexAskPrice"`
	IexAskSize            int     `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YtdChange             float64 `json:"ytdChange"`
}

// RecommendationTrends struct
type RecommendationTrends []struct {
	ConsensusEndDate            int64   `json:"consensusEndDate"`
	ConsensusStartDate          int64   `json:"consensusStartDate"`
	CorporateActionsAppliedDate int64   `json:"corporateActionsAppliedDate"`
	RatingBuy                   int     `json:"ratingBuy"`
	RatingHold                  int     `json:"ratingHold"`
	RatingNone                  int     `json:"ratingNone"`
	RatingOverweight            int     `json:"ratingOverweight"`
	RatingScaleMark             float64 `json:"ratingScaleMark"`
	RatingSell                  int     `json:"ratingSell"`
	RatingUnderweight           int     `json:"ratingUnderweight"`
}

// SectorPerformance struct
type SectorPerformance []struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Performance float64 `json:"performance"`
	LastUpdated int64   `json:"lastUpdated"`
}

// Splits struct
type Splits []struct {
	ExDate       string  `json:"exDate"`
	DeclaredDate string  `json:"declaredDate"`
	Ratio        float64 `json:"ratio"`
	ToFactor     int     `json:"toFactor"`
	FromFactor   int     `json:"fromFactor"`
	Description  string  `json:"description"`
}

// SystemEvent struct
type SystemEvent struct {
	SystemEvent string `json:"systemEvent"`
	Timestamp   int64  `json:"timestamp"`
}

// TechnicalIndicator struct
type TechnicalIndicator struct {
	Indicator [][]float64 `json:"indicator"`
	Charts    []Chart     `json:"chart"`
}

// TechnicalIndicatorParams struct
type TechnicalIndicatorParams struct {
	// Range should match allowed ranges in historical prices
	Range  string `url:"range"`
	input1 int    `url:"input1,omitempty"`
	input2 int    `url:"input2,omitempty"`
	input3 int    `url:"input3,omitempty"`
	input4 int    `url:"input4,omitempty"`
}

// Trades struct
type Trades []struct {
	Price                 float64 `json:"price"`
	Size                  int     `json:"size"`
	TradeID               int     `json:"tradeId"`
	IsISO                 bool    `json:"isISO"`
	IsOddLot              bool    `json:"isOddLot"`
	IsOutsideRegularHours bool    `json:"isOutsideRegularHours"`
	IsSinglePriceCross    bool    `json:"isSinglePriceCross"`
	IsTradeThroughExempt  bool    `json:"isTradeThroughExempt"`
	Timestamp             int64   `json:"timestamp"`
}

// UpcomingEvents struct
type UpcomingEvents struct {
	IPOS     IPOCalendar `json:"ipos,omitempty"`
	Earnings []struct {
		ActualEPS            float64 `json:"actualEPS"`
		ConsensusEPS         float64 `json:"consensusEPS"`
		AnnounceTime         string  `json:"announceTime"`
		NumberOfEstimates    int     `json:"numberOfEstimates"`
		EPSSurpriseDollar    float64 `json:"EPSSurpriseDollar"`
		EPSReportDate        string  `json:"EPSReportDate"`
		FiscalPeriod         string  `json:"fiscalPeriod"`
		FiscalEndDate        string  `json:"fiscalEndDate"`
		YearAgo              float64 `json:"yearAgo"`
		YearAgoChangePercent float64 `json:"yearAgoChangePercent"`
	} `json:"earnings"`
	Dividends Dividends `json:"dividends"`
	Splits    Splits    `json:"splits"`
}

// UpcomingEarnings struct
type UpcomingEarnings []struct {
	Symbol     string `json:"symbol"`
	ReportDate string `json:"reportDate"`
}

// VolumeByVenue struct
type VolumeByVenue []struct {
	Volume           int     `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	MarketPercent    float64 `json:"marketPercent"`
	AvgMarketPercent float64 `json:"avgMarketPercent"`
	Date             string  `json:"date"`
}

// NewStock return new Stock
func NewStock(token, version string, base *url.URL, httpClient *http.Client) *Stock {
	apiurl, err := url.Parse("stock/")
	if err != nil {
		panic(err)
	}
	return &Stock{
		iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// Token return token string
func (s *Stock) Token() string {
	return s.token
}

// Version return version string
func (s *Stock) Version() string {
	return s.version
}

// URL return URL base
func (s *Stock) URL() *url.URL {
	return s.url
}

// APIURL return APIURL
func (s *Stock) APIURL() *url.URL {
	return s.apiurl
}

// Client return HTTP client
func (s *Stock) Client() *http.Client {
	return s.client
}

// AdvancedStats GET /stock/{symbol}/advanced-stats
func (s *Stock) AdvancedStats(symbol string) (advstat *AdvancedStat, err error) {
	err = get(s, &advstat, symbol+"/advanced-stats", nil)
	return
}

// BalanceSheet GET /stock/{symbol}/balance-sheet
func (s *Stock) BalanceSheet(symbol string, params *BalanceSheetParams) (balsheet *BalanceSheet, err error) {
	if params == nil {
		endpoint := fmt.Sprintf("%s/balance-sheet", symbol)
		err = get(s, &balsheet, endpoint, params)
		return
	}
	p := *params
	if p.Last == 0 {
		p.Last = 1
	}
	endpoint := fmt.Sprintf("%s/balance-sheet", symbol)
	err = get(s, &balsheet, endpoint, p)
	return
}

// Batch GET /stock/{symbol}/batch?types=quote,news,chart&range=1m&last=1
func (s *Stock) Batch(symbol string, params interface{}) (batch *Batch, err error) {
	endpoint := fmt.Sprintf("%s/batch", symbol)
	err = get(s, &batch, endpoint, params)
	return
}

// Book GET /stock/{symbol}/book
func (s *Stock) Book(symbol string) (book *Book, err error) {
	endpoint := fmt.Sprintf("%s/book", symbol)
	err = get(s, &book, endpoint, nil)
	return
}

// CashFlow GET /stock/{symbol}/cash-flow
func (s *Stock) CashFlow(symbol string, params interface{}) (cashflow *CashFlow, err error) {
	endpoint := fmt.Sprintf("%s/cash-flow", symbol)
	err = get(s, &cashflow, endpoint, params)
	return
}

// Chart GET /stock/{symbol}/chart/{range}
func (s *Stock) Chart(symbol string, chartRange ChartRange, params *ChartQueryParams) (chart []Chart, err error) {
	endpoint := fmt.Sprintf("%s/chart/%s", symbol, chartRange)
	err = get(s, &chart, endpoint, params)
	return
}

// Collection GET /stock/market/collection/{collectionType}?collectionName=
func (s *Stock) Collection(collectionType string, params interface{}) (col Collection, err error) {
	endpoint := fmt.Sprintf("market/collection/%s", collectionType)
	err = get(s, &col, endpoint, params)
	return
}

// Company GET /stock/{symbol}/company
func (s *Stock) Company(symbol string) (com Company, err error) {
	endpoint := fmt.Sprintf("%s/company", symbol)
	err = get(s, &com, endpoint, nil)
	return
}

// DelayedQuote GET /stock/{symbol}/delayed-quote
func (s *Stock) DelayedQuote(symbol string) (dq *DelayedQuote, err error) {
	endpoint := fmt.Sprintf("%s/delayed-quote", symbol)
	err = get(s, &dq, endpoint, nil)
	return
}

// Dividends GET /stock/{symbol}/dividends/{range}
func (s *Stock) Dividends(symbol string, dividendRange DividendRange) (div Dividends, err error) {
	endpoint := fmt.Sprintf("%s/dividends/%s", symbol, dividendRange)
	err = get(s, &div, endpoint, nil)
	return
}

// Earnings GET /stock/{symbol}/earnings/{last}/{field}
func (s *Stock) Earnings(symbol string, params interface{}) (er *Earnings, err error) {
	endpoint := fmt.Sprintf("%s/earnings", symbol)
	err = get(s, &er, endpoint, params)
	return
}

// EarningsToday GET /stock/market/today-earnings
func (s *Stock) EarningsToday() (et *EarningsToday, err error) {
	err = get(s, &et, "market/today-earnings", nil)
	return
}

// Estimates GET /stock/{symbol}/estimates/{last}/{field}
func (s *Stock) Estimates(symbol string, opt ...interface{}) (est *Estimates, err error) {
	endpoint := fmt.Sprintf("%s/estimates", symbol)

	if len(opt) > 0 {
		last := opt[0].(int)
		endpoint = fmt.Sprintf("%s/%s", endpoint, strconv.Itoa(last))
	}
	if len(opt) > 1 {
		field := opt[1].(string)
		endpoint = fmt.Sprintf("%s/%s", endpoint, field)
	}
	err = get(s, &est, endpoint, nil)
	return
}

// Financials GET /stock/{symbol}/financials/{last}/{field}
func (s *Stock) Financials(symbol string, params interface{}, opt ...interface{}) (fin *Financials, err error) {
	endpoint := fmt.Sprintf("%s/financials", symbol)

	if len(opt) > 0 {
		last := opt[0].(int)
		endpoint = fmt.Sprintf("%s/%s", endpoint, strconv.Itoa(last))
	}
	if len(opt) > 1 {
		field := opt[1].(string)
		endpoint = fmt.Sprintf("%s/%s", endpoint, field)
	}
	err = get(s, &fin, endpoint, params)
	return
}

// FundOwnership GET /stock/{symbol}/fund-ownership
func (s *Stock) FundOwnership(symbol string) (fo FundOwnership, err error) {
	endpoint := fmt.Sprintf("%s/fund-ownership", symbol)
	err = get(s, &fo, endpoint, nil)
	return
}

// HistoricalPrices GET /stock/{symbol}/chart/{range}/{date}
func (s *Stock) HistoricalPrices(symbol string, chartRange ChartRange, params *ChartQueryParams) ([]Chart, error) {
	return s.Chart(symbol, chartRange, params)
}

// IncomeStatement GET /stock/{symbol}/income?{params}
func (s *Stock) IncomeStatement(symbol string, params interface{}) (incstmt *IncomeStatement, err error) {
	endpoint := fmt.Sprintf("%s/income", symbol)
	err = get(s, &incstmt, endpoint, params)
	return
}

// InsiderRoster GET /stock/{symbol}/insider-roster
func (s *Stock) InsiderRoster(symbol string) (ir InsiderRoster, err error) {
	endpoint := fmt.Sprintf("%s/insider-roster", symbol)
	err = get(s, &ir, endpoint, nil)
	return
}

// InsiderSummary GET /stock/{symbol}/insider-summary
func (s *Stock) InsiderSummary(symbol string) (is InsiderSummary, err error) {
	endpoint := fmt.Sprintf("%s/insider-summary", symbol)
	err = get(s, &is, endpoint, nil)
	return
}

// InsiderTransactions GET /stock/{symbol}/insider-transactions
func (s *Stock) InsiderTransactions(symbol string) (it InsiderTransactions, err error) {
	endpoint := fmt.Sprintf("%s/insider-transactions", symbol)
	err = get(s, &it, endpoint, nil)
	return
}

// InstitutionalOwnership GET /stock/{symbol}/institutional-ownership
func (s *Stock) InstitutionalOwnership(symbol string) (iop InstitutionalOwnership, err error) {
	endpoint := fmt.Sprintf("%s/institutional-ownership", symbol)
	err = get(s, &iop, endpoint, nil)
	return
}

// IntradayPrices GET /stock/{symbol}/intraday-prices
func (s *Stock) IntradayPrices(symbol string, params interface{}) (ip IntradayPrices, err error) {
	endpoint := fmt.Sprintf("%s/intraday-prices", symbol)
	err = get(s, &ip, endpoint, params)
	return
}

// TodayIPOS GET /stock/market/today-ipos
func (s *Stock) TodayIPOS() (ipo *IPOCalendar, err error) {
	err = get(s, &ipo, "market/today-ipos", nil)
	return
}

// KeyStats GET /stock/{symbol}/stats
func (s *Stock) KeyStats(symbol string) (ks *KeyStat, err error) {
	endpoint := fmt.Sprintf("%s/stats", symbol)
	err = get(s, &ks, endpoint, nil)
	return
}

// LargestTrades GET /stock/{symbol}/largest-trades
func (s *Stock) LargestTrades(symbol string) (lt LargestTrades, err error) {
	endpoint := fmt.Sprintf("%s/largest-trades", symbol)
	err = get(s, &lt, endpoint, nil)
	return
}

// List GET /stock/market/list/{list-type}
func (s *Stock) List(listType string, params interface{}) (list []*Quote, err error) {
	endpoint := fmt.Sprintf("market/list/%s", listType)
	err = get(s, &list, endpoint, params)
	return
}

// Logo GET /stock/{symbol}/logo
func (s *Stock) Logo(symbol string) (logo *Logo, err error) {
	endpoint := fmt.Sprintf("%s/logo", symbol)
	err = get(s, &logo, endpoint, nil)
	return
}

// MarketVolume GET /market
// ??? why did IEX make this endpoint here? doesn't have relative endpoint `stock`...
func (s *Stock) MarketVolume() (mkt MarketVolume, err error) {
	// NOTE: had to force absolute URL since for whatever reason IEX API didn't prepend
	// this endpoint with `stock
	endpoint := fmt.Sprintf("%s%s/market", s.url.String(), s.version)
	err = get(s, &mkt, endpoint, nil)
	return
}

// News GET /stock/{symbol}/news/last/{last}
func (s *Stock) News(symbol string, opt ...interface{}) (news News, err error) {
	endpoint := fmt.Sprintf("%s/news", symbol)
	if len(opt) > 0 {
		last := opt[0].(int)
		endpoint = fmt.Sprintf("%s/%s", endpoint, strconv.Itoa(last))
	}

	err = get(s, &news, endpoint, nil)
	return
}

// OHLC GET /stock/{symbol}/ohlc
func (s *Stock) OHLC(symbol string) (ohlc *OHLC, err error) {
	endpoint := fmt.Sprintf("%s/ohlc", symbol)
	err = get(s, &ohlc, endpoint, nil)
	return
}

// OpenClosePrice Refer to ohlc
func (s *Stock) OpenClosePrice(symbol string) (ohlc *OHLC, err error) {
	return s.OHLC(symbol)
}

// OptionDates GET /stock/{symbol}/options
// return available dates as string slice
func (s *Stock) OptionDates(symbol string) (dates []string, err error) {
	endpoint := fmt.Sprintf("%s/options", symbol)
	err = get(s, &dates, endpoint, nil)
	return
}

// Options GET /stock/{symbol}/options/{expiration}/{optionSide?}
func (s *Stock) Options(symbol, expiration string, opt ...interface{}) (options []*Option, err error) {
	endpoint := fmt.Sprintf("%s/options/%s", symbol, expiration)
	if len(opt) > 0 {
		optionSide := opt[0].(string)
		endpoint = fmt.Sprintf("%s/%s", endpoint, optionSide)
	}
	err = get(s, &options, endpoint, nil)
	return
}

// Peers GET /stock/{symbol}/peers
func (s *Stock) Peers(symbol string) (peers []string, err error) {
	endpoint := fmt.Sprintf("%s/peers", symbol)
	err = get(s, &peers, endpoint, nil)
	return
}

// PreviousDayPrice GET /stock/{symbol}/previous
func (s *Stock) PreviousDayPrice(symbol string) (prev *PreviousDayPrice, err error) {
	endpoint := fmt.Sprintf("%s/previous", symbol)
	err = get(s, &prev, endpoint, nil)
	return
}

// Price GET /stock/{symbol}/price
func (s *Stock) Price(symbol string) (price float64, err error) {
	endpoint := fmt.Sprintf("%s/price", symbol)
	err = get(s, &price, endpoint, nil)
	return
}

// PriceTarget GET /stock/{symbol}/price-target
func (s *Stock) PriceTarget(symbol string) (tgt *PriceTarget, err error) {
	endpoint := fmt.Sprintf("%s/price-target", symbol)
	err = get(s, &tgt, endpoint, nil)
	return
}

// Quote GET /stock/{symbol}/quote
func (s *Stock) Quote(symbol string, params interface{}) (quote *Quote, err error) {
	endpoint := fmt.Sprintf("%s/quote", symbol)
	err = get(s, &quote, endpoint, params)
	return
}

// RecommendationTrends GET /stock/{symbol}/recommendation-trends
func (s *Stock) RecommendationTrends(symbol string) (rt RecommendationTrends, err error) {
	endpoint := fmt.Sprintf("%s/recommendation-trends", symbol)
	err = get(s, &rt, endpoint, nil)
	return
}

// SectorPerformance GET /stock/market/sector-performance
func (s *Stock) SectorPerformance() (sp SectorPerformance, err error) {
	err = get(s, &sp, "market/sector-performance", nil)
	return
}

// Splits GET /stock/{symbol}/splits/{range}
func (s *Stock) Splits(symbol string, splitRange SplitRange) (sp Splits, err error) {
	endpoint := fmt.Sprintf("%s/splits/%s", symbol, splitRange)
	err = get(s, &sp, endpoint, nil)
	return
}

// TechnicalIndicator GET /stock/{symbol}/indicator/{indicatorName}?{range}
func (s *Stock) TechnicalIndicator(symbol string, indicatorName IndicatorName, params *TechnicalIndicatorParams) (ti *TechnicalIndicator, err error) {
	if params == nil {
		params = &TechnicalIndicatorParams{Range: "1d"}
	}
	endpoint := fmt.Sprintf("%s/indicator/%s", symbol, indicatorName)
	err = get(s, &ti, endpoint, params)
	return
}

// UpcomingDividends GET /stock/{symbol}/upcoming-dividends
func (s *Stock) UpcomingDividends(symbol string, params interface{}) (d Dividends, err error) {
	endpoint := fmt.Sprintf("%s/upcoming-dividends", symbol)
	err = get(s, &d, endpoint, params)
	return
}

// UpcomingEarnings GET /stock/{symbol}/upcoming-earnings
func (s *Stock) UpcomingEarnings(symbol string, params interface{}) (ue UpcomingEarnings, err error) {
	endpoint := fmt.Sprintf("%s/upcoming-earnings", symbol)
	err = get(s, &ue, endpoint, params)
	return
}

// UpcomingEvents GET /stock/{symbol}/upcoming-events
func (s *Stock) UpcomingEvents(symbol string, params interface{}) (ue *UpcomingEvents, err error) {
	endpoint := fmt.Sprintf("%s/upcoming-events", symbol)
	err = get(s, &ue, endpoint, params)
	return
}

// UpcomingIPOS GET /stock/market/upcoming-ipos
func (s *Stock) UpcomingIPOS(symbol string, params interface{}) (ipo *IPOCalendar, err error) {
	endpoint := fmt.Sprintf("%s/upcoming-ipos", symbol)
	err = get(s, &ipo, endpoint, params)
	return
}

// UpcomingSplits GET /stock/{symbol}/upcoming-splits
func (s *Stock) UpcomingSplits(symbol string, params interface{}) (spl Splits, err error) {
	endpoint := fmt.Sprintf("%s/upcoming-splits", symbol)
	err = get(s, &spl, endpoint, params)
	return
}

// VolumeByVenue GET /stock/{symbol}/volume-by-venue
func (s *Stock) VolumeByVenue(symbol string) (vbv VolumeByVenue, err error) {
	endpoint := fmt.Sprintf("%s/volume-by-venue", symbol)
	err = get(s, &vbv, endpoint, nil)
	return
}

// UnmarshalJSON helper
func (etd *EarningsTodayDTO) UnmarshalJSON(b []byte) error {
	var err error
	type alias EarningsTodayDTO
	aux := &struct {
		ConsensusEPS interface{} `json:"consensusEPS"`
		*alias
	}{
		alias: (*alias)(etd),
	}

	if err = json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if consensusEPS, ok := aux.ConsensusEPS.(string); ok {
		etd.ConsensusEPS, err = strconv.ParseFloat(consensusEPS, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (in IndicatorName) String() string {
	return [...]string{
		"abs",
		"acos",
		"ad",
		"add",
		"adosc",
		"adx",
		"adxr",
		"ao",
		"apo",
		"aroon",
		"aroonosc",
		"asin",
		"atan",
		"atr",
		"avgprice",
		"bbands",
		"bop",
		"cci",
		"ceil",
		"cmo",
		"cos",
		"cosh",
		"crossany",
		"crossover",
		"cvi",
		"decay",
		"dema",
		"di",
		"div",
		"dm",
		"dpo",
		"dx",
		"edecay",
		"ema",
		"emv",
		"exp",
		"fisher",
		"floor",
		"fosc",
		"hma",
		"kama",
		"kvo",
		"lag",
		"linreg",
		"linregintercept",
		"linregslope",
		"ln",
		"log10",
		"macd",
		"marketfi",
		"mass",
		"max",
		"md",
		"medprice",
		"mfi",
		"min",
		"mom",
		"msw",
		"mul",
		"natr",
		"nvi",
		"obv",
		"ppo",
		"psar",
		"pvi",
		"qstick",
		"roc",
		"rocr",
		"round",
		"rsi",
		"sin",
		"sinh",
		"sma",
		"sqrt",
		"stddev",
		"stoch",
		"stochrsi",
		"sub",
		"sum",
		"tan",
		"tanh",
		"tema",
		"todeg",
		"torad",
		"tr",
		"trima",
		"trix",
		"trunc",
		"tsf",
		"typprice",
		"ultosc",
		"var",
		"vhf",
		"vidya",
		"volatility",
		"vosc",
		"vwma",
		"wad",
		"wcprice",
		"wilders",
		"willr",
		"wma",
		"zlema",
	}[in]
}

func (cr ChartRange) String() string {
	return [...]string{
		"max",
		"5y",
		"2y",
		"1y",
		"ytd",
		"6m",
		"3m",
		"1m",
		"1d",
	}[cr]
}

func (dr DividendRange) String() string {
	return [...]string{
		"5y",
		"2y",
		"1y",
		"ytd",
		"6m",
		"3m",
		"1m",
		"next",
	}[dr]
}

func (sr SplitRange) String() string {
	return [...]string{
		"5y",
		"2y",
		"1y",
		"ytd",
		"6m",
		"3m",
		"1m",
		"next",
	}[sr]
}

func (pqp PeriodQueryParameter) String() string {
	return [...]string{
		"annual",
		"quarter",
	}[pqp]
}
