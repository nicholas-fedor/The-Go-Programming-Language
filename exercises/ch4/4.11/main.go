// Exercise 4.11
// Page 112
//
// Prompt:
// Build a tool that lets users create, read, update, and close GitHub issues from the command line,
// invoking their preferred text editor when substantial text input is required.
//
// Development Notes:
// Looked at using flags; however, using arguments was simpler.
//

// Program lets users work with GitHub issues from the command line.
// Invokes the user's preferred text editor, if needed.
//
// Example Usage:
// Search
// go run main.go search repo:golang/go is:open json decoder
//
// Create
// go run main.go create
//
// Read
// go run main.go read golang/go/issues/42571
// Request: https://api.github.com/repos/golang/go/issues/42571
//
// Update
// go run main.go update
//
// Close
// go run main.go close
package main

import (
	"log"
	"os"

	"gopl.io/exercises/ch4/4.11/github"
)

func main() {

	log.Printf("Program: %v", os.Args[0])
	// Input:
	// First argument specifies the command to run.
	if len(os.Args) < 2 {
		help()
	} else {
		switch os.Args[1] {

		// Search issues GitHub api get request.
		// Argument[1] = search
		case "search":
			// Dev ouput
			log.Println("Search Selected")
			search()

			// Create issue GitHub api post request.
			// Argument[1] = create
		case "create":

			// Dev output
			log.Println("Created Selected")

			// Read issue GitHub api get request.
			// Argument[1] = read
		case "read":
			// Dev output
			log.Println("Read Selected")
			read()

			// Update issue GitHub api get request.
			// Argument[1] = update
		case "update":
			// Dev output
			log.Println("Update Selected")

			// Close issue GitHub api get request.
			// Argument[1] = close
		case "close":
			// Dev output
			log.Println("Close Selected")

		default:
			help()
		}
	}

}

// help log.Printf() and os.Exit(1) basic listing of commands and their usage.
func help() {
	log.Fatalf("\nProgram usage:\nTerminal Command: go run main.go search repo:golang/go is:open json decoder\n")
}

// search takes one or more arguments to query the GitHub api and log.Printf's the results.
// Example usage: go run main.go search repo:golang/go is:open json decoder
func search() {
	result, err := github.SearchIssues(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		log.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// read queries the GitHub api and log.Printf's the results.
// Example usage: go run main.go read golang/go/issues/42571
func read() {
	input := os.Args[2]
	log.Printf("API Query: %s", input)

	result, err := github.ReadIssue(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("url: %s | #%-5d | status: %s\n", result.HTMLURL, result.Number, result.State)
	log.Printf("Title: %s\n", result.Title)
	log.Printf("Content:\n%v\n", result.Body)
}
