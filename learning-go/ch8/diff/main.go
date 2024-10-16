package main

import "fmt"

func main() {
	pair1 := Pair[StringInt]{
		Val1: StringInt(19),
		Val2: StringInt(5),
	}

	pair2 := Pair[StringInt]{
		Val1: StringInt(2),
		Val2: StringInt(10),
	}

	result := FindCloser(pair1, pair2)
	fmt.Println(result)
}

type StringInt int

func (si StringInt) String() string {
	return fmt.Sprint(int(si))
}

func (si StringInt) Diff(compare StringInt) float64 {
	return float64(int(si) - int(compare))
}

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1 Pair[T], pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}
