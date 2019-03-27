package main
// Print arguments using strings.Join

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  var start = time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
  fmt.Printf("took %d ns\n", time.Since(start).Nanoseconds())
}
