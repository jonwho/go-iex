package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// EconomicData struct to interface with /data-points endpoints
type EconomicData struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// NewEconomicData return new EconomicData
func NewEconomicData(token, version string, base *url.URL, httpClient *http.Client) *EconomicData {
	apiurl, err := url.Parse("data-points/")
	if err != nil {
		panic(err)
	}

	return &EconomicData{
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

// APIURL return APIURL
func (ed *EconomicData) APIURL() *url.URL {
	return ed.apiurl
}

// Client return HTTP client
func (ed *EconomicData) Client() *http.Client {
	return ed.client
}

// Token return token string
func (ed *EconomicData) Token() string {
	return ed.token
}

// URL return URL base
func (ed *EconomicData) URL() *url.URL {
	return ed.url
}

// Version return version string
func (ed *EconomicData) Version() string {
	return ed.version
}

func (ed *EconomicData) Do(req *Request) (*http.Response, error) {
	for i := 0; i < ed.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := ed.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if ed.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := ed.RetryPolicy(resp, err)
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

		remain := ed.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := ed.Backoff(ed.RetryWaitMin, ed.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, ed.RetryAttempts+1)
}

// EconomicPrices GET /data-points/market/{symbol}
func (ed *EconomicData) EconomicPrices(symbol string) (value interface{}, err error) {
	err = get(ed, &value, "market/"+symbol, nil)
	return
}
