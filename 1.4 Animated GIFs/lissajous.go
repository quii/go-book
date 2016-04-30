package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var (
	green   = color.RGBA{0x00, 0xAA, 0x00, 0xff}
	red     = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	palette = []color.Color{color.Black, green, red}
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // canvas size
		delay   = 8     // delay between frames in 10ms units
		nframes = 64
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			randomColorIndex := uint8(rand.Int31n(3) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
