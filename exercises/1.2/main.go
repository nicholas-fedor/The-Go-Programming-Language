// See page 8

// Prompt:
// Modify the echo program to print the index and value of each of its arguments, one per line.

// Echo prints the index and value of each argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}

// Example Usage:
//go run main.go hello world!
//
// Output:
// 0: hello
// 1: world!
