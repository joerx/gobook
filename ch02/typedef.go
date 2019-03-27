package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC = -273.15
	FreezingC     = Celsius(0)
	BoilingC      = Celsius(100)
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c Celsius) ToF() Fahrenheit {
	return CToF(c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}

func main() {
	fmt.Printf("%g°F\n", CToF(FreezingC))
	fmt.Printf("%g°F\n", CToF(BoilingC))
	fmt.Printf("%g°C\n", Celsius(10)+Celsius(20))
	fmt.Printf("%v\n", Celsius(100) == BoilingC)
	fmt.Printf("%v\n", Fahrenheit(BoilingC)) // !!
	fmt.Printf("%s\n", BoilingC)
	fmt.Printf("%s\n", Fahrenheit(100))
	fmt.Printf("%s\n", BoilingC.ToF())
	fmt.Printf("%s\n", Fahrenheit(100).ToC())

	// var c Celsius
	// var f Fahrenheit
	// fmt.Printf("%g °C\n", Celsius(10) + Fahrenheit(20)) // doesn't compile
	// fmt.Printf("%v\n", c == f) // doesn't compile
	// fmt.Printf("%v\n", BoilingC == Fahrenheit(100)) // doesn't compile
}
