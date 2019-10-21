GOCMD=go
GOTEST=$(GOCMD) test

all: test

.PHONY: test
test:
	$(GOTEST) ./go-iex-test/...

.PHONY: example
example:
	$(GOCMD) run examples/main.go
