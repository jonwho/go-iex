package main

import (
	"fmt"
	"log"
	"os"

	iex "github.com/jonwho/go-iex/v5"
)

// The IEX message usage for TechnicalIndicators is pretty high. Can potentially save on messages
// for API by implementing these indicators yourself.
func main() {
	token := os.Getenv("IEX_SECRET_TOKEN")
	client, err := iex.NewClient(token, iex.SetClientRetry())
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the RSI for a 14 day period.
	charts, err := client.Chart("aapl", iex.ChartRangeOneMonth, "", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// N.B. If chart history is less than 14 days then return 0.0 for RSI
	data := charts[len(charts)-14:]
	sumGain := 0.0
	sumLoss := 0.0
	for _, chart := range data[:len(data)-1] {
		sumGain += max(0, chart.Change)
		sumLoss += min(0, chart.Change)
	}
	period := 14.0
	avgGain := sumGain / period
	avgLoss := abs(sumLoss / period)
	relativeStrength := avgGain / avgLoss
	relativeStrengthIndicator := 100 - 100/(relativeStrength+1)
	fmt.Println("RSI is", relativeStrengthIndicator)
	if relativeStrengthIndicator >= 70.0 {
		fmt.Println("Overbought!")
	} else if relativeStrengthIndicator <= 30.0 {
		fmt.Println("Oversold!")
	} else {
		fmt.Println("Meh.")
	}

	os.Exit(0)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}
