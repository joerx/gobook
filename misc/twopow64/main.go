package main

import (
  "fmt"
  // "math"
)

// var arr []int64

func init() {
  // arr = make([]int64, )
  fmt.Println("init")
}

func main() {
  toobig := int32((4294967296-1)>>1) // 2^32-1/2
  fmt.Println(toobig)

  // arr := make([]byte, toobig) // takes very long an likely runs out of memory at some point
  // fmt.Println(arr)

  twopow10 := 2 << 9
  twopow32 := 2 << 31
  twopow12 := 2 << 11
  yay      := 2 << 31 >> 10 >> 10

  fmt.Printf("2^10 (%T) \t\t\t= %10[1]d\n", twopow10)
  fmt.Printf("2^32 (%T) \t\t\t= %10[1]d\n", twopow32)
  fmt.Printf("2^12 (%T) \t\t\t= %10[1]d\n", twopow12)
  fmt.Printf("2^32 / 2^10 (%T) \t\t= %10[1]d\n", twopow32 / twopow10)
  fmt.Printf("2^32 / 2^10 / 2^10 (%T) \t= %10[1]d\n", twopow32 / twopow10 / twopow10)
  fmt.Printf("2^32 >> 20 (%T) \t\t= %10[1]d\n", twopow32 >> 20)
  fmt.Printf("2^32 >> 20 (%T) \t\t= %10[1]d\n", yay)
}
