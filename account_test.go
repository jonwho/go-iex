package goiex

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestNewAccount(t *testing.T) {
	acc := NewAccount("test_token", "", sandboxURL, nil)

	expected = "https://sandbox.iexapis.com/"
	actual = acc.URL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "account/"
	actual = acc.APIURL().String()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}

	expected = "test_token"
	actual = acc.Token()
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestMetadata(t *testing.T) {
	rec, err := recorder.New("cassettes/account/metadata")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewAccount(testToken, DefaultVersion, sandboxURL, httpClient)

	metadata, err := cli.Metadata()
	if err != nil {
		t.Error(err)
	}
	if metadata == nil {
		t.Errorf("\nExpected metadata to be not nil\n")
	}
	expected = "alnnua"
	actual = metadata.SubscriptionTermType
	if expected != actual {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestUsage(t *testing.T) {
	rec, err := recorder.New("cassettes/account/usage")
	if err != nil {
		log.Fatal(err)
	} else {
		rec.SetMatcher(matchWithoutToken)
		httpClient = &http.Client{Transport: rec}
	}
	rec.AddFilter(removeToken)
	defer rec.Stop()
	cli := NewAccount(testToken, DefaultVersion, sandboxURL, httpClient)

	usage, err := cli.Usage()
	if err != nil {
		t.Error(err)
	}
	if usage.Messages.MonthlyUsage < 1 {
		t.Errorf("\nExpected MonthlyUsage > 0 but got: %d\n", usage.Messages.MonthlyUsage)
	}
	expected = "4"
	if actual = usage.Messages.DailyUsage.(map[string]interface{})["20191021"]; actual != expected {
		t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

// TODO: fix mock server usage

// func TestPayasyougo(t *testing.T) {
//   _, err := mockClient.Payasyougo(map[string]string{
//     "token": "my_token",
//     "allow": "false",
//   })
//   if err != nil {
//     t.Error(err)
//   }
// }
//
// func TestMessageBudget(t *testing.T) {
//   _, err := mockClient.MessageBudget(map[string]string{
//     "token":         "my_token",
//     "totalMessages": "500000",
//   })
//   if err != nil {
//     t.Error(err)
//   }
// }
