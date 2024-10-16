.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./dependency_injection ./dependency_injection
	go build -o ./embedded ./embedded
	go build -o ./interfaces ./interfaces
	go build -o ./type_assertion_and_switch ./type_assertion_and_switch
	go build -o ./exercises ./exercises

clean:
	go clean ./...
