package main

import (
	"fmt"

	"github.com/joerx/gobook/exercises/2.1/pkg/tempconv"
)

func main() {
	var c tempconv.Celsius = 10

	fmt.Println(tempconv.CToF(c))                      // 50 F
	fmt.Println(tempconv.KToC(0))                      // -237.5 C
	fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC)) // 0 K
	fmt.Println(tempconv.KToF(0))                      // -459.67
	fmt.Println(tempconv.FToK(0))                      // 255.37
}
