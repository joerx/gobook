package main

import (
	"fmt"
)

func main() {
	x := "hello!"
	for _, x := range x {
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()
}
