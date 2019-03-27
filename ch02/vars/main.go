package main

import (
	"fmt"
)

var foo1 string = "bar1" 	// full declaration, but redundant
var foo2 = "bar2" 				// type inference
var foo4 string 					// type only, defaults to zero-type (empty string)

const foo5 = "bar5" 			// constant

func main() {
	foo3 := "bar3" // short form, only for local vars
	fmt.Println(foo1)
	fmt.Println(foo2)
	fmt.Println(foo3)
	fmt.Println(len(foo4))
	fmt.Println(foo5)
}
