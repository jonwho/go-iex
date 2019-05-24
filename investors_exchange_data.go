package goiex

import (
	"net/http"
	"net/url"
)

// InvestorsExchangeData struct to interface with InvestorsExchangeData endpoints
type InvestorsExchangeData struct {
	iex
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

// NewInvestorsExchangeData return new InvestorsExchangeData
func NewInvestorsExchangeData(token, version string, base *url.URL, httpClient *http.Client) *InvestorsExchangeData {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}
	return &InvestorsExchangeData{
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

// TOPS GET /tops?symbols=snap
func (ied *InvestorsExchangeData) TOPS(params interface{}) (tops TOPS, err error) {
	get(ied, &tops, "tops", params)
	return
}
