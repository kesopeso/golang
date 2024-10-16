package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	i := 20
	var k float64 = float64(i)

	fmt.Println(i)
	fmt.Println(k)
}

func exercise2() {
	const value = 10
	var i int = value
	var k float64 = value

	fmt.Println(i)
	fmt.Println(k)
}

func exercise3() {
	var (
		b      byte   = 255
		smallI int32  = 2_147_483_647
		bigI   uint64 = 18_446_744_073_709_551_615
	)

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	b = b + 1
	b += 1

	smallI = smallI + 1
	smallI += 1

	bigI = bigI + 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
