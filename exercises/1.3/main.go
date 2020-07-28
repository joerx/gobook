package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	args := []string{"foo", "bar", "baz"}

	benchmark("echo3", func() {
		echo3(args)
	})

	benchmark("echo2", func() {
		echo2(args)
	})

	benchmark("echo1", func() {
		echo1(args)
	})
}

func benchmark(name string, fn func()) {
	start := time.Now().UnixNano()
	fn()
	end := time.Now().UnixNano()
	fmt.Printf("%s: took %d ns\n", name, end-start)
}

func echo1(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args, " ")
}
