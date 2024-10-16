package main

import (
	"fmt"
)

func main() {
	s := []string{"first", "second", "last"}
	fmt.Println(s, len(s))

	clear(s)
	fmt.Println(s, len(s))
}
