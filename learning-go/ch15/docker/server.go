package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stopServerCh := make(chan os.Signal, 1)
	signal.Notify(stopServerCh, os.Interrupt, syscall.SIGTERM)

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK!"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe err: %v\n", err)
		}
	}()

	fmt.Println("Server started on :8080")

	<-stopServerCh

	fmt.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("Shutdown err: %v\n", err)
	}

	fmt.Println("Server was shutdown successfully.")
}
