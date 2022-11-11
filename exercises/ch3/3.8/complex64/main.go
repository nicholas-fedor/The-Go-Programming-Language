// Exercise 3.8
// Page 63
//
// Prompt:
// Rendering fractals at high zoom levels demands great arithmetic precision.
// Implement the same fractal using four different representations of numbers:
// 	* complex64
// 	* complex128
// 	* big.Float
// 	* big.Rat
// (The latter two types are found in the math/big package.)
// Float uses arbitrary but bounded-precision floating-point.
// Rat uses unbounded-precision rational numbers.
//
// How do they compare in performance and memory usage?
// At what zoom levels do rendering artifacts become visible?

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// The example function has been updated to output via a basic web server.
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

func main() {
	webAddress := "localhost:3000"
	log.Printf("Webserver starting: http://%s", webAddress)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(webAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")

	generateFractal(w)
	log.Println("Outputting Mandelbrot.png")
}

func generateFractal(out io.Writer) {
	const (
		width, height          = 1024, 1024
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
	)

	xScale := float64(xMax - xMin)
	yScale := float64(yMax - yMin)

	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	for py := 0; py < height; py++ {
		y := float64(py)/height*(yScale) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xScale) + xMin
			z := complex(x, y)

			img.Set(px, py, mandelbrot64(z))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Printf("PNG encode error: %v", err)
	}
}

func mandelbrot64(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex64 // complex64 implementation

	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z) // complex64 implementation
		if cmplx.Abs(complex128(v)) > 2 {  // complex64 implementation
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
