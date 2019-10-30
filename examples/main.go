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
