package tempconv

import (
	"fmt"
)

const (
	AbsoluteZeroC = -273.15
	FreezingC = Celsius(0)
	BoilingC = Celsius(100)
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c Celsius) ToF() Fahrenheit {
	return CToF(c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}
