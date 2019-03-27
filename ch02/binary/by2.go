package main

import (
	"fmt"
)

func main() {
	var val int
	for i := 1; i <= 50; i++ {
		val = i << 1
		fmt.Println(val)
	}
}
