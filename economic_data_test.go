package goiex

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewEconomicData(t *testing.T) {
	ed := NewEconomicData("test_token", "stable", sandboxURL, nil)

	expected = "data-points/"
	actual = ed.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = true
	actual = ed.Client() == nil
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = ed.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "https://sandbox.iexapis.com/"
	actual = ed.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "stable"
	actual = ed.Version()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestEconomicPrices(t *testing.T) {
	rec, err := recorder.New("cassettes/data_point/economic_prices")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewEconomicData(testToken, DefaultVersion, sandboxURL, httpClient)

	// Mortgage Rates
	prices, err := cli.EconomicPrices("MORTGAGE30US")
	if err != nil {
		t.Error(err)
	}
	expected = 3.89
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Federal Fund Rates
	prices, err = cli.EconomicPrices("FEDFUNDS")
	if err != nil {
		t.Error(err)
	}
	expected = 2.12
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Credit Card Interest Rate
	prices, err = cli.EconomicPrices("TERMCBCCALLNS")
	if err != nil {
		t.Error(err)
	}
	expected = 15.2
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// CD Rates
	prices, err = cli.EconomicPrices("MMNRNJ")
	if err != nil {
		t.Error(err)
	}
	expected = 0.16
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Real GDP
	prices, err = cli.EconomicPrices("A191RL1Q225SBEA")
	if err != nil {
		t.Error(err)
	}
	expected = 2.0
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Industrial Production Index
	prices, err = cli.EconomicPrices("INDPRO")
	if err != nil {
		t.Error(err)
	}
	expected = 112.8888
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Consumer Price Index
	prices, err = cli.EconomicPrices("CPIAUCSL")
	if err != nil {
		t.Error(err)
	}
	expected = 259.705
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Total Payrolls
	prices, err = cli.EconomicPrices("PAYEMS")
	if err != nil {
		t.Error(err)
	}
	expected = 154599.0
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Total Housing Starts
	prices, err = cli.EconomicPrices("HOUST")
	if err != nil {
		t.Error(err)
	}
	expected = 1285.0
	actual = prices.(float64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Unemployment Rate
	prices, err = cli.EconomicPrices("UNRATE")
	if err != nil {
		t.Error(err)
	}
	expected = 3.6
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Total Vehicle Sales
	prices, err = cli.EconomicPrices("TOTALSA")
	if err != nil {
		t.Error(err)
	}
	expected = 18.045
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// US Recession Probabilities
	prices, err = cli.EconomicPrices("RECPROUSM156N")
	if err != nil {
		t.Error(err)
	}
	expected = 0.49
	actual = prices
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}
