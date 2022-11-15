// Exercise 4.2
// Page 84
//
// Prompt:
// Write a program that prints the SHA256 hash of its standard input by default,
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.

// Program takes a string and outputs the SHA256 hash.
// Optional flags for SHA384 and SHA512.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	// Flag Handling
	hashType := flag.String("Hash", "sha256", "Options: sha256, sha384, and sha512")
	// Runtime handling of flags.
	flag.Parse()

	// Command-Line Handling
	scanner := bufio.NewScanner(os.Stdin)

	// Loop reading the inputs.
	for scanner.Scan() {
		// Input entry
		input := scanner.Text()

		// Program exit.
		if input == "q" {
			os.Exit(0)
		}

		// Hashing and Output
		switch *hashType {
		// Usage: -type=sha512
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
		// Usage: -type=sha384
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
		// Usage: No flag.
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
		}
	}

	// Basic error handling.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
