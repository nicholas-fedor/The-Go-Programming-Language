// Exercise 3.10
// Page 74
//
// Prompt:
// Write a non-recursive version of comma, using bytes.Buffer
// instead of string concatenation.

// Original function:
/*
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
*/

// Used https://github.com/torbiak/gopl/blob/master/ex3.10/main.go to help complete.
package main

import (
	"bytes"
	"fmt"
)

func main() {

	intString := "12345"
	fmt.Println(comma(intString))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)

	switch {
	case n <= 3:
		fmt.Fprint(&buf, s)

	case n > 3:
		// commaPosition goes from left to right.
		commaPosition := len(s) % 3
		// Moves commaPosition to after first three digits
		// if the string has 6, 9, 12, etc digits.
		// Example: 123,456 instead of ,123,456
		if commaPosition == 0 {
			commaPosition = 3
		}
		// Output first triplet to buffer.
		fmt.Fprint(&buf, s[:commaPosition])
		// Iterates between remaining digits.
		for i := commaPosition; i < len(s); i += 3 {
			fmt.Fprint(&buf, ",")
			fmt.Fprint(&buf, s[i:i+3])
		}
	}
	// Output back to main.
	return buf.String()
}
