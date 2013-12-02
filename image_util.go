package main

import (
	"fmt"
)

func Downsample(img *Image, factor int) (*Image, error) {
	width := img.Width
	height := img.Height
	if width%factor != 0 || height%factor != 0 {
		return nil, fmt.Errorf("Bad image dimensions for supersampling rate %d\n", factor)
	}
	dWidth := width / factor
	dHeight := height / factor
	dImg := NewImage(dWidth, dHeight)
	colors := []Color{}
	for dx := 0; dx < dWidth; dx++ {
		for dy := 0; dy < dHeight; dy++ {
			startX := dx * factor
			startY := dy * factor
			colors = colors[:0]
			for x := startX; x < startX+factor; x++ {
				for y := startY; y < startY+factor; y++ {
					colors = append(colors, img.At(x, y))
				}
			}
			dImg.Set(dx, dy, ColorAvg(colors))
		}
	}
	return dImg, nil
}

func ColorAvg(colors []Color) Color {
	var r, g, b float64
	for _, c := range colors {
		r += c.R
		g += c.G
		b += c.B
	}
	n := float64(len(colors))
	return Color{
		R: r / n,
		G: g / n,
		B: b / n,
	}
}
