package tempconv

import (
	"fmt"
)

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}
