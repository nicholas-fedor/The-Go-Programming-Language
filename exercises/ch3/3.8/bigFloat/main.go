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
	"math"
	"math/big"
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

			img.Set(px, py, mandelbrotBigFloat(z))
		}
	}
	err := png.Encode(out, img)
	if err != nil {
		log.Printf("PNG encode error: %v", err)
	}
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200

	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}

	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}

		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))

		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))

				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
