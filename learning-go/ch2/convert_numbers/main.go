package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.2

	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)

	var b byte = 100

	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b

	fmt.Println(sum1, sum2, sum3, sum4)
}
