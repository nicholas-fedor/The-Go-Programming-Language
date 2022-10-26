// Exercise 1.6
// Page 15

// Prompt:
// Modify the Lissajous program to produce images in multiple colors
// by adding more values to palette and then displaying them by
// changing the third argument of Set-ColorIndex in some
// interesting way.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Added file I/O, in lieu of invoking in the terminal.
	// If the file doesn't exist, create it, or append to the file.
	f, err := os.OpenFile("lissajous.gif", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lissajous(f)
}

// Lissajous generates a random lissajous GIF.
func lissajous(out io.Writer) {
	// Image sequence is deterministic unless rand is seeded with the time.
	rand.Seed(time.Now().UTC().UnixNano())

	const (
		cycles  = 5     // the number of complete x oscillator revolutions
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
			paletteIndex := uint8(len(palette) - 1)
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
