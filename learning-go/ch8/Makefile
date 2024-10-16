.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./non_generic_tree ./non_generic_tree
	go build -o ./generic_tree ./generic_tree
	go build -o ./diff ./diff
	go build -o ./map_reduce ./map_reduce
	go build -o ./exercises ./exercises

clean:
	go clean ./...
