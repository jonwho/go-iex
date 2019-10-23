GOCMD=go
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOPHERBADGER=$(HOME)/go/bin/gopherbadger

all: test coverage

.PHONY: test
test:
	$(GOTEST) -v -cover -count=1

.PHONY: coverage
coverage:
	$(GOPHERBADGER) -md="README.md"

.PHONY: example
example:
	$(GORUN) examples/main.go
