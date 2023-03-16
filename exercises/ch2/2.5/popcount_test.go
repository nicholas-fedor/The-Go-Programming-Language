// Exercise 2.5
// Page 45
//
// Prompt:
// The expression x&(x-1) clears the rightmost non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact, 
// and assess its performance

// Package popcount
package popcount

import (
	"testing"
)

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	bench(b, PopCount)
}
