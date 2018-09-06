package goiex

import (
	"fmt"
	"reflect"
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

func TestChart(t *testing.T) {
	_, err := client.Chart("aapl", "6y")

	if err == nil {
		t.Error("err should not be nil!")
	}

	chart, _ := client.Chart("aapl", "1d")

	fmt.Printf("%+v\n\n", chart)

	if len(chart.Charts) == 0 {
		t.Error("charts shouldn't be empty")
	}

	if chart.Charts[0].Minute == "" {
		t.Error("minute should be non-empty string for 1d range")
	}
}

func TestRefDataSymbols(t *testing.T) {
	rds, _ := client.RefDataSymbols()

	if len(rds.Symbols) == 0 {
		t.Error("NANI?")
	}

	firstSymbol := rds.Symbols[0]

	if firstSymbol.Symbol == "" {
		t.Error("should not be zero-val")
	}

	if firstSymbol.Date == "" {
		t.Error("should not be zero-val")
	}

	if firstSymbol.Name == "" {
		t.Error("should not be zero-val")
	}

	if firstSymbol.IsEnabled == false {
		t.Error("should not be zero-val")
	}

	if firstSymbol.Type == "" {
		t.Error("should not be zero-val")
	}

	if firstSymbol.IexId == 0 {
		t.Error("should not be zero-val")
	}

	lastSymbol := rds.Symbols[len(rds.Symbols)-1]
	if lastSymbol.IexId == 0 {
		t.Error("should not be zero-val")
	}

	fmt.Println(firstSymbol.IexId)
	fmt.Println(lastSymbol.IexId)
	fmt.Println(reflect.TypeOf(firstSymbol.IexId))
	fmt.Println(reflect.TypeOf(lastSymbol.IexId))
}
