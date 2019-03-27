package main

import (
  "fmt"
)

func main() {
  var apples uint16 = 65535
  var oranges uint32 = 4294967295

  fmt.Println(byte(apples)) // 2^8
  fmt.Println(apples & 255) // 2^8 (bitmask)
  fmt.Println(uint16(oranges)) // 2^16
  fmt.Println(int16(apples))  // -1
  fmt.Println(apples + 1)  // 0 (overflows)
}
