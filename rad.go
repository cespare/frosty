package main

import (
	"encoding/json"
	"math"
)

// A Rad is an angular measurement of radians.
type Rad float64

// UnmarshalJSON unmarshals from JSON text specifying a quantity in degrees.
func (r *Rad) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	*r = Rad(f * math.Pi / 180.0)
	return nil
}
