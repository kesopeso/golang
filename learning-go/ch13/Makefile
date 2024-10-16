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
	go build -o ./count_letters ./count_letters
	go build -o ./time ./time
	go build -o ./encode_decode ./encode_decode
	go build -o ./custom_json ./custom_json
	go build -o ./net_http ./net_http
	go build -o ./server ./server
	go build -o ./exercises ./exercises

clean:
	go clean ./...
