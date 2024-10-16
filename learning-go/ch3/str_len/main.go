package main

import "fmt"

func main() {
	var s string = "Hello 🌞"
	var sliceByte = []byte(s)
	var sliceRune = []rune(s)
	sliceByte[0] = 92
	sliceRune[1] = 65
	newFromByte := string(sliceByte)
	newFromRune := string(sliceRune)
	fmt.Println("length", len(s))
	fmt.Println("byte changed", sliceByte)
	fmt.Println("rune changed", sliceRune)
	fmt.Println("string", s)
	fmt.Println("byte string", newFromByte)
	fmt.Println("rune string", newFromRune)
}
