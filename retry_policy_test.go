package goiex

import (
	"errors"
	"net/http"
	"testing"
)

func TestDefaultRetryPolicy(t *testing.T) {
	var ayroar = errors.New("ohno")
	actual, err := DefaultRetryPolicy(nil, ayroar)
	if ayroar != err {
		t.Errorf("\nExpected: %v\nActual: %v\n", ayroar, err)
	}
	expected = true
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	resp := &http.Response{StatusCode: 0}
	actual, err = DefaultRetryPolicy(resp, nil)
	if err != nil {
		t.Errorf("\nExpected: %v\nActual: %v\n", nil, err)
	}
	expected = true
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	resp = &http.Response{StatusCode: 429}
	actual, err = DefaultRetryPolicy(resp, nil)
	if err != nil {
		t.Errorf("\nExpected: %v\nActual: %v\n", nil, err)
	}
	expected = true
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	resp = &http.Response{StatusCode: 500}
	actual, err = DefaultRetryPolicy(resp, nil)
	if err != nil {
		t.Errorf("\nExpected: %v\nActual: %v\n", nil, err)
	}
	expected = true
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	resp = &http.Response{StatusCode: 501}
	actual, err = DefaultRetryPolicy(resp, nil)
	if err != nil {
		t.Errorf("\nExpected: %v\nActual: %v\n", nil, err)
	}
	expected = true
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	resp = &http.Response{StatusCode: 200}
	actual, err = DefaultRetryPolicy(resp, nil)
	if err != nil {
		t.Errorf("\nExpected: %v\nActual: %v\n", nil, err)
	}
	expected = false
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
