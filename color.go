package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type Color color.RGBA

func (c *Color) UnmarshalJSON(b []byte) error {
	if len(b) < 2 {
		return fmt.Errorf("Bad color string: %s", string(b))
	}
	if b[0] == '"' {
		b = b[1:]
	}
	if b[len(b)-1] == '"' {
		b = b[:len(b)-1]
	}
	col, err := ParseHexColor(string(b))
	if err != nil {
		return err
	}
	*c = col
	return nil
}

// ParseHexColor parses a CSS-style hex color (e.g. #123abc, #fff)
func ParseHexColor(c string) (Color, error) {
	if strings.HasPrefix(c, "#") {
		c = c[1:]
	}
	if len(c) == 3 {
		c = c[:1] + c[:1] + c[1:2] + c[1:2] + c[2:] + c[2:]
	}
	rgba := Color{}
	if len(c) == 6 {
		rgb, err := strconv.ParseUint(c, 16, 32)
		if err != nil {
			return rgba, err
		}
		rgba.R = uint8(rgb >> 16)
		rgba.G = uint8((rgb >> 8) & 0xFF)
		rgba.B = uint8(rgb & 0xFF)
		rgba.A = uint8(0xFF)
		return rgba, nil
	}
	return rgba, fmt.Errorf("Bad color string: %s", c)
}
