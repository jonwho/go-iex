package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	iex "github.com/jonwho/go-iex/v4"
)

func main() {
	token := os.Getenv("IEX_SECRET_TOKEN")
	// client will have all currently supported IEX APIs implemented by go-iex
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
	stock := iex.NewStock(token, iex.DefaultVersion, baseURL, http.DefaultClient)

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
	// query parameters are supported by `https://github.com/google/go-querystring`
	queryParams := &struct {
		DisplayPercent bool `url:"displayPercent,omitempty"`
	}{true}
	err = client.Get("stock/aapl/quote", anonstruct, queryParams)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Symbol", anonstruct.Symbol)

	// some functions have defined parameters for path or query parameters
	charts, err := stock.Chart("aapl", iex.ChartRangeOneMonth, &iex.ChartQueryParams{ChartCloseOnly: true})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Charts %#v\n", charts)
}
