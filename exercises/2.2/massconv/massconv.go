package massconv

import "fmt"

const poundToKilo = 2.205

// Kilogram is the mass of an object on the only scale that makes sense
type Kilogram float64

// Pound is still used by some underdeveloped island tribes
type Pound float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%.2f kg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.2f lb", p)
}

// KgtoLb converts Kilograms to Pounds
func KgtoLb(k Kilogram) Pound {
	return Pound(k * poundToKilo)
}

// LbtoKg converts Pounds to Kilograms
func LbtoKg(p Pound) Kilogram {
	return Kilogram(p / poundToKilo)
}
