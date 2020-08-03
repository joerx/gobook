package popcount4

// PopCount calculates the population count of x by iteratively clearing the rightmost non-zero bit
// until we cleared all non-zero bits (i.e. the remaining value of x is zero)
func PopCount(x uint64) (pc int) {
	for pc = 0; x > 0; pc++ {
		x = x & (x - 1)
	}
	return pc
}
