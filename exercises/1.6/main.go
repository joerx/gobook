package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"

	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	if err := lissajous(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func makePalette(max int) []color.Color {
	const (
		sat    = 0.84
		bright = 0.79
	)

	step := 360.0 / float64(max)
	p := make([]color.Color, max)
	p[0] = color.Black

	for i := 0; i < max-1; i++ {
		hue := float64(i) * step
		c := colorful.Hsv(hue, 0.84, 0.79)
		p[i+1] = c
	}

	return p
}

func lissajous(out io.Writer) error {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100 // image size
		nframes = 64  // num frames
		delay   = 8   // delay in ms
	)

	palette := makePalette(nframes)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colIdx := uint8(i % (len(palette) - 1))

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colIdx+1)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return gif.EncodeAll(out, &anim)
}
