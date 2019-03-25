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

## Supported endpoints
| Endpoint                     | Version |
| ---------------------------- | ------- |
| /stock/<ticker>/book         | &#9745; |
| /stock/<ticker>/chart        | &#9745; |
| /stock/<ticker>/earnings     | &#9745; |
| /stock/market/today-earnings | &#9745; |
| /stock/<ticker>/quote        | &#9745; |
| /stock/<ticker>/stats        | &#9745; |
| /ref-data/symbols            | &#9745; |
| etc...                       | &#9744; |
