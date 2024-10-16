package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Result struct {
	Id    string
	Value int
}

type Input struct {
	Id   string
	Op   string
	Val1 int
	Val2 int
}

func parser(data []byte) (Input, error) {
	// parse the data
	lines := bytes.Split(data, []byte("\n"))
	// each entry is line 1 id, line 2 operator, line 3 num 1, line 4 num2

	if len(lines) != 4 {
		return Input{}, errors.New("lines count is not 4")
	}

	id := string(lines[0])
	op := string(lines[1])
	val1, err := strconv.Atoi(string(lines[2]))
	if err != nil {
		return Input{}, err
	}
	val2, err := strconv.Atoi(string(lines[3]))
	if err != nil {
		return Input{}, err
	}
	return Input{
		Id:   id,
		Op:   op,
		Val1: val1,
		Val2: val2,
	}, nil
}

func DataProcessor(in <-chan []byte, out chan<- Result) {
	for data := range in {
		input, err := parser(data)
		if err != nil {
			continue
		}
		var calc int
		switch input.Op {
		case "+":
			calc = input.Val1 + input.Val2
		case "-":
			calc = input.Val1 - input.Val2
		case "*":
			calc = input.Val1 * input.Val2
		case "/":
			calc = input.Val1 / input.Val2
		default:
			continue
		}
		// sum numbers in the data
		// write to another channel
		result := Result{
			Id:    input.Id,
			Value: calc,
		}
		out <- result
	}
	close(out)
}

func WriteData(in <-chan Result, w io.Writer) {
	for r := range in {
		// write the output data to writer
		// each line is id:result
		w.Write([]byte(fmt.Sprintf("%s:%d\n", r.Id, r.Value)))
	}
}

func NewController(out chan []byte) http.Handler {
	var numSent int
	var numRejected int
	var l sync.Mutex
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Lock()
		numSent++
		l.Unlock()
		// take in data
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Input"))
			return
		}
		// write it to the queue in raw format
		select {
		case out <- data:
			// success!
		default:
			// if the channel is backed up, return an error
			l.Lock()
			numRejected++
			l.Unlock()
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Too Busy: " + strconv.Itoa(numRejected)))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("OK: " + strconv.Itoa(numSent)))
	})
}

type MainRunner struct {
	ch1         chan []byte
	ch2         chan Result
	resultsPath string
	addr        string
}

func (mr MainRunner) Run() error {
	go DataProcessor(mr.ch1, mr.ch2)
	f, err := os.Create(mr.resultsPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	go WriteData(mr.ch2, f)

	err = http.ListenAndServe(mr.addr, NewController(mr.ch1))
	return err
}

func main() {
	MainRunner{
		make(chan []byte, 100),
		make(chan Result, 100),
		"results.txt",
		":8080",
	}.Run()
}
