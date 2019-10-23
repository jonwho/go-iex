package goiex

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewDataAPI(t *testing.T) {
	u, _ := url.Parse(SandboxBaseURL)
	da := NewDataAPI("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = da.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = ""
	actual = da.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = da.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestDataPoints(t *testing.T) {
	rec, err := recorder.New("cassettes/data_api/data_points")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()

	u, _ := url.Parse(SandboxBaseURL)
	cli := NewDataAPI(testToken, DefaultVersion, u, httpClient)
	datapoints, err := cli.DataPoints("aapl")
	if err != nil {
		t.Error(err)
	}
	if len(datapoints) == 0 {
		t.Errorf("\nExpected datapoints to be not empty\n")
	}
}

func TestDataPoint(t *testing.T) {
	rec, err := recorder.New("cassettes/data_api/data_points_aapl_accountspayable")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()

	u, _ := url.Parse(SandboxBaseURL)
	cli := NewDataAPI(testToken, DefaultVersion, u, httpClient)
	datapoint, err := cli.DataPoint("aapl", "ACCOUNTSPAYABLE")
	if err != nil {
		t.Error(err)
	}
	expected = 37661668256.0
	actual = datapoint.(float64)
	if expected != actual {
		t.Errorf("\nExpected: %f\nActual: %f\n", expected, actual)
	}
}
