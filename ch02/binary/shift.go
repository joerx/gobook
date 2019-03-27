package main

import (
	"fmt"
	"strconv"
)

func main() {
	// res1 := 16 >> 5
	fmt.Printf("%v\n", 16 >> 5)
	fmt.Printf("%v\n", byte(32))
	fmt.Printf("%v\n", 288 >> 8) 				// 32 
	fmt.Printf("%v\n", byte(288 >> 8)) 	// 1
	fmt.Printf("%v\n", byte(2413 >> 8)) 	// 1

	num := int64(412353451) // 0001 1000 1001 0100 0000 0011 1010 1011
	fmt.Printf("%v; %s; %v\n", num, strconv.FormatInt(num, 2), byte(num >> (1*8))) 	// 1
}
