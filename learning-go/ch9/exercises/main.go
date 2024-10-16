package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID      = regexp.MustCompile(`\w{4}-\d{3}`)
	ErrInvalidId = errors.New("invalid ID (SENTINEL ERROR)")
)

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v \n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		if err != nil {
			switch err := err.(type) {
			case interface{ Unwrap() []error }:
				var messages []string
				for _, e := range err.Unwrap() {
					messages = append(messages, processError(e, emp))
				}
				message += ", multiple errors: " + strings.Join(messages, ", ")
			default:
				message += ", error: " + processError(err, emp)
			}
		}
		fmt.Println(message)
	}
}

type EmptyFieldErr struct {
	FieldName string
}

func (efe EmptyFieldErr) Error() string {
	return efe.FieldName
}

func processError(err error, emp Employee) string {
	if errors.Is(err, ErrInvalidId) {
		return fmt.Sprintf("invalid ID '%s'", emp.ID)
	}
	var emptyFieldError EmptyFieldErr
	if errors.As(err, &emptyFieldError) {
		return fmt.Sprintf("field '%s' is empty", emptyFieldError)
	}
	return fmt.Sprintf("generic error - %v", err)
}

func ValidateEmployee(e Employee) error {
	var errs []error

	if len(e.ID) == 0 {
		errs = append(errs, errors.New("missing ID"))
	}

	if !validID.MatchString(e.ID) {
		errs = append(errs, ErrInvalidId)
	}

	if len(e.FirstName) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "FirstName"})
	}

	if len(e.LastName) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "LastName"})
	}

	if len(e.Title) == 0 {
		errs = append(errs, EmptyFieldErr{FieldName: "Title"})
	}

	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errors.Join(errs...)
	}
}
