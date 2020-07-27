package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Account struct to interface with /account endpoints
type Account struct {
	iex

	RetryWaitMin  time.Duration // Minimum time to wait on HTTP request retry
	RetryWaitMax  time.Duration // Maximum time to wait on HTTP request retry
	RetryAttempts int           // Maximum number of HTTP request retries
	RetryPolicy   RetryPolicy   // Defines when to retry a HTTP request
	Backoff       Backoff       // Defines wait time between HTTP request retries
}

// Metadata struct
type Metadata struct {
	PayAsYouGoEnabled    bool   `json:"payAsYouGoEnabled"`
	EffectiveDate        int64  `json:"effectiveDate"`
	EndDateEffective     int64  `json:"endDateEffective"`
	SubscriptionTermType string `json:"subscriptionTermType"`
	TierName             string `json:"tierName"`
	MessageLimit         int    `json:"messageLimit"`
	MessagesUsed         int    `json:"messagesUsed"`
}

// Usage struct
type Usage struct {
	Messages struct {
		MonthlyUsage      int         `json:"monthlyUsage"`
		MonthlyPayAsYouGo int         `json:"monthlyPayAsYouGo"`
		DailyUsage        interface{} `json:"dailyUsage"`
		TokenUsage        interface{} `json:"tokenUsage"`
		KeyUsage          interface{} `json:"keyUsage"`
	} `json:"messages"`
}

// NewAccount return new Account
func NewAccount(token, version string, base *url.URL, httpClient *http.Client) *Account {
	apiurl, err := url.Parse("account/")
	if err != nil {
		panic(err)
	}

	return &Account{
		RetryWaitMin:  defaultRetryWaitMin,
		RetryWaitMax:  defaultRetryWaitMax,
		RetryAttempts: defaultRetryAttempts,
		RetryPolicy:   DefaultRetryPolicy,
		Backoff:       DefaultBackoff,
		iex: iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// Token return token string
func (a *Account) Token() string {
	return a.token
}

// Version return version string
func (a *Account) Version() string {
	return a.version
}

// URL return URL base
func (a *Account) URL() *url.URL {
	return a.url
}

// APIURL return APIURL
func (a *Account) APIURL() *url.URL {
	return a.apiurl
}

// Client return HTTP client
func (a *Account) Client() *http.Client {
	return a.client
}

func (a *Account) Do(req *Request) (*http.Response, error) {
	for i := 0; i < a.RetryAttempts; i++ {
		// Rewind the request body
		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		// Attempt request
		resp, err := a.iex.client.Do(req.Request)

		// No RetryPolicy policy set so return right away
		if a.RetryPolicy == nil {
			return resp, err
		}

		// Check for retry
		checkOK, checkErr := a.RetryPolicy(resp, err)
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}
			return resp, err
		}

		// Perform retry
		if err == nil {
			drainBody(resp.Body)
		}

		remain := a.RetryAttempts - i
		if remain == 0 {
			break
		}
		wait := a.Backoff(a.RetryWaitMin, a.RetryWaitMax, i, resp)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("%s %s request failed after %d attempts", req.Method, req.URL, a.RetryAttempts+1)
}

// Metadata GET /account/metadata
func (a *Account) Metadata() (metadata *Metadata, err error) {
	err = get(a, &metadata, "metadata", nil)
	return
}

// Usage GET /account/usage
// No support for GET /account/usage/{type}
func (a *Account) Usage() (usage *Usage, err error) {
	err = get(a, &usage, "usage", nil)
	return
}

// Payasyougo POST /account/payasyougo
func (a *Account) Payasyougo(params interface{}) (ifc interface{}, err error) {
	err = post(a, &ifc, "payasyougo", params.(map[string]interface{}))
	return
}

// MessageBudget POST /account/messagebudget
func (a *Account) MessageBudget(params interface{}) (ifc interface{}, err error) {
	err = post(a, &ifc, "messagebudget", params.(map[string]interface{}))
	return
}
