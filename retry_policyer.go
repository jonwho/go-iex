package goiex

import (
	"net/http"
)

type RetryPolicyer interface {
	RetryPolicy(resp *http.Response, err error) (bool, error)
}
