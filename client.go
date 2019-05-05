package goiex

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	// SandboxBaseURL use this as URL base if you do not want your account
	// message limits affected on IEX cloud
	SandboxBaseURL string = "https://sandbox.iexapis.com/"
	// DefaultBaseURL default base URL
	DefaultBaseURL string = "https://cloud.iexapis.com/"
	// DefaultVersion default IEX API version
	DefaultVersion string = "stable"

	maxIdleConnections int = 10
	requestTimeout     int = 5
)

var (
	// DefaultHTTPClient default http client to use
	DefaultHTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
)

type iex struct {
	token, version string
	url, apiurl    *url.URL
	client         *http.Client
}

type iexapi interface {
	Token() string
	Version() string
	URL() *url.URL
	APIURL() *url.URL
	Client() *http.Client
}

// Client API struct to IEX
type Client struct {
	iex

	*Account
	*DataAPI
	*Stock
}

// Option is a func that operates on *Client
type Option func(*Client) error

// NewClient creates client interface to IEX Cloud APIs
func NewClient(token string, options ...Option) (*Client, error) {
	client := &Client{}
	SetToken(token)(client)
	SetVersion(DefaultVersion)(client)
	SetURL(DefaultBaseURL)(client)
	SetHTTPClient(DefaultHTTPClient)(client)

	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}

	if client.Account == nil {
		SetAccount(client.iex.token, client.iex.version, client.iex.url, client.iex.client)(client)
	}
	if client.DataAPI == nil {
		SetDataAPI(client.iex.token, client.iex.version, client.iex.url, client.iex.client)(client)
	}
	if client.Stock == nil {
		SetStock(client.iex.token, client.iex.version, client.iex.url, client.iex.client)(client)
	}
	return client, nil
}

// SetToken assigns secret token
func SetToken(token string) Option {
	return func(c *Client) error {
		c.iex.token = token
		return nil
	}
}

// SetHTTPClient assigns HTTP client
func SetHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.iex.client = httpClient
		return nil
	}
}

// SetURL assigns URL base
func SetURL(rawurl string) Option {
	return func(c *Client) error {
		baseURL, err := url.Parse(rawurl)
		if err != nil {
			return err
		}

		c.iex.url = baseURL
		return nil
	}
}

// SetVersion set IEX version
func SetVersion(version string) Option {
	return func(c *Client) error {
		c.iex.version = version
		return nil
	}
}

// SetAccount set new Account
func SetAccount(token, version string, url *url.URL, httpClient *http.Client) Option {
	return func(c *Client) error {
		c.Account = NewAccount(token, version, url, httpClient)
		return nil
	}
}

// SetDataAPI set new DataAPI
func SetDataAPI(token, version string, url *url.URL, httpClient *http.Client) Option {
	return func(c *Client) error {
		c.Account = NewAccount(token, version, url, httpClient)
		return nil
	}
}

// SetStock set new Stock
func SetStock(token, version string, url *url.URL, httpClient *http.Client) Option {
	return func(c *Client) error {
		c.Stock = NewStock(token, version, url, httpClient)
		return nil
	}
}

// TODO: delete this func later
func (c Client) Get(endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	return nil, nil
}

func get(api iexapi, response interface{}, endpoint string, params interface{}) error {
	relurl, _ := url.Parse(endpoint)
	iexurl := baseURL(api).ResolveReference(relurl)
	req, err := http.NewRequest("GET", iexurl.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "jonwho/goiex")

	q := url.Values{}
	q.Set("token", api.Token())
	moreq, err := query.Values(params)
	if err != nil {
		return err
	}
	rawQuery := fmt.Sprintf("%s&%s", q.Encode(), moreq.Encode())
	req.URL.RawQuery = rawQuery

	resp, err := api.Client().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%v: %v", resp.Status, string(respBody))
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	return err
}

func baseURL(api iexapi) *url.URL {
	versionURL, _ := url.Parse(api.Version() + "/")
	return api.URL().ResolveReference(versionURL).ResolveReference(api.APIURL())
}
