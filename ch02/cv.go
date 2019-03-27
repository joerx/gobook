package main

import (
	"gobook/ch02/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("%g°F\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("%g°F\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("%g°C\n", tempconv.Celsius(10)+tempconv.Celsius(20))
	fmt.Printf("%v\n", tempconv.Celsius(100) == tempconv.BoilingC)
	fmt.Printf("%v\n", tempconv.Fahrenheit(tempconv.BoilingC)) // !!
	fmt.Printf("%s\n", tempconv.BoilingC)
	fmt.Printf("%s\n", tempconv.Fahrenheit(100))
	fmt.Printf("%s\n", tempconv.BoilingC.ToF())
	fmt.Printf("%s\n", tempconv.Fahrenheit(100).ToC())
}
