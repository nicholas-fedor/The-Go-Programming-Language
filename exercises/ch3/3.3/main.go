// Exercise 3.3
// Page 60
//
// Prompt:
// Color each polygon based on its height,
// so that the peaks are colored red (#ff0000) and valleys blue (#0000ff).

// This exercise adds color gradients to the book's surface example in chapter 3.
// Solution mirrored from https://github.com/kdama/gopl/blob/master/ch03/ex03/main.go.
package main

import (
	"fmt"
	"io"
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

const (
	Red  = "#FF0000" // Peaks
	Blue = "#0000FF" // Valleys
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// zFunc is used for the visualization function requested by the user.
type zFunc func(x, y float64) float64

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	switch r.FormValue("visualization") {
	case "eggbox":
		fmt.Println("Outputting eggbox svg")
		printSVG(w, eggbox)
	case "moguls":
		fmt.Println("Outputting moguls svg")
		printSVG(w, moguls)
	case "saddle":
		fmt.Println("Outputting saddle svg")
		printSVG(w, saddle)
	case "original":
		fmt.Println("Outputting original svg")
		printSVG(w, defaultF)
	default:
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<div><h1>Home</h1><h2>Use the ?visualization= html form request to access eggbox, moguls, saddle, or original svg options.</h2></div>")
	}
}

func printSVG(out io.Writer, r zFunc) {
	
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: gray; fill: white; stroke-width: 0.7' "+
	"width='%d' height='%d'>", width, height)
	
	maxHeight, minHeight := getMaxMinHeight(r)

	// Loops through the cells.
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// Calculations for the SVG polygon point values.
			ax, ay := getPolygonPoints(i+1, j, r)
			bx, by := getPolygonPoints(i, j, r)
			cx, cy := getPolygonPoints(i, j+1, r)
			dx, dy := getPolygonPoints(i+1, j+1, r)

			// Calculates the color for the SVG polygon points.
			_, _, z := getPoints(i, j, r) // ignoring x and y
			color := getColor(z, maxHeight, minHeight)

			// Checks from exercise 3.1 to skip invalid polygon points.
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				// Polygon color based on height.
				fmt.Fprintf(out, "<polygon stroke='%s' ", color)
				// Polygon points.
				fmt.Fprintf(out, "points='%g, %g, %g, %g, %g, %g, %g, %g '/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(out, "</svg>") // terminate SVG tag.
}

// getPoints returns the x, y, z values.
func getPoints(i, j int, r zFunc) (float64, float64, float64) {
	// Computes x and y based on i and j.
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	
	// Computes z based on the selected visualization.
	z := r(x, y)
	
	return x, y, z
}

// Compute and return sx, sy.
func getPolygonPoints(i, j int, r zFunc) (float64, float64) {
	x, y, z := getPoints(i, j, r)
	// Compute sx, xy
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale // accounts for surface height.
	
	return sx, sy
}

// Calculates the max and min of the surface height(z).
func getMaxMinHeight(r zFunc) (float64, float64) {
	// maxHeight is the maximum surface height.
	maxHeight := math.NaN()
	// minHeight is the minimum surface height.
	minHeight := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := getPoints(i, j, r) // ignoring x, y

			if isFinite(z) {
				if math.IsNaN(maxHeight) || maxHeight < z {
					maxHeight = z
				}
				if math.IsNaN(minHeight) || minHeight > z {
					minHeight = z
				}
			}
		}
	}
	return maxHeight, minHeight
}

// getColor returns the color values as hex-coded strings.
func getColor(height, maxHeight, minHeight float64) string {
	if !isFinite(height) || !isFinite(maxHeight) || !isFinite(minHeight) {
		return Blue
	}
	// n is the (canvas height - minimum surface height) / (maximum surface height - minimum surface height) 
	// scaled (multiplied) by the max color scale value (255)
	n := int((height - minHeight) / (maxHeight - minHeight) * 255)

	// Calculations for the RGB hex values.
	rr := fmt.Sprintf("%02x", n)
	gg := "00"
	bb := fmt.Sprintf("%02x", 255-n)

	return fmt.Sprintf("#%s%s%s", rr, gg, bb)
}

// isFinite returns true if f is a number that is not infinite.
func isFinite(f float64) bool {
	if math.IsInf(f, 0) {
		return false
	}
	if math.IsNaN(f) {
		return false
	}
	return true
}

// Surface height functions
// defaultF returns the original visualization from the example.
func defaultF(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

// eggbox returns the eggbox visualization from exercise 3.2.
func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

// saddle returns the saddle visualization from exercise 3.2.
func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

// moguls returns the moguls visualization from exercise 3.2.
func moguls(x, y float64) float64 {
	return math.Abs(math.Sin(x) * math.Sin(y))
}
