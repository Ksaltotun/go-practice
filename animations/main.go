package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 65, 255}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.1
		sizeY   = 600
		sizeX   = 800
		nframes = 64
		delay   = 8
	)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*sizeX+1, 2*sizeY+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += 0.0001 {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(sizeX+int(x*sizeX+0.5), sizeY+int(y*sizeY+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
