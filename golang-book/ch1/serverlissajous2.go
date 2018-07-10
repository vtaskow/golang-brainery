package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var DEFAULT_CYCLES int = 20

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0xdd, 0xff},
	color.RGBA{0xff, 0x99, 0x00, 0xff},
	color.RGBA{0x00, 0x23, 0x4c, 0xff},
	color.RGBA{0xff, 0x99, 0x99, 0xff},
	color.RGBA{0x00, 0x93, 0xdd, 0xff},
	color.RGBA{0xff, 0x23, 0x00, 0xff},
	color.RGBA{0x0f, 0xdd, 0x49, 0xff},
	color.RGBA{0xff, 0x23, 0x93, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.FormValue("cycles"))
	if err != nil {
		cycles = DEFAULT_CYCLES
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		randIndex := 1 + rand.Intn(len(palette)-2)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
