package main

import (
	"fmt"
	"strconv"
)

func main() {
	type opFunc = func(int, int) int
	opMap := map[string]opFunc{
		"+": func(i int, j int) int { return i + j },
		"-": func(i int, j int) int { return i - j },
		"*": func(i int, j int) int { return i * j },
		"/": func(i int, j int) int { return i / j },
	}
	expressions := [][]string{
		{"3", "+", "2"},
		{"21", "-", "4"},
		{"212", "*", "2"},
		{"212", "*+", "2"},
		{"20", "/", "2"},
		{"432", "/", "2", "not ok"},
	}
	for _, expression := range expressions {
		fmt.Println("these are arguments", expression)

		if len(expression) != 3 {
			fmt.Println("Invalid slice size")
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println("Can't convert the first value to number")
			continue
		}

		op, ok := opMap[expression[1]]
		if !ok {
			fmt.Println("Can't find the expression you're looking for")
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println("Can't convert the second value to number")
			continue
		}

		fmt.Println("Result:", op(p1, p2))
	}
}
