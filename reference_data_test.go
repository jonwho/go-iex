package goiex

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewReferenceData(t *testing.T) {
	u, _ := url.Parse(SandboxBaseURL)
	rd := NewReferenceData("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = rd.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "ref-data/"
	actual = rd.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = rd.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	symbols, err := cli.Symbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestIEXSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/iex_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	symbols, err := cli.IEXSymbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestRegionSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/region_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	symbols, err := cli.RegionSymbols("ca")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestExchangeSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/exchange_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	symbols, err := cli.ExchangeSymbols("tse")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestInternationalExchanges(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/international_exchanges")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	exchanges, err := cli.InternationalExchanges()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(exchanges) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUSExchanges(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/us_exchanges")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	exchanges, err := cli.USExchanges()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(exchanges) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUSHolidaysAndTradingDates(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/us_holidays_and_trading_dates")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	dates, err := cli.USHolidaysAndTradingDates("trade", "next")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	dates, err = cli.USHolidaysAndTradingDates("trade", "last")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	dates, err = cli.USHolidaysAndTradingDates("holiday", "next")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	dates, err = cli.USHolidaysAndTradingDates("holiday", "last")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	dates, err = cli.USHolidaysAndTradingDates("trade", "next", 1)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	dates, err = cli.USHolidaysAndTradingDates("holiday", "last", 1, "20190101")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSectors(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/sectors")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	sectors, err := cli.Sectors()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sectors) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestTags(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/tags")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	tags, err := cli.Tags()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(tags) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestMutualFundSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/mutual_fund_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	funds, err := cli.MutualFundSymbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(funds) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestOTCSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/otc_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	otc, err := cli.OTCSymbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(otc) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestFXSymbols(t *testing.T) {
	rec, err := recorder.New("cassettes/reference_data/fx_symbols")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	u, _ := url.Parse(SandboxBaseURL)
	retry, _ := NewRetry(httpClient, SetRetryAttempts(7))
	cli := NewReferenceData(testToken, DefaultVersion, u, httpClient, SetRetry(retry))

	fx, err := cli.FXSymbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fx.Currencies) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: fix implementation
// func TestOptionsSymbols(t *testing.T) {
//   os, err := cli.OptionsSymbols()
//   if err != nil {
//     t.Error(err)
//   }
// }
