package main

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

//go:generate stringer -type=Direction
func main() {
	fmt.Println(East.String())
}
