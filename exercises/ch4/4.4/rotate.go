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
// Example:
// input  = [0, 1, 2, 3, 4, 5]
// output = [2, 3, 4, 5, 0, 1]
package main

import "fmt"

func Rotate[T comparable](s[] T) {
	first := 0
	last := len(s) - 1

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}

	fmt.Println(s)
}

func main() {
	input := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(input)

	// TODO:
	// Still need to figure out algorithm that 
	// conforms to the exercise requirements.
	Rotate(input[:])
}
