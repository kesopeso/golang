package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte("hello " + name + "\n"))
	})

	mux.HandleFunc("/goodbye/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte("goodbye " + name + "\n"))
	})

	passwordMux := http.NewServeMux()
	passwordMux.HandleFunc("/password", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is /auth/password path"))
	})

	mux.Handle("/auth/", http.StripPrefix("/auth", passwordMux)) // trailing / is mandatory, it denotes that this is a "sub" handler

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  2 * time.Second,
		Handler:      mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
