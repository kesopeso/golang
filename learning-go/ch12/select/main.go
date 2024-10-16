package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	mainValue := 2
	routineValue := 1

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			wg.Done()
		}()
		ch1 <- routineValue
		fmt.Println("Wrote in routine to ch1")
		mainValueInRoutine := <-ch2
		fmt.Println("Read from ch2 in routine", mainValueInRoutine)
	}()

	routineValue = 45

	ch2Wrote := false
	ch1Read := false

	for !ch2Wrote || !ch1Read {
		select {
		case ch2 <- mainValue:
			fmt.Println("Wrote in main to ch2")
			ch2Wrote = true
		case readRoutineValue := <-ch1:
			fmt.Println("Read from ch1 in main", readRoutineValue)
			ch1Read = true
		}
	}

	fmt.Println("Waiting to read from ch2")
	wg.Wait()

	fmt.Println("All done")
}
