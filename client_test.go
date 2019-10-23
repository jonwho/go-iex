package goiex

import (
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
	prodToken        = os.Getenv("IEX_SECRET_TOKEN")
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
	if cli.AlternativeData == nil {
		t.Error("Should have a default value")
	}
	if cli.APISystemMetadata == nil {
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
	return nil
}
