package mockiex

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

// Server returns a mock IEX server
func Server() *httptest.Server {
	var resp []byte

	return httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		log.Println("logging request URI:", r.RequestURI)
		switch r.RequestURI {
		case "/stock/aapl/quote":
			resp = read("mock-iex/responses/quote/aapl.json")
		case "/stock/aapl/quote?displayPercent=true":
			resp = read("mock-iex/responses/quote/aapl_with_display_percent_true.json")
		case "/stock/aapl/chart/1d":
			resp = read("mock-iex/responses/chart/aapl_1d.json")
		case "/ref-data/symbols":
			resp = read("mock-iex/responses/ref_data_symbols.json")
		case "/stock/market/today-earnings":
			resp = read("mock-iex/responses/earnings_today.json")
		case "/stock/aapl/earnings":
			resp = read("mock-iex/responses/earnings/aapl.json")
		case "/stock/aapl/stats":
			resp = read("mock-iex/responses/key_stats/aapl.json")
		case "/stock/aapl/book":
			resp = read("mock-iex/responses/book/aapl.json")
		case "/ref-data/daily-list/corporate-actions":
			resp = read("mock-iex/responses/ref_data_corporate_actions.json")
		case "/ref-data/daily-list/corporate-actions/sample":
			resp = read("mock-iex/responses/ref_data_corporate_actions_sample.json")
		case "/ref-data/daily-list/dividends":
			resp = read("mock-iex/responses/ref_data_dividends.json")
		case "/ref-data/daily-list/dividends/sample":
			resp = read("mock-iex/responses/ref_data_dividends_sample.json")
		case "/ref-data/daily-list/next-day-ex-date":
			resp = read("mock-iex/responses/ref_data_next_day_ex_date.json")
		case "/ref-data/daily-list/next-day-ex-date/sample":
			resp = read("mock-iex/responses/ref_data_next_day_ex_date_sample.json")
		case "/ref-data/daily-list/symbol-directory":
			resp = read("mock-iex/responses/ref_data_symbol_directory.json")
		case "/ref-data/daily-list/symbol-directory/sample":
			resp = read("mock-iex/responses/ref_data_symbol_directory_sample.json")
		default:
			http.Error(w, "not found", http.StatusNotFound)
		}
		w.Write(resp)
	}))
}

func read(path string) []byte {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	return bytes
}
