package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//deferOne()
	deferTwo()
}

func deferTwo() int {
	a := 10

	defer func(i int) {
		fmt.Println("this is first defer:", i)
	}(a)

	a = 20

	defer func(i int) {
		fmt.Println("this is second defer:", i)
	}(a)

	a = 30
	fmt.Println("inside function:", a)

	return a
}

func deferOne() {
	if len(os.Args) < 2 {
		log.Fatal("No file specified!")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// this will defer the f.Close() execution until the surrounding function (in this case main ends execution
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
