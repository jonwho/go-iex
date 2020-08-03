package goiex

import (
	"net/http"
	"time"
)

// Backoffer is the interface to define a wait time in Backoff method.
type Backoffer interface {
	Backoff(min, max time.Duration, attempts int, retry *http.Response) time.Duration
}
