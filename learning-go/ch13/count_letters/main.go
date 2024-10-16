// package main starts the count letters function
package main

import (
	"compress/gzip"
	//	"embed"
	"fmt"
	"io"
	"os"
)

// //go:embed my_data.txt.gz
//var fs embed.FS

func main() {
	//reader, closeReader, err := getGzipReader(fs)
	reader, closeReader, err := getGzipReader()
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	defer closeReader()
	result, err := countLetters(reader)
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Println("result", result)
}

func getGzipReader() (*gzip.Reader, func(), error) {
	//func getGzipReader(fs embed.FS) (*gzip.Reader, func(), error) {
	//r, err := fs.Open("my_data.txt.gz")
	r, err := os.Open("./my_data.txt.gz")
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
