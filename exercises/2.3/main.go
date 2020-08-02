package main

import (
	"fmt"

	"github.com/joerx/gobook/exercises/2.3/popcount1"
)

func main() {
	numbers := []uint64{1, 2, 3, 7, 8, 100, 123, 4530}
	for _, num := range numbers {
		fmt.Printf("PopCount(%d) = %d\n", num, popcount1.PopCount(num))
	}
}
