package main

import (
  "fmt"
  "os"
)

const fname = "/Users/jhenning/.bash_profile"

func main() {
  // open file
  f, err := os.Open(fname)
  if err != nil {
    fmt.Fprintf(os.Stderr, "An error occured: %v\n", err)
    return
  }

  // get file size
  stats, err := f.Stat()
  if err != nil {
    fmt.Fprintf(os.Stderr, "An error occured: %v\n", err)
    return
  }

  // allocate buffer given size, read file into buffer
  buf := make([]byte, stats.Size())
  _, err = f.Read(buf)

  if err != nil {
    fmt.Fprintf(os.Stderr, "An error occured: %v\n", err)
    return
  }

  // convert to string
  fmt.Println(string(buf))
}
