package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountViaLoop(x uint64) int {
	var popCount int

	for i := 0; i < 8; i++ {
		res := x >> (i * 8)
		popCount += int(pc[byte(res)])
	}

	return popCount
}

func PopCountByRightShift(x uint64) int {
	var popCount int
	shift := uint64(1)

	for i := 0; i < 64; i++ {
		if x&shift > 0 {
			popCount++
		}
		x = x >> shift
	}

	return popCount
}

func PopCountByRightMostBit(x uint64) int {
	var popCount int

	for x != 0 {
		x = x&(x-1)
		popCount++
	}

	return popCount
}