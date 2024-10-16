package main

import (
	"fmt"
	"strconv"
)

func main() {
	in := []string{"1", "2", "3", "4", "5"}

	mapOut := Map(in, func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	})
	fmt.Println(mapOut)

	reduceOut := Reduce(mapOut, 0, func(acc int, i int) int {
		return acc + i
	})
	fmt.Println(reduceOut)
}

func Map[TIn, TOut any](in []TIn, f func(TIn) TOut) []TOut {
	out := make([]TOut, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func Reduce[TIn, TOut any](in []TIn, initializer TOut, f func(TIn, TOut) TOut) TOut {
	out := initializer
	for _, v := range in {
		out = f(v, out)
	}
	return out
}
