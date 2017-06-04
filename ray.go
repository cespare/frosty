package main

// A Ray is defined by a starting point, V, and an offset vector, D.
type Ray struct {
	V Vec3
	D Vec3
}

// At returns the point p at distance d along r from r.V.
func (r Ray) At(d float64) Vec3 {
	return r.D.Mul(d).Add(r.V)
}
