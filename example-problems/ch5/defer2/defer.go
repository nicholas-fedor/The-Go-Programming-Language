// see page 151

// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/* Output:
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
goroutine 1 [running]:
main.printStack()
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:19 +0x39
panic({0x48cba0, 0x523ef0})
        /usr/local/go/src/runtime/panic.go:884 +0x212
main.f(0x4b9298?)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:24 +0x113
main.f(0x1)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.f(0x2)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.f(0x3)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.main()
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:14 +0x45
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x4b9298?)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:24 +0x113
main.f(0x1)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.f(0x2)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.f(0x3)
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:26 +0xf5
main.main()
        /home/nick/Github/The-Go-Programming-Language/ch5/defer2/defer.go:14 +0x45 */