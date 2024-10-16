package main

import (
	"errors"
	"fmt"
)

func main() {
	err := AFunctionThatReturnsAnError()
	if err != nil {
		var myErr MyErr
		if errors.As(err, &myErr) {
			fmt.Println("first error handler")
			fmt.Println(myErr.Codes)
		}

		var coder Coder
		if errors.As(err, &coder) {
			fmt.Println("second error handler")
			fmt.Println(coder.CodesVal())
		}
	}
}

type Coder interface {
	CodesVal() []int
}

type MyErr struct {
	Codes []int
}

func (me MyErr) CodesVal() []int {
	return me.Codes
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func AFunctionThatReturnsAnError() error {
	return MyErr{Codes: []int{1, 2, 3}}
}
