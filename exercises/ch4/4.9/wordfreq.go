// Exercise 4.9
// Page 99
//
// Prompt:
// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.
//
// Development Notes:
// Used https://www.geeksforgeeks.org/counting-number-of-repeating-words-in-a-golang-string/
// to help with mapping.

// wordfreq counts the frequency of each word in an input.txt file.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Static text file.
	filename := "input.txt"
	
	// File handling.
	fmt.Printf("Input File: %s\n", filename) 
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	
	// Input buffer handling.
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	// Mapping input to count unique words.
	// Leading and trailing punctuation affects individual words.
	// i.e. "hello, world" = "hello," and "world"
	wc := make(map[string]int)
	for input.Scan() {
		words := strings.Fields(input.Text())
		for _, word := range words {
			_, matched := wc[word]
			if matched {
				wc[word] += 1
				} else {
					wc[word] = 1
				}
			}
		}

	// Output results.
	fmt.Printf("Word:\tCount:\n")
	for w, c := range wc {
			fmt.Printf("%s\t%d\n", w, c)
	}
}
