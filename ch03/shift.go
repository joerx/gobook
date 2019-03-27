package main

import (
  "fmt"
)

func main() {
  var val int32 = 1
  var flag byte = 3
  fmt.Println(val << flag) // works, shifts don't need type to match
  // fmt.Println(val & flag) // error
  // fmt.Println(val + flag) // error
}
