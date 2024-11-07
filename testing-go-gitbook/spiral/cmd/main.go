package main

import (
	"log"
	"os"
	"spiral"
)

func main() {
	ish := spiral.NewImageSpiralHandler(99, 100, os.Stdout)
	sd := spiral.NewSpiralData(1, 1, 1)
	err := spiral.WriteSpiral(ish, sd)
	if err != nil {
		log.Fatal("error occured", err)
	}

}
