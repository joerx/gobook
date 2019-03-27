package main

import (
  "fmt"
)

func f() {}

var g = "g"

func main() {
  f := "f"

  fmt.Printf("f: %v\n", f)
  fmt.Printf("g: %v\n", g)
}
