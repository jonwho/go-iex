package goiex

import (
	"fmt"
	"net/http"
	"time"
)

var (
	// Default retry configuration
	defaultRetryWaitMin  = 1 * time.Second
	defaultRetryWaitMax  = 30 * time.Second
	defaultRetryAttempts = 4
)

// Retry implements Retryer
type Retry struct {
	*http.Client
	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// RetryOption is a func that operates on *Retry.
type RetryOption func(*Retry) error

// SetRetryWaitMinMax sets the min and max time.Duration that a retry request is bounded by.
func SetRetryWaitMinMax(min, max time.Duration) RetryOption {
	return func(r *Retry) error {
		if min <= 0 || max <= 0 {
			return fmt.Errorf("The value of min: %d or max: %d cannot be less than or equal to 0", min, max)
		}
		if min > max {
			return fmt.Errorf("The value of min: %d cannot be greater than max: %d", min, max)
		}
		r.RetryWaitMin = min
		r.RetryWaitMax = max
		return nil
	}
}

// SetRetryAttempts sets the max number of retry attempts.
func SetRetryAttempts(attempts int) RetryOption {
	return func(r *Retry) error {
		if attempts <= 0 {
			return fmt.Errorf("The value of attempts: %d cannot be less than or equal to 0", attempts)
		}
		r.RetryAttempts = attempts
		return nil
	}
}

// SetRetryPolicy sets the RetryPolicy that defines when to retry a HTTP request.
func SetRetryPolicy(rp RetryPolicy) RetryOption {
	return func(r *Retry) error {
		r.RetryPolicy = rp
		return nil
	}
}

// SetBackoff sets the Backoff that defines how long to wait between HTTP request retries.
func SetBackoff(b Backoff) RetryOption {
	return func(r *Retry) error {
		r.Backoff = b
		return nil
	}
}

// NewRetry returns a struct that can retry a HTTP request defined by the RetryPolicy and Backoff.
func NewRetry(httpClient *http.Client, options ...RetryOption) (*Retry, error) {
	retry := &Retry{
		Client:        httpClient,
		RetryWaitMin:  defaultRetryWaitMin,
		RetryWaitMax:  defaultRetryWaitMax,
		RetryAttempts: defaultRetryAttempts,
		RetryPolicy:   DefaultRetryPolicy,
		Backoff:       DefaultBackoff,
	}

	for _, option := range options {
		err := option(retry)
		if err != nil {
			return nil, err
		}
	}

	return retry, nil
}

// Do will perform the Request and attempt it again if the RetryPolicy passes
func (r *Retry) Do(req *Request) (*http.Response, error) {
	for i := 0; i < r.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := r.Client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if r.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := r.RetryPolicy(resp, err)
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

		remain := r.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := r.Backoff(r.RetryWaitMin, r.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, r.RetryAttempts+1)
}
