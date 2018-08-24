package goiex

import (
	"fmt"
	"testing"
)

var client = NewClient()

func TestEarningsToday(t *testing.T) {
	earnings := client.EarningsToday()

	fmt.Println(earnings)

	if earnings != "ok!" {
		t.Error("wrong string!")
	}
}
