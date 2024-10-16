package main

import (
	"fmt"
	"math"
	"sync"
)

func createMap() map[int]float64 {
	result := map[int]float64{}
	for i := 0; i < 100_000; i++ {
		result[i] = math.Sqrt(float64(i))
	}
	return result
}

var cachedMapInitializer = sync.OnceValue(createMap)

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise3() {
	myMap := cachedMapInitializer()
	for i := 0; i < 100000; i += 1000 {
		fmt.Println("value", i, myMap[i])
	}
}

func exercise2() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for done := 0; done < 2; {
		select {
		case value, ok := <-ch1:
			if !ok {
				done++
				ch1 = nil
				continue
			}
			fmt.Println("First go routine wrote the value", value, ok)
		case value, ok := <-ch2:
			if !ok {
				done++
				ch2 = nil
				continue
			}
			fmt.Println("Second go routine wrote the value", value, ok)
		}
	}

	fmt.Println("Done!")
}

func exercise1() {
	ch := make(chan int, 20)

	var wg1 sync.WaitGroup
	wg1.Add(2)

	addValuesToChannel := func(values []int) {
		defer wg1.Done()
		for _, v := range values {
			ch <- v
		}
	}

	go addValuesToChannel([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	go addValuesToChannel([]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	go func() {
		wg1.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg2.Wait()
	fmt.Println("Done!")
}
