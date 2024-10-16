// Package main runs the go_routine part of the chapter 12 exercises.
package main

import (
	"fmt"
)

var goRoutinesBufferSize = 5

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := processConcurrently(x)
	fmt.Println("This is the final result: ", result)
}

func process(val int) int {
	return val * 2
}

func processConcurrently(inVals []int) []int {
	// create the channels
	in := make(chan int, goRoutinesBufferSize)
	out := make(chan int, goRoutinesBufferSize)

	// launch goRoutinesCount number of go routines
	for i := 0; i < goRoutinesBufferSize; i++ {
		fmt.Println("looping, i: ", i)
		go func() {
			for v := range in {
				fmt.Println("Processing value and writing to out, value: ", v)
				out <- process(v)
			}
		}()
	}

	// load the data into the in channel in another go routine
	go func() {
		for _, v := range inVals {
			fmt.Println("Loading data into in, value: ", v)
			in <- v
		}
		fmt.Println("Closing in")
		close(in)
	}()

	// read the data
	outVals := make([]int, len(inVals))
	for i := 0; i < len(inVals); i++ {
		outVals[i] = <-out
		fmt.Println("Reading from out to the out values", outVals[i])
	}

	return outVals
}
