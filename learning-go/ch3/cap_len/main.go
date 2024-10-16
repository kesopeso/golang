package main

import (
	"fmt"
)

func main() {
	capacity()
}

func capacity() {
	x := []int{1, 2, 3}
	fmt.Println(x, len(x), cap(x))
	fmt.Println(x == nil)
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	fmt.Println(x == nil)
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
