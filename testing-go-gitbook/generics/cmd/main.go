package main

import (
	"fmt"
	"generics"
)

func main() {
	first := generics.Sum([]int{1, 2, 3})
	fmt.Println(first)
	second := generics.SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{5, 4, 3, 2, 1})
	fmt.Println(second)
}
