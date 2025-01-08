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
	"net/url"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	fmt.Println("Web server has been created!")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.URL.String())
		q := u.Query()
		number, _ := strconv.Atoi(q.Get("cycles"))
		//fmt.Fprintf(w, "query paran:", q.Get("cycles"))
		//fmt.Println(q.Get("cycles"))
		mu.Lock()
		//sum := 0
		//for i := 0; i < 1000000; i++ {
		//	sum += 1
		//}
		//fmt.Fprintf(w, "count: %d\n", sum)
		lissajous(w, number)
		mu.Unlock()
	})
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.Black, color.RGBA{0, 255, 65, 255}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles int) {

	const (
		res     = 0.1
		sizeY   = 200
		sizeX   = 200
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += 0.0001 {
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
