package main

import (
	"encoding/json"
	"math"
)

// Vector ops are modeled after math/big.

type Vec3 struct {
	X, Y, Z float64
}

// V returns a newly initialized zero vector.
func V() *Vec3 { return &Vec3{} }

func (w *Vec3) Add(u, v *Vec3) *Vec3 {
	w.X = u.X + v.X
	w.Y = u.Y + v.Y
	w.Z = u.Z + v.Z
	return w
}

func (w *Vec3) Sub(u, v *Vec3) *Vec3 {
	w.X = u.X - v.X
	w.Y = u.Y - v.Y
	w.Z = u.Z - v.Z
	return w
}

// Div sets v to be u / x for some scalar x and returns v.
func (v *Vec3) Div(u *Vec3, x float64) *Vec3 {
	v.X = u.X / x
	v.Y = u.Y / x
	v.Z = u.Z / x
	return v
}

// Mul sets v to be u * x for some scalar x and returns v.
func (v *Vec3) Mul(u *Vec3, x float64) *Vec3 {
	v.X = u.X * x
	v.Y = u.Y * x
	v.Z = u.Z * x
	return v
}

// Mag returns the magnitude of v.
func (v *Vec3) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normal sets v to be the normalized (unit) vector of u and returns v.
func (v *Vec3) Normal(u *Vec3) *Vec3 {
	return v.Div(u, u.Mag())
}

func (v *Vec3) UnmarshalJSON(b []byte) error {
	a := [3]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y, v.Z = a[0], a[1], a[2]
	return nil
}
