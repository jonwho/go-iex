package goiex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	defaultBaseURL     string = "https://api.iextrading.com/1.0"
	maxIdleConnections int    = 10
	requestTimeout     int    = 5
)

var (
	chartRanges = make(map[string]bool)
)

// Option is a func that operates on *Client
type Option func(*Client) error

func init() {
	log.Println("Running init() for goiex")
	initChartRanges()
}

// NewClient creates interface to IEX
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
		client.baseURL = defaultBaseURL
	}
	return client, nil
}

// SetHTTPClient lets you assign your own HTTP client
func SetHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// SetBaseURL pretty much only useful for testing
func SetBaseURL(url string) Option {
	return func(c *Client) error {
		c.baseURL = url
		return nil
	}
}

// Get is a helper to perform request to IEX
// TODO: ran into limitation here where for /stock/market/batch?types=quote&symbols=aapl,fb
//       the escaped URL string could not be built correctly
// TODO: create func (c Client) GetRaw(rawURL string) which could just call this func perhaps --
//       test later
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

// Batch call to /stock/<symbol>/batch
func (c *Client) Batch(symbol string, params map[string]string) (*Batch, error) {
	endpoint := "stock/" + symbol + "/batch"

	res, err := c.Get(endpoint, params, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Batch call")
	}

	batch := new(Batch)
	err = json.NewDecoder(res.Body).Decode(&batch)

	if err != nil {
		return nil, err
	}

	return batch, nil
}

// Book call to /book
func (c *Client) Book(symbol string) (*Book, error) {
	endpoint := "stock/" + symbol + "/book"

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Invalid Book call")
	}

	book := new(Book)
	err = json.NewDecoder(res.Body).Decode(&book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

// Chart call to /chart
func (c *Client) Chart(symbol, chartRange string) (*Chart, error) {
	if !chartRanges[chartRange] {
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
		return nil, errors.New("Invalid Chart call")
	}

	chart := new(Chart)
	err = json.NewDecoder(res.Body).Decode(&chart)

	if err != nil {
		return nil, err
	}

	return chart, nil
}

// Earnings call to /earnings
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
		return nil, errors.New("Invalid Earnings call")
	}

	err = json.NewDecoder(res.Body).Decode(&earnings)

	if err != nil {
		return nil, err
	}

	return earnings, nil
}

// EarningsToday call to /market/todays-earnings
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

// News call to /news
func (c *Client) News(symbol string, last int) (*News, error) {
	endpoint := fmt.Sprintf("stock/%s/news/last/%d", symbol, last)
	news := new(News)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

// Quote call to /quote
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
		return nil, errors.New("Invalid Quote call")
	}

	err = json.NewDecoder(res.Body).Decode(&quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

// RefDataSymbols call to /ref-data/symbols
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

// RefDataCorporateActions call to /ref-data/daily-list/corporate-actions
func (c *Client) RefDataCorporateActions(date string) (*RefDataCorporateActions, error) {
	var endpoint string
	if date != "" {
		endpoint = "ref-data/daily-list/corporate-actions/" + date
	} else {
		endpoint = "ref-data/daily-list/corporate-actions"
	}

	refDataCorporateActions := new(RefDataCorporateActions)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not finish request for /ref-data/daily-list/corporate-actions")
	}

	err = json.NewDecoder(res.Body).Decode(&refDataCorporateActions)
	if err != nil {
		return nil, err
	}

	return refDataCorporateActions, nil
}

// RefDataDividends call to /ref-data/daily-list/dividends
func (c *Client) RefDataDividends(date string) (*RefDataDividends, error) {
	var endpoint string
	if date != "" {
		endpoint = "ref-data/daily-list/dividends/" + date
	} else {
		endpoint = "ref-data/daily-list/dividends"
	}

	refDataDividends := new(RefDataDividends)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not finish request for /ref-data/daily-list/dividends")
	}

	err = json.NewDecoder(res.Body).Decode(&refDataDividends)
	if err != nil {
		return nil, err
	}

	return refDataDividends, nil
}

// RefDataNextDayExDates call to /ref-data/daily-list/next-day-ex-date
func (c *Client) RefDataNextDayExDates(date string) (*RefDataNextDayExDates, error) {
	var endpoint string
	if date != "" {
		endpoint = "ref-data/daily-list/next-day-ex-date/" + date
	} else {
		endpoint = "ref-data/daily-list/next-day-ex-date"
	}

	refDataNextDayExDates := new(RefDataNextDayExDates)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not finish request for /ref-data/daily-list/next-day-ex-date")
	}

	err = json.NewDecoder(res.Body).Decode(&refDataNextDayExDates)
	if err != nil {
		return nil, err
	}

	return refDataNextDayExDates, nil
}

// RefDataSymbolDirectories call to /ref-data/daily-list/symbol-directory
func (c *Client) RefDataSymbolDirectories(date string) (*RefDataSymbolDirectories, error) {
	var endpoint string
	if date != "" {
		endpoint = "ref-data/daily-list/symbol-directory/" + date
	} else {
		endpoint = "ref-data/daily-list/symbol-directory"
	}

	refDataSymbolDirectories := new(RefDataSymbolDirectories)

	res, err := c.Get(endpoint, nil, nil)

	// use defer only if http.Get is successful
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not finish request for /ref-data/daily-list/symbol-directory")
	}

	err = json.NewDecoder(res.Body).Decode(&refDataSymbolDirectories)
	if err != nil {
		return nil, err
	}

	return refDataSymbolDirectories, nil
}

// KeyStat call to /stats
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
		return nil, errors.New("Invalid KeyStat call")
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
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
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
		chartRanges[s] = true
	}
}
