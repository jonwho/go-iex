package goiextest

import (
	"testing"
)

func TestStatus(t *testing.T) {
	status, err := iexSandboxClient.Status()
	if err != nil {
		t.Error(err)
	}
	expected = `up`
	actual = status.Status
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
