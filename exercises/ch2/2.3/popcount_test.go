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

func BenchmarkTableLookup(b *testing.B) {
	bench(b, PopCountTableLoop)
}
