// see page 228

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func () {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break // chnnel was closed and drained
			}
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}