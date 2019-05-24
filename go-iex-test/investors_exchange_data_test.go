package goiextest

import (
	"testing"
)

func TestTOPS(t *testing.T) {
	tops, err := iexSandboxClient.TOPS(nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(tops) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
