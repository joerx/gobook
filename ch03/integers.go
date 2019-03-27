package main

import (
  "fmt"
)

const str = "%v\t%08b\n"

func main() {
  // Representation of integers in binary
  // Computers internally use 2's complement to represent binary numbers
  //
  // I.e. i = -5 = 0000 0101
  //     ^i = 1111 1010,
  //     ^i+1 = 1111 1011
  //
  // But printing the binary representation of i is odd:

  var i int8 = -5
  fmt.Printf(str, i, i) // prints -101, not 1111 1011

  // However, converting it into an unsigned int, gives 1111 1011, so
  // internally it was represented in 2's complement all along?

  var u uint8 = uint8(i)
  fmt.Printf(str, u, u) // prints 1111 1011 (which is 251 in this case)


  var i2 int8 = 127
  fmt.Println(i2, i2+1, i2*i2)

  var i3 = 127
  fmt.Println(i3*i3, i3<<7-i3) // 127*127 = 127*(128-1) = 127*2^7 - 127 = 127<<7 - 127
}
