package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Println(PopCount(100000))
}

func PopCount(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
