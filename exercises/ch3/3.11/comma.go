// Exercise 3.11
// Page 74
//
// Prompt:
// Enhance comma so that it deals correctly with
// floating-point numbers and an optional sign.

// Comma takes a floating point value and adds commas.
// Outputs as a string value.
// Development note: There are several type conversions
// and this program could be optimized.
package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// %f will round to precision 6
	float := -100000.50
	fmt.Println(comma(float))
}

func comma(f float64) string {

	// Variables are here for sanity.
	var buf bytes.Buffer
	var sign string
	var inputIntF, inputFloat float64
	var inputIntI int
	var inputIntS string
	var outputInt string
	var outputFloat string

	// TODO:
	// *Validate input (if IsInf or isNaN)
	switch math.IsInf(f, 0) || math.IsNaN(f) {
	// ** End program if float is NaN or infite.
	case true:
		break
		// ** Otherwise parse float64
	case false:
		// Separate out integer and float values of input.
		inputIntF, inputFloat = math.Modf(f)

		// Convert inputIntF to integer type.
		inputIntI = int(inputIntF)

		// Remove negatives from inputIntI and inputFloat and save sign
		if math.Signbit(f) {
			sign = "-"
			inputIntI = inputIntI * -1
			inputFloat = inputFloat * -1
		}

		// Write sign to buffer.
		buf.WriteString(sign) // "-"

		// Handle commas
		// Convert inputIntI to type string.
		inputIntS = strconv.Itoa(inputIntI)
		n := len(inputIntS)

		switch {
		case n <= 3:
			fmt.Fprint(&buf, inputIntS)
		case n > 3:
			commaPosition := len(inputIntS) % 3

			if commaPosition == 0 {
				commaPosition = 3
			}
			outputInt = inputIntS[:commaPosition]
			for i := commaPosition; i < len(inputIntS); i += 3 {
				outputInt += ","
				outputInt += inputIntS[i : i+3]
			}
		}

	}

	// Convert inputFloat to string and trim "0." prefix.
	outputFloat = fmt.Sprintf("%g", inputFloat)
	outputFloat = strings.TrimPrefix(outputFloat, "0.")

	// Write both parts to buffer
	fmt.Fprintf(&buf, "%s.", outputInt)    // 1000.
	fmt.Fprintf(&buf, "%s\n", outputFloat) // 499999

	// Current Output: -1000.499999

	// Return string to main
	return buf.String()
}
