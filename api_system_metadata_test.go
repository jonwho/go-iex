package goiex

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewAPISystemMetadata(t *testing.T) {
	u, _ := url.Parse(SandboxBaseURL)
	asm := NewAPISystemMetadata("test_token", "", u, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = asm.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = ""
	actual = asm.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = asm.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestStatus(t *testing.T) {
	rec, err := recorder.New("cassettes/api_system_metadata/status")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()

	u, _ := url.Parse(SandboxBaseURL)
	cli := NewAPISystemMetadata(testToken, DefaultVersion, u, httpClient)

	status, err := cli.Status()
	if err != nil {
		t.Error(err)
	}
	expected = `up`
	actual = status.Status
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
