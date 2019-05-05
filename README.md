# go-iex

[![GoDoc](https://godoc.org/github.com/jonwho/go-iex?status.svg)](http://godoc.org/github.com/jonwho/go-iex)

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
	"os"
	"fmt"

	iex "github.com/jonwho/go-iex"
)

func main() {
	token := os.Getenv("IEX_SECRET_TOKEN")
	// client will have all currently supported IEX APIs
	client = iex.NewClient(token)

	quote, err := client.Quote("aapl")
	if err != nil {
		fmt.Fatalln(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)

	// if you only want to test against sandbox build a custom client
	// get Stocks only API client for sandbox testing
	token = os.Getenv("IEX_TEST_SECRET_TOKEN")
	baseURL, _ := url.Parse(iex.SandboxBaseURL)
	stock := iex.NewStock(token, iex.DefaultVersion, baseURL, iex.DefaultHTTPClient)

	quote, err = stock.Quote("aapl")
	if err != nil {
		fmt.Fatalln(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)
}
```

## SUPPORTED ENDPOINTS
### Account
- [x] Metadata
### Data API
- [x] DataPoints
### Stocks
- [x] AdvancedStats
- [x] BalanceSheet
- [x] Batch
- [x] Book
- [x] CashFlow
- [x] Chart
- [x] Collection
- [x] Company
- [x] DelayedQuote
- [x] Dividends
- [x] Earnings
- [x] EarningsToday
- [x] EffectiveSpread
- [x] Quote

## DEV NOTES
* Use this [online json to struct converter](https://mholt.github.io/json-to-go/) to save time
