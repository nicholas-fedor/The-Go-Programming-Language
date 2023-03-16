// Exercise 2.2
// Page 44
//
// Prompt:
// Write a general-purpose unit-conversion program analogous to cF
// that reads numbers from its command-line arguments or from the
// standard input if there are no arguments, and converts each
// number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms,
// and the like.

// Program outputs numeric arguments into a variety of temperature, weight, and length formats.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	tempconv "gopl.io/exercises/ch2/2.2/tempconv"
	weightconv "gopl.io/exercises/ch2/2.2/weightconv"
	lenconv "gopl.io/exercises/ch2/2.2/lenconv"
)

func main() {
	x := len(os.Args[1:])
	switch {
	case x > 0:
		for _, arg := range os.Args[1:] {
			printTemperatures(arg)
			printWeights(arg)
			printLengths(arg)
		}
	default:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter Your Input (number) Below:")
		for scanner.Scan() {
			printTemperatures(scanner.Text())
			printWeights(scanner.Text())
			printLengths(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Reading standard input: ", err)
		}
	}
}

func printTemperatures(arg string) {
	// Converts args into float64
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad input: %v\n", err)
		return
	}

	// Horizontal line separator
	sep := strings.Repeat("-", 40)

	// Converts input as temperature variables.
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)

	// Temperature outputs
	fmt.Printf("%s\n", sep)
	fmt.Printf("\tInput: %.2f\n", t)
	fmt.Printf("%s\n", sep)
	fmt.Println("\tTemperature Conversions")
	fmt.Printf("%s\n", sep)
	fmt.Printf("Fahrenheit: %s\n", f)
	fmt.Printf("\tCelsius: %.2f°C\n", tempconv.FToC(f))
	fmt.Printf("\tKelvin: %.2f°K\n", tempconv.FToK(f))
	fmt.Printf("%s\n", sep)
	fmt.Printf("Celsius: %s\n", c)
	fmt.Printf("\tFahrenheit: %.2f°F\n", tempconv.CToF(c))
	fmt.Printf("\tKelvin: %.2f°K\n", tempconv.CToK(c))
	fmt.Printf("%s\n", sep)
	fmt.Printf("Kelvin: %s\n", k)
	fmt.Printf("\tFahrenheit: %.2f°F\n", tempconv.KToF(k))
	fmt.Printf("\tCelsius: %.2f°C\n", tempconv.KToC(k))
	fmt.Printf("%s\n", sep)
}

func printWeights(arg string) {
	// Convert args into float64
	w, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad Input: %v\n", err)
		return
	}

	// Horizontal line separator
	sep := strings.Repeat("-", 40)

	// Converts input as weight variables.
	l := weightconv.Pounds(w)
	k := weightconv.Kilograms(w)

	// Weight outputs
	fmt.Printf("%s\n", sep)
	fmt.Printf("\tWeight Conversions\n")
	fmt.Printf("%s\n", sep)
	fmt.Printf("Pounds: %.2flbs\n", l)
	fmt.Printf("\tKilograms: %.2fkgs\n", weightconv.LbsToKgs(l))
	fmt.Printf("%s\n", sep)
	fmt.Printf("Kilograms: %.2fkgs\n", k)
	fmt.Printf("\tPounds: %.2flbs\n", weightconv.KgsToLbs(k))
	fmt.Printf("%s\n", sep)
}

func printLengths(arg string) {
	// Convert args into float64
	l, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad Input: %v\n", err)
		return
	}

	// Horizontal line separator
	sep := strings.Repeat("-", 40)

	// Converts input as length variables.
	f := lenconv.Feet(l)
	m := lenconv.Meters(l)

	// Length Outputs
	fmt.Printf("%s\n", sep)
	fmt.Printf("\tLength Conversions\n")
	fmt.Printf("%s\n", sep)
	fmt.Printf("Feet: %.2fft\n", f)
	fmt.Printf("\tMeters: %.2fmeters\n", lenconv.FtToMeters(f))
	fmt.Printf("%s\n", sep)
	fmt.Printf("Meters: %.2fmeters\n", m)
	fmt.Printf("\tFeet: %.2fft\n", lenconv.MetersToFt(m))
	fmt.Printf("%s\n", sep)
}