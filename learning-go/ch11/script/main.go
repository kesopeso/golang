package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Running directory check")

	executable, err := os.Executable()
	if err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}

	exeDir := filepath.Dir(executable)
	fmt.Printf("Dir of the current executable: %s", exeDir)
}
