package main

import (
	"math"
)

// A Camera is defined by a Ray (for the location and orientation of the image
// plane), the width of the image plane, a horizontal angle of view, and an
// aspect ratio. For now, there's no rotation specified so the top and bottom
// edges of the image plane are always parallel to the xz plane.
type Camera struct {
	Loc    Ray     // Location; Loc.V1 is the center of the image plane
	Width  float64 // Width of the image plane
	Haov   Rad     // Horizontal angle of view in degrees
	Aspect float64 // Aspect ratio: height / width
}

// Vantage returns the vantage point for the camera. This is the origin of rays.
// It is situated behind the center of the image plane.
func (c *Camera) Vantage() *Vec3 {
	// First find how far behind the image plane the vp is.
	// tan(half of aspect ratio) = (half of image plane width) / distance
	d := (0.5 * c.Width) / math.Tan(float64(0.5*c.Haov))

	// Construct the vector of magnitude d and the right direction to go
	// from the image plane center to the vantage point.
	v := c.Loc.D.Copy()
	return v.Normalize(v).Mul(v, -d).Add(c.Loc.V, v)
}
