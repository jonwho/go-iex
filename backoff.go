package goiex

import (
	"math"
	"net/http"
	"time"
)

// Backoff defines the wait time between retry requests
type Backoff func(min, max time.Duration, attempts int, resp *http.Response) time.Duration

// DefaultBackoff defines the wait time as a exponential backoff
func DefaultBackoff(min, max time.Duration, attempts int, resp *http.Response) time.Duration {
	return exponentialBackoff(min, max, attempts, resp)
}

func exponentialBackoff(min, max time.Duration, attempts int, resp *http.Response) time.Duration {
	expo := math.Pow(2, float64(attempts)) * float64(min)
	sleep := time.Duration(expo)
	if float64(sleep) != expo || sleep > max {
		sleep = max
	}
	return sleep
}
