.DEFAULT_GOAL := build

.PHONY: fmt get lint generate build clean test

fmt:
	go fmt ./...

get: fmt
	go get -t ./...

lint: get
	go vet ./...
	staticcheck ./...
	golangci-lint run

generate: get
	go generate ./...

build: generate
	go build -o ./docker/server ./docker

clean:
	go clean ./...

test:
	go test -count=1 ./...

test_file_parser:
	go test -count=1 -v -cover -coverprofile=./file_parser/c.out ./file_parser
	go tool cover -html=./file_parser/c.out
