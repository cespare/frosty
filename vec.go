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

func (v *Vec3) Copy() *Vec3 { return &Vec3{v.X, v.Y, v.Z} }

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

// Normalize sets v to be the normalized (unit) vector of u and returns v.
func (v *Vec3) Normalize(u *Vec3) *Vec3 {
	return v.Div(u, u.Mag())
}

// Cross returns the cross product of u and v as a newly allocated vector.
// (This function does not follow the math/big pattern because it wouldn't work if the result vector were also
// one of the operands.)
func (u *Vec3) Cross(v *Vec3) *Vec3 {
	return &Vec3{
		X: u.Y*v.Z - u.Z*v.Y,
		Y: u.Z*v.X - u.X*v.Z,
		Z: u.X*v.Y - u.Y*v.X,
	}
}

// Dot returns the dot product of u and v.
func (u *Vec3) Dot(v *Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func (v *Vec3) UnmarshalJSON(b []byte) error {
	a := [3]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y, v.Z = a[0], a[1], a[2]
	return nil
}
