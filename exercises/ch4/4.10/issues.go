// Exercise 4.10
// Page 112
//
// Prompt:
// Modify issues to report the results in age categories,
// say less than a month old, less than a year old, and more than a year old.

// Development notes:
// Used solution from here: 
// https://github.com/linehk/gopl/blob/main/ch4/exercise4.10/main.go
// 
// I got stuck, prior to reviewing the above solution, with filtering and 
// storing the data for later output handling.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/example-problems/ch4/github"
)

type class string

const (
	LTOM class = "less than one month"
	MTOM class = "more than one month"
	LTOY class = "less than one year"
	MTOY class = "more than one year"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	issueClass := make(map[class][]github.Issue)
	cy, cm, _ := time.Now().Date()

	for _, item := range result.Items {
		item := *item
		y, m, _ := item.CreatedAt.Date()

		switch {
		// issue is <= 1 month.
		case cm-m <= time.Month(1):
			issueClass[LTOM] = append(issueClass[LTOM], item)
			
			// issue is > 1 month.
		case cm-m > time.Month(1):
			issueClass[MTOM] = append(issueClass[MTOM], item)
			
			// issue is <= 1 year.
		case cy-y <= 1:
			issueClass[LTOY] = append(issueClass[LTOY], item)
			
			// issue is > 1 year.
		case cy-y > 1:
			issueClass[MTOY] = append(issueClass[MTOY], item)
		}
	}

	for class, issues := range issueClass {
		fmt.Printf("Age Class: %s\n", class)
		fmt.Printf("Item #\tUser:\tItem Summary:\n")
		for _, item := range issues {
			fmt.Printf("#%-5d\t%9.9s\t%.55s\n", item.Number, item.User.Login, item.Title)
		}
		fmt.Println()
	}
}
