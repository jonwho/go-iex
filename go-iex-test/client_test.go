package goiextest

import (
	"os"
	"testing"

	iex "github.com/jonwho/go-iex"
)

func TestNewClient(t *testing.T) {
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	cli, err := iex.NewClient(token, iex.SetURL(iex.SandboxBaseURL))
	if err != nil {
		t.Error(err)
	}

	if cli.Account == nil {
		t.Error("Should have a default value")
	}
	if cli.AlternativeData == nil {
		t.Error("Should have a default value")
	}
	if cli.APISystemMetadata == nil {
		t.Error("Should have a default value")
	}
	if cli.DataAPI == nil {
		t.Error("Should have a default value")
	}
	if cli.Forex == nil {
		t.Error("Should have a default value")
	}
	if cli.InvestorsExchangeData == nil {
		t.Error("Should have a default value")
	}
	if cli.ReferenceData == nil {
		t.Error("Should have a default value")
	}
	if cli.Stock == nil {
		t.Error("Should have a default value")
	}
}

func TestGet(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	cli, err := iex.NewClient(token, iex.SetURL(iex.SandboxBaseURL))
	if err != nil {
		t.Error(err)
	}

	anonstruct := &struct {
		Status string `json:"status,omitempty"`
	}{}
	err = cli.Get("status", anonstruct, nil)
	if err != nil {
		t.Error(err)
	}
	expected = `up`
	actual = anonstruct.Status
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
