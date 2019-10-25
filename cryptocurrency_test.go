package goiex

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewCryptocurrency(t *testing.T) {
	cry := NewCryptocurrency("test_token", "", sandboxURL, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = cry.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "crypto/"
	actual = cry.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = cry.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestCryptoBook(t *testing.T) {
	rec, err := recorder.New("cassettes/cryptocurrency/book")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewCryptocurrency(testToken, DefaultVersion, sandboxURL, httpClient)

	book, err := cli.CryptoBook("btcusd")
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(book.Bids) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = false
	actual = len(book.Asks) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCryptoPrice(t *testing.T) {
	rec, err := recorder.New("cassettes/cryptocurrency/price")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewCryptocurrency(testToken, DefaultVersion, sandboxURL, httpClient)

	price, err := cli.CryptoPrice("btcusd")
	if err != nil {
		t.Error(err)
	}
	expected = "BTCUSD"
	actual = price.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = 7734.85
	actual = price.Price
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestCryptoQuote(t *testing.T) {
	rec, err := recorder.New("cassettes/cryptocurrency/quote")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewCryptocurrency(testToken, DefaultVersion, sandboxURL, httpClient)

	quote, err := cli.CryptoQuote("btcusd")
	if err != nil {
		t.Error(err)
	}
	expected = "BTCUSD"
	actual = quote.Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = 7563.71
	actual = quote.LatestPrice
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
