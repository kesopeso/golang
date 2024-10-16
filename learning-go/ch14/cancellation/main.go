package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ch := make(chan string)
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	go requestRandomDelay(ctx, cancel, ch)
	go requestRandomStatus(ctx, cancel, ch)
	waitForRequestsToFinish(ctx, ch)
}

func waitForRequestsToFinish(ctx context.Context, resCh <-chan string) {
loop:
	for {
		select {
		case s := <-resCh:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			err := context.Cause(ctx)
			if err != nil {
				fmt.Println("finished with error:", err)
			}
			fmt.Println("finished without errors!")
			break loop
		}
	}
}

func requestRandomDelay(ctx context.Context, cancelCtx context.CancelCauseFunc, ch chan<- string) {
	for {
		res, err := makeRequest(ctx, "http://httpbin.org/delay/1")
		if err != nil {
			cancelCtx(fmt.Errorf("errror in delay: %w", err))
			return
		}

		select {
		case ch <- "success from delay: " + res.Header.Get("date"):
		case <-ctx.Done():
			return
		}
	}
}

func requestRandomStatus(ctx context.Context, cancelCtx context.CancelCauseFunc, ch chan<- string) {
	for {
		res, err := makeRequest(ctx, "http://httpbin.org/status/200,200,200,500")
		if err != nil {
			fmt.Println("error occured in random status request", err)
			cancelCtx(fmt.Errorf("error in status %w", err))
			return
		}
		if res.StatusCode == http.StatusInternalServerError {
			cancelCtx(errors.New("error occured in random status request - internal server error"))
			return
		}

		select {
		case ch <- "success from random status":
		case <-ctx.Done():
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
