// Exercise 4.3
// Page 93
//
// Prompt:
// Rewrite reverse to use an array pointer instead of a slice.

// Development Notes:
// Using an array pointer look like it will require specifying an array length.
//
// https://eli.thegreenplace.net/2021/generic-functions-on-slices-with-go-type-parameters/
// This article highlights potential improvments for practical implementations of this function via
// the use of generics, particularly when it comes to the underlying data types.
// https://go.dev/blog/why-generics
//
// Resources:
// https://www.geeksforgeeks.org/how-to-pass-an-array-to-a-function-in-golang/
// https://www.geeksforgeeks.org/golang-pointer-to-an-array-as-function-argument/

package main

import "fmt"

// Reverse uses an array pointer to reverse an array.
func reverse(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("Input")
	fmt.Println(a) // "[0, 1, 2, 3, 4, 5]"
	reverse(&a)
	fmt.Println("Output")
	fmt.Println(a) // "[5, 4, 3, 2, 1, 0]"
}
