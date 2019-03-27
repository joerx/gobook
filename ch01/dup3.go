package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			counts[line]++
		}

		for line, cnt := range counts {
			if cnt > 1 {
				fmt.Printf("%d\t%s\n", cnt, line)
			}
		}
	}
}
