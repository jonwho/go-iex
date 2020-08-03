package main

import (
	"log"
	"os"
	"sync"

	iex "github.com/jonwho/go-iex/v4"
)

func main() {
	sandboxToken := os.Getenv("IEX_TEST_SECRET_TOKEN")
	client, err := iex.NewSandboxClient(sandboxToken, iex.SetClientRetry())
	if err != nil {
		log.Fatalln(err)
	}

	iexSymbols, err := client.IEXSymbols()
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	var mtx = &sync.Mutex{}
	hmap := make(map[string]interface{})

	for _, iexSymbol := range iexSymbols {
		hmap[iexSymbol.Symbol] = nil
		wg.Add(1)

		go func(symbol string) {
			defer wg.Done()

			quote, err := client.Quote(symbol, nil)
			if err != nil {
				log.Printf("Error for %s: %s\n", symbol, err.Error())
				mtx.Lock()
				delete(hmap, symbol)
				mtx.Unlock()
				return
			}

			if quote == nil {
				log.Printf("Quote for %s was nil\n", symbol)
			}

			mtx.Lock()
			if quote.LatestPrice < 5.00 {
				hmap[symbol] = quote
			} else {
				delete(hmap, symbol)
			}
			mtx.Unlock()
		}(iexSymbol.Symbol)
	}

	wg.Wait()

	log.Println("Number of IEX Symbols:", len(iexSymbols))
	log.Println("Number of IEX Symbols after rate limited filter run:", len(hmap))

	os.Exit(0)
}
