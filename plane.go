package main

import (
	"fmt"
)

type Plane struct {
	q      *Vec3 // A point on the plane
	normal *Vec3 // A vector normal to the plane
}

type PlaneObject struct {
	*Plane `json:"-"`
	Mat    *Material `json:"-"`

	V1, V2, V3 *Vec3
	MatName    string `json:"mat"`
}

func (p *PlaneObject) Initialize(materials map[string]*Material) error {
	m, ok := materials[p.MatName]
	if !ok {
		return fmt.Errorf("Cannot find material %s", p.MatName)
	}
	p.Mat = m
	// Construct two vectors, v2v1 and v3v1; their cross product is normal to the plane.
	l1 := V().Sub(p.V2, p.V1)
	l2 := V().Sub(p.V3, p.V1)
	p.Plane = &Plane{
		q:      p.V1,
		normal: l1.Cross(l2),
	}
	//spew.Dump(p)
	return nil
}

// Intersect determines the intersection of r with p.
//
// A point p is on the plane if normal·(p - q) = 0.
// Points on the ray are of the form P(t) = ray.V + t*ray.D for t >= 0.
// Thus intersections have the solution
//           normal · (q - ray.V)
//       t = ---------------------
//             ray.D · normal
// If t < 0, then the intersection is behind the vantage point and doesn't count.
func (p *PlaneObject) Intersect(r Ray) (float64, *Material, *Vec3, *Vec3, bool) {
	denom := r.D.Dot(p.normal)
	if denom == 0 {
		// Ray is parallel to the plane
		return 0, nil, nil, nil, false
	}
	num := p.normal.Dot(V().Sub(p.q, r.V))
	t := num / denom
	if t < minDistance {
		// Intersection is behind the vantage point
		return 0, nil, nil, nil, false
	}
	normal := p.normal
	if denom > 0 {
		// If the ray is less than 90 degrees from the normal line
		// (dot product > 0), we're hitting the 'back' of the plane
		// and we want to return the opposite normal.
		normal = p.normal.Copy()
		normal.Mul(normal, -1)
	}
	return t, p.Mat, r.At(t), normal, true
}
