package goiex

import (
	"net/http"
	"net/url"
)

// Forex struct to interface with Forex / Currencies endpoints
type Forex struct {
	iex
}

// ExchangeRates struct
type ExchangeRates struct {
	Date         string  `json:"date"`
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Rate         float64 `json:"rate"`
}

// NewForex return new Forex
func NewForex(token, version string, base *url.URL, httpClient *http.Client) *Forex {
	apiurl, err := url.Parse("fx/")
	if err != nil {
		panic(err)
	}
	return &Forex{
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
func (f *Forex) Token() string {
	return f.token
}

// Version return version string
func (f *Forex) Version() string {
	return f.version
}

// URL return URL base
func (f *Forex) URL() *url.URL {
	return f.url
}

// APIURL return APIURL
func (f *Forex) APIURL() *url.URL {
	return f.apiurl
}

// Client return HTTP client
func (f *Forex) Client() *http.Client {
	return f.client
}

// ExchangeRates GET /fx/rate/{from}/{to}
func (f *Forex) ExchangeRates(from, to string) (er *ExchangeRates, err error) {
	err = get(f, &er, "rate/"+from+"/"+to, nil)
	return
}
