package main

import (
	"fmt"
	"sync"
)

var slowParser SlowParser
var once sync.Once

var initAnotherSlowParser = sync.OnceValue(InitParser)

func main() {
	result := ParseReallySlowly("some data, ")
	fmt.Println(result)
	result = ParseReallySlowly("something new")
	fmt.Println(result)
}

type SlowParser interface {
	Parse(string) string
}

type ReallySlowParser struct {
}

func (rsp ReallySlowParser) Parse(data string) string {
	return data + data
}

func InitParser() ReallySlowParser {
	return ReallySlowParser{}
}

func ParseReallySlowly(data string) string {
	once.Do(func() {
		fmt.Println("Initializing really slow parser")
		slowParser = InitParser()
	})
	newParser := initAnotherSlowParser()
	newData := newParser.Parse(data)
	return slowParser.Parse(newData)
}
