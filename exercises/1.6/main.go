// Exercise 1.6
// Page 15

// Prompt:
// Modify the Lissajous program to produce images in multiple colors
// by adding more values to palette and then displaying them by
// changing the third argument of Set-ColorIndex in some
// interesting way.

// Lissajous generates GIF animations of random Lissajous figures.
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

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	// Copied from https://github.com/torbiak/gopl/blob/master/ex1.6/main.go.
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{uint8(55 + 200*scale), uint8(55 + 200*scale), uint8(55 + 200*scale), 255}
		palette = append(palette, c)
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paletteIndex := uint8(i%(len(palette)-1) + 1)
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
