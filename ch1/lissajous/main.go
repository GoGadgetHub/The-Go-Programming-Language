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
	"os"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{255, 0, 0, 255}}

const (
	whiteIndex = 0
	redIndex   = 1
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handle := func(w http.ResponseWriter, request *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handle)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的 x 震荡器变化个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 画布包含 [-size, +size]
		nframes = 64    // 帧数
		delay   = 8     // 10ms单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y震荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // 内容容器
		img := image.NewPaletted(rect, palette)      //
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Cos(t) + math.Sin(t*freq+phase)
			if y >= 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
				img.SetColorIndex(size+int(x*size+0.5), size+int(-y*size+0.5), redIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
