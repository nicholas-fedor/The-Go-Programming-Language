// see page 348

// Search is a demo of the params.Unpack function.
package main

import (
	"fmt"
	"log"
	"net/http"

	"gopl.io/example-problems/ch12/params"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"1"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
