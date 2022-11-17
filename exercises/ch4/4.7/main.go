// Exercise 4.7
// Page 93
// 
// Prompt:
// Modify reverse to reverse the characters 
// of a []byte slice that represents a 
// UTF-8 encoded string, in place. 
// Can you do it without allocating new memory?

package main

import "fmt"

func main() {
	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

// reverse reverses a slice of ints in a place.
func reverse(s []int)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}