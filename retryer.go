package goiex

import (
	"net/http"
)

// Retryer is the interface that defines methods to retry a HTTP request.
type Retryer interface {
	Backoffer
	RetryPolicyer
	Do(*Request) (*http.Response, error)
}
