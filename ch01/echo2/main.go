package main
// Like echo1, but with fancier syntax
// 
// variable declarations:
// s := ""           - short variable declaration, type inferred from value, not on package level
// var s string      - explicit type declaration, value will be ""
// var s = ""        - type inference
// var s string = "" - cod, necessary if value is not of declared type

import (
  "os"
  "fmt"
)

func main() {
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
}
