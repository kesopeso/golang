package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	child, cancel2 := context.WithTimeout(parent, 15*time.Second)
	defer cancel()
	defer cancel2()

	start := time.Now()
	<-child.Done()
	elapsed := time.Since(start)
	fmt.Println("Everything finished", elapsed)
}
