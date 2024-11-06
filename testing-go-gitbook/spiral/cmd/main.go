package main

import (
	"log"
	"os"
	"spiral"
)

func main() {
	ish := spiral.NewImageSpiralHandler(500, 90, os.Stdout)
	sd := spiral.NewSpiralData(1, 500, 3)
	err := spiral.WriteSpiral(ish, sd)
	if err != nil {
		log.Fatal("error occured", err)
	}

}
