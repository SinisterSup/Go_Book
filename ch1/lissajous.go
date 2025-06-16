package ch1

// package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	// "os"
)

var wbPalette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// func main() {
// 	// lissajous(os.Stdout)
// 	lissajousGreen(os.Stdout)
// }

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, wbPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1 // increment phase for next frame
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
}

/* Exercise 1.5: Change the Lissajous program's color palette to green on black, for added autheniticity. \
* To create the web color #rrggbb, use color.RGBA{0xrr, 0xgg, 0xbb, 0xff}. */

func lissajousGreen(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	greenColor := color.RGBA{R: 0, G: 255, B: 0, A: 255} // green
	greenBlackPalette := []color.Color{greenColor, color.Black}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, greenBlackPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1 // increment phase for next frame
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
}

/* Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more values
* to palette and then displaying them by changing the third argument of Set-ColorIndex in some interesting way. */
// func AmuseLissajous(out io.Writer) {
// 	const (
// 		cycles  = 5     // number of complete x oscillator revolutions
// 		res     = 0.001 // angular resolution
// 		size    = 100   // image canvas covers [-size..+size]
// 		nframes = 64    // number of animation frames
// 		delay   = 8     // delay between frames in 10ms units
// 	)
// 	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
// 	anim := gif.GIF{LoopCount: nframes}
// 	phase := 0.0 // phase difference
// 	for i := 0; i < nframes; i++ {
// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
// 		img := image.NewPaletted(rect, wbPalette)
// 		for t := 0.0; t < cycles*2*math.Pi; t += res {
// 			x := math.Sin(t)
// 			y := math.Sin(t*freq + phase)
// 			colorIndex := uint8((int(x*size+size) + int(y*size+size)) % len(wbPalette))
// 			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
// 		}
// 		phase += 0.1 // increment phase for next frame
// 		anim.Delay = append(anim.Delay, delay)
// 		anim.Image = append(anim.Image, img)
// 	}
// 	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
// }
