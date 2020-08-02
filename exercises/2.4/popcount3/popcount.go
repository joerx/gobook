package popcount3

// PopCount counts the non-zero bits in x by right-shifting x and checking the value of the last bit
func PopCount(x uint64) int {
	var cnt int
	for x > 0 {
		cnt += int(x & 1)
		x = x >> 1
	}
	return cnt
}
