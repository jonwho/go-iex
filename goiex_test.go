package goiex

import (
	"fmt"
	"testing"
)

var client = NewClient()

func TestEarningsToday(t *testing.T) {
	earningsToday, _ := client.EarningsToday()

	fmt.Printf("%+v\n\n", earningsToday)

	// TODO: really need mock lib because this endpoint does not work after market is over
	// it'll just return [] for both BTO and AMC
	// if len(earningsToday.Bto) < 1 {
	//   t.Error("fetch broke!")
	// }
	//
	// if len(earningsToday.Amc) < 1 {
	//   t.Error("fetch broke!")
	// }
}

func TestEarnings(t *testing.T) {
	earnings, _ := client.Earnings("aapl")

	// TODO: use a mock lib later to stub the response
	// only need to check struct works
	fmt.Printf("%+v\n\n", earnings)

	if earnings.Symbol != "AAPL" {
		t.Error("wrong string!")
	}

	if earnings.Earnings[0].SymbolId != 11 {
		t.Error("wrong value!")
	}
}

func TestQuote(t *testing.T) {
	// TODO: use mock lib to test optional arg for displayPercent=true
	// expect changePercent from 0.00919 -> 0.919 as an example
	quote, _ := client.Quote("aapl", false)

	fmt.Printf("%+v\n\n", quote)

	if quote.Symbol != "AAPL" {
		t.Error("wrong string!")
	}

	if quote.CompanyName != "Apple Inc." {
		t.Error("wrong string!")
	}

	_, err := client.Quote("fakesymbol", false)

	if err == nil {
		t.Error("err should not be nil!")
	}
}
