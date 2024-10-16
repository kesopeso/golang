package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	data := `
        {"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
    `

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	dec := json.NewDecoder(strings.NewReader(data))

	type personJson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var el personJson
	allInputs := make([]personJson, 0, 3)

	for {
		err := dec.Decode(&el)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println("Decoded a line", el)
		allInputs = append(allInputs, el)
	}

	for _, input := range allInputs {
		err := enc.Encode(input)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Encoded back everything")
	fmt.Print(b.String())
}
