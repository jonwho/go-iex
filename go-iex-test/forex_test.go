package goiextest

import (
	"testing"
)

// TODO: @mock
func TestExchangeRates(t *testing.T) {
	er, err := iexSandboxClient.ExchangeRates("eur", "usd")
	if err != nil {
		t.Error(err)
	}
	expected = "EUR"
	actual = er.FromCurrency
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
