package goiex

import (
	"net/http"
	"net/url"
)

// Cryptocurrency struct to interface with / endpoints
type Cryptocurrency struct {
	iex
}

// CryptoBook struct
type CryptoBook struct {
	Bids []struct {
		Price     string `json:"price"`
		Size      string `json:"size"`
		Timestamp int64  `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     string `json:"price"`
		Size      string `json:"size"`
		Timestamp int64  `json:"timestamp"`
	} `json:"asks"`
}

// CryptoPrice struct
type CryptoPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

// CryptoQuote struct
type CryptoQuote struct {
	Symbol           string  `json:"symbol"`
	Sector           string  `json:"sector"`
	CalculationPrice string  `json:"calculationPrice"`
	LatestPrice      float64 `json:"latestPrice,string"`
	LatestSource     string  `json:"latestSource"`
	LatestUpdate     int64   `json:"latestUpdate"`
	LatestVolume     float64 `json:"latestVolume,string"`
	BidPrice         float64 `json:"bidPrice,string"`
	BidSize          float64 `json:"bidSize,string"`
	AskPrice         float64 `json:"askPrice,string"`
	AskSize          float64 `json:"askSize,string"`
	High             float64 `json:"high,string"`
	Low              float64 `json:"low,string"`
	PreviousClose    float64 `json:"previousClose,string"`
}

// NewCryptocurrency returns new Cryptocurrency
func NewCryptocurrency(token, version string, base *url.URL, httpClient *http.Client) *Cryptocurrency {
	apiurl, err := url.Parse("crypto/")
	if err != nil {
		panic(err)
	}
	return &Cryptocurrency{
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
func (c *Cryptocurrency) Token() string {
	return c.token
}

// Version return version string
func (c *Cryptocurrency) Version() string {
	return c.version
}

// URL return URL base
func (c *Cryptocurrency) URL() *url.URL {
	return c.url
}

// APIURL return APIURL
func (c *Cryptocurrency) APIURL() *url.URL {
	return c.apiurl
}

// Client return HTTP client
func (c *Cryptocurrency) Client() *http.Client {
	return c.client
}

// CryptoBook GET /crypto/{symbol}/book
func (c *Cryptocurrency) CryptoBook(symbol string) (cb *CryptoBook, err error) {
	err = get(c, &cb, symbol+"/book", nil)
	return
}

// CryptoPrice GET /crypto/{symbol}/price
func (c *Cryptocurrency) CryptoPrice(symbol string) (cp *CryptoPrice, err error) {
	err = get(c, &cp, symbol+"/price", nil)
	return
}

// CryptoQuote GET /crypto/{symbol}/quote
func (c *Cryptocurrency) CryptoQuote(symbol string) (cq *CryptoQuote, err error) {
	err = get(c, &cq, symbol+"/quote", nil)
	return
}
