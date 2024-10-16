package main

import (
	"fmt"
	"strings"
)

func main() {
	namedAndOptionalOpts(NamedOptionalOpts{firstName: "Janez", lastName: "Jansa"})
}

type NamedOptionalOpts struct {
	firstName string
	lastName  string
	age       int
}

func namedAndOptionalOpts(opts NamedOptionalOpts) {
	fmt.Println("first name:", opts.firstName)
	fmt.Println("last name:", opts.lastName)
	fmt.Println("age:", opts.age)
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println(strings.Repeat("=", len(info)))
}
