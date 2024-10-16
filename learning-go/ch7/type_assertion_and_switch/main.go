package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(MyInt)
	if !ok {
		if !ok {
			fmt.Println("unexpected type for i2")
			return
		}
	}

	fmt.Println(i2 + 1)

	switch i := i.(type) {
	case int:
		fmt.Println("type of i is an int!", i)
	case MyInt:
		fmt.Println("type of i is MyInt", i)
	default:
		fmt.Println("no type found for i", i)
	}

	i3, ok := i.(int)
	if !ok {
		fmt.Println("unexpected type for i3")
		return
	}
	fmt.Println(i3 + 1)
}
