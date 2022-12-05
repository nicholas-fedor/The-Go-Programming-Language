// Exercise 4.10
// Page 112
//
// Prompt:
// Modify issues to report the results in age categories,
// say less than a month old, less than a year old, and more than a year old.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/example-problems/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Current Date/Time:", time.Now())

	for _, item := range result.Items {
		// Age in days
		age := time.Since(item.CreatedAt).Hours()/24
		fmt.Printf("#%-5d %9.9s %s %.2f %.55s \n",
			item.Number, item.User.Login, item.CreatedAt.Format("2006-01-02 15:05"), age, item.Title)

	}

}
 