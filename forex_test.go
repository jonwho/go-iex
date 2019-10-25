package goiex

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewForex(t *testing.T) {
	fx := NewForex("test_token", "", sandboxURL, nil)

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

func TestLatestCurrencyRates(t *testing.T) {
	rec, err := recorder.New("cassettes/forex/latest_currency_rates")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewForex(testToken, DefaultVersion, sandboxURL, httpClient)

	lcr, err := cli.LatestCurrencyRates(struct {
		Symbols string `url:"symbols"`
	}{Symbols: "USDCAD,USDGBP,USDJPY"})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(lcr) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCurrencyConversion(t *testing.T) {
	rec, err := recorder.New("cassettes/forex/currency_conversion")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewForex(testToken, DefaultVersion, sandboxURL, httpClient)

	cc, err := cli.CurrencyConversion(struct {
		Symbols string `url:"symbols"`
		Amount  int    `url:"amount"`
	}{Symbols: "USDCAD,USDGBP,USDJPY", Amount: 73})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(cc) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestHistoricalDaily(t *testing.T) {
	rec, err := recorder.New("cassettes/forex/historical_daily")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewForex(testToken, DefaultVersion, sandboxURL, httpClient)

	hd, err := cli.HistoricalDaily(struct {
		Symbols string `url:"symbols"`
		From    string `url:"from"`
		To      string `url:"to"`
		Last    int    `url:"last"`
	}{Symbols: "USDCAD,USDGBP,USDJPY", From: "2019-10-24", To: "2019-10-25", Last: 5})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(hd) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
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
	cli := NewForex(testToken, DefaultVersion, sandboxURL, httpClient)

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
