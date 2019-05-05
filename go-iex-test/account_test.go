package goiextest

import (
	"net/url"
	"os"
	"testing"

	iex "github.com/jonwho/go-iex"
	// "github.com/jonwho/go-iex/mock-iex"
)

func TestNewAccount(t *testing.T) {
	u, _ := url.Parse(iex.SandboxBaseURL)
	acc := iex.NewAccount("test_token", "", u, nil)

	expected := "https://sandbox.iexapis.com/"
	actual := acc.URL().String()

	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = acc.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestMetadata(t *testing.T) {
	token := os.Getenv("IEX_SECRET_TOKEN")
	u, _ := url.Parse(iex.DefaultBaseURL)
	acc := iex.NewAccount(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	metadata, err := acc.Metadata()
	if err != nil {
		t.Error(err)
	}
	if metadata == nil {
		t.Errorf("\nExpected metadata to be not nil\n")
	}
}
