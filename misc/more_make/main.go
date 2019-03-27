package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 0, 8)
	max := cap(slice) + 1

	fmt.Printf("len(slice)=%d\n", len(slice)) // 0
	fmt.Printf("cap(slice)=%d\n", cap(slice)) // 8

	for i := 0; i < max; i++ {
		slice = append(slice, i)
	}

	fmt.Printf("len(slice)=%d\n", len(slice)) // 9
	fmt.Printf("len(slice)=%d\n", cap(slice)) // 16 (initial length * 2)
}
