package main

import (
	"errors"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func NewPressureGauge(limit int) *PressureGauge {
	return &PressureGauge{ch: make(chan struct{}, limit)}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		return errors.New("No more capacity")
	}
}

func doSomeLongTask(text string) string {
	time.Sleep(2 * time.Second)
	return text
}

func main() {
	pg := NewPressureGauge(5)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doSomeLongTask("Testiram ce dela")))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
