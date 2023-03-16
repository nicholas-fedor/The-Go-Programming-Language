// Exercise 1.11
// Page 19
//
// Prompt:
// Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com.
// How does the program behave if a web site just doesn't respond?
// (Section 8.9 describes mechanisms for coping in such cases.)

// Answer:
// https://alexa.com/ was retired on May 1, 2022.
// See https://hackertarget.com/top-million-site-list-download/ for similar resources.
// Code below modified to accept a txt file of links.
// Fetchall will not timeout if a web site doesn't respond.
//
// Section 8.9 addresses cancelling a goroutine, with the example using a separate channel to 
// signal the program to terminate.

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	// "io/ioutil" deprecated as of Go 1.16. Use the "io" stnd library instead.
	// https://pkg.go.dev/io/ioutil
)

func main() {
	// Opens links.txt.
	file, err := os.Open("links.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	// Reads the file.
	urls := bufio.NewScanner(file)
	// Loops through each line for text.
	for urls.Scan() {
		// Adds text to url variable.
		url := urls.Text()

		// skip blank lines
		if len(url) == 0 {
			continue
		}

		// Checks if url has "http://" prefix
		// Adds "https://", if not.
		if !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}

		start := time.Now()
		ch := make(chan string)
		// start a goroutine
		go fetch(string(url), ch)

		// receive from channel ch
		fmt.Println(<-ch) 

		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
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
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}