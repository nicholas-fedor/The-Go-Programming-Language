// Exercise 4.6
// Page 93
//
// Prompt:
// Write an in-place function that squashes
// each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8 encoded
// []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Input
	s := []byte("a\tb\tc\td\te")
	fmt.Printf("%s\n", s) // "a       b       c       d       e"
	// Loop through each element of the slice.
	for i := 0; i < len(s); i++ {
		// Converts each element into a rune for unicode.IsSpace
		v := rune(s[i])
		// If the rune is a unicode space character, i.e. "/t"
		if unicode.IsSpace(v) {
			// Replaces the element with a single space.
			s[i] = byte(' ')
		}
	}
	fmt.Printf("%s\n", s) // "a b c d e"
}
