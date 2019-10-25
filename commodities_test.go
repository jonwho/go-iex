package goiex

import (
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewCommodities(t *testing.T) {
	cm := NewCommodities("test_token", "stable", sandboxURL, nil)

	expected = "data-points/"
	actual = cm.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = true
	actual = cm.Client() == nil
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = cm.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "https://sandbox.iexapis.com/"
	actual = cm.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "stable"
	actual = cm.Version()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestCommoditiesPrices(t *testing.T) {
	rec, err := recorder.New("cassettes/data_point/commodities_prices")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewCommodities(testToken, DefaultVersion, sandboxURL, httpClient)

	// Oil Prices
	prices, err := cli.CommoditiesPrices("DCOILWTICO")
	if err != nil {
		t.Error(err)
	}
	expected = 53.93
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Natural Gas Prices
	prices, err = cli.CommoditiesPrices("DHHNGSP")
	if err != nil {
		t.Error(err)
	}
	expected = 2.22
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Heating Oil Prices
	prices, err = cli.CommoditiesPrices("DHOILNYH")
	if err != nil {
		t.Error(err)
	}
	expected = 1.98
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Jet Fuel Prices
	prices, err = cli.CommoditiesPrices("DJFUELUSGULF")
	if err != nil {
		t.Error(err)
	}
	expected = 1.942
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Diesel Prices
	prices, err = cli.CommoditiesPrices("GASDESW")
	if err != nil {
		t.Error(err)
	}
	expected = 3.14
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Gas Prices
	prices, err = cli.CommoditiesPrices("GASREGCOVW")
	if err != nil {
		t.Error(err)
	}
	expected = 2.541
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	// Propane Prices
	prices, err = cli.CommoditiesPrices("DPROPANEMBTX")
	if err != nil {
		t.Error(err)
	}
	expected = 0.482
	actual, _ = strconv.ParseFloat(string(prices.([]uint8)), 64)
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}
