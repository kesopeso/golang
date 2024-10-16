package main

import (
	"fmt"
	"strings"
)

func main() {
	printFunctionInfo("for switch break!")

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8:
			fmt.Println(i, "even")
		case 1, 3, 5, 9:
			fmt.Println(i, "odd")
		case 7:
			fmt.Println("should exit all!")
			break loop
		default:
		}
	}
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
