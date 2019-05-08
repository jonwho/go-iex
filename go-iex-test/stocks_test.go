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
//     t.Error(err)
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
		t.Error(err)
	}
	expected = "AAPL"
	actual = balanceSheet.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", nil, 5)
	if err != nil {
		t.Error(err)
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = stock.BalanceSheet("aapl", nil, 5, "annual")
	if err != nil {
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
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

// TODO: @mock
func TestBook(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	book, err := stock.Book("aapl")
	if err != nil {
		t.Error(err)
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
		t.Error(err)
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
	if err == nil {
		t.Error("Expected err to not be nil")
	}
	expected = `Received invalid date range for chart`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	chart, err = stock.Chart("aapl", "max", nil)
	if err != nil {
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
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
	if err == nil {
		t.Error("Expected err to not be nil")
	}
	expected = `Received invalid date range for dividend`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	div, err = stock.Dividends("aapl", "5y")
	if err != nil {
		t.Error(err)
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
		t.Error(err)
	}

	er, err = stock.Earnings("aapl", struct {
		Last int `url:"last,omitempty"`
	}{2})
	if err != nil {
		t.Error(err)
	}
	expected = 2
	actual = len(er.Earnings)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestEarningsToday(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	et, err := stock.EarningsToday()
	if err != nil {
		t.Error(err)
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
		t.Error(err)
	}
	expected = false
	actual = len(es) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestEstimates(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	est, err := stock.Estimates("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(est.Estimates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	est, err = stock.Estimates("aapl", 2)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(est.Estimates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	est, err = stock.Estimates("aapl", 2, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(est.Estimates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestFinancials(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	fin, err := stock.Financials("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = stock.Financials("aapl", nil, 2)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = stock.Financials("aapl", nil, 2, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = stock.Financials("aapl", struct {
		Period string `url:"period,omitempty"`
	}{"quarterly"}, 2, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestFundOwnership(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	fo, err := stock.FundOwnership("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fo) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestIncomeStatement(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	stmt, err := stock.IncomeStatement("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(stmt.Income) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestInsiderRoster(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ir, err := stock.InsiderRoster("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ir) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestInsiderSummary(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	is, err := stock.InsiderSummary("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(is) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestInsiderTransactions(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	it, err := stock.InsiderTransactions("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(it) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestInsitutionalOwnership(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	iop, err := stock.InstitutionalOwnership("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(iop) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestIntradayPrices(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ip, err := stock.IntradayPrices("aapl", struct {
		chartIEXOnly    bool
		chartReset      bool
		chartSimplify   bool
		chartInterval   int
		changeFromClose bool
		chartLast       int
	}{true, true, true, 5, true, 10})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ip) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUpcomingIPOS(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ipo, err := stock.UpcomingIPOS()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ipo.RawData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(ipo.ViewData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestTodayIPOS(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ipo, err := stock.TodayIPOS()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ipo.RawData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(ipo.ViewData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestKeyStats(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ks, err := stock.KeyStats("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = "Apple, Inc."
	actual = ks.CompanyName
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestLargestTrades(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	lt, err := stock.LargestTrades("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(lt) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestList(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	list, err := stock.List("gainers", struct {
		displayPercent bool
	}{true})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(list) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestLogo(t *testing.T) {
	var actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	logo, err := stock.Logo("aapl")
	if err != nil {
		t.Error(err)
	}
	actual = logo.URL
	if actual == "" {
		t.Errorf("Expected non-empty string")
	}
}

func TestMarketVolume(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	mkt, err := stock.MarketVolume()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(mkt) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestNews(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	news, err := stock.News("aapl", 10)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(news) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestOHLC(t *testing.T) {
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	ohlc, err := stock.OHLC("aapl")
	if err != nil {
		t.Error(err)
	}
	if ohlc.Open.Price < 0 {
		t.Errorf("Expected open price greater than 0")
	}
}

// TODO: @mock
func TestOptionDates(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	dates, err := stock.OptionDates("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestOptions(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	options, err := stock.Options("aapl", "201912")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(options) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	options, err = stock.Options("aapl", "201912", "call")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(options) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestPeers(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	peers, err := stock.Peers("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(peers) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestPreviousDayPrice(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	prev, err := stock.PreviousDayPrice("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = "AAPL"
	actual = prev.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestPrice(t *testing.T) {
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	price, err := stock.Price("aapl")
	if err != nil {
		t.Error(err)
	}
	if price < 0 {
		t.Errorf("Expected price greater than 0")
	}
}

func TestPriceTarget(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	tgt, err := stock.PriceTarget("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = "AAPL"
	actual = tgt.Symbol
	if expected != actual {
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
		t.Error(err)
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
		t.Error(err)
	}
	expected = "AAPL"
	actual = quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestRecommendationTrends(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	rt, err := stock.RecommendationTrends("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(rt) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSectorPerformance(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	sp, err := stock.SectorPerformance()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestSplits(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	sp, err := stock.Splits("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	sp, err = stock.Splits("aapl", "5y")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	sp, err = stock.Splits("aapl", "outofrange")
	if err == nil {
		t.Error("Expected err to be not nil")
	}
	expected = `Received invalid date range for splits`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: @mock
func TestUpcoming(t *testing.T) {
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	if _, err := stock.Upcoming("aapl", "events", nil); err != nil {
		t.Error(err)
	}
	if _, err := stock.Upcoming("aapl", "earnings", nil); err != nil {
		t.Error(err)
	}
	if _, err := stock.Upcoming("aapl", "dividends", nil); err != nil {
		t.Error(err)
	}
	if _, err := stock.Upcoming("aapl", "splits", nil); err != nil {
		t.Error(err)
	}
	if _, err := stock.Upcoming("aapl", "ipos", nil); err != nil {
		t.Error(err)
	}
}

func TestVolumeByVenue(t *testing.T) {
	var expected, actual interface{}
	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	u, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, u, iex.DefaultHTTPClient)

	vbv, err := stock.VolumeByVenue("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(vbv) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
