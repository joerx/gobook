package popcount1

var pc [256]byte

// For given int i, i*2 is equivalent to a binary left shift by one bit
// After left shift, the rightmost bit is 0, so no additional `1` is added
//
// For an even number i, we can trivially say: popcount(i) == popcount(i/2)
// For an odd number i, popcount(i) == popcount(i/2) + 1 since the last bit is `1`
// In other words: i/2 is a bitwise right shift, discarding the last bit. To revert this, we need to add it back
// We use a bitwise AND to determine if 1 is even or odd: 100 & 1 == 0, 101 & 1 == 1, etc.
//
// Putting it all together: popcount(i) = popcount(i/2) + i&1

// i byte(i) 	pc[i/2] 	i&1 	pc[i]		int(pc[i])
// 0 00000000 	00000000	0		000000000	0
// 1 00000001  	00000000   	1		000000001	1
// 2 00000010	00000001	0		000000001	1
// 3 00000011	00000001	1		000000010	2
// 4 00000100	00000001	0		000000001	1			-- 4 = 2*2+0 = 2<<1 + 4&1 - i.e. same popcount as 2
// 5 00000101	00000001	1		000000010	2			-- 5 = 2*2+1 = 2<<1 + 5&1 - i.e. popcount of 2 + 1
// 6 00000110	00000010	0		000000011	2			-- 6 = 3*2+0 = 3<<1 + 6&1
// 7 00000111	00000010	1		000000011	3			-- 7 = 3*2+1 = 3<<1 + 7&1
// 8 00001000	00000001	0		000000001	1			-- 8 = 4*2+0 = 4<<1 + 8&1 - i.e. popcount of

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count(number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}
