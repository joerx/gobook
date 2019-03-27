package main

import "fmt"

// Currency represents a currency name
type Currency int

// Currency name
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
	}
	fmt.Println(RMB, symbol[RMB])
}
