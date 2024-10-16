.DEFAULT_GOAL := build

.PHONY: clean fmt vet get vuln generate build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...
	staticcheck ./...
	revive -config built_in.toml ./...
	golangci-lint run

get: vet
	go get ./...

vuln: get
	govulncheck ./...

generate: vuln
	go generate ./...

build: generate
	go build -o ./script ./script
	go build -o ./go_embed_comments ./go_embed_comments
	go build -o ./help_system ./help_system

clean:
	go clean ./...
