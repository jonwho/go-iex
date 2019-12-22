# go-iex

[![GoDoc](https://godoc.org/github.com/jonwho/go-iex?status.svg)](http://godoc.org/github.com/jonwho/go-iex)
[![releases](https://img.shields.io/github/release/jonwho/go-iex.svg)](https://github.com/jonwho/go-iex/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonwho/go-iex)](https://goreportcard.com/report/github.com/jonwho/go-iex)
![](https://github.com/jonwho/go-iex/workflows/tests/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>

## ATTRIBUTION
[Data provided by IEX](https://iexcloud.io)

## DESCRIPTION
Client interface to IEX trading API.

## ENV
### TEST
* Grab your test/real tokens from [https://iexcloud.io/console/](https://iexcloud.io/console/)
* Set ENV VAR
```sh
export IEX_TEST_SECRET_TOKEN=Tsk_ahsvyao12u4u0ausvn1o3rhw988120yf_FAKE
export IEX_TEST_PUBLISHABLE_TOKEN=Tpk_la091720ihakbso128uihotbfao_FAKE
```
### PRODUCTION
```sh
export IEX_SECRET_TOKEN=Tsk_ahsvyao12u4u0ausvn1o3rhw988120yf_REAL
export IEX_PUBLISHABLE_TOKEN=Tpk_la091720ihakbso128uihotbfao_REAL
```

## DEV NOTES
* Use this [online json to struct converter](https://mholt.github.io/json-to-go/) to save time
* Use `make test` to run test suite
* Use `make coverage` to update `README.md` coverage badge

## USAGE
```go
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	iex "github.com/jonwho/go-iex/v2"
)

func main() {
	token := os.Getenv("IEX_SECRET_TOKEN")
	// client will have all currently supported IEX APIs
	client, err := iex.NewClient(token)
	if err != nil {
		log.Fatalln(err)
	}

	quote, err := client.Quote("aapl", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)

	// if you only want to test against sandbox build a custom client
	token = os.Getenv("IEX_TEST_SECRET_TOKEN")
	baseURL, _ := url.Parse(iex.SandboxBaseURL)
	// get Stocks only API client for sandbox testing
	stock := iex.NewStock(token, iex.DefaultVersion, baseURL, iex.DefaultHTTPClient)

	quote, err = stock.Quote("aapl", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)

	// you can also use the Get helper on client to unmarshal to your own custom struct
	anonstruct := &struct {
		Symbol string `json:"symbol,omitempty"`
	}{}
	err = client.Get("stock/aapl/quote", anonstruct, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Symbol", anonstruct.Symbol)
}
```

## SUPPORTED ENDPOINTS
### Introduction
- [x] Batch Requests

### Account
- [x] Metadata
- [x] Usage
- [ ] Pay as you go
- [ ] Message Budget
- [ ] Signed Requests
- [ ] Setting up signed token
- [ ] Getting the secret for a signed token

### API System Metadata
- [x] Status

### Data APIS
- [x] Data Points
- [ ] Data Tables
- [ ] Time Series

### Stock Prices
- [x] Book
- [x] Charts
- [x] Delayed Quote
- [x] Extended Hours Quote (included as part of Quote response)
- [x] Historical Prices
- [x] Intraday Prices
- [x] Largest Trades
- [x] Open / Close Price
- [x] OHLC
- [x] Previous Day Price
- [x] Price Only
- [x] Quote
- [ ] Real-time Quote (included as part of Quote response)
- [x] Volume by Venue

### Stock Profiles
- [x] Company
- [x] Insider Roster
- [x] Insider Summary
- [x] Insider Transactions
- [x] Logo
- [x] Peer Groups

### Stock Fundamentals
- [x] Balance Sheet
- [x] Cash Flow
- [x] Dividends (Basic)
- [x] Earnings
- [x] Financials
- [ ] Financials As Reported
- [x] Income Statement
- [ ] SEC Filings
- [x] Splits (Basic)

### Stock Research
- [x] Advanced Stats
- [ ] Analyst Recommendations (see RecommendationTrends)
- [x] Estimates
- [x] Fund Ownership
- [x] Institutional Ownership
- [x] Key Stats
- [x] Price Target
- [ ] Technical Indicators

### Corporate Actions
- [ ] Bonus Issue
- [ ] Distribution
- [ ] Dividends
- [ ] Return of Capital
- [ ] Rights Issue
- [ ] Right to Purchase
- [ ] Security Reclassification
- [ ] Security Swap
- [ ] Spinoff
- [ ] Splits

### Market Info
- [x] Collections
- [x] Earnings Today
- [x] IPO Calendar
- [x] List
- [x] Market Volume (U.S.)
- [x] Sector Performance
- [x] Upcoming Events

### News
- [x] News
- [ ] Streaming News
- [ ] Historical News

### Cryptocurrency
- [x] Cryptocurrency Book
- [ ] Cryptocurrency Event
- [x] Cryptocurrency Price
- [x] Cryptocurrency Quote

### Forex / Currencies
- [ ] Real-time Streaming
- [x] Latest Currency Rates
- [x] Currency Conversion
- [x] Historical Daily

### Options
- [x] End of Day Options

### Social Sentiment
- [ ] Social Sentiment

### CEO Compensation
- [ ] CEO Compensation

### Treasures
- [ ] Daily Treasury Rates

### Commodities
- [x] Oil Prices
- [x] Natural Gas Price
- [x] Heating Oil Prices
- [x] Jet Fuel Prices
- [x] Diesel Prices
- [x] Diesel Price
- [x] Gas Prices
- [x] Propane Prices

### Economic Data
- [x] CD Rates
- [x] Consumer Price Index
- [x] Credit Card Interest Rate
- [x] Federal Fund Rates
- [x] Real GDP
- [ ] Institutional Money Funds
- [ ] Initial Claims
- [x] Industrial Production Index
- [x] Mortgage Rates
- [x] Total Housing Starts
- [x] Total Payrolls
- [x] Total Vehicle Sales
- [ ] Retail Money Funds
- [x] Unemployment Rate
- [x] US Recession Probabilities

### Reference Data
- [ ] Search
- [ ] Cryptocurrency Symbols
- [x] FX Symbols
- [x] IEX Symbols
- [x] International Symbols
- [x] International Exchanges
- [ ] ISIN Mapping
- [x] Mutual Fund Symbols
- [ ] Options Symbols
- [x] OTC Symbols
- [x] Sectors
- [x] Symbols
- [x] Tags
- [x] U.S. Exchanges
- [x] U.S. Holidays and Trading Dates

### Investors Exchange Data
- [x] DEEP
- [ ] DEEP Auction
- [ ] DEEP Book
- [ ] DEEP Operational Halt Status
- [ ] DEEP Official Price
- [ ] DEEP Security Event
- [ ] DEEP Short Sale Price Test Status
- [ ] DEEP System Event
- [ ] DEEP Trades
- [ ] DEEP Trade Break
- [ ] DEEP Trading Status
- [x] Last
- [ ] Listed Regulation SHO Threshold Securities List
- [ ] Listed Short Interest List In Dev
- [ ] Stats Historical Daily In Dev
- [ ] Stats Historical Summary
- [ ] Stats Intraday
- [ ] Stats Recent
- [ ] Stats Records
- [x] TOPS
