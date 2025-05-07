package main

import (
	"acceptance_tests/acceptancetests"
	"log"
	"net/http"
)

func main() {
	httpServer := &http.Server{Addr: ":8081", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
