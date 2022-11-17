// Exercise 4.4
// Page 93
//
// Prompt:
// Write a version of rotate that operates in a single pass.

// Development Notes:
// https://www.geeksforgeeks.org/c-program-cyclically-rotate-array-one/
// Tried a few different iterations to get the type generics for slices to work.
// The article below helped to resolve all the issues that I was having.
// https://gosamples.dev/generics-slice-contains/

// Rotate rotates a slice left by n elements.
// Example: Left Rotate (i.e. array elements shifted left and )
// input  = [0, 1, 2, 3, 4, 5]
// output = [2, 3, 4, 5, 0, 1]
package main

import "fmt"

// Rotates left two elements.
func Rotate[T comparable](s []T) {
	t := s[:2]
	s = append(s[2:], t...)
	fmt.Println(s)
}

func main() {
	input := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(input)
	Rotate(input[:])
}

// Input: 
// [0 1 2 3 4 5]

// Output: 
// [2 3 4 5 0 1]