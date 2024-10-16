package main

import (
	"fmt"
)

func main() {
	valOne := OrderableInt(5)
	valTwo := OrderableInt(6)
	valThree := OrderableInt(-1)
	valFour := OrderableInt(34)
	valFive := OrderableInt(8)

	var t *Tree
	t = t.Insert(valOne)
	t = t.Insert(valTwo)
	t = t.Insert(valThree)
	t = t.Insert(valFour)
	t = t.Insert(valFive)

	fmt.Println(t)
}

type OrderableInt int

func (oi OrderableInt) Order(o any) int {
	if oi < o.(OrderableInt) {
		return -1
	}
	if oi > o.(OrderableInt) {
		return 1
	}
	return 0
}

type Orderable interface {
	Order(o any) int
}

type Tree struct {
	left  *Tree
	right *Tree
	val   Orderable
}

func (t *Tree) Insert(v Orderable) *Tree {
	if t == nil {
		return &Tree{val: v}
	}

	switch comp := t.val.Order(v); {
	case comp > 0:
		t.left = t.left.Insert(v)
	case comp < 0:
		t.right = t.right.Insert(v)
	}

	return t
}
