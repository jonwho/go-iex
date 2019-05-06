package goiex

import (
	"net/http"
	"net/url"
)

// APISystemMetadata struct to interface with / endpoints
type APISystemMetadata struct {
	iex
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
		iex{
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

// Status GET /status
func (a *APISystemMetadata) Status() (status *Status, err error) {
	err = get(a, &status, "status", nil)
	return
}
