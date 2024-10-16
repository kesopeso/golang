.DEFAULT_GOAL := test

.PHONY: fuzz_test test

fuzz_test:
	go test -fuzz=FuzzParser

test_race:
	go test -race

test:
	go test -v -cover -coverprofile=simplewebapp.out
	go tool cover -html=simplewebapp.out
	rm -rf simplewebapp.out
	rm results.txt
