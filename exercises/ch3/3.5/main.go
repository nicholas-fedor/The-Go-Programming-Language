// Exercise 3.5
// Page 62
//
// Prompt:
// Implement a full-color Mandelbrot set using the function
// image.NewRGBA and the type color.RGBA or color.YCbCr.

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// This has been modified to output the image via a webserver.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
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
// grayscale raster image representing the -2 to +2 portion
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
func generateMandelbrot(out io.Writer) {
	// Separated out rect from img.
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, colorMandelbrot(z))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Printf("PNG Encode Error: %v", err)
	}
}

func colorMandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// Implements the ability to manipulate the foreground (exterior) color.
			r := uint8(0)
			g := uint8(255 - contrast*n)
			b := uint8(255)
			exteriorColor := color.RGBA{r, g, b, 255}
			return exteriorColor
		}
	}

	// normal for most visualizations to have dark interior.
	interiorColor := color.Black
	return interiorColor
}
