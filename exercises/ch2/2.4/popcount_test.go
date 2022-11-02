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

import (
	"testing"
)

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkTableShift(b *testing.B) {
	bench(b, PopCountTableShift)
}
