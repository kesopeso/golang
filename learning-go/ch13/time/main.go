package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.ParseDuration("2h45m")
	if err != nil {
		fmt.Println("cannot parse time")
		return
	}
	fmt.Println("parsed time", t)
}
