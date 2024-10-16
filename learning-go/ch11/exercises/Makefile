.DEFAULT_GOAL := build

.PHONY: fmt vet static_check build clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

static_check: vet
	staticcheck ./...

build: static_check
	go build
	GOOS=windows GOARCH=arm64 go build

clean:
	go clean
