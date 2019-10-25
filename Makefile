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
	$(GOTEST) -coverprofile=coverage.out && $(GOTOOL) cover -func=coverage.out

.PHONY: example
example:
	$(GORUN) examples/main.go
