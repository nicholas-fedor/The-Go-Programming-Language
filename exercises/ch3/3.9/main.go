// Exercise 3.9
// Page 63
//
// Prompt:
// Write a web server that renders fractals and writes the image data to the client.
// Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.

// Mandelbrot emits a PNG image of the Mandelbrot fractal to a web server.
// In this current iteration, the x, y, and zoom values will stretch and skew the mandelbrot fractal image.
// This is far from a perfect implementation, but will allow the user to see the image with the default values,
// if no input is provided.
// To keep things simple, I opted to limit user inputs to single x, y variables, instead of xMin/xMax and yMin/yMax.
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
	// Takes user's URL request.
	// i.e. http://localhost:3000/?x=4&y=4&zoom=100
	// If the request is malformed, it will respond with the error.
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Setting defaults to string "0" if any form value is missing.
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

	// Parsing into Float64 allows for requests to include
	// values with negative numbers and decimals.
	// Malformed values will respond with the error.
	xValue, err := strconv.ParseFloat(xRequest, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	yValue, err := strconv.ParseFloat(yRequest, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	zoomValue, err := strconv.ParseFloat(zoomRequest, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sanity check via printing out the requests in the terminal.
	log.Printf("xRequest: %f, yRequest: %f, zoomRequest: %f", xValue, yValue, zoomValue)

	// Passing the values to the renderer.
	renderFractal(w, xValue, yValue, zoomValue)
}

func renderFractal(w http.ResponseWriter, xValue, yValue, zoomValue float64) {
	const (
		width, height = 1024, 1024
		yMax, yMin    = 2, -2
		xMax, xMin    = 2, -2
	)

	// Initialized variables outside of the loops for sanity.
	var px, py float64
	var x, y float64
	var z complex128

	// Initializing rect object outside of img.NewRGBA for sanity.
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	// Main rendering loop.
	for py = 0; py < height; py++ {
		// Uses default yMax and yMin if yValue is 0.
		// This can be improved, as prevents user from using 0.
		if yValue == 0 {
			y = py/height*(yMax-yMin) + yMin
		} else {
			y = py/height*(yValue*2) + (yValue * -1)
		}

		for px = 0; px < width; px++ {
			// Uses default xMax and xMin if xValue is 0.
			// This can be improved, as prevents user from using 0.
			if xValue == 0 {
				x = px/width*(xMax-xMin) + xMin
			} else {
				x = px/width*(xValue*2) + (xValue * -1)
			}

			// If zoomValue omitted, then zoomValue variable not used in
			// calculating z.
			if zoomValue == 0 {
				z = complex(x, y)
			} else {
				// Otherwise, both x and y are scaled.
				// value + (value * zoomValue%)
				// If zoomValue is negative, then will reduce overall value.
				z = complex(x+x*(zoomValue/100), y+y*(zoomValue/100))
			}

			img.Set(int(px), int(py), mandelbrot(z))
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
