package tempconv

import "fmt"

// Celsius represents temperature on the celsius scale
type Celsius float64

// Fahrenheit represents temperature on the celsius scale
type Fahrenheit float64

// Kelvin represents temperature on the Kelvin scale
type Kelvin float64

const (
	// AbsoluteZeroC is the absolute zero on the Celsius scale
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing point of water on the Celsius scale at zero altitude
	FreezingC Celsius = 0
	// BoilingC is the boiling point of water on the Celsius scale at zero altitude
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f °K", k)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f °F", f)
}

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC) // i.e. -273.14 C = 0 K
}

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

// FToK convert Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9) - Kelvin(AbsoluteZeroC)
}

// KToF convert Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(Celsius(k)+AbsoluteZeroC)*9/5 + 32
}

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
