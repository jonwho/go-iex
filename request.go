package goiex

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Request struct {
	body io.ReadSeeker // used to rewind request data between retries
	*http.Request
}

func NewRequest(method, url string, body io.ReadSeeker) (*Request, error) {
	var rc io.ReadCloser
	if body != nil {
		rc = ioutil.NopCloser(body)
	}

	req, err := http.NewRequest(method, url, rc)
	if err != nil {
		return nil, err
	}

	return &Request{body, req}, nil
}
