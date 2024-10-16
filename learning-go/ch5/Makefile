.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./functions ./functions
	go build -o ./calculator ./calculator
	go build -o ./calculator_anon ./calculator_anon
	go build -o ./defer ./defer
	go build -o ./exercises ./exercises

clean:
	go clean ./...
