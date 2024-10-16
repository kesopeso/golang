package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	values := exercise1()
	exercise2(values)
	exercise3()
}

func exercise1() []int {
	printFunctionInfo("100 random numbers")
	values := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		values = append(values, rand.Intn(101))
	}
	fmt.Println("values:", values)
	return values
}

func exercise2(values []int) {
	printFunctionInfo("output values and break on divisable by 6")
loop:
	for _, v := range values {
		fmt.Println("Value to check:", v)
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
			break loop
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Nevermind!")
		}
	}
}

func exercise3() {
	printFunctionInfo("total check")
	var total int
	for i := 0; i < 10; i++ {
		// this is a BUG! := should be just =
		total := total + i
		fmt.Println("total in", total)
	}
	fmt.Println("total out", total)
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
