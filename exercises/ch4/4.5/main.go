// Exercise 4.5
// Page 93
//
// Prompt:
// Write an in-place function to eliminate
// adjacent duplicates in a []string slice.

// Development Notes:
// Looping through the slice is easy, but I ended up
// using the following resource for the algorithm to
// delete and truncate the slice.
// https://yourbasic.org/golang/delete-element-slice/

// Program loops through a slice from left to right,
// checks if the next element is the same value,
// and erases and truncates, if it is the same.
// The order of elements is otherwise maintained.
package main

import "fmt"

func main() {
	// Input
	s := []string{"a", "a", "b", "b", "c", "c", "d", "d"}
	// Loops through the slice.
	for n := 0; n < len(s)-1; n++ {
		// If the element at index n is the same value as element at n+1.
		if s[n] == s[n+1] {
			// Copy last element to index n.
			copy(s[n:], s[n+1:])
			// Erase last element
			s[len(s)-1] = ""
			// Truncate slice.
			s = s[:len(s)-1]
		}
	}
	fmt.Println(s) // "[a b c d]"
}
