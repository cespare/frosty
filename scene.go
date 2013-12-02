package main

import "math"

type Scene struct {
	Camera *Camera
	Lights []*Light

	// Some kinds of objects have convenient representations for input.
	RPrisms []*RPrism

	// The computed list of objects over which the tracer iterates.
	objects []Object
}

// A Light is a light source. White only for now.
type Light struct {
	Pos *Vec3
}

// An Object is any object in the scene.
type Object interface {
	// If the ray intersects the object, return the distance to the nearest intersection (from ray.V, the eye
	// point), the color at that point, and true. Otherwise ok is false.
	Intersect(Ray) (nearest float64, color Color, ok bool)
}

// After loading the scene from file, load all objects into the objects slice.
func (s *Scene) Initialize() {
	for _, rp := range s.RPrisms {
		s.objects = append(s.objects, rp)
	}
}

// Trace traces a single ray through the scene.
func (s *Scene) Trace(r Ray) Color {
	nearest := math.MaxFloat64
	color := Black // Background
	for _, obj := range s.objects {
		d, c, ok := obj.Intersect(r)
		if ok && d < nearest {
			nearest = d
			color = c
		}
	}
	return color
}
