// Exercise 2.3
// Page 45
//
// Prompt:
// Rewrite PopCount to use a loop instead of a single expression.
// Compare the performance of the two versions.
// (Section 11.4 shows how to compare the performance of different
// implementations systematically.)

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

// PopCount returns the population count (number of set bits) of x.
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

// PopCountTableLoop implements a loop 
func PopCountTableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}