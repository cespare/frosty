package main

// A Camera is defined by a Ray (for the location and orientation of the image plane), a horizontal angle of
// view, and an aspect ratio. For now, there's no rotation specified so the top and bottom edges of the image
// plane are always parallel to the xz plane.
type Camera struct {
	Loc    Ray     // Location; Loc.Pos is the center of the image plane
	Haov   float64 // Horizontal angle of view in degrees
	Aspect float64 // Aspect ratio: height / width
}

// A Light is a light source. White only for now.
type Light struct {
	Pos Vec3
}

// A RPrism is a rectangular prism. It's defined by two vectors which are opposite corners.
type RPrism struct {
	Corner1, Corner2 Vec3
	Color            Color
}
