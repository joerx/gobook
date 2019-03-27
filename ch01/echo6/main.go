package main

// Print args, with index and one per line

import (
	"fmt"
	"os"
)

func main() {
	s, nl := "", ""
	for idx, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s %d: %s", nl, idx, arg)
		nl = "\n"
	}
	fmt.Println(s)
}
