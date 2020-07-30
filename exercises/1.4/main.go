package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stringSet map[string]nix

type nix struct{}

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string]stringSet)

	files := os.Args[1:]

	for _, arg := range files {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		} else {
			countLines(file, counts, foundIn)
			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s (%s)\n", n, line, strings.Join(toSlice(foundIn[line]), ", "))
		}
	}
}

func toSlice(s stringSet) []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

// insert is like append() but for string sets
func insert(s stringSet, elem string) stringSet {
	if s == nil {
		s = make(stringSet)
	}
	s[elem] = nix{}
	return s
}

func countLines(file *os.File, counts map[string]int, foundIn map[string]stringSet) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			continue
		}
		counts[line]++
		foundIn[line] = insert(foundIn[line], file.Name())
	}
}
