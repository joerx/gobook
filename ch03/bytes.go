package main

import (
  "fmt"
)

func main() {
  var str = "%08b = %2d\n"
  var x byte = 1<<1 | 1<<5 // 00100010
  var y byte = 1<<1 | 1<<2 // 00000110 (same as 3<<1?)

  fmt.Printf(str, x, x)       // 00100010 34
  fmt.Printf(str, y, y)       // 00000110  6
  fmt.Printf(str, x&y, x&y)   // 00000010  2
  fmt.Printf(str, x|y, x|y)   // 00100110 38
  fmt.Printf(str, x^y, x^y)   // 00100100 36
  fmt.Printf(str, x&^y, x&^y) // 00100000 32
  fmt.Printf(str, x<<1, x<<1) // 01000100 68
  fmt.Printf(str, x>>1, x>>1) // 00010001 17

  // membership test
  for i := uint8(0); i < 8; i++ {
    if x&(1<<i) > 0 {
      fmt.Println(i)
    }
  }
}
