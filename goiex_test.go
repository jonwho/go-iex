package goiex

import (
	"fmt"
	"testing"
)

var client = NewClient()

// func TestEarningsToday(t *testing.T) {
//   earnings := client.EarningsToday()
//
//   fmt.Println(earnings)
//
//   if earnings != "ok!" {
//     t.Error("wrong string!")
//   }
// }

func TestEarnings(t *testing.T) {
	earnings, _ := client.Earnings("aapl")

	fmt.Printf("%+v\n", earnings)

	if earnings.Symbol != "AAPL" {
		t.Error("wrong string!")
	}

	if earnings.Earnings[0].ActualEPS != 2.34 {
		t.Error("wrong value!")
	}
}
