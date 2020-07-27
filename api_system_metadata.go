package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// APISystemMetadata struct to interface with / endpoints
type APISystemMetadata struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// Status struct
type Status struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Time    int64  `json:"time"`
}

// NewAPISystemMetadata return new APISystemMetadata
func NewAPISystemMetadata(token, version string, base *url.URL, httpClient *http.Client) *APISystemMetadata {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}
	return &APISystemMetadata{
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
func (a *APISystemMetadata) Token() string {
	return a.token
}

// Version return version string
func (a *APISystemMetadata) Version() string {
	return a.version
}

// URL return URL base
func (a *APISystemMetadata) URL() *url.URL {
	return a.url
}

// APIURL return APIURL
func (a *APISystemMetadata) APIURL() *url.URL {
	return a.apiurl
}

// Client return HTTP client
func (a *APISystemMetadata) Client() *http.Client {
	return a.client
}

func (a *APISystemMetadata) Do(req *Request) (*http.Response, error) {
	for i := 0; i < a.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := a.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if a.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := a.RetryPolicy(resp, err)
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

		remain := a.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := a.Backoff(a.RetryWaitMin, a.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, a.RetryAttempts+1)
}

// Status GET /status
func (a *APISystemMetadata) Status() (status *Status, err error) {
	err = get(a, &status, "status", nil)
	return
}
