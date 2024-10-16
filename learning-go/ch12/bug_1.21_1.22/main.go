package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8}
	ch := make(chan int, len(a))
	for _, v := range a {
		fmt.Println("testiram", v)
		go func(cur int) {
			fmt.Println("writing", cur)
			ch <- cur * 2
			fmt.Println("wrote", cur)
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
