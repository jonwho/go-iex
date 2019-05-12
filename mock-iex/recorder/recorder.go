package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/jonwho/go-iex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(6)

	go recordAccount()
	go recordAlternativeData()
	go recordAPISystemMetadata()
	go recordDataAPIS()
	go recordForex()
	go recordStock()

	wg.Wait()
}

func recordAccount() {
	defer wg.Done()

	os.Remove("account.yaml")
	// Start our recorder
	r, err := recorder.New("account")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token, goiex.SetHTTPClient(httpClient))
	iex.Metadata()
}

func recordAlternativeData() {
	defer wg.Done()

	os.Remove("alternative_data.yaml")
	// Start our recorder
	r, err := recorder.New("alternative_data")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token,
		goiex.SetURL(goiex.SandboxBaseURL),
		goiex.SetHTTPClient(httpClient),
	)
	iex.Crypto("btcusdt")
	iex.SocialSentiment("aapl")
	iex.CEOCompensation("aapl")
}

func recordAPISystemMetadata() {
	defer wg.Done()

	os.Remove("api_system_metadata.yaml")
	// Start our recorder
	r, err := recorder.New("api_system_metadata")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token,
		goiex.SetURL(goiex.SandboxBaseURL),
		goiex.SetHTTPClient(httpClient),
	)
	iex.Status()
}

func recordForex() {
	defer wg.Done()

	os.Remove("forex.yaml")
	// Start our recorder
	r, err := recorder.New("forex")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token,
		goiex.SetURL(goiex.SandboxBaseURL),
		goiex.SetHTTPClient(httpClient),
	)
	iex.ExchangeRates("eur", "usd")
}

func recordStock() {
	defer wg.Done()

	os.Remove("stocks.yaml")
	// Start our recorder
	r, err := recorder.New("stocks")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token,
		goiex.SetURL(goiex.SandboxBaseURL),
		goiex.SetHTTPClient(httpClient),
	)
	if err != nil {
		log.Fatal(err)
	}
	iex.BalanceSheet("aapl", nil)
	iex.BalanceSheet("aapl", nil, 5)
	iex.BalanceSheet("aapl", nil, 5, "annual")
	iex.BalanceSheet("aapl", struct {
		Period string `url:"period,omitempty"`
	}{"annual"}, 5, "annual")
	iex.BalanceSheet("aapl", struct {
		Period string `url:"period,omitempty"`
		Last   int    `url:"last,omitempty"`
	}{"quarter", 12})
	iex.Batch("aapl", struct {
		Types string `url:"types,omitempty"`
		Range string `url:"range,omitempty"`
		Last  int    `url:"last,omitempty"`
	}{"quote,news,chart", "1m", 1})
	iex.Book("aapl")
	iex.CashFlow("aapl", nil)
	iex.Chart("aapl", "outofrange", nil)
	iex.Chart("aapl", "max", nil)
	iex.Collection("sector", struct {
		CollectionName string `url:"collectionName,omitempty"`
	}{"Technology"})
	iex.Company("aapl")
	iex.DelayedQuote("aapl")
	iex.Dividends("aapl", "outofrange")
	iex.Dividends("aapl", "5y")
	iex.Earnings("aapl", nil)
	iex.Earnings("aapl", struct {
		Last int `url:"last,omitempty"`
	}{2})
	iex.EarningsToday()
	iex.EffectiveSpread("aapl")
	iex.Estimates("aapl")
	iex.Estimates("aapl", 2)
	iex.Estimates("aapl", 2, "annual")
	iex.Financials("aapl", nil)
	iex.Financials("aapl", nil, 2)
	iex.Financials("aapl", nil, 2, "annual")
	iex.Financials("aapl", struct {
		Period string `url:"period,omitempty"`
	}{"quarterly"}, 2, "annual")
	iex.FundOwnership("aapl")
	iex.IncomeStatement("aapl", nil)
	iex.InsiderRoster("aapl")
	iex.InsiderSummary("aapl")
	iex.InsiderTransactions("aapl")
	iex.InstitutionalOwnership("aapl")
	iex.IntradayPrices("aapl", struct {
		chartIEXOnly    bool
		chartReset      bool
		chartSimplify   bool
		chartInterval   int
		changeFromClose bool
		chartLast       int
	}{true, true, true, 5, true, 10})
	iex.UpcomingIPOS()
	iex.TodayIPOS()
	iex.KeyStats("aapl")
	iex.LargestTrades("aapl")
	iex.List("gainers", struct {
		displayPercent bool
	}{true})
	iex.Logo("aapl")
	iex.MarketVolume()
	iex.News("aapl", 10)
	iex.OHLC("aapl")
	iex.OptionDates("aapl")
	iex.Options("aapl", "201912")
	iex.Options("aapl", "201912", "call")
	iex.Peers("aapl")
	iex.PreviousDayPrice("aapl")
	iex.Price("aapl")
	iex.PriceTarget("aapl")
	iex.Quote("aapl", nil)
	iex.Quote("aapl", struct {
		DisplayPercent bool `url:"displayPercent,omitempty"`
	}{true})
	iex.RecommendationTrends("aapl")
	iex.SectorPerformance()
	iex.Splits("aapl")
	iex.Splits("aapl", "5y")
	iex.Splits("aapl", "outofrange")
	iex.Upcoming("aapl", "events", nil)
	iex.Upcoming("aapl", "earnings", nil)
	iex.Upcoming("aapl", "dividends", nil)
	iex.Upcoming("aapl", "splits", nil)
	iex.Upcoming("aapl", "ipos", nil)
	iex.VolumeByVenue("aapl")
}

func recordDataAPIS() {
	defer wg.Done()

	os.Remove("data_apis.yaml")
	// Start our recorder
	r, err := recorder.New("data_apis")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(removeToken)
	defer r.Stop() // Make sure recorder is stopped once done with it

	token := os.Getenv("IEX_TEST_SECRET_TOKEN")
	httpClient := &http.Client{Transport: r}
	iex, err := goiex.NewClient(token,
		goiex.SetURL(goiex.SandboxBaseURL),
		goiex.SetHTTPClient(httpClient),
	)
	if err != nil {
		log.Fatal(err)
	}
	iex.DataPoints("aapl")
}

func removeToken(i *cassette.Interaction) error {
	u, err := url.Parse(i.Request.URL)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Del("token")
	u.RawQuery = q.Encode()
	i.Request.URL = u.String()
	return nil
}
