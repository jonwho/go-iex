package goiextest

import (
	"testing"
)

func TestAdvancedStats(t *testing.T) {
	advancedStat, err := mockClient.AdvancedStats("aapl")
	if err != nil {
		t.Error(err)
	}
	expected := "Apple Inc."
	actual := advancedStat.CompanyName
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestBalanceSheet(t *testing.T) {
	balanceSheet, err := mockClient.BalanceSheet("aapl", nil)
	if err != nil {
		t.Error(err)
	}
	expected = "AAPL"
	actual = balanceSheet.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = iexSandboxClient.BalanceSheet("aapl", nil, 5)
	if err != nil {
		t.Error(err)
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = iexSandboxClient.BalanceSheet("aapl", nil, 5, "annual")
	if err != nil {
		t.Error(err)
	}
	expected = 5
	actual = len(balanceSheet.BalanceSheet)
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	balanceSheet, err = iexSandboxClient.BalanceSheet("aapl", struct {
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

	balanceSheet, err = iexSandboxClient.BalanceSheet("aapl", struct {
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
	batch, err := iexSandboxClient.Batch("aapl", struct {
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

// // TODO: @mock
// func TestBook(t *testing.T) {
//   book, err := iexSandboxClient.Book("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = true
//   actual = len(book.Asks) != 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = true
//   actual = len(book.Bids) != 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = "AAPL"
//   actual = book.Quote.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = true
//   actual = len(book.Trades) != 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = struct{}{}
//   actual = book.SystemEvent
//   if expected == actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestCashFlow(t *testing.T) {
//   cashflow, err := iexSandboxClient.CashFlow("aapl", nil)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = cashflow.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = false
//   actual = len(cashflow.CashFlow) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestChart(t *testing.T) {
//   chart, err := iexSandboxClient.Chart("aapl", "outofrange", nil)
//   if err == nil {
//     t.Error("Expected err to not be nil")
//   }
//   expected = `Received invalid date range for chart`
//   actual = err.Error()
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   chart, err = iexSandboxClient.Chart("aapl", "max", nil)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(chart) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestCollection(t *testing.T) {
//   col, err := iexSandboxClient.Collection("sector", struct {
//     CollectionName string `url:"collectionName,omitempty"`
//   }{"Technology"})
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(col) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestCompany(t *testing.T) {
//   com, err := iexSandboxClient.Company("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = com.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestDelayedQuote(t *testing.T) {
//   dq, err := iexSandboxClient.DelayedQuote("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = dq.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestDividends(t *testing.T) {
//   div, err := iexSandboxClient.Dividends("aapl", "outofrange")
//   if err == nil {
//     t.Error("Expected err to not be nil")
//   }
//   expected = `Received invalid date range for dividend`
//   actual = err.Error()
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   div, err = iexSandboxClient.Dividends("aapl", "5y")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(div) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestEarnings(t *testing.T) {
//   er, err := iexSandboxClient.Earnings("aapl", nil)
//   if err != nil {
//     t.Error(err)
//   }
//
//   er, err = iexSandboxClient.Earnings("aapl", struct {
//     Last int `url:"last,omitempty"`
//   }{2})
//   if err != nil {
//     t.Error(err)
//   }
//   expected = 2
//   actual = len(er.Earnings)
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestEarningsToday(t *testing.T) {
//   et, err := iexSandboxClient.EarningsToday()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(et.BTO) != 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = false
//   actual = len(et.AMC) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestEffectiveSpread(t *testing.T) {
//   es, err := iexSandboxClient.EffectiveSpread("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(es) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestEstimates(t *testing.T) {
//   est, err := iexSandboxClient.Estimates("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(est.Estimates) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   est, err = iexSandboxClient.Estimates("aapl", 2)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(est.Estimates) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   est, err = iexSandboxClient.Estimates("aapl", 2, "annual")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(est.Estimates) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestFinancials(t *testing.T) {
//   fin, err := iexSandboxClient.Financials("aapl", nil)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fin.Financials) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   fin, err = iexSandboxClient.Financials("aapl", nil, 2)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fin.Financials) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   fin, err = iexSandboxClient.Financials("aapl", nil, 2, "annual")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fin.Financials) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   fin, err = iexSandboxClient.Financials("aapl", struct {
//     Period string `url:"period,omitempty"`
//   }{"quarterly"}, 2, "annual")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fin.Financials) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestFundOwnership(t *testing.T) {
//   fo, err := iexSandboxClient.FundOwnership("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(fo) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestIncomeStatement(t *testing.T) {
//   stmt, err := iexSandboxClient.IncomeStatement("aapl", nil)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(stmt.Income) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestInsiderRoster(t *testing.T) {
//   ir, err := iexSandboxClient.InsiderRoster("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(ir) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestInsiderSummary(t *testing.T) {
//   is, err := iexSandboxClient.InsiderSummary("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(is) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestInsiderTransactions(t *testing.T) {
//   it, err := iexSandboxClient.InsiderTransactions("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(it) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestInsitutionalOwnership(t *testing.T) {
//   iop, err := iexSandboxClient.InstitutionalOwnership("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(iop) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestIntradayPrices(t *testing.T) {
//   ip, err := iexSandboxClient.IntradayPrices("aapl", struct {
//     chartIEXOnly    bool
//     chartReset      bool
//     chartSimplify   bool
//     chartInterval   int
//     changeFromClose bool
//     chartLast       int
//   }{true, true, true, 5, true, 10})
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(ip) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestUpcomingIPOS(t *testing.T) {
//   ipo, err := iexSandboxClient.UpcomingIPOS()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(ipo.RawData) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = false
//   actual = len(ipo.ViewData) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestTodayIPOS(t *testing.T) {
//   ipo, err := iexSandboxClient.TodayIPOS()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(ipo.RawData) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//   expected = false
//   actual = len(ipo.ViewData) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestKeyStats(t *testing.T) {
//   ks, err := iexSandboxClient.KeyStats("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "Apple, Inc."
//   actual = ks.CompanyName
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestLargestTrades(t *testing.T) {
//   lt, err := iexSandboxClient.LargestTrades("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(lt) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestList(t *testing.T) {
//   list, err := iexSandboxClient.List("gainers", struct {
//     displayPercent bool
//   }{true})
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(list) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestLogo(t *testing.T) {
//   logo, err := iexSandboxClient.Logo("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   actual = logo.URL
//   if actual == "" {
//     t.Errorf("Expected non-empty string")
//   }
// }
//
// func TestMarketVolume(t *testing.T) {
//   mkt, err := iexSandboxClient.MarketVolume()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(mkt) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestNews(t *testing.T) {
//   news, err := iexSandboxClient.News("aapl", 10)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(news) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestOHLC(t *testing.T) {
//   ohlc, err := iexSandboxClient.OHLC("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   if ohlc.Open.Price < 0 {
//     t.Errorf("Expected open price greater than 0")
//   }
// }
//
// // TODO: @mock
// func TestOptionDates(t *testing.T) {
//   dates, err := iexSandboxClient.OptionDates("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(dates) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestOptions(t *testing.T) {
//   options, err := iexSandboxClient.Options("aapl", "201912")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(options) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   options, err = iexSandboxClient.Options("aapl", "201912", "call")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(options) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestPeers(t *testing.T) {
//   peers, err := iexSandboxClient.Peers("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(peers) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestPreviousDayPrice(t *testing.T) {
//   prev, err := iexSandboxClient.PreviousDayPrice("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = prev.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestPrice(t *testing.T) {
//   price, err := iexSandboxClient.Price("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   if price < 0 {
//     t.Errorf("Expected price greater than 0")
//   }
// }
//
// func TestPriceTarget(t *testing.T) {
//   tgt, err := iexSandboxClient.PriceTarget("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = tgt.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestQuote(t *testing.T) {
//   quote, err := iexSandboxClient.Quote("aapl", nil)
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = quote.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   quote, err = iexSandboxClient.Quote("aapl", struct {
//     DisplayPercent bool `url:"displayPercent,omitempty"`
//   }{true})
//   if err != nil {
//     t.Error(err)
//   }
//   expected = "AAPL"
//   actual = quote.Symbol
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestRecommendationTrends(t *testing.T) {
//   rt, err := iexSandboxClient.RecommendationTrends("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(rt) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// func TestSectorPerformance(t *testing.T) {
//   sp, err := iexSandboxClient.SectorPerformance()
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(sp) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestSplits(t *testing.T) {
//   sp, err := iexSandboxClient.Splits("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(sp) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   sp, err = iexSandboxClient.Splits("aapl", "5y")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(sp) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
//
//   sp, err = iexSandboxClient.Splits("aapl", "outofrange")
//   if err == nil {
//     t.Error("Expected err to be not nil")
//   }
//   expected = `Received invalid date range for splits`
//   actual = err.Error()
//   if expected != actual {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
//
// // TODO: @mock
// func TestUpcoming(t *testing.T) {
//   if _, err := iexSandboxClient.Upcoming("aapl", "events", nil); err != nil {
//     t.Error(err)
//   }
//   if _, err := iexSandboxClient.Upcoming("aapl", "earnings", nil); err != nil {
//     t.Error(err)
//   }
//   if _, err := iexSandboxClient.Upcoming("aapl", "dividends", nil); err != nil {
//     t.Error(err)
//   }
//   if _, err := iexSandboxClient.Upcoming("aapl", "splits", nil); err != nil {
//     t.Error(err)
//   }
//   if _, err := iexSandboxClient.Upcoming("aapl", "ipos", nil); err != nil {
//     t.Error(err)
//   }
// }
//
// func TestVolumeByVenue(t *testing.T) {
//   vbv, err := iexSandboxClient.VolumeByVenue("aapl")
//   if err != nil {
//     t.Error(err)
//   }
//   expected = false
//   actual = len(vbv) == 0
//   if actual.(bool) {
//     t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
//   }
// }
