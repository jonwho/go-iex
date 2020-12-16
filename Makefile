GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOTOOL=$(GOCMD) tool
GOPHERBADGER=$(shell go env GOPATH)/bin/gopherbadger

.PHONY: all
all: test cover coverage

.PHONY: test
test:
	$(GOTEST) -v -cover -count=1 -mod=vendor

.PHONY: cover
cover:
	$(GOTEST) ./... -coverprofile=coverage.out && $(GOTOOL) cover -html=coverage.out -o=coverage.html

.PHONY: coverage
coverage:
	$(GOPHERBADGER) -md="README.md"

.PHONY: funcoverage
funcoverage:
	$(GOTEST) -mod=vendor -coverprofile=coverage.out && $(GOTOOL) cover -func=coverage.out

.PHONY: example
example:
	$(GORUN) examples/main.go -mod=vendor

.PHONY: example-rate-limit
example-rate-limit:
	$(GORUN) examples/ratelimit/main.go -mod=vendor

.PHONY: noutpsse
noutpsse:
	curl --header 'Accept: text/event-stream' https://cloud-sse.iexapis.com/stable/stocksUSNoUTP\?symbols\=spy\&token\=$(IEX_SECRET_TOKEN)
