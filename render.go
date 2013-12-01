package main

import (
	"fmt"
	"image"
	"image/color"
)

type Rendering struct {
	*Scene
	HPixels int
}

// TODO: Also return a progress chan
func (r *Rendering) Render() *image.RGBA {
	w := r.HPixels
	h := int(float64(w) * r.Camera.Aspect)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	fmt.Printf("%#v\n", r.Scene.Camera.Vantage())
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.SetRGBA(x, y, color.RGBA{0, 0xff, 0, 0xff})
		}
	}
	return img
}
