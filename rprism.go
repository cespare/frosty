package main

import (
	"fmt"
	"math"
)

// A RPrism is a rectangular prism with sides parallel to the axis planes. It is defined by a point and
// extends in the positive X, Y, and Z dimensions as given by Dim.
type RPrism struct {
	Pos     *Vec3      // The corner with smallest X, Y, Z
	Dim     [3]float64 // X, Y, Z
	Mat     *Material  `json:"-"`
	MatName string     `json:"mat"`
}

func (p *RPrism) Initialize(materials map[string]*Material) error {
	m, ok := materials[p.MatName]
	if !ok {
		return fmt.Errorf("Cannot find material %s", p.MatName)
	}
	p.Mat = m
	return nil
}

// a, b, and c are dimensions (i.e. each is one of X, Y, Z)
// For example, for the unit cube, to find an intersection with the side on the YZ plane (x=0), you might call
// with arguments:
//   bcPlane = 0
//   minB    = 0
//   maxB    = 1
//   minC    = 0
//   maxC    = 1
// v and d are r.V and r.D vectors in the order [a, b, c].
func rprismIntersects(q *rprismIntersectQ) (float64, bool) {
	t := (q.bcPlane - q.v[0]) / q.d[0]
	if t < 0 {
		return 0, false
	}
	b := q.v[1] + t*q.d[1]
	c := q.v[2] + t*q.d[2]
	if b >= q.minB && b <= q.maxB && c >= q.minC && c <= q.maxC {
		return t, true
	}
	return 0, false
}

type rprismIntersectQ struct {
	v, d                            [3]float64
	bcPlane, minB, maxB, minC, maxC float64
	normal                          *Vec3
}

// P(t) = r.V.X + t*r.D.X = x1
// t = (x1 - r.V.X) / r.D.X
func (p *RPrism) Intersect(r Ray) (float64, *Material, *Vec3, *Vec3, bool) {
	queries := []*rprismIntersectQ{
		{
			[3]float64{r.V.X, r.V.Y, r.V.Z},
			[3]float64{r.D.X, r.D.Y, r.D.Z},
			p.Pos.X, p.Pos.Y, p.Pos.Y + p.Dim[1], p.Pos.Z, p.Pos.Z + p.Dim[2],
			&Vec3{-1, 0, 0},
		},
		{
			[3]float64{r.V.X, r.V.Y, r.V.Z},
			[3]float64{r.D.X, r.D.Y, r.D.Z},
			p.Pos.X + p.Dim[0], p.Pos.Y, p.Pos.Y + p.Dim[1], p.Pos.Z, p.Pos.Z + p.Dim[2],
			&Vec3{1, 0, 0},
		},
		{
			[3]float64{r.V.Y, r.V.Z, r.V.X},
			[3]float64{r.D.Y, r.D.Z, r.D.X},
			p.Pos.Y, p.Pos.Z, p.Pos.Z + p.Dim[2], p.Pos.X, p.Pos.X + p.Dim[0],
			&Vec3{0, -1, 0},
		},
		{
			[3]float64{r.V.Y, r.V.Z, r.V.X},
			[3]float64{r.D.Y, r.D.Z, r.D.X},
			p.Pos.Y + p.Dim[1], p.Pos.Z, p.Pos.Z + p.Dim[2], p.Pos.X, p.Pos.X + p.Dim[0],
			&Vec3{0, 1, 0},
		},
		{
			[3]float64{r.V.Z, r.V.X, r.V.Y},
			[3]float64{r.D.Z, r.D.X, r.D.Y},
			p.Pos.Z, p.Pos.X, p.Pos.X + p.Dim[0], p.Pos.Y, p.Pos.Y + p.Dim[1],
			&Vec3{0, 0, -1},
		},
		{
			[3]float64{r.V.Z, r.V.X, r.V.Y},
			[3]float64{r.D.Z, r.D.X, r.D.Y},
			p.Pos.Z + p.Dim[2], p.Pos.X, p.Pos.X + p.Dim[0], p.Pos.Y, p.Pos.Y + p.Dim[1],
			&Vec3{0, 0, 1},
		},
	}
	nearest := math.MaxFloat64
	found := false
	var normal *Vec3
	for _, q := range queries {
		d, ok := rprismIntersects(q)
		if ok {
			if d > minDistance && d < nearest {
				found = true
				nearest = d
				normal = q.normal
			}
		}
	}

	if !found {
		return 0, nil, nil, nil, false
	}
	pt := r.At(nearest)
	return nearest, p.Mat, pt, normal, found
}
