package main

import (
	"encoding/json"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v *Vec3) UnmarshalJSON(b []byte) error {
	a := [3]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y, v.Z = a[0], a[1], a[2]
	return nil
}

type Ray struct {
	Pos, Dir Vec3
}
