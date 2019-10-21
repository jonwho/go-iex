GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOPHERBADGER=$(HOME)/go/bin/gopherbadger

all: test coverage

.PHONY: test
test:
	$(GOTEST) -cover -count=1 ./go-iex-test/...

.PHONY: coverage
coverage:
	$(GOPHERBADGER) -md="README.md"

.PHONY: example
example:
	$(GORUN) examples/main.go
