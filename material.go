package main

type Material struct {
	// Object Color
	Color    Color // Ambient/Diffuse color
	Specular Color

	// Phong parameters
	Ka    float64 // Ambient light
	Kd    float64 // Diffuse/Lambertian reflection
	Ks    float64 // Specular reflection
	Alpha float64 // Exponent for specular highlight (shininess constant)
}
