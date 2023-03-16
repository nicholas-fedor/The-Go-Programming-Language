// Exercise 2.4
// Page 45
//
// Prompt:
// Write a version of PopCount that counts bits
// by shifting its argument through 64 bit positions,
// testing the rightmost bit each time.
// Compare its performance to the table-lookup version.

// Package popcount
package popcount

// pc[i] is the population count of i.
var pc [256]byte

// Precomputes a table of results, pc, for each possible 8-bit value
// so that the PopCount function needn't take 64 steps, but can just
// return the sum of eight table lookups.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTable returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountTableShift
func PopCountTableShift(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}
