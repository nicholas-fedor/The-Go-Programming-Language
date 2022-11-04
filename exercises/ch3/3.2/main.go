// Exercise 3.2
// Page 60
//
// Prompt:
// Experiment with visualizations of other functions from the math package.
// Can you produce an egg box, moguls, or a saddle?

// Surface computes an SVG rendering of a 3-D surface
// and outputs it via a webserver on localhost:3000.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvase size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// zFunc type allows the program to utilize different functions for generating the x, y coordinates.
type zFunc func(x, y float64) float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// handler takes the http request and generates an output.
// defaults to the original example SVG, including if the request is malformed.
func handler(w http.ResponseWriter, r *http.Request) {
	// Setting the header allows the browser to interpret the output as a SVG image.
	w.Header().Set("Content-Type", "image/svg+xml")
	// HTTP requests via "localhost:3000?visualization="
	switch r.FormValue("visualization") {
	// "localhost:3000?visualization=eggbox"
	case "eggbox":
		fmt.Println("eggbox")
		printSVG(w, eggbox)
	// "localhost:3000?visualization=moguls"
	case "moguls":
		fmt.Println("moguls")
		printSVG(w, moguls)
	// "localhost:3000?visualization=saddle"
	case "saddle":
		fmt.Println("saddle")
		printSVG(w, saddle)
	// "localhost:3000" defaults to original surface plot of sin(r)/r
	default:
		fmt.Println("default")
		printSVG(w, defaultF)
	}
}

// printSVG outputs the requested SVG to the io.Writer for the http handler.
func printSVG(out io.Writer, r zFunc) {

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, r)
			bx, by := corner(i, j, r)
			cx, cy := corner(i, j+1, r)
			dx, dy := corner(i+1, j+1, r)

			// Prints polygons if points are valid numbers.
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Fprintf(out, "<polygon points='%g, %g, %g, %g, %g, %g, %g, %g '/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

// corner translates 3-d coordinates to 2-d. 
// variable r represents the function used for generating x, y inputs.
func corner(i, j int, r zFunc) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := r(x, y)

	// Project (x, y, z) isometrically onto a 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// isFinite returns true if values are not invalid or infinite numbers.
func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

// defaultF provides the coordinates to printSVG for generating the original example svg.
func defaultF(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r)    // r
}

// eggbox provides the coordinates to printSVG for generating the eggbox svg.
func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

// saddle provides the coordinates to printSVG for generating the saddle svg.
func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

// moguls provides the coordinates to printSVG for generating the moguls svg
func moguls(x, y float64) float64 {
	return math.Abs(math.Sin(x) * math.Sin(y))
}