package main

import (
	_ "embed"
	"fmt"
	"os"
)

var (
	//go:embed french_rights.txt
	frenchTranslation string

	//go:embed english_rights.txt
	englishTranslation string

	//go:embed spanish_rights.txt
	spanishTranslation string
)

var languageMapper = map[string]string{
	"french":  frenchTranslation,
	"english": englishTranslation,
	"spanish": spanishTranslation,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage ./<prog_name> <language=french|english|spanish>")
		os.Exit(0)
	}

	translation, ok := languageMapper[os.Args[1]]
	if !ok {
		fmt.Printf("Invalid language. Got: %s, allowed: english|spanish|french\n", os.Args[1])
		os.Exit(0)
	}

	fmt.Println(translation)
}
