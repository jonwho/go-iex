package goiex

import (
	"net/http"
	"net/url"
)

// Commodities struct to interface with /data-points endpoints
type Commodities struct {
	iex
}

// NewCommodities return new Commodities
func NewCommodities(
	token, version string,
	base *url.URL,
	httpClient *http.Client,
	options ...IEXOption,
) *Commodities {
	apiurl, err := url.Parse("data-points/")
	if err != nil {
		panic(err)
	}

	comm := &Commodities{
		iex: iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}

	for _, option := range options {
		err := option(&comm.iex)
		if err != nil {
			return nil
		}
	}

	return comm
}

// APIURL return APIURL
func (c *Commodities) APIURL() *url.URL {
	return c.apiurl
}

// Client return HTTP client
func (c *Commodities) Client() *http.Client {
	return c.client
}

// Token return token string
func (c *Commodities) Token() string {
	return c.token
}

// URL return URL base
func (c *Commodities) URL() *url.URL {
	return c.url
}

// Version return version string
func (c *Commodities) Version() string {
	return c.version
}

// Retry return Retry struct that implements Retryer
func (c *Commodities) Retry() *Retry {
	return c.iex.Retry
}

// CommoditiesPrices GET /data-points/market/{symbol}
func (c *Commodities) CommoditiesPrices(symbol string) (value interface{}, err error) {
	err = getRaw(c, &value, "market/"+symbol, nil)
	return
}
