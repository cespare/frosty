package main

// A Ray is defined by a starting point, V, and an offset vector, D.
type Ray struct {
	V, D *Vec3
}

// At returns the point p at distance d along r from r.V.
func (r Ray) At(d float64) *Vec3 {
	p := V().Mul(r.D, d)
	return p.Add(p, r.V)
}
