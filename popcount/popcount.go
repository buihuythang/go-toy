// Count the number of set bits (bits whole value is 1) within an 64-bit unsigned integer.
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Count the number of set bits using Loopup Table.
func PopCountLUT(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Count the number of set bits using shift.
func PopCountShift(x uint64) int {
	cnt := 0
	for i := 0; i < 64; i++ {
		cnt += int(byte((x >> i) & 1))
	}

	return cnt
}

// The expression x&(x-1) clears the rightmost non-zero bit of x.
// This version of PopCount use that fact.
func PopCountRNZ(x uint64) int {
	cnt := 0
	for x != 0 {
		x &= x - 1
		cnt++
	}

	return cnt
}
