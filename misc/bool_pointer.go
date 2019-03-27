package main

import (
	"fmt"
)

func main() {
	val := true
	p := &val

	fmt.Printf("%t\n", val)
	fmt.Printf("%t\n", *p)
	fmt.Printf("%t\n", !*p)
}
