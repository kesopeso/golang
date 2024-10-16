package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed help
var helpNoHidden embed.FS

//go:embed help/*
var helpRootHidden embed.FS

//go:embed all:help
var helpAllHidden embed.FS

func main() {
	selection := ""
	filePath := ""

	if len(os.Args) >= 2 {
		selection = os.Args[1]
	}

	if len(os.Args) >= 3 {
		filePath = os.Args[2]
	}

	fsSelection := map[string]embed.FS{
		"none": helpNoHidden,
		"root": helpRootHidden,
		"all":  helpAllHidden,
	}

	selectedFs, ok := fsSelection[selection]
	if !ok {
		fmt.Printf("Invalid selection: %s. Supported: none/root/all\n", selection)
		os.Exit(1)
	}

	printFile(selectedFs, filePath)
}

func printFile(t embed.FS, filePath string) {
	if filePath == "" {
		printHelpFiles(t)
		os.Exit(0)
	}
	data, err := t.ReadFile("help/" + filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func printHelpFiles(t embed.FS) {
	fmt.Println("Contents:")
	_ = fs.WalkDir(t, "help", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			_, fileName, _ := strings.Cut(path, "/")
			fmt.Println(fileName)
		}
		return nil
	})
}
