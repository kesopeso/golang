package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	exercise1()

	filesize, error := exercise2("./main.go")
	fmt.Println("Exercise 2 results, filesize/error:", filesize, error)

	exercise3()
}

func exercise1() {
	fmt.Println("Exercise 1")

	opMap := map[string]func(int, int) (int, error){
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"3", "+", "2"},
		{"21", "-", "4"},
		{"212", "*", "2"},
		{"212", "*+", "2"},
		{"21", "/", "2"},
		{"432", "/", "2", "not ok"},
		{"432", "/", "0"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Not a valid expression length")
			continue
		}

		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		opFunc, ok := opMap[exp[1]]
		if !ok {
			fmt.Println("No such operation")
			continue
		}

		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, error := opFunc(p1, p2)
		if error != nil {
			fmt.Println("Error occured:", error)
			continue
		}
		fmt.Println("This is the result:", result)
	}
}

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

func exercise2(filename string) (int, error) {
	fmt.Println("Exercise 2")

	f, error := os.Open(filename)
	if error != nil {
		return 0, error
	}

	defer f.Close()

	data, filesize := make([]byte, 2048), 0

	for {
		count, error := f.Read(data)
		filesize += count
		if error != nil {
			if error != io.EOF {
				return 0, error
			}
			break
		}
	}

	return filesize, nil
}

func exercise3() {
	fmt.Println("Exercise 3")

	prefixer := func(prefix string) func(string) string {
		return func(value string) string {
			return prefix + " " + value
		}
	}

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should output Hello Bob
	fmt.Println(helloPrefix("Maria")) // should output Hello Maria
}
