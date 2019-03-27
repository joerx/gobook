package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		mod2 := i % 2
		and1 := i & 1
		fmt.Printf("%d mod 2 = %d, %d & 1 = %d, %b\n", i, mod2, i, and1, mod2 - and1 + 1)
	}
}
