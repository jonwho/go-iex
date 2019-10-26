package goiex

import (
	"net/http"
	"net/url"
)

// EconomicData struct to interface with /data-points endpoints
type EconomicData struct {
	iex
}

// NewEconomicData return new EconomicData
func NewEconomicData(token, version string, base *url.URL, httpClient *http.Client) *EconomicData {
	apiurl, err := url.Parse("data-points/")
	if err != nil {
		panic(err)
	}

	return &EconomicData{
		iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// APIURL return APIURL
func (ed *EconomicData) APIURL() *url.URL {
	return ed.apiurl
}

// Client return HTTP client
func (ed *EconomicData) Client() *http.Client {
	return ed.client
}

// Token return token string
func (ed *EconomicData) Token() string {
	return ed.token
}

// URL return URL base
func (ed *EconomicData) URL() *url.URL {
	return ed.url
}

// Version return version string
func (ed *EconomicData) Version() string {
	return ed.version
}

// EconomicPrices GET /data-points/market/{symbol}
func (ed *EconomicData) EconomicPrices(symbol string) (value interface{}, err error) {
	err = get(ed, &value, "market/"+symbol, nil)
	return
}
