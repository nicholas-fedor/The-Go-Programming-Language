// Exercise 4.12
// Page 113
//
// Prompt:
// The popular web comic xkcd has a JSON interface.
// For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of many favorites.
// Download each URL (once!) and build an offline index.
// Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.
//
// Development Notes:
// Used bard.google.com to help, but certainly not perfect.

// Program xkcd uses a offline index to print the URL and transcript of each comic that matches a search term provided via the command line.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type ComicIndex struct {
	Comics []Comic
}

func downloadComics() (ComicIndex, error) {
	// Check if the index exists
	if _, err := os.Stat("xkcd.idx"); os.IsNotExist(err) {
		// Index does not exist, so download the comics
		fmt.Println("Index does not exist, downloading comics...")

		// Get the total number of comics
		resp, err := http.Get("https://xkcd.com/info.json")
		if err != nil {
			return ComicIndex{}, err
		}

		defer resp.Body.Close()

		// Decode the response
		var info struct {
			TotalComics int `json:"total_comics"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
			return ComicIndex{}, err
		}

		// Create a new map to store the comics
		var comics []Comic

		// Iterate over all of the comics
		for i := 1; i <= info.TotalComics; i++ {
			// Make a request to the JSON interface
			resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", i))
			if err != nil {
				return ComicIndex{}, err
			}

			defer resp.Body.Close()

			// Decode the response
			var comic Comic

			if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
				return ComicIndex{}, err
			}

			// Add the comic to the map
			comics = append(comics, comic)
		}

		// Save the index to a file
		f, err := os.Create("xkcd.idx")
		if err != nil {
			return ComicIndex{}, err
		}

		defer f.Close()

		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "  ")
		encoder.Encode(ComicIndex{Comics: comics})

		return ComicIndex{Comics: comics}, nil
	} else {
		// Index exists, so load it from the file
		fmt.Println("Index exists, loading...")

		// Open the index file
		f, err := os.Open("xkcd.idx")
		if err != nil {
			return ComicIndex{}, err
		}

		defer f.Close()

		// Decode the index
		decoder := json.NewDecoder(f)
		comics := make([]Comic, 0)

		if err := decoder.Decode(&comics); err != nil {
			return ComicIndex{}, err
		}

		return ComicIndex{Comics: comics}, nil
	}
}

func needsUpdate(indexFile, comicsWebsite string) bool {
	// Get the last modified date of the index file
	indexLastModified, err := os.Stat(indexFile)
	if err != nil {
		return false
	}

	// Get the last modified date of the comics website
	resp, err := http.Head(comicsWebsite)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	// Get the last modified date from the Date header
	comicsLastModified, err := time.Parse(http.TimeFormat, resp.Header.Get("Date"))
	if err != nil {
		return false
	}

	// Get the last modified time of the index file
	indexLastModifiedTime := indexLastModified.ModTime()

	// Compare the last modified dates
	return comicsLastModified.After(indexLastModifiedTime)
}

func searchComics(index ComicIndex, searchTerm string) []Comic {
	// Search for the comic in the index
	results := []Comic{}
	for _, comic := range index.Comics {
		// Check if the comic matches the search term
		if strings.Contains(comic.Title, searchTerm) {
			results = append(results, comic)
		}
	}

	return results
}

func main() {
	// Get the search term from the command line
	searchTerm := strings.TrimSpace(os.Args[1])

	// Check if the index needs to be updated
	if needsUpdate("xkcd.idx", "https://xkcd.com/") {
		// Download the comics
		index, err := downloadComics()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Search for the comic in the index
	results := searchComics(index, searchTerm)

	// Print the results
	for _, comic := range results {
		fmt.Println(comic.Image, comic.Title)
	}
}
