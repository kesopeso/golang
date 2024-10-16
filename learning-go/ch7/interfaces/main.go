package main

import (
	"fmt"
)

type BasicReadWriter struct {
	ReadReturnValue  string
	WillWriteSucceed bool
}

func (brw BasicReadWriter) Read(value string) string {
	return brw.ReadReturnValue
}

func (brw BasicReadWriter) Write(value string) bool {
	return brw.WillWriteSucceed
}

func main() {
	brw := BasicReadWriter{
		ReadReturnValue:  "true",
		WillWriteSucceed: true,
	}

	RunReadWrite(brw)
}

type Reader interface {
	Read(value string) string
}

type Writer interface {
	Write(value string) bool
}

type ReadWriter interface {
	Reader
	Writer
}

func RunReadWrite(rw ReadWriter) {
	output := rw.Read("This is just a test")
	fmt.Println("is output string 'true'", output == "true")

	writeSuccessful := rw.Write("add something")
	fmt.Println("was write successful", writeSuccessful)
}
