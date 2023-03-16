// Exercise 3.12
// Page 74
//
// Prompt:
// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.

// Other attempted solutions:
// https://github.com/torbiak/gopl/blob/master/ex3.12/anagram.go
// * Compares runes.
// https://github.com/Julineo/golang1training/blob/master/3/3.12/main.go
// * This solution looks at string reversal, as opposed to parsing for anagrams.
// https://github.com/kdama/gopl/blob/master/ch03/ex12/main.go
// * Maps runes and then has a custom rune map comparison.
// https://siongui.github.io/2017/05/06/go-check-if-two-string-are-anagram/
// * Sorts the runes for each string and compares the strings.

// Program inspects if two strings are ananagrams of each other.
// Anagrams: https://en.wikipedia.org/wiki/Anagram
//
// Example:
// inputString1 = "Tom Marvolo Riddle"
// inputString2 = "I am Lord Voldemort"
// output = true
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Takes two string inputs, runs the result function,
// and prints the boolean result.
func main() {
	// Input strings
	inputString1 := "Tom Marvolo Riddle"
	inputString2 := "I am Lord Voldemort"

	// Check if the two inputs are anagrams of each other.
	result := checkAnagram(inputString1, inputString2)

	// Print boolean result
	fmt.Println(result)
}

// Takes two string inputs and returns boolean
// if character counts are equivalent.
func checkAnagram(inputString1, inputString2 string) bool {

	// Sets case for input strings to lowercase.
	// Otherwise, will impact character mapping.
	s1 := strings.ToLower(inputString1)
	s2 := strings.ToLower(inputString2)

	// Map the character counts for each string.
	// https://www.geeksforgeeks.org/comparing-maps-in-golang/
	charMap1 := countStringChars(s1)
	charMap2 := countStringChars(s2)

	// Compare if the two character maps are equal.
	result := reflect.DeepEqual(charMap1, charMap2)

	return result
}

// Maps and counts string characters.
// Returns a map of the characters and the count.
// i.e. "satan" = map[a:2 n:1 s:1 t:1]
// Used solution from https://code-maven.com/slides/golang/solution-count-characters-sort-by-frequency.
func countStringChars(s string) map[string]int {
	stringCharMap := make(map[string]int)
	for _, c := range s {
		// Ignores spaces while counting characters.
		if string(c) != " " {
			stringCharMap[string(c)]++
		}
	}
	return stringCharMap
}
