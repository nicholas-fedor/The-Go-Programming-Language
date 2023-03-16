// see page 229

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func ()  {
		for x := 0; x < 100; x++ { // runs until 100
			naturals <- x	// sends the result to the channel
		}
		close(naturals) // closes the channel after the last value is received
	}()

	// Squarer
	go func ()  {
		for x := range naturals {
			squares <- x * x // sends results to the squares channel
		}
		}()
		close(squares)

	// Printer (in the main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}