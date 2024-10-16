.DEFAULT_GOAL := build

.PHONY: fmt lint get generate build clean

fmt:
	go fmt ./...

lint: fmt
	go vet ./...
	staticcheck ./...
	golangci-lint run
	revive -config built_in.toml ./...

get: lint
	go get ./...

generate: get
	go generate ./...

build: generate
	go build -o ./cancellation ./cancellation
	go build -o ./ctx_timeout ./ctx_timeout
	go build -o ./exercises ./exercises

clean:
	go clean ./...
