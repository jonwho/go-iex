package goiextest

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	iex "github.com/jonwho/go-iex"
)

var (
	expected, actual            interface{}
	sandboxURL, _               = url.Parse(iex.SandboxBaseURL)
	defaultURL, _               = url.Parse(iex.DefaultBaseURL)
	iexClient, iexSandboxClient *iex.Client
	err                         error
)

func init() {
	iexSandboxClient, err = iex.NewClient("",
		iex.SetVersion(iex.DefaultVersion),
		iex.SetURL(iex.SandboxBaseURL),
		iex.SetAccount("", iex.DefaultVersion, defaultURL, accountClient()),
		iex.SetAlternativeData("", iex.DefaultVersion, sandboxURL, alternativeDataClient()),
		iex.SetAPISystemMetadata("", iex.DefaultVersion, sandboxURL, apiSystemMetadataClient()),
		iex.SetDataAPI("", iex.DefaultVersion, sandboxURL, dataAPIClient()),
		iex.SetForex("", iex.DefaultVersion, sandboxURL, forexClient()),
		iex.SetStock("", iex.DefaultVersion, sandboxURL, stockClient()),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func accountClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/account"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func alternativeDataClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/alternative_data"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func apiSystemMetadataClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/api_system_metadata"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func dataAPIClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/data_apis"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func forexClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/forex"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func stockClient() (cli *http.Client) {
	if rec, err := recorder.New("../mock-iex/recorder/stocks"); err != nil {
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rec.SetMatcher(matchWithoutToken)
		cli = &http.Client{Transport: rec}
	}
	return
}

func matchWithoutToken(req *http.Request, i cassette.Request) bool {
	u := req.URL
	q := u.Query()
	q.Del("token")
	u.RawQuery = q.Encode()
	req.URL = u
	return u.String() == i.URL
}
