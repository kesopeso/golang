package main

import "fmt"

const x int64 = 10

const (
	idKey  = "id"
	idName = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1            // will not compile
	//y = "some new value" // will not compile

	fmt.Println(x)
	fmt.Println(y)
}
