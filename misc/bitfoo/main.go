package main

import (
	"fmt"
)

func main() {
	cnt := 0
	val := 2423
	for n := val ; n > 0; n = n >> 1 {
		cnt += n & 1
	}

	fmt.Printf("popcount(%b) = %d\n", val, cnt)

	tests := []int {4, 12, 349}

	for _, x := range tests {
		fmt.Printf("%b & %b = %b\n", x, x-1, x&(x-1))
	}

	cnt2 := 0
	for n := val; n > 0; n = n&(n-1) {
		cnt2++
	}

	fmt.Printf("popcount2(%b) = %d\n", val, cnt2)
}
