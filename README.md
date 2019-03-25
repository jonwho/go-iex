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
`aapl` is just a placeholder ticker example

| Endpoint                     | Version |
| ---------------------------- | ------- |
| /stock/aapl/book             | ✅      |
| /stock/aapl/chart            | ✅      |
| /stock/aapl/earnings         | ✅      |
| /stock/market/today-earnings | ✅      |
| /stock/aapl/quote            | ✅      |
| /stock/aapl/stats            | ✅      |
| /ref-data/symbols            | ✅      |
| etc...                       | ❌      |
