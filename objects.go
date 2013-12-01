package main

import (
	"math"
)

// A Camera is defined by a Ray (for the location and orientation of the image plane), the width of the image
// plane, a horizontal angle of view, and an aspect ratio. For now, there's no rotation specified so the top
// and bottom edges of the image plane are always parallel to the xz plane.
type Camera struct {
	Loc    *Ray    // Location; Loc.V1 is the center of the image plane
	Width  float64 // Width of the image plane
	Haov   Rad     // Horizontal angle of view in degrees
	Aspect float64 // Aspect ratio: height / width
}

// A Light is a light source. White only for now.
type Light struct {
	Pos *Vec3
}

// A RPrism is a rectangular prism. It's defined by two vectors which are opposite corners.
type RPrism struct {
	Corner1, Corner2 *Vec3
	Color            Color
}

// Vantage returns the vantage point for the camera. This is the origin of all the rays. It is situated behind
// the center of the image plane at an appropriate distance
func (c *Camera) Vantage() *Vec3 {
	// First find how far behind the image plane the vp is.
	// tan(half of aspect ratio) = (half of image plane width) / distance
	d := (0.5 * c.Width) / math.Tan(float64(0.5*c.Haov))

	// Construct the vector of magnitude d and the right direction to go from the image plane center to the
	// vantage point.
	v := c.Loc.Vec()
	v.Normal(v)
	v.Mul(v, -d)
	return v.Add(c.Loc.V1, v)
}
