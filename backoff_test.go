package goiex

import (
	"testing"
	"time"
)

func TestDefaultBackoff(t *testing.T) {
	dur := DefaultBackoff(1*time.Second, 5*time.Second, 4, nil)
	if dur > 5*time.Second {
		t.Error("Duration is too long")
	}

	dur = DefaultBackoff(1*time.Second, 5*time.Second, 200, nil)
	if dur != 5*time.Second {
		t.Errorf("\nExpected: %v\nActual: %v", 5*time.Second, dur)
	}
}
