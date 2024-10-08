// see page 181

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/example-problems/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

// Usage: go build gopl.io/ch7/tempflag
// ./tempflag
// Output: 20°C

// ./tempflag -temp -18C
// Output: -18°C
