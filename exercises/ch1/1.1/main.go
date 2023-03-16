// Exercise 1.1
// Page 8

// Prompt:
// Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.

// Echo prints the name of the command that invoked it and its arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// The command that invoked the program.
	fmt.Println(os.Args[0])
	// The arguments
	fmt.Println(strings.Join(os.Args[1:], " "))
}
