package main
// Print args with square brackets

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println(os.Args[1:])
}
