.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o complex ./complex
	go build -o const ./const
	go build -o convert_numbers ./convert_numbers
	go build -o exercises ./exercises

clean:
	go clean ./...
