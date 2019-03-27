package main

import (
  "fmt"
  "math"
)

func main() {
  var f float32 = 1 << 24
  fmt.Println(f)
  fmt.Println(f == f+1)  // true!

  for x := 0; x < 8; x++ {
    fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
  }

  var z float64
  fmt.Println(z, -z, 1/z, -1/z, z/z)

  var nan = math.NaN()
  fmt.Println(nan == nan, nan < nan, nan > nan) // all is false
}
