package main

// Prints its commandline arguments into a single line

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var start = time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
  fmt.Printf("took %dns\n", time.Since(start).Nanoseconds())
}
