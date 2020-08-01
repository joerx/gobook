package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"

	"github.com/lucasb-eyer/go-colorful"
)

func mkPalette(max int) []color.Color {
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

type lissParams struct {
	cycles  int
	size    int // image size
	nframes int // num frames
	delay   int // delay in ms
}

func lissajous(out io.Writer, p lissParams) error {

	const res = 0.001
	palette := mkPalette(p.nframes)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0

	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		colIdx := uint8(i % (len(palette) - 1))

		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5), colIdx+1)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}

	return gif.EncodeAll(out, &anim)
}
