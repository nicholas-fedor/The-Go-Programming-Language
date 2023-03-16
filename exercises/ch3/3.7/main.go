// Exercise 3.7
// Page 62
//
// Prompt:
// Another simple fractal uses Newton's method to find complex
// solutions to a function, such as z⁴ - 1 = 0.
// * Shade each starting point by the number of iterations required
//   to get close to one of the four roots.
// * Color each point by the root it approaches.

// Additional information.
// Newton's Method
// https://en.wikipedia.org/wiki/Newton%27s_method
//
// Root finding algorithm which produces successively better
// approximations to the roots (or zeros) of a function.
//
// xₙ₊₁ = xₙ - ( ƒ(xₙ) / ƒ'(xₙ) )

// Program emits a PNG image of a fractal.
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

func main() {
	webAddress := "localhost:3000"
	log.Printf("Webserver starting at address: http://%s", webAddress)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(webAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	generateFractal(w)
	log.Println("Outputting Mandelbrot.png")
}

func generateFractal(out io.Writer) {
	// Image dimensions and boundaries for x, y plane.
	const (
		width, height          = 1024, 1024
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
	)

	// Boundary scales for x, y plane.
	xScale := float64(xMax - xMin)
	yScale := float64(yMax - yMin)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yScale) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xScale) + xMin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonColor(z))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Printf("PNG Encoding Error: %v", err)
	}
}

// newtonColor implements Newton's method:
// zₙ₊₁ = zₙ - ( ƒ(zₙ) / ƒ'(zₙ) )
func newtonColor(z complex128) color.Color {
	// Increasing these will increase compute cycles to output the fractal.
	const (
		iterations   = 200
		contrast     = 30
		significance = 1e-8
	)

	for i := 0; i < iterations; i++ {
		// The syntatic sugar of (z-=) initially threw me off.
		// This is equivalent to "z = z - ( zFunc(z) / zFuncPrime(z) )".
		z -= zFunc(z) / zFuncPrime(z)
		if cmplx.Abs(zFunc(z)) < significance {
			return color.Gray{255 - contrast*uint8(i)}
		}
	}
	return color.Black
}

// ƒ(z) = z⁴ - 1.
func zFunc(z complex128) complex128 {
	return cmplx.Pow(z, 4) - 1
}

// ƒ'(z) = 4z³
func zFuncPrime(z complex128) complex128 {
	return 4 * cmplx.Pow(z, 3)
}
