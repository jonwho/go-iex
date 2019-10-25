package goiex

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonwho/go-iex/mock-iex"
)

var (
	expected, actual interface{}
	testToken        = os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient       *http.Client
	mockClient, _    = NewClient("", SetURL(mockiex.Server().URL))
	sandboxURL, _    = url.Parse(SandboxBaseURL)
)

func TestNewClient(t *testing.T) {
	token := "test_token"
	cli, err := NewClient(token, SetURL(SandboxBaseURL))
	if err != nil {
		t.Error(err)
	}

	if cli.Account == nil {
		t.Error("Should have a default value")
	}
	if cli.APISystemMetadata == nil {
		t.Error("Should have a default value")
	}
	if cli.Cryptocurrency == nil {
		t.Error("Should have a default value")
	}
	if cli.DataAPI == nil {
		t.Error("Should have a default value")
	}
	if cli.Forex == nil {
		t.Error("Should have a default value")
	}
	if cli.InvestorsExchangeData == nil {
		t.Error("Should have a default value")
	}
	if cli.ReferenceData == nil {
		t.Error("Should have a default value")
	}
	if cli.Stock == nil {
		t.Error("Should have a default value")
	}
}

func TestGet(t *testing.T) {
	rec, err := recorder.New("cassettes/misc/client_get")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, SetURL(SandboxBaseURL), SetHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
	}

	anonstruct := &struct {
		Status string `json:"status,omitempty"`
	}{}
	err = cli.Get("status", anonstruct, nil)
	if err != nil {
		t.Error(err)
	}
	expected = `up`
	actual = anonstruct.Status
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestPost(t *testing.T) {
	rec, err := recorder.New("cassettes/misc/client_post")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli, err := NewClient(testToken, SetURL(SandboxBaseURL), SetHTTPClient(httpClient))
	if err != nil {
		t.Error(err)
	}

	anonstruct := &struct {
		OverageCircuitBreaker int `json:"overageCircuitBreaker,omitempty"`
	}{}
	err = cli.Post("account/messagebudget", anonstruct, map[string]interface{}{
		"token":         testToken,
		"totalMessages": 100_000,
	})
	if err != nil {
		t.Error(err)
	}
	expected = 102759
	actual = anonstruct.OverageCircuitBreaker
	if expected != actual {
		t.Errorf("\nExpected: %f\nActual: %f\n", expected, actual)
	}
}

func matchWithoutToken(req *http.Request, i cassette.Request) bool {
	u := req.URL
	q := u.Query()
	q.Del("token")
	u.RawQuery = q.Encode()
	req.URL = u
	return u.String() == i.URL
}

func removeToken(i *cassette.Interaction) error {
	u, err := url.Parse(i.Request.URL)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Del("token")
	u.RawQuery = q.Encode()
	i.Request.URL = u.String()

	originalBody := []byte(i.Request.Body)
	var unmarshalBody map[string]interface{}
	if err = json.Unmarshal(originalBody, &unmarshalBody); err != nil {
		// sometimes the data isn't JSON so ignore unmarshal error
		return nil
	}
	delete(unmarshalBody, "token")
	bodyWithoutToken, err := json.Marshal(unmarshalBody)
	i.Request.Body = string(bodyWithoutToken)
	return nil
}
