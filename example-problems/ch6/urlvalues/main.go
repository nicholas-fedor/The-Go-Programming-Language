// see page 160

// The urlvalues command demonstrates a map type with methods
package main

import (
	"fmt"
	"net/url"
)

func main() {
	//!+main
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"
	fmt.Println(m["item"])     // "[1 2]"

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map
	//!-main
}
