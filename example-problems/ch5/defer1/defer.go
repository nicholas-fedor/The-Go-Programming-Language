// see page 150

// Defer1 demonstrates a defered call being invoked during a panic.
package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

// Output:
// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.f(0x4b91b8?)
//         /home/nick/Github/The-Go-Programming-Language/ch5/defer1/main.go:13 +0x113
// main.f(0x1)
//         /home/nick/Github/The-Go-Programming-Language/ch5/defer1/main.go:15 +0xf5
// main.f(0x2)
//         /home/nick/Github/The-Go-Programming-Language/ch5/defer1/main.go:15 +0xf5
// main.f(0x3)
//         /home/nick/Github/The-Go-Programming-Language/ch5/defer1/main.go:15 +0xf5
// main.main()
//         /home/nick/Github/The-Go-Programming-Language/ch5/defer1/main.go:9 +0x1e
// exit status 2
