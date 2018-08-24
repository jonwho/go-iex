package goiex

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	BaseURL            string = "https://api.iextrading.com/1.0/"
	MaxIdleConnections int    = 10
	RequestTimeout     int    = 5
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: createHTTPClient()}
}

// make a wrapper for http.Get
// NOTE: (c Client) without pointer just means you cannot modify the receiver
// (c *Client) has a pointer so you can mutate the public variables on the receiver c
func (c Client) Get(endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	iexURL := BaseURL + endpoint

	// 3rd arg is for the optional body.
	// body is expected type io.Reader
	req, err := http.NewRequest("GET", iexURL, nil)

	if err != nil {
		// %v modifier uses the default format for the data type
		log.Fatalln("Could not build request for " + iexURL)
	}

	// get query instance
	q := req.URL.Query()

	// set query params onto q
	for key, val := range params {
		q.Add(key, val)
	}

	// set query params onto request url
	req.URL.RawQuery = q.Encode()

	// perform request and return the response, error
	return c.httpClient.Do(req)
}

func (c Client) EarningsToday() string {
	endpoint := "stock/market/today-earnings"

	res, err := c.Get(endpoint, nil, nil)

	if err != nil {
		log.Fatalln("An error occurred for EarningsToday()")
	}
	// make sure to either close the response body or read the stream entirely
	// eitherwise the http client cannot be re-used as the connection is held up
	defer res.Body.Close()

	bodyBytes, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatalln("An error occurred for EarningsToday()")
	}

	return string(bodyBytes)
}

/***************************************************************************************************
 * PRIVATE BELOW
 **************************************************************************************************/

// return a configured http client for re-use
func createHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
}
