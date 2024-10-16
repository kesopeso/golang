package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Gender int

//go:generate stringer -type=Gender
const (
	Male Gender = iota
	Female
)

type Person struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      Gender  `json:"gender"`
	Probability float64 `json:"probability"`
}

func (g Gender) MarshalJSON() ([]byte, error) {
	switch g {
	case Male:
		return []byte(`"male"`), nil
	case Female:
		return []byte(`"female"`), nil
	default:
		return nil, errors.New("unsupported gender")
	}
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var err error

	switch string(b) {
	case `"male"`:
		*g = Male
	case `"female"`:
		*g = Female
	default:
		err = errors.New("unsupported gender")
	}
	return err
}

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	checkName := "Luka"
	endpoint := "https://api.genderize.io/?name=" + checkName

	request, err := http.NewRequest(http.MethodGet, endpoint, nil /* body reader */)
	if err != nil {
		fmt.Println("Error occured while creating request", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error occured while getting the response", err)
		return
	}

	defer response.Body.Close()

	person, err := readResponse(response.Body)
	if err != nil {
		fmt.Println("Error occured while reading the response body", err)
		return
	}
	fmt.Println("This is the response", person.Count, person.Name, person.Gender, person.Probability)

	encodedPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Print("Cannot encode!")
		return
	}
	fmt.Println("Encoded back to", string(encodedPerson))
}

func readResponse(r io.Reader) (Person, error) {
	var person Person

	err := json.NewDecoder(r).Decode(&person)
	if err != nil {
		var zero Person
		return zero, err
	}

	return person, nil
}
