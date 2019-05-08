package goiextest

import (
	"net/url"
	"os"
	"testing"

	iex "github.com/jonwho/go-iex"
)

func TestStatus(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	sys := iex.NewAPISystemMetadata(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	status, err := sys.Status()
	if err != nil {
		t.Error(err)
	}
	expected = `up`
	actual = status.Status
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
