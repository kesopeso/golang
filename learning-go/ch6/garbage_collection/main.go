package main

import (
	"fmt"
	"runtime"
	"time"
)

type A struct {
	b *B
}

type B struct {
	c *C
}

type C struct {
	field string
}

func makeAPointer() *A {
	// the data pointed to by a, a.b, and a.b.c are allocated on the heap
	a := &A{&B{&C{"hello"}}}
	// this function specifies a function to run when something is garbage collected
	runtime.SetFinalizer(a.b.c, func(c *C) { fmt.Println("a.b.c with value", c.field, "is garbage collected") })
	return a
}

func main() {
	aPointer := makeAPointer()
	// force a garbage collection
	// aPointer is still pointing to the data on the heap, so there's no garbage yet
	runtime.GC()
	// give the finalizer a chance to run (it won't, because there's no garbage yet)
	time.Sleep(20)
	fmt.Println(aPointer)
	// setting aPointer to nil makes the data that was pointed to by aPointer into garbage
	aPointer = nil
	fmt.Println(aPointer)
	// force a garbage collection
	// aPointer no longer points to the data on the heap, so it's now garbage
	runtime.GC()
	// give the finalizer a chance to run (it will)
	time.Sleep(20)
}
