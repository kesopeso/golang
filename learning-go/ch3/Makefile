.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o ./cap_len ./cap_len
	go build -o ./make_slice ./make_slice
	go build -o ./clear_slice ./clear_slice
	go build -o ./slicing_slices ./slicing_slices
	go build -o ./slice_share_storage ./slice_share_storage
	go build -o ./confusing_slices ./confusing_slices
	go build -o ./str_len ./str_len
	go build -o ./maps ./maps
	go build -o ./structs ./structs
	go build -o ./exercises ./exercises

clean:
	go clean ./...
