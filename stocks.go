package goiex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var (
	// ChartRanges allowed for Chart API
	ChartRanges = map[string]bool{
		"max": true,
		"5y":  true,
		"2y":  true,
		"1y":  true,
		"ytd": true,
		"6m":  true,
		"3m":  true,
		"1m":  true,
		"1d":  true,
	}
	// DividendRanges allowed for Dividends API
	DividendRanges = map[string]bool{
		"5y":   true,
		"2y":   true,
		"1y":   true,
		"ytd":  true,
		"6m":   true,
		"3m":   true,
		"1m":   true,
		"next": true,
	}
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
	Amount       float64 `json:"amount"`
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

// EffectiveSpread struct
type EffectiveSpread struct {
	Volume           int     `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	EffectiveSpread  float64 `json:"effectiveSpread"`
	EffectiveQuoted  float64 `json:"effectiveQuoted"`
	PriceImprovement float64 `json:"priceImprovement"`
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

// KeyStat struct
type KeyStat struct {
	CompanyName         string  `json:"companyName"`
	Marketcap           int64   `json:"marketcap"`
	Week52High          float64 `json:"week52high"`
	Week52Low           float64 `json:"week52low"`
	Week52Change        float64 `json:"week52change"`
	SharesOutstanding   int64   `json:"sharesOutstanding"`
	Float               int64   `json:"float"`
	Symbol              string  `json:"symbol"`
	Avg10Volume         int     `json:"avg10Volume"`
	Avg30Volume         int     `json:"avg30Volume"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	Employees           int     `json:"employees"`
	TtmEPS              float64 `json:"ttmEPS"`
	TtmDividendRate     float64 `json:"ttmDividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	NextDividendDate    string  `json:"nextDividendDate"`
	ExDividendDate      string  `json:"exDividendDate"`
	NextEarningsDate    string  `json:"nextEarningsDate"`
	PeRatio             int     `json:"peRatio"`
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

// SystemEvent struct
type SystemEvent struct {
	SystemEvent string `json:"systemEvent"`
	Timestamp   int64  `json:"timestamp"`
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
	err = get(s, &advstat, "data-points/"+symbol, nil)
	return
}

// BalanceSheet GET /stock/{symbol}/balance-sheet/{last}/{field}
func (s *Stock) BalanceSheet(symbol string, params interface{}, opt ...interface{}) (balsheet *BalanceSheet, err error) {
	endpoint := fmt.Sprintf("%s/balance-sheet", symbol)
	if len(opt) > 0 {
		last := opt[0].(int)
		endpoint = fmt.Sprintf("%s/%s", endpoint, strconv.Itoa(last))
	}
	if len(opt) > 1 {
		field := opt[1].(string)
		endpoint = fmt.Sprintf("%s/%s", endpoint, field)
	}

	err = get(s, &balsheet, endpoint, params)
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
func (s *Stock) Chart(symbol, chartRange string, params interface{}) (chart []*Chart, err error) {
	if !ChartRanges[chartRange] {
		err = fmt.Errorf("Received invalid date range for chart")
		return
	}

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
func (s *Stock) Dividends(symbol, divRange string) (div Dividends, err error) {
	if !DividendRanges[divRange] {
		err = fmt.Errorf("Received invalid date range for dividend")
		return
	}

	endpoint := fmt.Sprintf("%s/dividends/%s", symbol, divRange)
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

// EffectiveSpread GET /stock/{symbol}/effective-spread
func (s *Stock) EffectiveSpread(symbol string) (spreads []EffectiveSpread, err error) {
	endpoint := fmt.Sprintf("%s/effective-spread", symbol)
	err = get(s, &spreads, endpoint, nil)
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

// Quote GET /stock/{symbol}/quote
func (s *Stock) Quote(symbol string, params interface{}) (quote *Quote, err error) {
	endpoint := fmt.Sprintf("%s/quote", symbol)
	err = get(s, &quote, endpoint, params)
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
	}

	return nil
}
