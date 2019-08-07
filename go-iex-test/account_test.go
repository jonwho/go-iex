package goiextest

import (
	"net/url"
	"testing"

	iex "github.com/jonwho/go-iex"
)

func TestNewAccount(t *testing.T) {
	u, _ := url.Parse(iex.SandboxBaseURL)
	acc := iex.NewAccount("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = acc.URL().String()

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
	metadata, err := iexSandboxClient.Metadata()
	if err != nil {
		t.Error(err)
	}
	if metadata == nil {
		t.Errorf("\nExpected metadata to be not nil\n")
	}
}

func TestUsage(t *testing.T) {
	usage, err := iexSandboxClient.Usage()
	if err != nil {
		t.Error(err)
	}
	if usage.Messages.MonthlyUsage < 1 {
		t.Errorf("\nExpected MonthlyUsage > 0 but got: %d\n", usage.Messages.MonthlyUsage)
	}
	expected = "6"
	if actual = usage.Messages.DailyUsage.(map[string]interface{})["20190801"]; actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}
