package main

import (
	"acceptance_tests"
	"acceptance_tests/acceptancetests"
	"context"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", acceptancetests.SlowHandler)
	ctx := context.Background()
	httpServer := &http.Server{Addr: ":8080", Handler: mux}

	myServ := acceptance_tests.NewServer(httpServer)
	if err := myServ.ListenAndServe(ctx); err != nil {
		// this will typically happen if our responses aren't written before the ctx deadline, not much can be done here
		log.Fatalf("uh oh, didn't shutdown gracefully, some responses may have been lost %v", err)
	}

	log.Println("shutdown gracefully! all responses were sent.")
}
