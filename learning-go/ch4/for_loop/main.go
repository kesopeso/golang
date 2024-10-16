package main

import (
	"fmt"
	"strings"
)

func main() {
	completeForLoop()
	conditionOnlyForLoop()
	infiniteForLoop()
	rangeForLoop()
	onlyKeyRangeMap()
	forRangeIsACopy()
}

func forRangeIsACopy() {
	printFunctionInfo("for range is a copy")
	vals := []int{1, 2, 3}
	for _, v := range vals {
		v *= 2
	}
	fmt.Println(vals)
}

func onlyKeyRangeMap() {
	printFunctionInfo("only key range loop")
	myMap := map[string]bool{
		"John":    true,
		"Bob":     false,
		"Raymond": true,
	}

	for k := range myMap {
		fmt.Println("Just print the name key:", k)
	}
}

func rangeForLoop() {
	printFunctionInfo("range for loop")
	vals := []int{2, 3, 4, 5, 6}
	for i, x := range vals {
		fmt.Println("index:", i, "value:", x)
	}
}

func infiniteForLoop() {
	printFunctionInfo("infinite for loop")
	// use it when you want to execute the loop at least once
	i := 0
	for {
		fmt.Println("infinite hello!")
		i++
		if i > 9 {
			break
		}
	}
	fmt.Println("done executing infinite for")
}

func conditionOnlyForLoop() {
	printFunctionInfo("condition only - while like!")
	// this is same as while statement in other languages
	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 2
	}
}

func completeForLoop() {
	printFunctionInfo("complete for ")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
