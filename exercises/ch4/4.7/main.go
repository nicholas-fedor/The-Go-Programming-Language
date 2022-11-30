// Exercise 4.7
// Page 93
//
// Prompt:
// Modify reverse to reverse the characters
// of a []byte slice that represents a
// UTF-8 encoded string, in place.
// Can you do it without allocating new memory?

// Development notes:
// Modifying the array, in place, means that the
// underlying input array is modified during the call,
// as opposed to moving the data into a new output array.
//
// This exercise builds upon example ch4/rev in that the
// datatype is a UFT-8 encoded string, instead of integers.
//
// Used solution from https://github.com/kdama/gopl/blob/master/ch04/ex07/main.go

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Input string
	s := "Lorem ipsum"

	// Remember to re-encode using string().
	fmt.Println(string(reverseUTF8([]byte(s))))
	// Output:
	// "muspi meroL"
}

// reverseUTF8 iterates through a UTF-8 encoded string
// and decodes each rune before reversing the string.
func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	return b
}

// reverse reverses a slice of bytes in place.
func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
