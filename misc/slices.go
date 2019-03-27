package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]string, 3)

	fmt.Println(len(slice))

	for i := 0; i < 5; i++ {
		// slice[i] = strconv.Itoa(i) // runs out of range!
		slice = append(slice, strconv.Itoa(i)) // reslices or allocates new array
	}

	fmt.Println(len(slice))
}
