package main

import (
	"fmt"
	"strings"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	printFunctionInfo("Exercise 1")

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subsliceOne := greetings[:2]
	subsliceTwo := greetings[1:4]
	subsliceThree := greetings[3:]

	fmt.Println(greetings, subsliceOne, subsliceTwo, subsliceThree)
}

func exercise2() {
	printFunctionInfo("Exercise 2")

	message := "Hi 👩 and 👨"
	messageAsRune := []rune(message)

	fmt.Println(message, string(messageAsRune[3]))
}

func exercise3() {
	printFunctionInfo("Exercise 3")

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	bob := Employee{
		"Bob",
		"Myers",
		1,
	}

	andrew := Employee{
		firstName: "Andrew",
		lastName:  "Beggar",
		id:        2,
	}

	var mathew Employee
	mathew.firstName = "Mathew"
	mathew.lastName = "Moore"
	mathew.id = 3

	fmt.Println(bob, andrew, mathew)
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
