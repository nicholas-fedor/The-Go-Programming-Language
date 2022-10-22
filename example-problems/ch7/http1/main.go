// see page 191

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// Usage:
// go build gopl.io/ch7/http1
// ./http1 &
// * the & flag allows the process to run without locking out the terminal.
// go build gopl.io/ch1/fetch
// ./fetch http://localhost:8000

// Output:
// shoes: $50.00
// socks: $5.00

// To stop the server:
// Find the process using the "jobs" command
// End the process using the "kill" command - ex) kill -s SIGINT %1
