package main

import (
	"fmt"
	"gopl.io/example-problems/ch2/tempconv"
)

func Example_one() {
	//!+arith
	t := 32

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	
	fmt.Printf("%s = %s, %s = %s", f, tempconv.FToC(f), c, tempconv.CToF(c))
	//!-arith

	// Output: 
	// 32°F = 0°C, 32°C = 89.6°F
}