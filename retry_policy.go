// See https://medium.com/@nitishkr88/http-retries-in-go-e622e51d249f for implementation guide.

package goiex

import (
	"net/http"
)

// RetryPolicy is a func type that defines when to retry a http request
type RetryPolicy func(resp *http.Response, err error) (bool, error)

// DefaultRetryPolicy defines a retry on status codes 0, 429, and >= 500
func DefaultRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		return true, err
	}

	if resp.StatusCode == 0 || resp.StatusCode == 429 || resp.StatusCode >= 500 {
		return true, nil
	}

	return false, nil
}
