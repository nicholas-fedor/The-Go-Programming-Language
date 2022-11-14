// Exercise 4.1
// Page 84
//
// Prompt:
// Write a function that counts the number of bits
// that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2.)

/* 
// popcount
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] | byte(i&1)
	}
}

// Popcount returns the population count 
// (number of set bits) of x.
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
 */

// Compare takes two SHA256 hashes and 
// counts the number of bits that are different.
package main

func main() {
	
}