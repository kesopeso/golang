package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
)

var pallete = []color.Color{color.White, color.Black, color.RGBA{0, 255, 0, 1}, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{255, 255, 0, 1}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
	redIndex   = 3
	blueIndex  = 4
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}))

	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 128
		delay   = 1
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)

		for j := 0; j < 2*size+1; j++ {
			for k := 0; k < 2*size+1; k++ {
				img.SetColorIndex(j, k, blackIndex)
			}
		}

		for t := float64(0); t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			var colorIndex uint8
			switch {
			case i%2 == 0:
				colorIndex = whiteIndex
			case i%3 == 0:
				colorIndex = greenIndex
			case i%5 == 0:
				colorIndex = redIndex
			case i%7 == 0:
				colorIndex = blueIndex
			default:
				colorIndex = whiteIndex
			}

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
