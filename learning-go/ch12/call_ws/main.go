package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// TODO: function that calls 3 web services. first send data to 2 services, take result of those 2 and send them to the 3rd one

func main() {
	data, err := handleRequests()
	if err != nil {
		fmt.Println("Some error occured while processing", err)
		os.Exit(1)
	}
	fmt.Println("This is the data we got back", data)
}

func handleRequests() (string, error) {
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	go func() {
		values := getValues()
		time.Sleep(20 * time.Millisecond)
		ch <- strings.Join(values, ",")
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("Timeout achieved!")
	case result := <-ch:
		return result, nil
	}
}

func getValues() []string {
	ch := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
		}()
		// send request to "https://localhost:8080"
		time.Sleep(45 * time.Millisecond)
		someDataOne := "one"
		ch <- someDataOne
	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		// send request to "https://localhost:8080"
		time.Sleep(20 * time.Millisecond)
		someDataTwo := "two"
		ch <- someDataTwo
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// this will wait until first two are done
	values := make([]string, 0, 2)
	for v := range ch {
		values = append(values, v)
	}

	return values
}
