package goiex

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	DefaultBaseURL     string = "https://api.iextrading.com/1.0"
	MaxIdleConnections int    = 10
	RequestTimeout     int    = 5
)

var (
	// pseudo constant
	// you cannot make a constant out of a map because the expression is not a constant expression
	ChartRanges map[string]bool = make(map[string]bool)
)

type Option func(*Client) error

func init() {
	log.Println("Running init() for goiex")
	initChartRanges()
}

// Client interface to IEX.
func NewClient(options ...Option) (*Client, error) {
	client := &Client{}

	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}

	if client.httpClient == nil {
		client.httpClient = createHTTPClient()
	}

	if client.baseURL == "" {
		client.baseURL = DefaultBaseURL
	}
	return client, nil
}

// Assign your own http client.
func SetHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// Pretty much only useful for testing.
func SetBaseURL(url string) Option {
	return func(c *Client) error {
		c.baseURL = url
		return nil
	}
}

// Helper to perform request to IEX.
func (c Client) Get(endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	iexURL := c.baseURL + "/" + endpoint

	log.Println("Making request to " + iexURL)

	// 3rd arg is for the optional body.
	// body is expected type io.Reader
	req, err := http.NewRequest("GET", iexURL, nil)

	if err != nil {
		return nil, err
	}

	// get query instance
	q := req.URL.Query()

	// set query params onto q
	for key, val := range params {
		q.Add(key, val)
	}

	// set query params onto request url
	req.URL.RawQuery = q.Encode()

	log.Println("Built request with parameters is " + req.URL.String())

	// perform request and return the response, error
	return c.httpClient.Do(req)
}

// Returns Chart for the given symbol.
func (c *Client) Chart(symbol, chartRange string) (*Chart, error) {
	if !ChartRanges[chartRange] {
		return nil, errors.New("Received invalid date range for chart")
	}

	endpoint := "stock/" + symbol + "/chart/" + chartRange
	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Symbol")
	}

	chart := new(Chart)
	err = json.NewDecoder(res.Body).Decode(&chart)

	if err != nil {
		return nil, err
	}

	return chart, nil
}

// Returns Earnings for the given symbol.
func (c *Client) Earnings(symbol string) (*Earnings, error) {
	endpoint := "stock/" + symbol + "/earnings"
	earnings := new(Earnings)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Symbol")
	}

	err = json.NewDecoder(res.Body).Decode(&earnings)

	if err != nil {
		return nil, err
	}

	return earnings, nil
}

// Returns EarningsToday for the given symbol.
func (c *Client) EarningsToday() (*EarningsToday, error) {
	endpoint := "stock/market/today-earnings"
	earningsToday := new(EarningsToday)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&earningsToday)
	if err != nil {
		return nil, err
	}

	return earningsToday, nil
}

func (c *Client) Quote(symbol string, displayPercent bool) (*Quote, error) {
	endpoint := "stock/" + symbol + "/quote"
	quote := new(Quote)

	if displayPercent {
		endpoint = endpoint + "?displayPercent=true"
	}

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Symbol")
	}

	err = json.NewDecoder(res.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (c *Client) RefDataSymbols() (*RefDataSymbols, error) {
	endpoint := "ref-data/symbols"
	refDataSymbols := new(RefDataSymbols)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not finish request for /ref-data/symbols")
	}

	err = json.NewDecoder(res.Body).Decode(&refDataSymbols)
	if err != nil {
		return nil, err
	}

	return refDataSymbols, nil
}

func (c *Client) KeyStat(symbol string) (*KeyStat, error) {
	endpoint := "stock/" + symbol + "/stats"
	keyStat := new(KeyStat)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Symbol")
	}

	err = json.NewDecoder(res.Body).Decode(&keyStat)
	if err != nil {
		return nil, err
	}

	return keyStat, nil
}

/***************************************************************************************************
 * PRIVATE BELOW
 **************************************************************************************************/

// return a configured http client for re-use
func createHTTPClient() *http.Client {
	log.Println("Creating default http.Client for goiex")

	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
}

func initChartRanges() {
	allowedRanges := []string{
		"5y",
		"2y",
		"1y",
		"ytd",
		"6m",
		"3m",
		"1m",
		"1d",
	}

	for _, s := range allowedRanges {
		ChartRanges[s] = true
	}
}
