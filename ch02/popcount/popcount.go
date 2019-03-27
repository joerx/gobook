package main

import(
	"fmt"
	"strconv"
)

// pc[i] is population count of i
var pc [256]byte

func init() {
	// Multiply by 2 shifts all bits to the left by one, not affecting the number of ones or zeroes.
	// For every odd number, the last bit is `1` so we add one to the population count.
	// In reverse, the number of ones for `i` is the same as for half `i/2` with one added for odd
	// numbers. (i&1 == 1 for odd numbers)
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountExp(num uint64) uint8 {
	// Cast `byte(num>>n)` has similar effect as `num>>n&&255` - cuts off all bits left of the last 
	// 8 binary digits. This only works if num is a variable, if it'd be a constant the code wouldn't 
	// compile
	res := 
		pc[byte(num>>(0*8))] + 
		pc[byte(num>>(1*8))] + 
		pc[byte(num>>(2*8))] + 
		pc[byte(num>>(3*8))] + 
		pc[byte(num>>(4*8))] + 
		pc[byte(num>>(5*8))] + 
		pc[byte(num>>(6*8))] + 
		pc[byte(num>>(7*8))]
	return res
}

func PopCountLoop(num uint64) uint8 {
	var i, res uint8
	for i = 0; i < 8; i++ {
		res += pc[byte(num>>(i*8))]
	}
	return res
}
 
func printPopCount(num uint64) {
	fmt.Printf("%s: %v\n", strconv.FormatInt(int64(num), 2), PopCountExp(num))
}

func main() {
	printPopCount(16)
	printPopCount(25)
	printPopCount(2413)
	printPopCount(412353451)
}
