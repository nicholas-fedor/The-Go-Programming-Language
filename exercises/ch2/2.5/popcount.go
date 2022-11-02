// Exercise 2.5
// Page 45
//
// Prompt:
// The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact,
// and assess its performance

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
func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
