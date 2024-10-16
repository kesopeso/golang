.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./custom_as_errors ./custom_as_errors
	go build -o ./errors ./errors
	go build -o ./exercises ./exercises
	go build -o ./wrap_error ./wrap_error

clean:
	go clean ./...
