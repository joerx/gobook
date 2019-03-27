package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("line %d\n", i)
	}

	i := 0
	for i < 10 {
		i++
		fmt.Printf("line %d\n", i)
	}
}
