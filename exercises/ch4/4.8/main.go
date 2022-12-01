// Exercise 4.8
// Page 99
//
// Prompt:
// Modify charcount to count letters, digits, and
// so on in their Unicode categories, using
// functions like unicode.IsLetter.

// Developer Notes:
// Unicode categories within Go can be found here:
// https://pkg.go.dev/unicode#pkg-variables
//
// The requested unicode functions can be found here:
// https://pkg.go.dev/unicode#pkg-functions
//
// Used solution from here:
// https://github.com/linehk/gopl/blob/main/ch4/exercise4.8/main.go

// Charcount computes counts of Unicode characters.
//
// Usage: echo -n "hello" | go run main.go
// -n flag removes the newline.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type class string

const (
	letter  class = "letter"
	number  class = "number"
	graphic class = "graphic"
	space   class = "space"
	symbol  class = "symbol"
)

func main() {
	// classCount is the count of Unicode characters for each class.
	classCount := make(map[class]int, 5)
	// Buffered input via standard input.
	in := bufio.NewReader(os.Stdin)

	for {
		// returns rune, nbytes, error. ignoring nbytes
		r, _, err := in.ReadRune()

		// Ends program loop once end of Stdin reached.
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		// Implementation of count per class
		switch {
		case unicode.IsLetter(r):
			classCount[letter]++
		case unicode.IsNumber(r):
			classCount[number]++
		case unicode.IsGraphic(r):
			classCount[graphic]++
		case unicode.IsSpace(r):
			classCount[space]++
		case unicode.IsSymbol(r):
			classCount[symbol]++
		}
	}

	fmt.Printf("class\tcount\n")
	for class, count := range classCount {
		fmt.Printf("%s\t%d\n", class, count)
	}
}
