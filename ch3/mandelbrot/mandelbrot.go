// package mandelbrot
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) corresponds to complex number z.
			img.Set(px, py, mandelbrotColor(z))
		}
	}
	png.Encode(os.Stdout, img) // Note: ignoring errors
}

func mandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 10

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n} // Color based on escape time
		}
	}
	return color.Black
}

/* Exercise 3.5: Implement a full-color version of Mandelbrot set using the function image.NewRGBA
* and the type color.RGBA or color.YCbCr. */

/* Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the
* color value at several points within each pixel and taking the average. The simplest method is
* to divide each pixel into four "subpixels". Implement it. */

/* Exercise 3.7: Another simple fractal uses Newton's method to find complex solutions to a
* function such as z^4 - 1 = 0. Shade each starting point by the number of iterations required to
* get close to one of the four roots. Color each point by the root it approaches. */

/* Exercise 3.9: Write a web server that renders fractals and writes the image data to the client.
* Allow the client to specify the x,y, and zoom values as parameters to the HTTP request. */
