package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// make is a built in for creating maps, slices and channels (and nothing else)
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
