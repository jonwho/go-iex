package goiex

import (
	"net/http"
)

// RetryPolicyer interface defines methods for when to retry
type RetryPolicyer interface {
	RetryPolicy(resp *http.Response, err error) (bool, error)
}
