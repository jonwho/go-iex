package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DataAPI struct to interface with DataAPI endpoints
type DataAPI struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// DataPoint struct
type DataPoint struct {
	Key         string    `json:"key"`
	Weight      int       `json:"weight"`
	Description string    `json:"description"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// NewDataAPI return new DataAPI
func NewDataAPI(token, version string, base *url.URL, httpClient *http.Client) *DataAPI {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}
	return &DataAPI{
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
func (d *DataAPI) Token() string {
	return d.token
}

// Version return version string
func (d *DataAPI) Version() string {
	return d.version
}

// URL return URL base
func (d *DataAPI) URL() *url.URL {
	return d.url
}

// APIURL return APIURL
func (d *DataAPI) APIURL() *url.URL {
	return d.apiurl
}

// Client return HTTP client
func (d *DataAPI) Client() *http.Client {
	return d.client
}

func (d *DataAPI) Do(req *Request) (*http.Response, error) {
	for i := 0; i < d.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := d.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if d.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := d.RetryPolicy(resp, err)
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

		remain := d.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := d.Backoff(d.RetryWaitMin, d.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, d.RetryAttempts+1)
}

// DataPoints GET /data-points/{symbol}
func (d *DataAPI) DataPoints(symbol string) (datapoints []*DataPoint, err error) {
	err = get(d, &datapoints, "data-points/"+symbol, nil)
	return
}

// DataPoint GET /data-points/{symbol}/{datapoint}
func (d *DataAPI) DataPoint(symbol, datapoint string) (ifc interface{}, err error) {
	err = get(d, &ifc, "data-points/"+symbol+"/"+datapoint, nil)
	return
}
