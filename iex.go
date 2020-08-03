package goiex

import (
	"net/http"
	"net/url"
)

// iex is the struct holds all fields necessary to make HTTP requests against an
// IEX Cloud endpoint(s)
type iex struct {
	token, version string
	url, apiurl    *url.URL
	client         *http.Client

	Retry *Retry
}

// IEXOption is a func that operates on *iex
type IEXOption func(*iex) error

// SetRetry assigns the
func SetRetry(retry *Retry) IEXOption {
	return func(i *iex) error {
		i.Retry = retry
		return nil
	}
}
