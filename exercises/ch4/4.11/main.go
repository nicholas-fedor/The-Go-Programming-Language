// Exercise 4.11
// Page 112
//
// Prompt:
// Build a tool that lets users create, read, update, and close GitHub issues from the command line,
// invoking their preferred text editor when substantial text input is required.
//
// Development Notes:
//

// Program lets users work with GitHub issues from the command line.
// Invokes the user's preferred text editor, if needed.
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
	}

	switch os.Args[1] {
	case "search":
		// Search issues GitHub api get request.
		// Argument[1] = search
		result, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Found %d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			log.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	
	case "create":
		// Dev output
		log.Println("Created")
		
	case "read":
		// Dev output
		log.Println("Read")
		
	case "update":
		// Dev output
		log.Println("Updated")
		
	case "close":
		// Dev output
		log.Println("Closed")
		
	default:
		// Dev output
		help()
	}

	// Functions by command-line argument:
	// 
	// Search issues GitHub api get request.
	// // Argument[1] = search
	//
	// Create issue GitHub api post request.
	// // Argument[1] = create
	//
	// Read issue GitHub api get request.
	// // Argument[1] = read
	//
	// Update issue GitHub api get request.
	// // Argument[1] = update
	//
	// Close issue GitHub api get request.
	// Argument[1] = close
	//


	// Output: Response from GitHub API request.

}

// help log.Printf() and os.Exit(1) basic listing of commands and their usage.
func help() {
	log.Fatalf("\nProgram usage:\nTerminal Command: go run main.go search repo:golang/go is:open json decoder\n")
	
}

// Example Usage:
// Search
// go run main.go search repo:golang/go is:open json decoder
//
// Create
// go run main.go create
//
// Read
// go run main.go read
//
// Update
// go run main.go update
//
// Close
// go run main.go close
