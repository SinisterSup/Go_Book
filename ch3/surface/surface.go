// package surface
// Surface computes an SVG rendering of a 3-D surface function
// output of function sin(r) / r , where r is sqrt(x^2 + y^2)
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange...+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.2        // pixels per z unit
	angle         = math.Pi / 6         // angle for x and y axes (=30 degree)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30 degrees), cos(30 degrees)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey, fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // distance from (0, 0)
	return math.Sin(r) / r // surface function
}

/* Exercise 3.1: If the function f returns a non-finite float64 value, the SVG file will contain
* invalid <polygon> elements (although many SVG renderers handle this gracefully). Modify
* the program to skip invalid polygons. */

/* Exercise 3.2: Experiment with visualizations of other functions from the math package.
* Can you produce an egg box, moguls, or a saddle? */

/* Exercise 3.3: Color each polygon based on its height, so that the peaks are colored red
* (#ff0000) and the valleys are colored blue (#0000ff). You can use the color package to convert. */

/* Exercise 3.4: Following the approach of the Lissajous program (ch1/lissajous.go), construct a web
* server that computes surfaces and writes SVG data to client. The server must set the Content-Type
* header like -> w.Header().Set("Content-Type", "image/svg+xml"). */
