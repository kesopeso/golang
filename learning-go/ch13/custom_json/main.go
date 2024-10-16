package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type RFC822ZTime struct {
	time.Time
}

func (r RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := r.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (r *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	date, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}

	*r = RFC822ZTime{Time: date}
	return nil
}

// INFO: for me this approach is easier, since you also know in what format time is encoded
type OrderTimeSpecific struct {
	ID          string      `json:"id"`
	DateOrdered RFC822ZTime `json:"dateOrdered"`
}

// INFO: more complex, but you don't need to worry about the format of time when coding
type OrderTimeGeneral struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"dateOrdered"`
}

// INFO: Using Dup is necessary, because if we'd use OrderTimeGeneral inside tmp struct
// that would cause json.Marshal to do an infinite loop, since there is a custom MarshalJSON
// method defined on OrderTimeGeneral
func (o OrderTimeGeneral) MarshalJSON() ([]byte, error) {
	type Dup OrderTimeGeneral

	tmp := struct {
		DateOrdered string `json:"dateOrdered"`
		Dup
	}{
		Dup: Dup(o),
	}

	tmp.DateOrdered = o.DateOrdered.Format(time.RFC822Z)

	b, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// INFO: Using Dup is necessary, because if we'd use OrderTimeGeneral inside tmp struct
// that would cause json.Marshal to do an infinite loop, since there is a custom MarshalJSON
// method defined on OrderTimeGeneral
func (o *OrderTimeGeneral) UnmarshalJSON(b []byte) error {
	type Dup OrderTimeGeneral

	tmp := &struct {
		DateOrdered string `json:"dateOrdered"`
		*Dup
	}{
		Dup: (*Dup)(o),
	}

	err := json.Unmarshal(b, tmp)
	if err != nil {
		return err
	}

	o.DateOrdered, err = time.Parse(time.RFC822Z, tmp.DateOrdered)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	orders := `
        {"id": "1", "dateOrdered": "01 Jan 06 15:04 -0600"}
        {"id": "2", "dateOrdered": "02 Jan 07 15:04 -0600"}
        {"id": "3", "dateOrdered": "03 Jan 08 15:04 -0600"}
        {"id": "4", "dateOrdered": "04 Jan 09 15:04 -0600"}
        {"id": "5", "dateOrdered": "05 Jan 10 15:04 -0600"}
        {"id": "6", "dateOrdered": "06 Jan 11 15:04 -0600"}
    `

	reader := strings.NewReader(orders)
	dec := json.NewDecoder(reader)

	var order OrderTimeGeneral
	ordersDecoded := make([]OrderTimeGeneral, 0, 6)
	//var order OrderTimeSpecific
	//ordersDecoded := make([]OrderTimeSpecific, 0, 6)

	for dec.More() {
		err := dec.Decode(&order)
		if err != nil {
			fmt.Println("Cannot decode entry", err)
			continue
		}
		ordersDecoded = append(ordersDecoded, order)
	}

	fmt.Println("All orders have been decoded", ordersDecoded)

	res, err := json.Marshal(ordersDecoded)
	if err != nil {
		fmt.Println("Cannot encode orders", err)
		return
	}
	fmt.Println("Orders encoded", string(res))
}
