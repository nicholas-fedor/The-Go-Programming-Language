// Exercise 3.9
// Page 63
//
// Prompt:
// Write a web server that renders fractals and writes the image data to the client.
// Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.

// Mandelbrot emits a PNG image of the Mandelbrot fractal to a web server.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	address := "localhost:3000"
	log.Printf("Webserver starting: http://%s", address)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(address, mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Loads URL form requests.
	// Error handling for bad user inputs.
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Request values are strings.
	// If omitted by user, defaults to "0".
	xRequest := r.FormValue("x")
	if xRequest == "" {
		xRequest = "0"
	}
	yRequest := r.FormValue("y")
	if yRequest == "" {
		yRequest = "0"
	}
	zoomRequest := r.FormValue("zoom")
	if zoomRequest == "" {
		zoomRequest = "0"
	}

	// String conversion to int.
	// Error handling for bad user inputs.
	xValue, err := strconv.Atoi(xRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	yValue, err := strconv.Atoi(yRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	zoomValue, err := strconv.Atoi(zoomRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calls for fractal PNG.
	renderFractal(w, xValue, yValue, zoomValue)
}

func renderFractal(w http.ResponseWriter, xValue, yValue, zoomValue int) {
	const (
		xMin, yMin, xMax, yMax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("Outputting mandelbrot.png")
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
