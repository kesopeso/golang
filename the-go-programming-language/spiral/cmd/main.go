package main

import (
	"log"
	"os"
	"spiral"
)

func main() {
	sm := spiral.NewSpiralMatrix(150, 1000, 3)
	err := sm.Print(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
