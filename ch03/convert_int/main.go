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

  f := 3.14
  i := int(f)
  fmt.Println(f, i) // 3

  mask := 0666
  fmt.Printf("%#o %#[1]b %[1]d\n", mask)

  oe := 'ö'
  nl := '\n'
  fmt.Printf("%q %[1]d %#[1]x\n", oe)
  fmt.Printf("%q %[1]d %#[1]x\n", nl)
}
