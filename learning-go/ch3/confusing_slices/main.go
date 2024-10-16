package main

import "fmt"

func main() {
	unprotected()
	protected()
}

func unprotected() {
	x := make([]string, 0, 5)         // slice of strings with 0 elements and capacity 5
	x = append(x, "a", "b", "c", "d") // x = abcd_ len 4, cap 5
	y := x[:2]                        // y = ab___ len 2, cap 5
	z := x[2:]                        // z =   cd_ len 2, cap 3
	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, "i", "j", "k") // y = abijk, len 5, cap 5, x = abijk len 5, cap 5
	x = append(x, "x")           // x = abijx len 5 cap 5, y = abijx
	z = append(z, "y")           // z = cdz len 3 cap 3 x = y = abijz
	fmt.Println(len(x), cap(x), "x:", x)
	fmt.Println(len(y), cap(y), "y:", y)
	fmt.Println(len(z), cap(z), "z:", z)
}

func protected() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	y = append(y, "nov")
	z = append(z, "test")
	fmt.Println(x, y, z)
	fmt.Println(len(x), len(y), len(z))
	fmt.Println(cap(x), cap(y), cap(z))
}
