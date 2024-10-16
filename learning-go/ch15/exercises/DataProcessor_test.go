package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMainRunnerFailFileCreation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic did not happen")
		}
	}()
	ch1 := make(chan []byte, 1)
	ch2 := make(chan Result, 1)
	filePath := ""
	addr := ":8080"
	mr := MainRunner{ch1, ch2, filePath, addr}
	mr.Run()
}

// data flow
// 1. send request, http server writes request data to ch1
// 2. DataProcessor listens for ch1 in loop, parses data and writes it to ch2
// 3. WriteData listens on ch2 and writes data to a file
// 4. close ch1
// 5. DataProcessor loop ends and closes ch2
// 6. check the results.txt if it contains the data
func TestMainRunnerRun(t *testing.T) {
	go func() {
		main()
	}()

	// send a test request
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080", bytes.NewBuffer([]byte("1\n*\n5\n5")))
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		t.Fatal("Status code should be StatusAccepted, got", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Cannot read response body", err)
	}
	bodyText := string(body)
	if bodyText != "OK: 1" {
		t.Fatal("Bad response, got", bodyText, ", want", "OK: 1")
	}

	// open results.txt
	fileContent, err := os.ReadFile("results.txt")
	if err != nil {
		t.Fatal("Can't read results.txt", err)
	}
	stringFileContent := string(fileContent)
	if stringFileContent != "1:25\n" {
		t.Error("file content not ok, got", stringFileContent, ", want", "1:25\n")
	}
}

func TestNewControllerHappyPath(t *testing.T) {
	out := make(chan []byte, 1)
	c := NewController(out)

	input := "test data"

	go func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(input)))
		c.ServeHTTP(w, r)
		response := w.Result()
		defer response.Body.Close()
		responseText, err := io.ReadAll(response.Body)
		if err != nil {
			t.Error("cannot decode response", err)
		}
		if string(responseText) != "OK: 1" {
			t.Error("response not matching", string(responseText), input)
		}
	}()

	data := <-out
	if string(data) != input {
		t.Error("output not matching", string(data), input)
	}
}

type prematureEOF struct {
}

func (prematureEOF) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func readStatusCodeAndBody(t *testing.T, w *httptest.ResponseRecorder) (int, string) {
	t.Helper()

	response := w.Result()
	defer response.Body.Close()

	responseText, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error("cannot decode response", err)
	}

	return response.StatusCode, string(responseText)
}

func TestNewControllerConcurrency(t *testing.T) {
	out := make(chan []byte, 1)
	out <- []byte("full channel")
	c := NewController(out)

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 100; i++ {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("test")))
				c.ServeHTTP(w, r)
			}
		}()
	}

	wg.Wait()

	// check num rejected
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	c.ServeHTTP(w, r)

	responseStatusCode, responseText := readStatusCodeAndBody(t, w)
	if responseStatusCode != http.StatusServiceUnavailable {
		t.Error("response status code missmatch", responseStatusCode, http.StatusServiceUnavailable)
	}
	if responseText != "Too Busy: 10001" {
		t.Error("response not matching", responseText, "Too Busy: 10001")
	}

	// check num sent
	data := <-out
	if string(data) != "full channel" {
		t.Error("output not matching", string(data), "full channel")
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("testing ok")))
	c.ServeHTTP(w, r)

	responseStatusCode, responseText = readStatusCodeAndBody(t, w)
	if responseStatusCode != http.StatusAccepted {
		t.Error("response status code missmatch", responseStatusCode, http.StatusAccepted)
	}
	if responseText != "OK: 10002" {
		t.Error("response not matching", responseText, "OK: 10002")
	}

	data = <-out
	if string(data) != "testing ok" {
		t.Error("response not matching", string(data), "testing ok")
	}
}

func TestNewControllerFullChannel(t *testing.T) {
	out := make(chan []byte, 1)
	out <- []byte("full channel")
	c := NewController(out)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("test")))

		c.ServeHTTP(w, r)

		responseStatusCode, responseText := readStatusCodeAndBody(t, w)
		if responseStatusCode != http.StatusServiceUnavailable {
			t.Error("response status code missmatch", responseStatusCode, http.StatusServiceUnavailable)
		}
		if responseText != "Too Busy: 1" {
			t.Error("response not matching", responseText, "Too Busy: 1")
		}
	}()

	wg.Wait()

	data := <-out
	if string(data) != "full channel" {
		t.Error("output not matching", string(data), "full channel")
	}
}

func TestNewControllerBadRequest(t *testing.T) {
	out := make(chan []byte, 1)
	out <- []byte("value to read")
	c := NewController(out)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", prematureEOF{})
		c.ServeHTTP(w, r)

		responseStatusCode, responseText := readStatusCodeAndBody(t, w)
		if responseStatusCode != http.StatusBadRequest {
			t.Error("response status code missmatch", responseStatusCode, http.StatusBadRequest)
		}
		if responseText != "Bad Input" {
			t.Error("response not matching", responseText, "Bad Input")
		}
	}()

	wg.Wait()

	data := <-out
	if string(data) != "value to read" {
		t.Error("output not matching", string(data), "value to read")
	}
}

func FuzzParser(f *testing.F) {
	f.Add([]byte("1\n+\n1\n1"))
	f.Add([]byte("2\n-\n1\n1"))

	f.Fuzz(func(t *testing.T, in []byte) {
		out1, err := parser(in)
		if err != nil {
			t.Skip("handled error")
		}

		in = []byte(out1.Id + "\n" + out1.Op + "\n" + fmt.Sprintf("%d", out1.Val1) + "\n" + fmt.Sprintf("%d", out1.Val2))
		out2, err := parser(in)
		if err != nil {
			t.Error("fuzz error", err)
		}
		if diff := cmp.Diff(out1, out2); diff != "" {
			t.Error(diff)
		}
	})
}

func TestDataProcessorHappyPaths(t *testing.T) {
	data := []struct {
		name   string
		input  []byte
		result Result
	}{
		{"addition", []byte("1\n+\n1\n1"), Result{"1", 2}},
		{"substraction", []byte("2\n-\n1\n1"), Result{"2", 0}},
		{"multiply", []byte("3\n*\n4\n5"), Result{"3", 20}},
		{"division", []byte("4\n/\n10\n5"), Result{"4", 2}},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			in := make(chan []byte)
			out := make(chan Result)

			go DataProcessor(in, out)
			in <- d.input
			result := <-out
			close(in)

			if result.Id != d.result.Id {
				t.Errorf("ids dont match, want: %s, got: %s", d.result.Id, result.Id)
			}
			if result.Value != d.result.Value {
				t.Errorf("values dont match, want: %d, got: %d", d.result.Value, result.Value)
			}
		})
	}
}

func TestDataProcessorErrorPaths(t *testing.T) {
	data := []struct {
		name  string
		input []byte
	}{
		{"invalid_val_1", []byte("1\n+\ninvalid\n1")},
		{"invalid_val_2", []byte("2\n-\n1\ninvalid")},
		{"invalid_operator", []byte("3\n%\n4\n5")},
		{"invalid_line_count", []byte("3\n%\n4")},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			in := make(chan []byte)
			out := make(chan Result)

			go DataProcessor(in, out)
			in <- d.input
			close(in)

			result, ok := <-out
			if ok {
				t.Error("out channel is still opened")
			}
			var zero Result
			if diff := cmp.Diff(zero, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	data := []Result{
		{"1", 10},
		{"2", 5},
		{},
	}

	want := "1:10\n2:5\n:0\n"

	in := make(chan Result)
	var b bytes.Buffer

	go WriteData(in, &b)

	for _, d := range data {
		in <- d
	}

	close(in)

	got := b.String()
	if got != want {
		t.Errorf("write error, want: %s, got: %s", want, got)
	}
}
