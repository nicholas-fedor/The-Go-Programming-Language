// Exercise 1.9
// Page 17

// Prompt:
// Modify fetch to also print the HTTP status code, found in resp.Status.

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

	fmt.Println("Fetch")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Status Code: %d\n\n", resp.StatusCode)

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}
