.DEFAULT_GOAL := build

.PHONY: clean fmt vet get build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

get: vet
	go get ./...

build: get
#	go build -o ./todo ./todo

clean:
	go clean ./...
