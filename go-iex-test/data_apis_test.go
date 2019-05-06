package goiextest

import (
	"net/url"
	"os"
	"testing"

	iex "github.com/jonwho/go-iex"
)

func TestDataPoints(t *testing.T) {
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	data := iex.NewDataAPI(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	datapoints, err := data.DataPoints("aapl")
	if err != nil {
		t.Error(err)
	}
	if len(datapoints) == 0 {
		t.Errorf("\nExpected datapoints to be not empty\n")
	}
}
