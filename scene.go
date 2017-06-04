package main

import (
	"math"
)

type Scene struct {
	Camera  *Camera
	Ambient Color     // Ambient light
	PLights []*PLight // Point lights

	// Materials
	Materials map[string]*Material

	// Some kinds of objects have convenient representations for input.
	RPrisms []*RPrism
	Planes  []*PlaneObject

	// The computed list of objects over which the tracer iterates.
	objects []Object
}

// Don't consider it an intersection if the distance is less than this cutoff.
const minDistance = 0.0001

// An Object is any object in the scene.
type Object interface {
	Initialize(map[string]*Material) error
	// If the ray intersects the object, return the distance to the nearest
	// intersection (from ray.V, the eye point), the Material at that point,
	// the intersection point, the normal vector at that point, and true.
	// Otherwise ok is false.
	Intersect(Ray) (d float64, mat *Material, p, normal Vec3, ok bool)
}

// After loading the scene from file, load all objects into the objects slice.
func (s *Scene) Initialize() error {
	for _, rp := range s.RPrisms {
		s.objects = append(s.objects, rp)
	}
	for _, p := range s.Planes {
		s.objects = append(s.objects, p)
	}
	for _, o := range s.objects {
		if err := o.Initialize(s.Materials); err != nil {
			return err
		}
	}
	return nil
}

// Trace traces a single ray through the scene.
func (s *Scene) Trace(r Ray) (c Color) {
	nearest := math.MaxFloat64
	found := false
	var (
		mat  *Material
		p    Vec3
		norm Vec3
	)
	for _, obj := range s.objects {
		d, m, pt, n, ok := obj.Intersect(r)
		if ok && d < nearest {
			found = true
			nearest = d
			mat = m
			p = pt
			norm = n
		}
	}
	color := Black
	if !found {
		return color
	}
	// ambient term
	la := s.Ambient.Mul(mat.Color)     // La, the ambient light * ambient object color
	color = color.Add(la.MulS(mat.Ka)) // ambient term is ka * La

	// For further calculations it's nice to normalize all vectors.
	norm = norm.Normalize()
	// For each light, compute diffuse and specular components
lights:
	for _, light := range s.PLights {
		// Compute the shadow ray
		shadow := light.Pos
		shadow = shadow.Sub(p)
		if shadow.Dot(norm) < 0 {
			// Light is behind the surface.
			continue
		}
		d := shadow.Mag() // distance from the point to the light
		shadow = shadow.Normalize()
		for _, obj := range s.objects {
			if d2, _, _, _, ok := obj.Intersect(Ray{p, shadow}); ok {
				if d2 < d-minDistance {
					// An object blocks the shadow raw (i.e., this point is in shadow),
					// so skip the specular and diffuse terms for this light.
					continue lights
				}
			}
		}
		// Point lights fall off according to the inverse square law.
		intensity := light.Color.MulS(1.0 / (d * d))
		// For the diffuse term, Li is the diffuse object color * light source.
		li := intensity.Mul(mat.Color)
		diffuse := shadow.Dot(norm)
		li = li.MulS(diffuse)
		color = color.Add(li.MulS(mat.Kd))
	}
	return color
}
