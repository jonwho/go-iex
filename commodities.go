package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Commodities struct to interface with /data-points endpoints
type Commodities struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// NewCommodities return new Commodities
func NewCommodities(token, version string, base *url.URL, httpClient *http.Client) *Commodities {
	apiurl, err := url.Parse("data-points/")
	if err != nil {
		panic(err)
	}

	return &Commodities{
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
func (c *Commodities) APIURL() *url.URL {
	return c.apiurl
}

// Client return HTTP client
func (c *Commodities) Client() *http.Client {
	return c.client
}

// Token return token string
func (c *Commodities) Token() string {
	return c.token
}

// URL return URL base
func (c *Commodities) URL() *url.URL {
	return c.url
}

// Version return version string
func (c *Commodities) Version() string {
	return c.version
}

func (c *Commodities) Do(req *Request) (*http.Response, error) {
	for i := 0; i < c.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := c.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if c.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := c.RetryPolicy(resp, err)
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

		remain := c.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := c.Backoff(c.RetryWaitMin, c.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, c.RetryAttempts+1)
}

// CommoditiesPrices GET /data-points/market/{symbol}
func (c *Commodities) CommoditiesPrices(symbol string) (value interface{}, err error) {
	err = getRaw(c, &value, "market/"+symbol, nil)
	return
}
