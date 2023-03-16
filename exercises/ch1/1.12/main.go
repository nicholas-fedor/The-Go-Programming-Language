// Exercise 1.12
// Page 22
//
// Modify the Lissajous server to read parameter values from the URL.
// For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20
// instead of the default 5.
// Use strconv.Atoi function to convert the string parameter into an
// integer.
// You can see its documentation with the command:
// go doc strconv.Atoi

// Results:
// Used solution from https://github.com/torbiak/gopl/blob/master/ex1.12/main.go
// Split handler function.

// Lissajous server displays the Lissajous GIF via a web server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Default of 5 x-oscillator revolutions unless specified via URL query.
	var cycles = 5
	cyclesStr := r.FormValue("cycles")
	if cyclesStr != "" {
		var err error
		cycles, err = strconv.Atoi(cyclesStr)
		if err != nil {
			fmt.Fprintf(w, "Error: Incorrect cycles parameter: %s", err)
			return
		}
	}

	lissajous(cycles, w)
}

// Lissajous generates a random lissajous GIF.
func lissajous(cycles int, out io.Writer) {
	// Image sequence is deterministic unless rand is seeded with the time.
	rand.Seed(time.Now().UTC().UnixNano())
	
	const (
		res     = 0.001 // angular resolution
		size    = 250   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	// Allocates for type []color.Color, length 0, capacity of nframes.
	palette := make([]color.Color, 0, nframes)
	// Appends the palette with black as first entry.
	palette = append(palette, color.RGBA{0, 0, 0, 255})

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// Loops through nframes.
	for i := 0; i < 255; i++ {

		// Randomizes the color.
		c := color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 255}
		palette = append(palette, c)

		// Generates the individual GIF images.
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paletteIndex := uint8(len(palette)-1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), paletteIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		return
	}
}
