package goiex

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewStock(t *testing.T) {
	stk := NewStock("test_token", "stable", sandboxURL, nil)

	expected = "stock/"
	actual = stk.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = true
	actual = stk.Client() == nil
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = stk.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "https://sandbox.iexapis.com/"
	actual = stk.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "stable"
	actual = stk.Version()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestAdvancedStats(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/advanced_stats")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	advancedStat, err := cli.AdvancedStats("aapl")
	if err != nil {
		t.Error(err)
	}
	expected := "Apple, Inc."
	actual := advancedStat.CompanyName
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestBalanceSheet(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/balance_sheet")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	balanceSheet, err := cli.BalanceSheet("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = "AAPL"
	actual = balanceSheet.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = cli.BalanceSheet("aapl", nil, 5)
	if err != nil {
		t.Error(err)
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = cli.BalanceSheet("aapl", nil, 5, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = cli.BalanceSheet("aapl", struct {
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

	balanceSheet, err = cli.BalanceSheet("aapl", struct {
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
	rec, err := recorder.New("cassettes/stock/batch")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	batch, err := cli.Batch("aapl", struct {
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

func TestBook(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/book")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	book, err := cli.Book("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(book.Asks) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(book.Bids) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AAPL"
	actual = book.Quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(book.Trades) == 0
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
	rec, err := recorder.New("cassettes/stock/cash_flow")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	cashflow, err := cli.CashFlow("aapl", nil)
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
	rec, err := recorder.New("cassettes/stock/chart")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	chart, err := cli.Chart("aapl", "outofrange", nil)
	if err == nil {
		t.Error("Expected err to not be nil")
	}
	expected = `Received invalid date range for chart`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	chart, err = cli.Chart("aapl", "max", nil)
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
	rec, err := recorder.New("cassettes/stock/collection")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	col, err := cli.Collection("sector", struct {
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
	rec, err := recorder.New("cassettes/stock/company")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	com, err := cli.Company("aapl")
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
	dq, err := mockClient.DelayedQuote("aapl")
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
	rec, err := recorder.New("cassettes/stock/dividends")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	div, err := cli.Dividends("aapl", "outofrange")
	if err == nil {
		t.Error("Expected err to not be nil")
	}
	expected = `Received invalid date range for dividend`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	div, err = cli.Dividends("aapl", "5y")
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
	rec, err := recorder.New("cassettes/stock/earnings")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	er, err := cli.Earnings("aapl", nil)
	if err != nil {
		t.Error(err)
	}

	er, err = cli.Earnings("aapl", struct {
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

func TestEarningsToday(t *testing.T) {
	et, err := mockClient.EarningsToday()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(et.BTO) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(et.AMC) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestEstimates(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/estimates")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	est, err := cli.Estimates("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(est.Estimates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	est, err = cli.Estimates("aapl", 2)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(est.Estimates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	est, err = cli.Estimates("aapl", 2, "annual")
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
	rec, err := recorder.New("cassettes/stock/financials")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	fin, err := cli.Financials("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = cli.Financials("aapl", nil, 2)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = cli.Financials("aapl", nil, 2, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(fin.Financials) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	fin, err = cli.Financials("aapl", struct {
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

func TestFundOwnership(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/fund_ownership")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	fo, err := cli.FundOwnership("aapl")
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
	rec, err := recorder.New("cassettes/stock/income_statement")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	stmt, err := cli.IncomeStatement("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(stmt.Income) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	stmt, err = cli.IncomeStatement("aapl", struct {
		Last   int    `url:"last"`
		Period string `url:"period"`
	}{Last: 2, Period: "annual"})
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(stmt.Income) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestInsiderRoster(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/insider_roster")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ir, err := cli.InsiderRoster("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ir) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestInsiderSummary(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/insider_summary")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	is, err := cli.InsiderSummary("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(is) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestInsiderTransactions(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/insider_transactions")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	it, err := cli.InsiderTransactions("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(it) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestInsitutionalOwnership(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/institutional_ownership")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	iop, err := cli.InstitutionalOwnership("aapl")
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
	rec, err := recorder.New("cassettes/stock/intraday_prices")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ip, err := cli.IntradayPrices("aapl", struct {
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

func TestTodayIPOS(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/today_ipos")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ipo, err := cli.TodayIPOS()
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
	rec, err := recorder.New("cassettes/stock/key_stats")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ks, err := cli.KeyStats("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = "Apple, Inc."
	actual = ks.CompanyName
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

// TODO: re-record this request during live trading hours for actual data
func TestLargestTrades(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/largest_trades")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	lt, err := cli.LargestTrades("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(lt) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestList(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/list")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	list, err := cli.List("gainers", struct {
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
	rec, err := recorder.New("cassettes/stock/logo")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	logo, err := cli.Logo("aapl")
	if err != nil {
		t.Error(err)
	}
	actual = logo.URL
	if actual == "" {
		t.Errorf("Expected non-empty string")
	}
}

func TestMarketVolume(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/market_volume")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	mkt, err := cli.MarketVolume()
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
	rec, err := recorder.New("cassettes/stock/news")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	news, err := cli.News("aapl", 10)
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
	rec, err := recorder.New("cassettes/stock/ohlc")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ohlc, err := cli.OHLC("aapl")
	if err != nil {
		t.Error(err)
	}
	if ohlc.Open.Price < 0 {
		t.Errorf("Expected open price greater than 0")
	}
}

func TestOpenClosePrice(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/open_close_price")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ohlc, err := cli.OpenClosePrice("aapl")
	if err != nil {
		t.Error(err)
	}
	if ohlc.Open.Price < 0 {
		t.Errorf("Expected open price greater than 0")
	}
}

func TestOptionDates(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/option_dates")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	dates, err := cli.OptionDates("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(dates) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestOptions(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/options")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	options, err := cli.Options("aapl", "201912")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(options) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	options, err = cli.Options("aapl", "201912", "call")
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
	rec, err := recorder.New("cassettes/stock/peers")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	peers, err := cli.Peers("aapl")
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
	rec, err := recorder.New("cassettes/stock/previous_day_price")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	prev, err := cli.PreviousDayPrice("aapl")
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
	rec, err := recorder.New("cassettes/stock/price")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	price, err := cli.Price("aapl")
	if err != nil {
		t.Error(err)
	}
	if price < 0 {
		t.Errorf("Expected price greater than 0")
	}
}

func TestPriceTarget(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/price_target")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	tgt, err := cli.PriceTarget("aapl")
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
	rec, err := recorder.New("cassettes/stock/quote")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	quote, err := cli.Quote("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = "AAPL"
	actual = quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	quote, err = cli.Quote("aapl", struct {
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
	rec, err := recorder.New("cassettes/stock/recommendation_trends")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	rt, err := cli.RecommendationTrends("aapl")
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
	rec, err := recorder.New("cassettes/stock/sector_performance")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	sp, err := cli.SectorPerformance()
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestSplits(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/splits")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	sp, err := cli.Splits("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	sp, err = cli.Splits("aapl", "5y")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(sp) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	sp, err = cli.Splits("aapl", "outofrange")
	if err == nil {
		t.Error("Expected err to be not nil")
	}
	expected = `Received invalid date range for splits`
	actual = err.Error()
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUpcomingDividends(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/upcoming_dividends")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	div, err := cli.UpcomingDividends("market", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(div) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	div, err = cli.UpcomingDividends("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(div) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUpcomingEarnings(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/upcoming_earnings")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ue, err := cli.UpcomingEarnings("market", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ue) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	ue, err = cli.UpcomingEarnings("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(ue) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUpcomingEvents(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/upcoming_events")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	evts, err := cli.UpcomingEvents("market", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(evts.IPOS.RawData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(evts.IPOS.ViewData) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	evts, err = cli.UpcomingEvents("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(evts.IPOS.RawData) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(evts.IPOS.ViewData) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestUpcomingIPOS(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/upcoming_ipos")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	ipo, err := cli.UpcomingIPOS("market", nil)
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

	ipo, err = cli.UpcomingIPOS("aapl", nil)
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

func TestUpcomingSplits(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/upcoming_splits")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	spls, err := cli.UpcomingSplits("market", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(spls) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	spls, err = cli.UpcomingSplits("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(spls) != 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestVolumeByVenue(t *testing.T) {
	rec, err := recorder.New("cassettes/stock/volume_by_venue")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewStock(testToken, DefaultVersion, sandboxURL, httpClient)

	vbv, err := cli.VolumeByVenue("aapl")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(vbv) == 0
	if actual.(bool) {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
