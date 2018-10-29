package goiex

import (
	"testing"

	"github.com/jonwho/go-iex/mock-iex"
)

// TODO: use during hours reponse stub
func TestEarningsToday(t *testing.T) {
	mockServer := mockiex.MockIEXServer()
	defer mockServer.Close()
	client, err := NewClient(SetBaseURL(mockServer.URL))
	if err != nil {
		t.Error(err)
	}

	earningsToday, _ := client.EarningsToday()

	if len(earningsToday.BTO) != 0 {
		t.Errorf("expected 0 but got %v", earningsToday.BTO)
	}

	if len(earningsToday.AMC) != 0 {
		t.Errorf("expected 0 but got %v", earningsToday.AMC)
	}
}

func TestEarnings(t *testing.T) {
	mockServer := mockiex.MockIEXServer()
	defer mockServer.Close()
	client, err := NewClient(SetBaseURL(mockServer.URL))
	if err != nil {
		t.Error(err)
	}

	earnings, _ := client.Earnings("aapl")

	if earnings.Symbol != "AAPL" {
		t.Errorf("expected AAPL but got %v", earnings.Symbol)
	}

	if earnings.Earnings[0].SymbolId != 11 {
		t.Errorf("expected 11 but got %v", earnings.Earnings[0].SymbolId)
	}
}

func TestQuote(t *testing.T) {
	mockServer := mockiex.MockIEXServer()
	defer mockServer.Close()
	client, err := NewClient(SetBaseURL(mockServer.URL))
	if err != nil {
		t.Error(err)
	}

	quote, _ := client.Quote("aapl", false)

	if quote.Symbol != "AAPL" {
		t.Errorf("expected AAPL but got %v", quote.Symbol)
	}

	if quote.CompanyName != "Apple Inc." {
		t.Errorf("expected Apple Inc. but got %v", quote.CompanyName)
	}

	if quote.ChangePercent != -0.01592 {
		t.Errorf("expected -0.01592 but got %v", quote.ChangePercent)
	}

	_, err = client.Quote("fakesymbol", false)

	if err == nil {
		t.Error("expected err but got nil")
	}

	quote, _ = client.Quote("aapl", true)

	if quote.ChangePercent != -0.01592*100 {
		t.Errorf("expected -1.592 but got %v", quote.ChangePercent)
	}
}

func TestChart(t *testing.T) {
	mockServer := mockiex.MockIEXServer()
	defer mockServer.Close()
	client, err := NewClient(SetBaseURL(mockServer.URL))
	if err != nil {
		t.Error(err)
	}

	_, err = client.Chart("aapl", "6y")

	if err == nil {
		t.Error("expected err but got nil")
	}

	chart, _ := client.Chart("aapl", "1d")

	if len(chart.Charts) == 0 {
		t.Error("charts shouldn't be empty")
	}

	if chart.Charts[0].Minute == "" {
		t.Error("minute should be non-empty string for 1d range")
	}
}

func TestRefDataSymbols(t *testing.T) {
	mockServer := mockiex.MockIEXServer()
	defer mockServer.Close()
	client, err := NewClient(SetBaseURL(mockServer.URL))
	if err != nil {
		t.Error(err)
	}

	rds, _ := client.RefDataSymbols()

	firstSymbol := rds.Symbols[0]

	if firstSymbol.Symbol != "A" {
		t.Errorf("expected A but got %v", firstSymbol.Symbol)
	}

	if firstSymbol.Date != "2018-10-26" {
		t.Errorf("expected 2018-10-26 but got %v", firstSymbol.Date)
	}

	if firstSymbol.Name != "Agilent Technologies Inc." {
		t.Errorf("expected Agilent Technologies Inc. but got %v", firstSymbol.Name)
	}

	if firstSymbol.IsEnabled != true {
		t.Errorf("expected true but got %v", firstSymbol.IsEnabled)
	}

	if firstSymbol.Type != "cs" {
		t.Errorf("expected cs but got %v", firstSymbol.Type)
	}

	if firstSymbol.IexId != 2 {
		t.Errorf("expected 2 but got %v", firstSymbol.IexId)
	}
}
