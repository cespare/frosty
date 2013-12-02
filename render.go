package main

import (
	"fmt"
	"image"
)

type Rendering struct {
	*Scene
	HPixels int
}

// TODO: Also return a progress chan
func (r *Rendering) Render() *image.RGBA {
	w := r.HPixels
	h := int(float64(w) * r.Camera.Aspect)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	scanner := NewPixelScanner(r.Camera, r.HPixels)
	for scanner.Scan() {
		x, y, v := scanner.Cur()
		fmt.Printf("\033[01;34m>>>> x: %v\x1B[m\n", x)
		fmt.Printf("\033[01;34m>>>> y: %v\x1B[m\n", y)
		fmt.Printf("\033[01;34m>>>> v: %v\x1B[m\n", v)
		fmt.Println("-------------------")
	}

	//for x := 0; x < w; x++ {
	//for y := 0; y < h; y++ {
	//img.SetRGBA(x, y, color.RGBA{0, 0xff, 0, 0xff})
	//}
	//}
	return img
}

// A PixelScanner yields each point in the rendered image along with a vector to represent its position on the
// image plane.
type PixelScanner struct {
	hPixels, vPixels int
	camera           *Camera

	pixelSize float64
	height    float64
	// across and down are unit vectors parallel to the x and y axes, respectively, of the rendered image. They
	// are both perpendicular to the camera ray. Across is parallel to the xz plane.
	across, down      *Vec3
	origin            *Vec3
	curX, curY        int
	started, finished bool
}

func NewPixelScanner(camera *Camera, hPixels int) *PixelScanner {
	scanner := &PixelScanner{
		hPixels: hPixels,
		camera:  camera,
	}
	// Compute vertical pixels
	scanner.vPixels = int(float64(hPixels) * camera.Aspect)
	// Compute the size of a pixel.
	scanner.pixelSize = camera.Width / float64(hPixels)
	// Compute useful vectors.
	// Across is the perpendicular to the camera ray and parallel to the xz plane => it is the cross product of
	// the camera ray and the y axis.
	yAxis := &Vec3{0, 1, 0}
	cam := camera.Loc.Vec()
	a := cam.Cross(yAxis)
	scanner.across = a.Normal(a)
	// Down is perpendicular to the camera ray and Across.
	d := V()
	d = cam.Cross(a)
	scanner.down = d.Normal(d)
	// Now we can easily compute the location of the image origin.
	height := camera.Width * camera.Aspect
	o := camera.Loc.V1.Copy()
	scanner.origin = o.Add(o, V().Mul(d, -0.5*height)).Add(o, V().Mul(a, -0.5*camera.Width))
	scanner.height = height
	return scanner
}

func (s *PixelScanner) Scan() bool {
	if s.finished {
		return false
	}
	if !s.started {
		s.started = true
		return true
	}
	s.curX++
	if s.curX >= s.hPixels {
		s.curX = 0
		s.curY++
		if s.curY >= s.vPixels {
			s.finished = true
			return false
		}
	}
	return true
}

func (s *PixelScanner) Cur() (x, y int, v *Vec3) {
	// The ray goes through the *center* of the pixel.
	xDist := s.camera.Width*(float64(s.curX)/float64(s.hPixels)) + 0.5*s.pixelSize
	yDist := s.height*(float64(s.curY)/float64(s.vPixels)) + 0.5*s.pixelSize
	v = s.origin.Copy()
	v.Add(v, V().Mul(s.across, xDist)).Add(v, V().Mul(s.down, yDist))
	return s.curX, s.curY, v
}
