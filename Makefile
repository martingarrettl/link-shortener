MODULE = $(shell go list -m)

.PHONY: generate build test lint
generate:
	go generate ./..

build: # build server
	go build -a -o shorturl-server $(MODULE)/cmd

test:
	go clean -testcache
	go test ./.. -v

lint:
	gofmt -l .