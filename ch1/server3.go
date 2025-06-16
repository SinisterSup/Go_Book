package ch1

// package main

import (
	"fmt"
	"image"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	// sync
)

func server3demo() {
	// func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// var (
// 	mu    sync.Mutex
// 	count int
// )

func homeHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Welcome to the home page! Your requested path is %q\n", r.URL.Path)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "This is a sample about page.! The number of calls to home page are: %d\n", count)
	mu.Unlock()
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissajousGif(w)
}

//
// var wbPalette = []color.Color{color.White, color.Black}
//
// const (
// 	whiteIndex = 0 // first color in palette
// 	blackIndex = 1 // next color in palette
// )

func lissajousGif(out io.Writer) {
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

/* Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
* For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20
* sets the number of cycles to 20 instead of the default 5. Use the strconv.Atio func
* to convert the string parameter into an integer. You can see its documentation with
* go do strconv.Atoi. */
