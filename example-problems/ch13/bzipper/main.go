// See page 365.

// Bzipper reads input, bzkip2-compresses it, and writes it out.
package main

import (
	"io"
	"log"
	"os"

	"gopl.io/example-problems/ch13/bzip"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close %v\n", err)
	}
}
