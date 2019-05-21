package goiextest

import (
	"testing"
)

func TestSymbols(t *testing.T) {
	symbols, err := iexSandboxClient.Symbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestIEXSymbols(t *testing.T) {
	symbols, err := iexSandboxClient.IEXSymbols()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestRegionSymbols(t *testing.T) {
	symbols, err := iexSandboxClient.RegionSymbols("ca")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestExchangeSymbols(t *testing.T) {
	symbols, err := iexSandboxClient.ExchangeSymbols("tse")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(symbols) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
