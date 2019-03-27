package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fibo(n)
	fmt.Printf("\rFibonacci of %d is %d\n", n, fibN)
}

func spinner(interval time.Duration) {
	for {
		for _, r := range `_\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(interval)
		}
	}
}

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
