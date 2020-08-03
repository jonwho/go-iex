package goiex

import (
	"net/http"
	"net/url"
)

// InvestorsExchangeData struct to interface with InvestorsExchangeData endpoints
type InvestorsExchangeData struct {
	iex
}

// TOPSParams required/optional query parameters
type TOPSParams struct {
	Symbols string `url:"symbols"`
}

// LastParams required/optional query parameters
type LastParams struct {
	Symbols string `url:"symbols"`
}

// DEEPParams required/optional query parameters
type DEEPParams struct {
	Symbols string `url:"symbols"`
}

// TOPS struct
type TOPS []struct {
	Symbol        string  `json:"symbol"`
	BidSize       int     `json:"bidSize"`
	BidPrice      float64 `json:"bidPrice"`
	AskSize       int     `json:"askSize"`
	AskPrice      float64 `json:"askPrice"`
	Volume        int     `json:"volume"`
	LastSalePrice float64 `json:"lastSalePrice"`
	LastSaleSize  int     `json:"lastSaleSize"`
	LastSaleTime  int64   `json:"lastSaleTime"`
	LastUpdated   int64   `json:"lastUpdated"`
	Sector        string  `json:"sector"`
	SecurityType  string  `json:"securityType"`
}

// Last struct
type Last []struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Size   int     `json:"size"`
	Time   int64   `json:"time"`
}

// DEEP struct
type DEEP struct {
	Symbol        string  `json:"symbol"`
	MarketPercent float64 `json:"marketPercent"`
	Volume        int     `json:"volume"`
	LastSalePrice float64 `json:"lastSalePrice"`
	LastSaleSize  int     `json:"lastSaleSize"`
	LastSaleTime  int64   `json:"lastSaleTime"`
	LastUpdated   int64   `json:"lastUpdated"`
	Bids          []struct {
		Price     float64 `json:"price"`
		Size      int     `json:"size"`
		Timestamp int64   `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     float64 `json:"price"`
		Size      int     `json:"size"`
		Timestamp int64   `json:"timestamp"`
	} `json:"asks"`
	SystemEvent struct {
		SystemEvent string `json:"systemEvent"`
		Timestamp   int64  `json:"timestamp"`
	} `json:"systemEvent"`
	TradingStatus struct {
		Status    string `json:"status"`
		Reason    string `json:"reason"`
		Timestamp int64  `json:"timestamp"`
	} `json:"tradingStatus"`
	OpHaltStatus struct {
		IsHalted  bool  `json:"isHalted"`
		Timestamp int64 `json:"timestamp"`
	} `json:"opHaltStatus"`
	SsrStatus struct {
		IsSSR     bool   `json:"isSSR"`
		Detail    string `json:"detail"`
		Timestamp int64  `json:"timestamp"`
	} `json:"ssrStatus"`
	SecurityEvent struct {
		SecurityEvent string `json:"securityEvent"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"securityEvent"`
	Trades []struct {
		Price                 float64 `json:"price"`
		Size                  int     `json:"size"`
		TradeID               int     `json:"tradeId"`
		IsISO                 bool    `json:"isISO"`
		IsOddLot              bool    `json:"isOddLot"`
		IsOutsideRegularHours bool    `json:"isOutsideRegularHours"`
		IsSinglePriceCross    bool    `json:"isSinglePriceCross"`
		IsTradeThroughExempt  bool    `json:"isTradeThroughExempt"`
		Timestamp             int64   `json:"timestamp"`
	} `json:"trades"`
	TradeBreaks []struct {
		Price                 float64 `json:"price"`
		Size                  int     `json:"size"`
		TradeID               int     `json:"tradeId"`
		IsISO                 bool    `json:"isISO"`
		IsOddLot              bool    `json:"isOddLot"`
		IsOutsideRegularHours bool    `json:"isOutsideRegularHours"`
		IsSinglePriceCross    bool    `json:"isSinglePriceCross"`
		IsTradeThroughExempt  bool    `json:"isTradeThroughExempt"`
		Timestamp             int64   `json:"timestamp"`
	} `json:"tradeBreaks"`
	Auction struct {
		AuctionType          string  `json:"auctionType"`
		PairedShares         int     `json:"pairedShares"`
		ImbalanceShares      int     `json:"imbalanceShares"`
		ReferencePrice       float64 `json:"referencePrice"`
		IndicativePrice      float64 `json:"indicativePrice"`
		AuctionBookPrice     float64 `json:"auctionBookPrice"`
		CollarReferencePrice float64 `json:"collarReferencePrice"`
		LowerCollarPrice     float64 `json:"lowerCollarPrice"`
		UpperCollarPrice     float64 `json:"upperCollarPrice"`
		ExtensionNumber      int     `json:"extensionNumber"`
		StartTime            string  `json:"startTime"`
		LastUpdate           int64   `json:"lastUpdate"`
	} `json:"auction"`
}

// NewInvestorsExchangeData return new InvestorsExchangeData
func NewInvestorsExchangeData(
	token, version string,
	base *url.URL,
	httpClient *http.Client,
	options ...IEXOption,
) *InvestorsExchangeData {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}

	ied := &InvestorsExchangeData{
		iex: iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}

	for _, option := range options {
		err := option(&ied.iex)
		if err != nil {
			return nil
		}
	}

	return ied
}

// Token return token string
func (ied *InvestorsExchangeData) Token() string {
	return ied.token
}

// Version return version string
func (ied *InvestorsExchangeData) Version() string {
	return ied.version
}

// URL return URL base
func (ied *InvestorsExchangeData) URL() *url.URL {
	return ied.url
}

// APIURL return APIURL
func (ied *InvestorsExchangeData) APIURL() *url.URL {
	return ied.apiurl
}

// Client return HTTP client
func (ied *InvestorsExchangeData) Client() *http.Client {
	return ied.client
}

// Retry return Retry struct that implements Retryer
func (ied *InvestorsExchangeData) Retry() *Retry {
	return ied.iex.Retry
}

// TOPS GET /tops?{params}
func (ied *InvestorsExchangeData) TOPS(params *TOPSParams) (tops TOPS, err error) {
	get(ied, &tops, "tops", params)
	return
}

// Last GET /tops/last?{params}
func (ied *InvestorsExchangeData) Last(params *LastParams) (l Last, err error) {
	get(ied, &l, "tops/last", params)
	return
}

// DEEP GET /deep?symbols={params}
func (ied *InvestorsExchangeData) DEEP(params *DEEPParams) (d *DEEP, err error) {
	get(ied, &d, "deep", params)
	return
}
