// Exercise 3.6
// Page 62
//
// Prompt:
// Supersampling is a technique to reduce the effect of pixelation
// by computing the color value at several points within each pixel
// and taking the average.
//
// The simplest method is to divide each pixel into four subpixels.
// Implement this method.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// This has been modified to output the image via a webserver.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
)

// Mandelbrot web server at http://localhost:3000
func main() {
	webAddress := "localhost:3000"
	log.Printf("Webserver starting at address: http://%s", webAddress)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(webAddress, nil))
}

// Response handler
func handler(w http.ResponseWriter, r *http.Request) {

	// Sets the response header.
	w.Header().Set("Content-Type", "image/png")

	// Initiates the mandelbrot generator and
	// outputs to the http.ResponseWriter.
	generateMandelbrot(w)
	log.Println("Outputting Mandelbrot.png")
}

// The two nested loops iterate over each point in a 1024x1024
// RGBA raster image representing the -2 to +2 portion
// of the complex plane.
//
// The program tests whether repeatedly squaring and adding
// the number that point represents eventually "escapes"
// the circle of radius 2.
//
// If so, the point is shaded by the number of iterations
// it took to escape.
//
// If not, the value belongs to the Mandelbrot set,
// and the point remains black.
//
// Finally, the program writes to the io.Writer
// the PNG-encoded images of the iconic fractal.

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func generateMandelbrot(out io.Writer) {

	// Separated out rect from img.
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	// Loop upwards y axis from bottom to top max of 1024.
	for py := 0; py < height; py++ {

		// Calculates y. If py = 0, then y = -2.
		// Clarified the math to indicate the positive increase of ymin value
		y := (float64(py) / height * (ymax - ymin)) + ymin

		// Loop across x value from left to right to width 1024.
		for px := 0; px < width; px++ {

			// Calculates x. If px = 0, then x = -2.
			// Clarified the math to indicate the positive increase of xmin value
			x := (float64(px) / width * (xmax - xmin)) + xmin

			// Supersampling by using 2x2 grid of each pixel (px, py) and
			// using their average later in img.Set().
			// Used solution from https://github.com/torbiak/gopl/blob/master/ex3.6/main.go.

			// (xmax - xmin) = scale between upper and lower bounds || +2 and -2 = 4
			// scale / width = scaled width || 4 / 1024 = 1/256 or 0.00390625
			offsetX := []float64{-1 * ((xmax - xmin) / width), ((xmax - xmin) / width)} // (-1/256), 1/256
			offsetY := []float64{-1 * ((ymax - ymin) / height), ((ymax - ymin) / height)}  // (-1/256), 1/256

			subPixels := make([]color.Color, 0)
			
			// i and j are limited to 2, otherwise runtime error: index out of range [2] with length 2.
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offsetX[i], y+offsetY[j]) // 
					subPixels = append(subPixels, colorMandelbrot(z))
				}
			}

			// Sets the color for the pixel at point (px, py).
			img.Set(px, py, colorAvg(subPixels))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Printf("PNG Encode Error: %v", err)
	}
}

// colorAvg takes a slice of colors and averages them.
func colorAvg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors { // NOTE: ignoring the index value.
		rA, gA, bA, aA := c.RGBA()
		// Calculates the average RGBA values for the n number of colors.
		r += uint16(rA / uint32(n))
		g += uint16(gA / uint32(n))
		b += uint16(bA / uint32(n))
		a += uint16(aA / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

// colorMandelbrot takes input of complex number z
// and computes the color value for that px, py location.
func colorMandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := 0; n < iterations; n++ {

		// Repeatedly squares v and adds the number that point represents.
		// Manipulating v to produce a larger value will result in the
		// image zooming outwards.
		v = v*v + z

		// Computes if the value escapes the circle radius of 2.
		if cmplx.Abs(v) > 2 {

			// Implements the ability to manipulate the foreground (exterior) color.
			// logColorScale is used to create a gradient.
			logColorScale := math.Log(float64(n)) / math.Log(float64(iterations))

			// Calculates the RGBA values used by a pixel.
			r := uint8(255 - uint8(logColorScale*255*iterations))
			g := uint8(255 - contrast*n)
			b := uint8(255 - uint8(logColorScale*255))
			a := uint8(255)

			exteriorColor := color.RGBA{r, g, b, a}
			return exteriorColor
		}
	}

	// Normal for most visualizations to have dark interior.
	interiorColor := color.Black
	return interiorColor
}
