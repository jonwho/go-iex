package goiex

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewInvestorsExchangeData(t *testing.T) {
	u, _ := url.Parse(SandboxBaseURL)
	ied := NewInvestorsExchangeData("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = ied.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = ""
	actual = ied.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = ied.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestTOPS(t *testing.T) {
	rec, err := recorder.New("cassettes/investors_exchange_data/tops")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	cli := NewInvestorsExchangeData(testToken, DefaultVersion, u, httpClient)

	tops, err := cli.TOPS(nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(tops) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	tops, err = cli.TOPS(&TOPSParams{Symbols: "SNAP,fb,AIG+"})
	if err != nil {
		t.Error(err)
	}
	expected = "SNAP"
	actual = tops[0].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "FB"
	actual = tops[1].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AIG+"
	actual = tops[2].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestLast(t *testing.T) {
	rec, err := recorder.New("cassettes/investors_exchange_data/last")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	cli := NewInvestorsExchangeData(testToken, DefaultVersion, u, httpClient)

	last, err := cli.Last(nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(last) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	last, err = cli.Last(&LastParams{Symbols: "SNAP,fb,AIG+"})
	if err != nil {
		t.Error(err)
	}
	expected = "SNAP"
	actual = last[0].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "FB"
	actual = last[1].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AIG+"
	actual = last[2].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestDEEP(t *testing.T) {
	rec, err := recorder.New("cassettes/investors_exchange_data/deep")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	cli := NewInvestorsExchangeData(testToken, DefaultVersion, u, httpClient)

	deep, err := cli.DEEP(&DEEPParams{Symbols: "SNAP"})
	if err != nil {
		t.Error(err)
	}
	expected = "SNAP"
	actual = deep.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
