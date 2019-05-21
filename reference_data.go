package goiex

import (
	"fmt"
	"net/http"
	"net/url"
)

// ReferenceData struct to interface with /ref-data endpoints
type ReferenceData struct {
	iex
}

// Symbols struct
type Symbols []struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IsEnabled bool   `json:"isEnabled"`
}

// IEXSymbols struct
type IEXSymbols []struct {
	Symbol    string `json:"symbol"`
	Date      string `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
}

// InternationalSybmols struct
type InternationalSybmols []struct {
	Symbol    string `json:"symbol"`
	Exchange  string `json:"exchange"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IsEnabled bool   `json:"isEnabled"`
}

// NewReferenceData return new ReferenceData
func NewReferenceData(token, version string, base *url.URL, httpClient *http.Client) *ReferenceData {
	apiurl, err := url.Parse("ref-data/")
	if err != nil {
		panic(err)
	}
	return &ReferenceData{
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
func (rd *ReferenceData) Token() string {
	return rd.token
}

// Version return version string
func (rd *ReferenceData) Version() string {
	return rd.version
}

// URL return URL base
func (rd *ReferenceData) URL() *url.URL {
	return rd.url
}

// APIURL return APIURL
func (rd *ReferenceData) APIURL() *url.URL {
	return rd.apiurl
}

// Client return HTTP client
func (rd *ReferenceData) Client() *http.Client {
	return rd.client
}

// Symbols GET /ref-data/symbols
func (rd *ReferenceData) Symbols() (s Symbols, err error) {
	err = get(rd, &s, "symbols", nil)
	return
}

// IEXSymbols GET /ref-data/iex/symbols
func (rd *ReferenceData) IEXSymbols() (is IEXSymbols, err error) {
	err = get(rd, &is, "iex/symbols", nil)
	return
}

// RegionSymbols GET /ref-data/region/{region}/symbols
func (rd *ReferenceData) RegionSymbols(region string) (is InternationalSybmols, err error) {
	endpoint := fmt.Sprintf("region/%s/symbols", region)
	err = get(rd, &is, endpoint, nil)
	return
}

// ExchangeSymbols GET /ref-data/exchange/{exchange}/symbols
func (rd *ReferenceData) ExchangeSymbols(exchange string) (is InternationalSybmols, err error) {
	endpoint := fmt.Sprintf("exchange/%s/symbols", exchange)
	err = get(rd, &is, endpoint, nil)
	return
}
