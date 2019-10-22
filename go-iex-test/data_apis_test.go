package goiextest

import (
	"testing"
)

func TestDataPoints(t *testing.T) {
	datapoints, err := mockClient.DataPoints("aapl")
	if err != nil {
		t.Error(err)
	}
	if len(datapoints) == 0 {
		t.Errorf("\nExpected datapoints to be not empty\n")
	}
}

func TestDataPoint(t *testing.T) {
	datapoint, err := mockClient.DataPoint("aapl", "ACCOUNTSPAYABLE")
	if err != nil {
		t.Error(err)
	}
	expected = 37294000000.0
	actual = datapoint.(float64)
	if expected != actual {
		t.Errorf("\nExpected: %f\nActual: %f\n", expected, actual)
	}
}
