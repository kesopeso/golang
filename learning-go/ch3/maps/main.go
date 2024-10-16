package main

import (
	"fmt"
	"maps"
)

func main() {
	commaOkIdiom()
	mapsEqual()
	simulateSets()
	simulateStructSets()
}

func simulateStructSets() {
	printFunctionInfo("Struct sets simulation")

	structSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 100, 3, 9, 1, 2, 10}
	for _, v := range vals {
		structSet[v] = struct{}{}
	}

	fmt.Println(len(vals), len(structSet))
	fmt.Println(structSet[5])
	fmt.Println(structSet[500])
	if _, ok := structSet[100]; ok {
		fmt.Println("100 is in the struct set")
	} else {
		fmt.Println("100 is not in the struct set")
	}
}

func simulateSets() {
	printFunctionInfo("Sets simulation")

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is not in the set")
	}
}

func printFunctionInfo(info string) {
	fmt.Println()
	fmt.Println(info)

	var delimiter string
	for range info {
		delimiter += "="
	}

	fmt.Println(delimiter)
}

func mapsEqual() {
	mapOne := map[string]int{
		"buffer": 1,
		"two":    2,
	}

	mapTwo := map[string]int{
		"two": 2,
		"one": 1,
	}

	fmt.Println("are maps equal:", maps.Equal(mapOne, mapTwo))
}

func commaOkIdiom() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}
