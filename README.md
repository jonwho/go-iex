# go-iex

[![GoDoc](https://godoc.org/github.com/jonwho/go-iex?status.svg)](http://godoc.org/github.com/jonwho/go-iex)

## ATTRIBUTION
Data provided for free by [IEX](https://iextrading.com/developer/). View [IEX’s](https://iextrading.com/api-exhibit-a/) Terms of Use.

## DESCRIPTION
Client interface to iex trading API.

## USAGE
```go
package main

import (
	"fmt"

	iex "github.com/jonwho/go-iex"
)

func main() {
	client = iex.NewClient()

	quote, err := client.Quote("aapl")
	if err != nil {
		fmt.Fatalln(err)
	}

	fmt.Println("Symbol", quote.Symbol, "Company Name", quote.CompanyName,
		"Current Price", quote.LatestPrice)
}
```

## SUPPORTED ENDPOINTS
`%s` - string parameter
`%d` - integer parameter

| Endpoint                               | Version |
| -------------------------------------- | ------- |
| /stock/%s/book                         |   ✅    |
| /stock/%s/chart                        |   ✅    |
| /stock/%s/earnings                     |   ✅    |
| /stock/market/today-earnings           |   ✅    |
| /stock/%s/news/%d                      |   ✅    |
| /stock/%s/quote                        |   ✅    |
| /stock/%s/stats                        |   ✅    |
| /ref-data/symbols                      |   ✅    |
| /ref-data/daily-list/corporate-actions |   ✅    |
| /ref-data/daily-list/dividends         |   ✅    |
| /ref-data/daily-list/next-day-ex-date  |   ✅    |
| /ref-data/daily-list/symbol-directory  |   ✅    |
| etc...                                 |   ❌    |

## DEV NOTES
* Use this [online json to struct converter](https://mholt.github.io/json-to-go/) to save time
