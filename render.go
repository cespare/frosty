package main

import (
	"sync"
)

type Rendering struct {
	*Scene
	HPixels int
}

// TODO: Also return a progress chan
func (r *Rendering) Render(parallelism int) *Image {
	w := r.HPixels
	h := int(float64(w) * r.Camera.Aspect)
	img := NewImage(w, h)
	scanner := NewLineScanner(r.Camera, r.HPixels)

	var wg sync.WaitGroup
	wg.Add(parallelism)
	lines := scanner.Scan()
	for i := 0; i < parallelism; i++ {
		go func() {
			for line := range lines {
				for x := line.xMin; x < line.xMax; x++ {
					img.Set(x, line.y, r.Trace(line.rays[x]))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return img
}

// A LineScanner yields each line of pixels in the rendered image along with the
// vectors to represent their position on the image plane.
type LineScanner struct {
	hPixels, vPixels int
	camera           *Camera

	pixelSize float64
	height    float64
	// across and down are unit vectors parallel to the x and y axes,
	// respectively, of the rendered image. They are both perpendicular to
	// the camera ray. Across is parallel to the xz plane.
	across, down Vec3
	origin       Vec3
	vantage      Vec3
}

func NewLineScanner(camera *Camera, hPixels int) *LineScanner {
	// Compute vertical pixels
	vPixels := int(float64(hPixels) * camera.Aspect)
	// Compute the size of a pixel.
	pixelSize := camera.Width / float64(hPixels)
	// Compute useful vectors.
	// a is the perpendicular to the camera ray and parallel to the xz
	// plane => it is the cross product of the camera ray and the y axis.
	yAxis := Vec3{0, 1, 0}
	cam := camera.Loc.D
	a := cam.Cross(yAxis).Normalize()
	// d is perpendicular to the camera ray and a.
	d := cam.Cross(a).Normalize()
	// Now we can easily compute the location of the image origin.
	height := camera.Width * camera.Aspect
	return &LineScanner{
		hPixels:   hPixels,
		vPixels:   vPixels,
		camera:    camera,
		pixelSize: pixelSize,
		height:    height,
		across:    a,
		down:      d,
		origin:    camera.Loc.V.Add(d.Mul(-0.5 * height)).Add(a.Mul(-0.5 * camera.Width)),
		vantage:   camera.Vantage(),
	}
}

type scanLine struct {
	y          int
	xMin, xMax int
	rays       []Ray
}

func (s *LineScanner) Scan() <-chan scanLine {
	ch := make(chan scanLine)
	go func() {
		for y := 0; y < s.vPixels; y++ {
			line := scanLine{
				y:    y,
				xMin: 0,
				xMax: s.hPixels,
				rays: make([]Ray, s.hPixels),
			}
			for x := range line.rays {
				// The ray goes through the *center* of the pixel.
				xDist := s.camera.Width*(float64(x)/float64(s.hPixels)) + 0.5*s.pixelSize
				yDist := s.height*(float64(y)/float64(s.vPixels)) + 0.5*s.pixelSize
				v := s.origin.Add(s.across.Mul(xDist)).Add(s.down.Mul(yDist))
				line.rays[x] = Ray{V: s.vantage, D: v.Sub(s.vantage)}
			}
			ch <- line
		}
		close(ch)
	}()
	return ch
}
