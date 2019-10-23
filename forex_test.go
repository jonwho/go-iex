package goiex

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewForex(t *testing.T) {
	u, _ := url.Parse(SandboxBaseURL)
	fx := NewForex("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = fx.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "fx/"
	actual = fx.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = fx.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestExchangeRates(t *testing.T) {
	rec, err := recorder.New("cassettes/forex/exchange_rates")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()

	u, _ := url.Parse(SandboxBaseURL)
	cli := NewForex(testToken, DefaultVersion, u, httpClient)
	er, err := cli.ExchangeRates("eur", "usd")
	if err != nil {
		t.Error(err)
	}
	expected = "RUE"
	actual = er.FromCurrency
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
