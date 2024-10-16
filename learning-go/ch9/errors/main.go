package main

import (
	"archive/zip"
	"fmt"
)

func main() {
	err := zip.ErrFormat
	fmt.Println(err)
}
