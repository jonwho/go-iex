package goiex

import (
	"fmt"
	"net/http"
	"net/url"
)

// AlternativeData struct to interface with AlternativeData endpoints
type AlternativeData struct {
	iex
}

// CryptoQuote struct
type CryptoQuote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              int64   `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             int64   `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          int64   `json:"latestUpdate"`
	LatestVolume          float64 `json:"latestVolume"`
	IexRealtimePrice      float64 `json:"iexRealtimePrice"`
	IexRealtimeSize       float64 `json:"iexRealtimeSize"`
	IexLastUpdated        int64   `json:"iexLastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      int64   `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     int64   `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	IexMarketPercent      float64 `json:"iexMarketPercent"`
	IexVolume             float64 `json:"iexVolume"`
	AvgTotalVolume        float64 `json:"avgTotalVolume"`
	IexBidPrice           float64 `json:"iexBidPrice"`
	IexBidSize            float64 `json:"iexBidSize"`
	IexAskPrice           float64 `json:"iexAskPrice"`
	IexAskSize            float64 `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YtdChange             float64 `json:"ytdChange"`
}

// SocialSentiment struct
type SocialSentiment struct {
	Sentiment   float64 `json:"sentiment"`
	TotalScores int     `json:"totalScores"`
	Positive    float64 `json:"positive"`
	Negative    float64 `json:"negative"`
	Minute      string  `json:"minute"`
}

// CEOCompensation struct
type CEOCompensation struct {
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	CompanyName         string  `json:"companyName"`
	Location            string  `json:"location"`
	Salary              float32 `json:"salary"`
	Bonus               float32 `json:"bonus"`
	StockAwards         float32 `json:"stockAwards"`
	OptionAwards        float32 `json:"optionAwards"`
	NonEquityIncentives float32 `json:"nonEquityIncentives"`
	PensionAndDeferred  float32 `json:"pensionAndDeferred"`
	OtherComp           float32 `json:"otherComp"`
	Total               float32 `json:"total"`
	Year                string  `json:"year"`
}

// NewAlternativeData return new AlternativeData
func NewAlternativeData(token, version string, base *url.URL, httpClient *http.Client) *AlternativeData {
	apiurl, err := url.Parse("")
	if err != nil {
		panic(err)
	}
	return &AlternativeData{
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
func (ad *AlternativeData) Token() string {
	return ad.token
}

// Version return version string
func (ad *AlternativeData) Version() string {
	return ad.version
}

// URL return URL base
func (ad *AlternativeData) URL() *url.URL {
	return ad.url
}

// APIURL return APIURL
func (ad *AlternativeData) APIURL() *url.URL {
	return ad.apiurl
}

// Client return HTTP client
func (ad *AlternativeData) Client() *http.Client {
	return ad.client
}

// Crypto GET /crypto/{symbol}/quote
func (ad *AlternativeData) Crypto(symbol string) (q *CryptoQuote, err error) {
	endpoint := fmt.Sprintf("%s%s/crypto/%s/quote", ad.url.String(), ad.version, symbol)
	err = get(ad, &q, endpoint, nil)
	return
}

// SocialSentiment GET /stock/{symbol}/sentiment
func (ad *AlternativeData) SocialSentiment(symbol string) (ss *SocialSentiment, err error) {
	endpoint := fmt.Sprintf("stock/%s/sentiment", symbol)
	err = get(ad, &ss, endpoint, nil)
	return
}

// CEOCompensation GET /stock/{symbol}/ceo-compensation
func (ad *AlternativeData) CEOCompensation(symbol string) (cc *CEOCompensation, err error) {
	endpoint := fmt.Sprintf("stock/%s/ceo-compensation", symbol)
	err = get(ad, &cc, endpoint, nil)
	return
}
