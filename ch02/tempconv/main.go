package main

import (
	"fmt"

	"github.com/joerx/gobook/ch02/tempconv/conv"
)

func main() {
	fmt.Printf("%g°F\n", conv.CToF(conv.FreezingC))
	fmt.Printf("%g°F\n", conv.CToF(conv.BoilingC))
	fmt.Printf("%g°C\n", conv.Celsius(10)+conv.Celsius(20))
	fmt.Printf("%v\n", conv.Celsius(100) == conv.BoilingC)
	fmt.Printf("%v\n", conv.Fahrenheit(conv.BoilingC)) // !!
	fmt.Printf("%s\n", conv.BoilingC)
	fmt.Printf("%s\n", conv.Fahrenheit(100))
	fmt.Printf("%s\n", conv.BoilingC.ToF())
	fmt.Printf("%s\n", conv.Fahrenheit(100).ToC())
}
