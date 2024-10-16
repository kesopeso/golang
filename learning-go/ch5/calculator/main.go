package main

import (
	"fmt"
	"strconv"
)

func main() {
	opMap := map[string]func(int, int) int{
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

		result := opFunc(p1, p2)
		fmt.Println("This is the result:", result)
	}
}

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}
