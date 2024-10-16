package main

import (
	"fmt"
	"strconv"
)

func main() {
	exercise1()
	exercise2()
	exercise31()
	exercise32()
}

type IntOrFloat interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uintptr | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func DoubleGeneric[T IntOrFloat](in T) T {
	return in * 2
}

func exercise1() {
	fmt.Println("Exercise 1")

	num1 := 10.23
	num2 := 100000000000.2321323
	num3 := 23
	num4 := -2

	fmt.Println(DoubleGeneric(num1))
	fmt.Println(DoubleGeneric(num2))
	fmt.Println(DoubleGeneric(num3))
	fmt.Println(DoubleGeneric(num4))
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

func PrintPrintable[T Printable](in T) {
	fmt.Println(in)
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return "number: " + strconv.Itoa(int(pi))
}

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return "number: " + fmt.Sprintf("%f", float64(pf))
}

func exercise2() {
	fmt.Println("Exercise 2")

	num1 := PrintableInt(1000)
	num2 := PrintableInt(35)
	num3 := PrintableFloat(34.34)
	num4 := PrintableFloat(454.22222)

	PrintPrintable(num1)
	PrintPrintable(num2)
	PrintPrintable(num3)
	PrintPrintable(num4)

}

type Node[T comparable] struct {
	fmt.Stringer
	value T
	next  *Node[T]
}

func (n *Node[T]) String() string {
	if n == nil {
		return ""
	}

	delimiter := ""
	if n.next != nil {
		delimiter = " "
	}

	return fmt.Sprint("[", n.value, "]", delimiter, n.next)
}

func (n *Node[T]) Add(v T) *Node[T] {
	if n == nil {
		return &Node[T]{value: v}
	}
	n.next = n.next.Add(v)
	return n
}

// we need to return pointer because we can get nil pointer
// if we modify the nil pointer, then the original will not update
// another solution would be to check for nil pointers!
func (n *Node[T]) Insert(v T, ci int, ai int) *Node[T] {
	// add index is not valid, 0 or bigger valid
	if ai < 0 {
		return n
	}

	// current index is smaller than add index
	if ci < ai {
		// node does not exist, create a node with default value
		if n == nil {
			var zero T
			n = &Node[T]{value: zero}
		}
		// increase the current index and try inserting on next node
		n.next = n.next.Insert(v, ci+1, ai)
		return n
	}

	// current index is add index. define pointer to the next node
	var nextN *Node[T]
	// if current node exists, me move it forward by defining it as a next node
	if n != nil {
		nextN = n
	}

	// for the current node, we return the node we just made
	return &Node[T]{value: v, next: nextN}
}

func (n *Node[T]) Index(v T, ci int) int {
	if n == nil {
		return -1
	}

	if n.value != v {
		return n.next.Index(v, ci+1)
	}

	return ci
}

type List[T comparable] struct {
	fmt.Stringer
	root *Node[T]
}

// add new element to the end of the list
func (l *List[T]) Add(v T) {
	l.root = l.root.Add(v)
}

// insert element at the specific index
func (l *List[T]) Insert(v T, ai int) {
	l.root = l.root.Insert(v, 0, ai)
}

// return the position of the element, -1 if it doesnt exist
func (l *List[T]) Index(v T) int {
	return l.root.Index(v, 0)
}

func (l List[T]) String() string {
	return fmt.Sprint(l.root)
}

func exercise31() {
	fmt.Println("Exercise 3.1")

	l := CreateList[int]()

	l.Add(5)
	l.Add(1)
	l.Add(23)
	l.Add(454)
	l.Add(-1)
	l.Add(200)

	l.Insert(55, 3)
	l.Insert(1000, 9)

	fmt.Println(l)

	fmt.Println(l.Index(5))
	fmt.Println(l.Index(23))
	fmt.Println(l.Index(201))
	fmt.Println(l.Index(55))
	fmt.Println(l.Index(1000))
}

func CreateList[T comparable]() List[T] {
	return List[T]{}
}

type AltNode[T comparable] struct {
	fmt.Stringer
	Value T
	Next  *AltNode[T]
}

func (an AltNode[T]) String() string {
	return fmt.Sprint(an.Value)
}

type AltTree[T comparable] struct {
	fmt.Stringer
	Head *AltNode[T]
	Tail *AltNode[T]
}

func (at AltTree[T]) String() string {
	if at.Head == nil {
		return "<empty list>"
	}

	resultString := ""
	for currentNode := at.Head; currentNode != nil; currentNode = currentNode.Next {
		delimiter := " "
		if currentNode == at.Tail {
			delimiter = ""
		}
		resultString += fmt.Sprint("[", currentNode.Value, "]", delimiter)
	}

	return resultString
}

func (at *AltTree[T]) Add(value T) {
	node := &AltNode[T]{Value: value}

	if at.Head == nil {
		at.Head = node
		at.Tail = node
		return
	}

	at.Tail.Next = node
	at.Tail = node
}

func (at *AltTree[T]) Index(value T) int {
	i := 0
	for currentNode := at.Head; currentNode != nil; currentNode = currentNode.Next {
		if currentNode.Value == value {
			return i
		}
		i++
	}
	return -1
}

// INFO: this is my implementation, it is a bit more complex
// func (at *AltTree[T]) Insert(value T, pos int) {
// 	if pos < 0 {
// 		pos = 0
// 	}

// 	var prevNode *AltNode[T]
// 	for i := 0; i <= pos; i++ {
// 		if i == pos {
// 			newNode := &AltNode[T]{Value: value}
// 			if at.Head == nil {
// 				at.Head = newNode
// 				at.Tail = newNode
// 				return
// 			}
// 			if prevNode == nil {
// 				newNode.Next = at.Head
// 				at.Head = newNode
// 				return
// 			}
// 			if prevNode == at.Tail {
// 				prevNode.Next = newNode
// 				at.Tail = newNode
// 				return
// 			}
// 			newNode.Next = prevNode.Next
// 			prevNode.Next = newNode
// 			return
// 		}
// 		if at.Head == nil {
// 			var zero T
// 			at.Head = &AltNode[T]{Value: zero}
// 			at.Tail = at.Head
// 			prevNode = at.Head
// 			continue
// 		}
// 		if prevNode == at.Tail {
// 			var zero T
// 			prevNode.Next = &AltNode[T]{Value: zero}
// 			at.Tail = prevNode.Next
// 		}
// 		if prevNode == nil {
// 			prevNode = at.Head
// 			continue
// 		}
// 		prevNode = prevNode.Next
// 	}
// }

func (at *AltTree[T]) Insert(value T, pos int) {
	// create new node
	node := &AltNode[T]{Value: value}

	// initialize the list if it is empty
	if at.Head == nil {
		// insert at beginning and finish
		if pos <= 0 {
			at.Head = node
			at.Tail = node
			return
		}

		// insert empty node
		var zero T
		at.Head = &AltNode[T]{Value: zero}
		at.Tail = at.Head
	}

	// handle insert at zero
	if pos <= 0 {
		node.Next = at.Head
		at.Head = node
		return
	}

	// iterate just before insert position
	currentNode := at.Head
	for i := 1; i < pos; i++ {
		// if at tail, insert empty node and update tail
		if currentNode == at.Tail {
			var zero T
			currentNode.Next = &AltNode[T]{Value: zero}
			at.Tail = currentNode.Next
		}
		// update currentNode
		currentNode = currentNode.Next
	}

	// insert the node
	node.Next = currentNode.Next
	currentNode.Next = node
	if currentNode == at.Tail {
		at.Tail = node
	}
}

func exercise32() {
	fmt.Println("Exercise 3.2")

	list := AltTree[int]{}

	list.Add(4)
	list.Insert(100, 5)
	list.Insert(2, 1)
	list.Insert(555, 0)
	list.Insert(555, 0)

	fmt.Println("This is my list:", list)
	fmt.Println("Index of 100", list.Index(100))
	fmt.Println("Index of 2", list.Index(2))
	fmt.Println("Index of 76", list.Index(76))
	fmt.Println("Index of 555", list.Index(555))
}
