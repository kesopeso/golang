.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./garbage_collection ./garbage_collection
	go build -gcflags="-m" -o ./exercises ./exercises

clean:
	go clean ./...
