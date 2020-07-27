package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Forex struct to interface with Forex / Currencies endpoints
type Forex struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
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
		RetryWaitMin:  defaultRetryWaitMin,
		RetryWaitMax:  defaultRetryWaitMax,
		RetryAttempts: defaultRetryAttempts,
		RetryPolicy:   DefaultRetryPolicy,
		Backoff:       DefaultBackoff,

		iex: iex{
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

func (f *Forex) Do(req *Request) (*http.Response, error) {
	for i := 0; i < f.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := f.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if f.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := f.RetryPolicy(resp, err)
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}
			return resp, err
		}

		// Perform retry
		if err == nil {
			drainBody(resp.Body)
		}

		remain := f.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := f.Backoff(f.RetryWaitMin, f.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, f.RetryAttempts+1)
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
