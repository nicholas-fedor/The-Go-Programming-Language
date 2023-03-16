// see page 179

// Sleep demonstrates the flag.Value interface for defining new notations for command-line flags.
package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// Usage:
// go build gopl.io/ch7/sleep
// ./sleep *(the default is 1second)
// arguments are passed via the "-period" flag.
// specify the amount of time

// Examples:
// ./sleep
// Output: Sleeping for 1s...
//
// ./sleep -period 50ms
// Output: Sleeping for 50ms