.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./blocks ./blocks
	go build -o ./for_loop ./for_loop
	go build -o ./switch ./switch

clean:
	go clean ./...
