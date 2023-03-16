// Exercise 1.10
// Page 18

// Prompt:
// Find a website that produces a large amount of data.
// Investigate caching by running fetchall twice in succession
// to see whether the reported time changes much.
// Do you get the same content each time?
// Modify fetchall to print its output to a file
// so it can be examined.

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	// "io/ioutil" deprecated as of Go 1.16. Use the "io" stnd library instead.
	// https://pkg.go.dev/io/ioutil
)

func main() {
	// Output file handling.
	// If file already exists, the output will be appended.
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start time prior to running fetch.
	programStart := time.Now()

	// Creating a channel for receiving output from fetch.
	ch := make(chan string)
	// Iterate through command line arguments (urls) as input into
	// the fetch function.
	// Notice that it outputs to ch (the channel).
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	// Outputting from the channel to results.
	for range os.Args[1:] {
		results := <-ch // receive from channel ch
		// Output to both the terminal and log.txt file.
		// No impact to the fetch times.
		fmt.Println(results)
		fmt.Fprintln(f, results)
	}

	// Sets endpoint for overall fetch time.
	// Remember that this will be the total elapsed time for 
	// all of the goroutines to run, not necessarily the sum of 
	// times for the goroutines.
	// This means it will be equivalent to the longest running goroutine.
	elapsedMS := time.Since(programStart).Milliseconds()
	fmt.Printf("Elapsed Time: %dms\n", elapsedMS)
	fmt.Fprintf(f, "Elapsed Time: %dms\n\n", elapsedMS)
}

func fetch(url string, ch chan<- string) {
	fetchStart := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	fetchTime := time.Since(fetchStart).Milliseconds()
	// Outputting fetchResult to ch
	ch <- fmt.Sprintf("URL: %s\t\tTime: %dms\t\tData:%7d bytes", url, fetchTime, nbytes)
}
