package main

type Plane struct {
	q      *Vec3 // A point on the plane
	normal *Vec3 // A vector normal to the plane
	color  Color
}

func PlaneFromVertices(v1, v2, v3 *Vec3, color Color) *Plane {
	// Construct two vectors, v2v1 and v3v1; their cross product is normal to the plane.
	l1 := V().Sub(v2, v1)
	l2 := V().Sub(v3, v1)
	return &Plane{
		q:      v1,
		normal: l1.Cross(l2),
		color:  color,
	}
}

// http://www.cl.cam.ac.uk/teaching/1999/AGraphHCI/SMAG/node2.html
// A point p is on the plane if normal·(p - q) = 0.
// Points on the ray are of the form P(t) = ray.V + t*ray.D for t >= 0.
// Thus intersections have the solution
//           normal · (q - ray.V)
//       t = ---------------------
//             ray.V · ray.D
// If t < 0, then the intersection is behind the vantage point and doesn't count.
func (p *Plane) Intersect(r Ray) (float64, Color, bool) {
	denom := r.V.Dot(r.D)
	if denom == 0 {
		// Ray is parallel to the plane
		return 0, p.color, false
	}
	num := p.normal.Dot(V().Sub(p.q, r.V))
	t := num / denom
	if t < 0 {
		// Intersection is behind the vantage point
		return 0, p.color, false
	}
	return t, p.color, true
}
