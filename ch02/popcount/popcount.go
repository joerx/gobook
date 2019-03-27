package main

import(
	"fmt"
	"strconv"
)

// Thought about big-O for here and figured it's all ∈ O(1) since they all take only a single input
// value and the value of that input has no real effect on the number of operations - except 
// PopCountSmartShift, but even there the max n of 64 is negligible.

// pc[i] is population count of i - we only need 8 bit numbers since larger numbers can be broken
// down into 8x8-bit chunks w/o affecting the number of non-zero bits.
var pc [256]byte

func init() {
	// Multiply by 2 shifts all bits to the left by one, not affecting the amount of non-zero bits.
	// For every odd number, the last bit is `1` so we add one to the population count.
	// In reverse: number of non-zero bits for `i` is the same as for `i/2` with one added for odd 
	// numbers. (i&1 means 'apply bitmask 1', i.e. get value of the rightmost bit)
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

// Same as PopCountExp but using a loop instead. Looks more elegant (dynamic, yeah) but is slower 
// than PopCountExp. Apparently little point in having a loop if the size is fixed.
func PopCountLoop(num uint64) uint8 {
	var i, res uint8
	for i = 0; i < 8; i++ {
		res += pc[byte(num>>(i*8))]
	}
	return res
}

// Shifts through each bit of the input, adds up non-zero bit. Always needs 64 operations. This is
// obviously the slowest approach.
func PopCountShift(num uint64) uint8 {
	var cnt uint8
	for ; num > 0; num = num >> 1 {
		cnt += byte(num & 1)
	}
	return cnt
}

// Taking advantage of x&(x-1) clearing the rightmost non-zero bit. Needs  64 ops at max. Faster 
// then PopCountShift but still slower than PopCountExp and PopCountLoop
func PopCountSmartShift(num uint64) uint8 {
	var cnt uint8
	for x := num; x > 0; x = x&(x-1) {
		cnt++
	}
	return cnt
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
