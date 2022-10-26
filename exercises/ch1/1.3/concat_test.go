// Exercise 1.3
// Page 8

// Prompt:
// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package, and Section 11.4 shows
// how to write benchmark tests for systematic performance evaluation.)

package concat_test

import (
	"strings"
	"testing"
)

var args = []string{"fizz", "bang", "foo", "bar"}

func concat(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}

/*
goos: linux
goarch: amd64
pkg: gopl.io/exercises/1.3
cpu: AMD Ryzen Threadripper 2950X 16-Core Processor
BenchmarkConcat-32    	 7267400	       158.6 ns/op	      56 B/op	       3 allocs/op
BenchmarkJoin-32      	20115945	        59.53 ns/op	      24 B/op	       1 allocs/op
PASS
coverage: [no statements]
ok  	gopl.io/exercises/1.3	2.590s
*/
