package goiextest

import (
	"testing"
)

func TestDataPoints(t *testing.T) {
	datapoints, err := iexSandboxClient.DataPoints("aapl")
	if err != nil {
		t.Error(err)
	}
	if len(datapoints) == 0 {
		t.Errorf("\nExpected datapoints to be not empty\n")
	}
}
