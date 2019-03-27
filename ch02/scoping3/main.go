package main

import (
  "fmt"
  "os"
)

const fname = "~/.bashrc"

func main() {
  // will not compile: f is scoped inside the if statement
  if f, err := os.Open(fname); err != nil {
    fmt.Println(err)
    return
  }
  // f is not visible here
  f.ReadByte()
  f.Close()
}
