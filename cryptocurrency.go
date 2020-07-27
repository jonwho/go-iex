package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Cryptocurrency struct to interface with / endpoints
type Cryptocurrency struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
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

func (c *Cryptocurrency) Do(req *Request) (*http.Response, error) {
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
