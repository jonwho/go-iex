package goiex

import (
	"net/http"
	"net/url"
)

// Forex struct to interface with Forex / Currencies endpoints
type Forex struct {
	iex
}

// LatestCurrencyRatesParams required/optional query parameters
type LatestCurrencyRatesParams struct {
	Symbols string `url:"symbols"`
}

// CurrencyConversionParams required/optional query parameters
type CurrencyConversionParams struct {
	Symbols string `url:"symbols"`
	Amount  int    `url:"amount"`
}

// HistoricalDailyParams required/optional query parameters
type HistoricalDailyParams struct {
	Symbols string `url:"symbols"`
	From    string `url:"from"`
	To      string `url:"to"`
	Last    int    `url:"last"`
}

// LatestCurrencyRates struct
type LatestCurrencyRates []struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
}

// CurrencyConversion struct
type CurrencyConversion []struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp int64   `json:"timestamp"`
	Amount    float64 `json:"amount"`
}

// HistoricalDaily struct
type HistoricalDaily [][]struct {
	Date      string  `json:"date"`
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
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

// LatestCurrencyRates GET /fx/latest?{params}
func (f *Forex) LatestCurrencyRates(params *LatestCurrencyRatesParams) (lcr LatestCurrencyRates, err error) {
	err = get(f, &lcr, "latest", params)
	return
}

// CurrencyConversion GET /fx/convert?{params}
func (f *Forex) CurrencyConversion(params *CurrencyConversionParams) (cc CurrencyConversion, err error) {
	err = get(f, &cc, "convert", params)
	return
}

// HistoricalDaily GET /fx/historical?{params}
func (f *Forex) HistoricalDaily(params *HistoricalDailyParams) (hd HistoricalDaily, err error) {
	err = get(f, &hd, "historical", params)
	return
}

// ExchangeRates GET /fx/rate/{from}/{to}
func (f *Forex) ExchangeRates(from, to string) (er *ExchangeRates, err error) {
	err = get(f, &er, "rate/"+from+"/"+to, nil)
	return
}
