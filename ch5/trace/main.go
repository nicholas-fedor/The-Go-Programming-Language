// see page 146

// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"time"
)

// !+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

// Usage:
// go build gopl.io/ch5/trace
// ./trace

// Output:
// 2022/09/14 17:06:30 enter bigSlowOperation
// 2022/09/14 17:06:40 exit bigSlowOperation (10.009502353s)
