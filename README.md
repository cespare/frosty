# Frosty

Frosty is an experimental raytracer, written in Go.

It was written for the experience of creating a raytracer, not as a serious project.

Status: It makes a picture. Pre-pre-alpha.

## TODO:

- MVP:
  - ~~Cubes~~
  - ~~Light source~~
  - ~~Shadows~~
  - ~~Diffuse lighting~~
  - ~~Color clamping~~
- ~~Antialiasing~~
- ~~Parallelism~~
- ~~Tone mapping~~ (naive version for now)
- ~~Background (planes)~~
- Specular highlights
- Spheres
- Reflection
- Refraction
- Video

## Ponies

- Depth of field
- Motion blur
  - Stochastic temporal sampling
- Adaptive supersampling for faster antialiasing
  - Stochastic supersampling
- Soft shadows
- More shapes
  - Cones
  - Cylinders
  - Toroids
  - Convex polyhedra
  - CSG
- Import models
  - STL?
- Photon mapping (caustics)
- Spatial subdivision (octrees, BSP, k-d trees)

## Links

* http://www.cs.cmu.edu/afs/cs/academic/class/15462-s09/www/lec/13/lec13.pdf
* http://www.rhythm.com/~ivan/ray.html
* https://github.com/marczych/RayTracer
* http://web.cs.wpi.edu/~emmanuel/courses/cs563/write_ups/zackw/realistic_raytracing.html
* http://www.cs.rit.edu/~jmg/courses/cgII/20033/slides/raytrace-assn-3.pdf
* http://www.cs.rit.edu/~jmg/courses/cgII/20041/slides/raytrace-assn-7.pdf
* http://www.cs.cornell.edu/courses/cs4620/2011fa/lectures/08raytracingWeb.pdf
* http://graphics.cs.cmu.edu/nsp/course/15-462/Spring04/slides/07-lighting.pdf
