package main

import (
	"fmt"
	"strings"
)

func main() {
	createSimpleStruct()
}

func createSimpleStruct() {
	printFunctionInfo("Create simple struct")

	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person

	fmt.Println("empty person struct fred:", fred)
	fmt.Println("name:", fred.name)
	fmt.Println("age:", fred.age)
	fmt.Println("pet:", fred.pet)
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
