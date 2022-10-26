// Exercise 1.8
// Page 17

// Prompt:
// Modify fetch to add the prefix http:// to each argument URL if it is missing.
// You might want to use strings.HasPrefix

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		switch {
		case strings.HasPrefix(url, "http://"):
			fetch(url)
		case strings.HasPrefix(url, "https://"):
			fetch(url)
		default:
			fetch("https://" + url)
		}
	}
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}
