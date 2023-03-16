// see page 161

// Coloredpoint demonstrates struct embedding.
package main

//!+decl
import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }
type Coloredpoint struct {
	Point
	Color color.RGBA
}

//!-decl

// used code from https://github.com/adonovan/gopl.io/blob/master/ch6/coloredpoint/main.go to complete.
func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	//!+main
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = Coloredpoint{Point{1, 1}, red}
	var q = Coloredpoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	//!-main
}

/*
//!+error
	p.Distance(q) // compile error: cannot use q (coloredPoint) as Point
//!-error
*/

func init() {
	//!+methodexpr
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "{2 4}"
	fmt.Printf("%T\n", distance) // "func(*Point, float64)"
	//!-methodexpr
}

func init()  {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//!+indirect
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point	// p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
	//!-indirect
}