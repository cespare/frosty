package main

type Ray struct {
	V1, V2 *Vec3
}

func (r *Ray) Vec() *Vec3 {
	return V().Sub(r.V2, r.V1)
}
