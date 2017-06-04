package main

import (
	"image"
	"image/color"
)

// An Image is a rectangular grid of Colors. I'm not using the image.Image
// interface for now because my needs are a little different.
type Image struct {
	Width, Height int
	Pix           []Color
}

func NewImage(width, height int) *Image {
	return &Image{
		Width:  width,
		Height: height,
		Pix:    make([]Color, width*height),
	}
}

func (i *Image) At(x, y int) Color {
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		panic("Out of bounds for At() on Image.")
	}
	return i.Pix[y*i.Width+x]
}

func (i *Image) Set(x, y int, c Color) {
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		panic("Out of bounds for Set() on Image.")
	}
	i.Pix[y*i.Width+x] = c
}

// ToneMap scales down the range of colors to 32-bit RGBA. Right now it uses a
// simplistic heuristic: it just scales the values linearly such that the most
// intense channel value is 0xFF.
func (i *Image) ToneMap() *image.RGBA {
	var max float64
	for j := 0; j < i.Width*i.Height; j++ {
		c := i.Pix[j]
		if c.R < 0 || c.G < 0 || c.B < 0 {
			panic("Negative color values")
		}
		if c.R > max {
			max = c.R
		}
		if c.G > max {
			max = c.G
		}
		if c.B > max {
			max = c.B
		}
	}
	factor := float64(0xFF) / max
	img := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))
	for x := 0; x < i.Width; x++ {
		for y := 0; y < i.Height; y++ {
			c := i.At(x, y)
			rgba := color.RGBA{
				R: uint8(c.R * factor),
				G: uint8(c.G * factor),
				B: uint8(c.B * factor),
				A: 0xFF,
			}
			img.SetRGBA(x, y, rgba)
		}
	}
	return img
}
