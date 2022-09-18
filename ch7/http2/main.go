// see page 192

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

// handler
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

// Usage:
// go build gopl.io/ch7/http2
// ./http2 &
// * the & flag allows the process to run without locking out the terminal.
// go build gopl.io/ch1/fetch
// ./fetch http://localhost:8000/list

// Output:
// shoes: $50.00
// socks: $5.00

// ./fetch http://localhost:8000/price?item=socks
// Output: $5.00

// To stop the server:
// Find the process using the "jobs" command
// End the process using the "kill" command - ex) kill -s SIGINT %1
