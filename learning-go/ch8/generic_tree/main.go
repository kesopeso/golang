package main

import (
	"cmp"
	"fmt"
)

func main() {
	t := NewTree(cmp.Compare[string])
	t.Add("treniraj")
	t.Add("migaj")
	t.Add("ok")
	t.Add("mahir")
	fmt.Println(t.Contains("ok"))

	tP := NewTree(Person.Compare)
	tP.Add(Person{"Johnny", 34})
	tP.Add(Person{"Blake", 22})
	fmt.Println(tP.Contains(Person{"Johnny", 34}))
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Compare(p1 Person) int {
	out := cmp.Compare(p.Name, p1.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, p1.Age)
	}
	return out
}

type OrderFunc[T any] func(T, T) int

type Node[T any] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) Add(f OrderFunc[T], val T) *Node[T] {
	if n == nil {
		return &Node[T]{val: val}
	}

	switch orderResult := f(n.val, val); {
	case orderResult <= -1:
		n.left = n.left.Add(f, val)
	case orderResult >= 1:
		n.right = n.right.Add(f, val)
	}

	return n
}

func (n *Node[T]) Contains(f OrderFunc[T], val T) bool {
	if n == nil {
		return false
	}

	switch orderResult := f(n.val, val); {
	case orderResult <= -1:
		return n.left.Contains(f, val)
	case orderResult >= 1:
		return n.right.Contains(f, val)
	}

	return true
}

type Tree[T any] struct {
	root *Node[T]
	f    OrderFunc[T]
}

func (t *Tree[T]) Add(val T) {
	t.root = t.root.Add(t.f, val)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.root.Contains(t.f, val)
}

func NewTree[T any](f OrderFunc[T]) *Tree[T] {
	return &Tree[T]{f: f}
}
