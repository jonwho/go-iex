package goiex

import (
	"net/http"
	"net/url"
)

type iexapi interface {
	APIURL() *url.URL
	Client() *http.Client
	Token() string
	URL() *url.URL
	Version() string
	Retry() *Retry
}
