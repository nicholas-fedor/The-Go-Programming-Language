// see page 91

// Nonempty is an example of an in-place slice algorithm
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

// Output:
// '["one" "three"]'
// '["one" "three" "three"]'

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
