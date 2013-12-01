package main

import (
	"encoding/json"
	"math"
)

// Radians
type Rad float64

// Unmarshal from degrees
func (r *Rad) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	*r = Rad(f * math.Pi / 180.0)
	return nil
}
