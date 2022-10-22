// see page 218

// spinner program demonstrates goroutines by computing the 45th Fibonacci number.
// Since it uses the terribly inefficient recursive algorithm, it runs for an appreciable
// amount of time, during which we'd like to provide the user with a visual indication
// that the program is still running, by displaying an animated textual "spinner".
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

// Output:
// Fibonacci(45) = 1134903170

// Explaination:
// After several seconds of animation, the fib(45) call returns
// and the main function prints it's result.
// The main function then returns. When this happens, all goroutines are
// abruptly terminated and the program exits.
// Other than by returning from main or exiting the program, there
// is no programmatic way for one goroutine to stop another, but as
// we will see later, there are ways to communicate with a goroutine
// to request that it stops itself.

// Notice how the program is expressed as the composition of to autonomous
// activities, spinning and Fibonacci computation.
// Each is written as a separate function but both make progress concurrently.
