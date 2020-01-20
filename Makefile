GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOTOOL=$(GOCMD) tool
GOPHERBADGER=$(HOME)/go/bin/gopherbadger

all: test coverage

.PHONY: test
test:
	$(GOTEST) -v -cover -count=1 -mod=vendor

.PHONY: coverage
coverage:
	$(GOPHERBADGER) -md="README.md"

.PHONY: funcoverage
funcoverage:
	$(GOTEST) -mod=vendor -coverprofile=coverage.out && $(GOTOOL) cover -func=coverage.out

.PHONY: example
example:
	$(GORUN) examples/main.go -mod=vendor

.PHONY: noutpsse
noutpsse:
	curl --header 'Accept: text/event-stream' https://cloud-sse.iexapis.com/stable/stocksUSNoUTP\?symbols\=spy\&token\=$(IEX_SECRET_TOKEN)
