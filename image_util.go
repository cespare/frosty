package main

import (
	"fmt"
	"image"
	"image/color"
)

func Downsample(img image.Image, factor int) (*image.RGBA, error) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	if width%factor != 0 || height%factor != 0 {
		return nil, fmt.Errorf("Bad image dimensions for supersampling rate %d\n", factor)
	}
	dWidth := width / factor
	dHeight := height / factor
	dImg := image.NewRGBA(image.Rect(0, 0, dWidth, dHeight))
	colors := []color.Color{}
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

func ColorAvg(colors []color.Color) color.RGBA {
	var r, g, b, a uint32
	for _, col := range colors {
		r2, g2, b2, a2 := col.RGBA()
		// No worries about overflowing here because these numbers are in [0, 0xFFFF].
		r += r2
		g += g2
		b += b2
		a += a2
	}
	n := uint32(len(colors))
	rgba := color.RGBA{
		R: uint8((r / n) >> 8),
		G: uint8((g / n) >> 8),
		B: uint8((b / n) >> 8),
		A: uint8((a / n) >> 8),
	}
	return rgba
}
