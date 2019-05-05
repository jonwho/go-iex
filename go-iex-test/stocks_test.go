package goiextest

import (
	"net/url"
	"os"
	"testing"

	iex "github.com/jonwho/go-iex"
)

// TODO: mock this because this API only available higher account tier
// func TestAdvancedStats(t *testing.T) {
//   token := os.Getenv("IEX_TEST_SECRET_TOKEN")
//   u, _ := url.Parse(iex.SandboxBaseURL)
//   stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)
//
//   advancedStat, err := stock.AdvancedStats("aapl")
//   if err != nil {
//     t.Error(err.Error())
//   }
//   expected := "AAPL"
//   actual := advancedStat.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }

func TestBalanceSheet(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	balanceSheet, err := stock.BalanceSheet("aapl", nil)
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = balanceSheet.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", nil, 5)
	if err != nil {
		t.Error(err.Error())
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", nil, 5, "annual")
	if err != nil {
		t.Error(err.Error())
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", struct {
		Period string `url:"period,omitempty"`
	}{"annual"}, 5, "annual")

	if err != nil {
		t.Error(err.Error())
	}
	expected = 4
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", struct {
		Period string `url:"period,omitempty"`
		Last   int    `url:"last,omitempty"`
	}{"quarter", 12})
	if err != nil {
		t.Error(err.Error())
	}
	expected = 12
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestBatch(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	batch, err := stock.Batch("aapl", struct {
		Types string `url:"types,omitempty"`
		Range string `url:"range,omitempty"`
		Last  int    `url:"last,omitempty"`
	}{"quote,news,chart", "1m", 1})
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = batch.Quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(batch.News) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(batch.Chart) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestBook(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	book, err := stock.Book("aapl")
	if err != nil {
		t.Error(err.Error())
	}
	expected = true
	actual = len(book.Asks) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = true
	actual = len(book.Bids) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AAPL"
	actual = book.Quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = true
	actual = len(book.Trades) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = struct{}{}
	actual = book.SystemEvent
	if expected == actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCashFlow(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	cashflow, err := stock.CashFlow("aapl", nil)
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = cashflow.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(cashflow.CashFlow) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestChart(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	chart, err := stock.Chart("aapl", "outofrange", nil)
	expected = `Received invalid date range for chart`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	chart, err = stock.Chart("aapl", "max", nil)
	if err != nil {
		t.Error(err.Error())
	}
	expected = false
	actual = len(chart) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCollection(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	col, err := stock.Collection("sector", struct {
		CollectionName string `url:"collectionName,omitempty"`
	}{"Technology"})
	if err != nil {
		t.Error(err.Error())
	}
	expected = false
	actual = len(col) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCompany(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	com, err := stock.Company("aapl")
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = com.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestDelayedQuote(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	dq, err := stock.DelayedQuote("aapl")
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = dq.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestDividends(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	div, err := stock.Dividends("aapl", "outofrange")
	expected = `Received invalid date range for dividend`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	div, err = stock.Dividends("aapl", "5y")
	if err != nil {
		t.Error(err.Error())
	}
	expected = false
	actual = len(div) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestEarnings(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	er, err := stock.Earnings("aapl", nil)
	if err != nil {
		t.Error(err.Error())
	}

	er, err = stock.Earnings("aapl", struct {
		Last int `url:"last,omitempty"`
	}{2})
	if err != nil {
		t.Error(err.Error())
	}
	expected = 2
	actual = len(er.Earnings)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: mock this request to get consistent results
func TestEarningsToday(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	et, err := stock.EarningsToday()
	if err != nil {
		t.Error(err.Error())
	}
	expected = false
	actual = len(et.BTO) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(et.AMC) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: mock this request to get consistent results
func TestEffectiveSpread(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	es, err := stock.EffectiveSpread("aapl")
	if err != nil {
		t.Error(err.Error())
	}
	expected = false
	actual = len(es) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestQuote(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	quote, err := stock.Quote("aapl", nil)
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	quote, err = stock.Quote("aapl", struct {
		DisplayPercent bool `url:"displayPercent,omitempty"`
	}{true})
	if err != nil {
		t.Error(err.Error())
	}
	expected = "AAPL"
	actual = quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
