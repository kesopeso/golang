package main

import (
	"acceptance_tests"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response dawg"))
	})
	ctx := context.Background()
	httpServer := &http.Server{Addr: ":8080", Handler: mux}

	myServ := acceptance_tests.NewServer(httpServer)
	if err := myServ.ListenAndServe(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gracefully shoutdown server")
}
