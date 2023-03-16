// Exercise 4.1
// Page 84
//
// Prompt:
// Write a function that counts the number of bits
// that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2.)

// Development notes:
// Solution and additional resources from
// https://stackoverflow.com/questions/34966658/compare-bits-in-two-arrays

// Compare takes two strings, hashes them using SHA256, and
// counts the number of bits that are different.
package main

import (
	"crypto/sha256"
	"fmt"
)

func CountBitDifference(h1, h2 *[sha256.Size]byte) int {
	count := 0
	// Loops through h1's index.
	for i := range h1 {
		// Var b = the 
		for b := h1[i] ^ h2[i]; b != 0; b &= b - 1 {
			count++
		}
	}
	return count
}

func main() {
	// Inputs
	s1 := "1"
	s2 := "2"

	// Compute hashes
	h1 := sha256.Sum256([]byte(s1))
	h2 := sha256.Sum256([]byte(s2))

	// Calculate Difference
	difference := CountBitDifference(&h1, &h2)

	// Output
	fmt.Printf("Input1: %s | Hash1: %X | Type: %T\n", s1, h1, h1)
	fmt.Printf("Input2: %s | Hash2: %X | Type: %T\n", s2, h2, h2)
	fmt.Printf("Number of different bits: %d\n", difference)
}
