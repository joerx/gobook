package main

import(
	"fmt"
	"flag"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "seperator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
	}
}

func echo(newline bool, seperator string, args []string) error {
	fmt.Fprint(out, strings.Join(args, seperator))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
