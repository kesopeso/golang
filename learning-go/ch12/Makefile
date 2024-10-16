.DEFAULT_GOAL := build

.PHONY: fmt lint get build clean

fmt:
	go fmt ./...

lint: fmt
	go vet ./...
	staticcheck ./...
	golangci-lint run
	revive ./...

get: lint
	go get ./...

build: get
	go build -o ./go_routine ./go_routine
	go build -o ./select ./select

clean:
	go clean ./...
