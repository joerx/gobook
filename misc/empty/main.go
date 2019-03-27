package main

import (
	"fmt"
)

func main() {
	arr1 := []string {}
	arr2 := []string {""}
	arr3 := []string {"foo", "bar", ""}

	fmt.Printf("len(%v) = %d\n", arr1, len(arr1)) // len([]) = 0
	fmt.Printf("len(%v) = %d\n", arr2, len(arr2)) // len([]) = 1
	fmt.Printf("len(%v) = %d\n", arr3, len(arr3)) // len([foo bar ]) = 1

	// better use %q:
	fmt.Printf("len(%q) = %d\n", arr2, len(arr2)) // len([""]) = 1
}
