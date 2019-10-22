# go-iex

[![GoDoc](https://godoc.org/github.com/jonwho/go-iex?status.svg)](http://godoc.org/github.com/jonwho/go-iex)
[![releases](https://img.shields.io/github/release/jonwho/go-iex.svg)](https://github.com/jonwho/go-iex/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/jonwho/go-iex)](https://goreportcard.com/report/github.com/jonwho/go-iex)
![](https://github.com/jonwho/go-iex/workflows/tests/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-70%25-brightgreen.svg?longCache=true&style=flat)</a>

## ATTRIBUTION
[Data provided by IEX](https://iexcloud.io)

## DESCRIPTION
Client interface to IEX trading API.

## ENV
### TESTING
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

## USAGE
```go
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	iex "github.com/jonwho/go-iex"
)

func main() {
	token := os.Getenv("IEX_SECRET_TOKEN")
	// client will have all currently supported IEX APIs
	client, err := iex.NewClient(token)
	if err != nil {
		log.Println(err)
	}

	quote, err := client.Quote("aapl", nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)

	// if you only want to test against sandbox build a custom client
	// get Stocks only API client for sandbox testing
	token = os.Getenv("IEX_TEST_SECRET_TOKEN")
	baseURL, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, baseURL, iex.DefaultHTTPClient)

	quote, err = stock.Quote("aapl", nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)

	// you can also use the Get helper on client to unmarshal to your own custom struct
	anonstruct := &struct {
		Symbol string `json:"symbol,omitempty"`
	}{}
	err = client.Get("stock/aapl/quote", anonstruct, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Symbol", anonstruct.Symbol)
}
```

## SUPPORTED ENDPOINTS
### Account
- [x] Metadata
- [x] Usage
- [ ] Pay as you go
- [ ] Message Budget
- [ ] Signed Requests
- [ ] Setting up signed token
- [ ] Getting the secret for a signed token
### Data APIS
- [x] Data Points
- [ ] Data Tables
- [ ] Time Series
### Stocks
- [x] Advanced Stats
- [x] Balance Sheet
- [x] Batch Requests
- [x] Book
- [x] Cash Flow
- [x] Chart
- [x] Collection
- [x] Company
- [x] Delayed Quote
- [x] Dividends
- [x] Earnings
- [x] Earnings Today
- [x] Effective Spread
- [x] Estimates
- [x] Financials
- [x] Fund Ownership
- [x] Historical Prices
- [x] Income Statement
- [x] Insider Roster
- [x] Insider Summary
- [x] Insider Transactions
- [x] Institutional Ownership
- [x] Intraday Prices
- [x] IPO Calendar
- [x] Key Stats
- [x] Largest Trades
- [x] List
- [x] Logo
- [x] Market Volume (U.S.)
- [x] News
- [x] OHLC
- [x] Open / Close Price
- [x] Options
- [x] Peers
- [x] Previous Day Price
- [x] Price
- [x] Price Target
- [x] Quote
- [x] Recommendation Trends
- [x] Sector Performance
- [x] Splits
- [x] Upcoming Events
- [x] Volume by Venue
### Alternative Data
- [x] Crypto
- [x] Social Sentiment
- [x] CEO Compensation
### Reference Data
- [x] Symbols
- [x] IEX Symbols
- [x] International Symbols
- [x] International Exchanges
- [x] U.S. Exchanges
- [x] U.S. Holiday and Trading Dates
- [x] Sectors
- [x] Tags
- [x] Mutual Fund Symbols
- [x] OTC Symbols
- [x] FX Symbols
- [ ] Options Symbols
- [ ] Commodities Symbols In Dev
- [ ] Bonds Symbols In Dev
- [ ] Crypto Symbols In Dev
### Forex / Currencies
- [x] Exchange Rates
### Investors Exchange Data
- [x] TOPS
- [x] Last
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
- [ ] Listed Regulation SHO Threshold Securities List
- [ ] Listed Short Interest List In Dev
- [ ] Stats Historical Daily In Dev
- [ ] Stats Historical Summary
- [ ] Stats Intraday
- [ ] Stats Recent
- [ ] Stats Records
### API System Metadata
- [x] Status

## DEV NOTES
* Use this [online json to struct converter](https://mholt.github.io/json-to-go/) to save time
