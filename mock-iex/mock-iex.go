package mockiex

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// Client returns a http.Client using the mock Server
func Client() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				log.Println(Server().Listener.Addr().String())
				return net.Dial(network, Server().Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// Server returns a mock IEX server
func Server() *httptest.Server {
	var resp []byte

	return httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		q, _ := url.ParseQuery(u.RawQuery)
		q.Del("token")
		u.RawQuery = q.Encode()
		log.Println("logging request URI:", u.RequestURI())
		switch u.RequestURI() {
		case "/stable/account/metadata":
			resp = read("../mock-iex/responses/account/metadata.json")
		case "/stable/account/usage":
			resp = read("../mock-iex/responses/account/usage.json")
		case "/stable/account/payasyougo":
			if r.Method == http.MethodPost {
				resp = read("../mock-iex/responses/account/payasyougo.json")
			} else {
				http.Error(w, "not found", http.StatusNotFound)
			}
		case "/stable/account/messagebudget":
			if r.Method == http.MethodPost {
				resp = read("../mock-iex/responses/account/message_budget.json")
			} else {
				http.Error(w, "not found", http.StatusNotFound)
			}
		case "/stable/data-points/aapl":
			resp = read("../mock-iex/responses/data-points/data_points_aapl.json")
		case "/stock/aapl/batch?types=quote":
			resp = read("mock-iex/responses/batch/aapl.json")
		case "/stock/aapl/batch?last=5&range=1m&types=quote%2Cnews%2Cchart":
			resp = read("mock-iex/responses/batch/aapl_many_params.json")
		case "/stock/aapl/quote":
			resp = read("mock-iex/responses/quote/aapl.json")
		case "/stock/aapl/quote?displayPercent=true":
			resp = read("mock-iex/responses/quote/aapl_with_display_percent_true.json")
		case "/stock/aapl/chart/1d":
			resp = read("mock-iex/responses/chart/aapl_1d.json")
		case "/stock/aapl/news/last/1":
			resp = read("mock-iex/responses/news/aapl_1.json")
		case "/stock/aapl/news/last/10":
			resp = read("mock-iex/responses/news/aapl_10.json")
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
