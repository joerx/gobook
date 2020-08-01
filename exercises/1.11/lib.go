package main

import (
	"fmt"
	"os"
)

func fatal(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
