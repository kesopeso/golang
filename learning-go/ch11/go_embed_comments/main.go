package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed passwords.txt
var passwords string

func main() {
	pwds := strings.Split(passwords, "\n")

	for i, v := range pwds {
		fmt.Printf("Password %d: %s\n", i+1, v)
	}
}
