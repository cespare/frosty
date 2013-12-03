package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/cespare/gomaxprocs"
	"github.com/davecgh/go-spew/spew"
)

func fatal(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func jsonError(raw []byte, err error) error {
	e, ok := err.(*json.SyntaxError)
	if !ok {
		return err
	}
	first := raw[:e.Offset]
	newlines := bytes.Count(first, []byte{'\n'})
	last := bytes.LastIndex(first, []byte{'\n'})
	return fmt.Errorf("JSON error at line %d, column %d: %s", newlines+1, len(first)-last-1, e)
}

func main() {
	gomaxprocs.SetToNumCPU()
	maxprocs := gomaxprocs.Get()
	var (
		sceneFile     = flag.String("scene", "scene.json", "The scene description json file")
		hpixels       = flag.Int("hpixels", 800, "Horizontal pixel size of the output image")
		out           = flag.String("out", "render.png", "Output png image")
		debug         = flag.Bool("debug", false, "Print verbose debugging information")
		supersampling = flag.Int("supersampling", 1, "Supersampling (antialiasing) factor")
		// Default value of 4 * num cpu is based on some ad hoc testing.
		parallelism = flag.Int("parallelism", 4*maxprocs, "Number of rays to compute in parallel")
	)
	flag.Parse()

	if *supersampling < 1 || *supersampling > 8 {
		fatal("Supersampling should be between 1 and 8; got %d\n", *supersampling)
	}
	*hpixels *= *supersampling

	if *parallelism < 1 {
		fatal("Bad value for parallelism (should be at least one): %d\n", *parallelism)
	}

	raw, err := ioutil.ReadFile(*sceneFile)
	if err != nil {
		fatal("Cannot open scene file %s: %s\n", *sceneFile, err)
	}

	fmt.Printf("Loading scene...")
	scene := &Scene{}
	if err := json.Unmarshal(raw, scene); err != nil {
		fatal("\nError loading scene: %s\n", jsonError(raw, err))
	}
	fmt.Println("done")

	fmt.Printf("Initializing primitives...")
	if err := scene.Initialize(); err != nil {
		fmt.Printf("\nError: %s\n", err)
	}
	fmt.Println("done")

	if *debug {
		spew.Dump(scene)
	}

	fmt.Printf("Rendering...")
	rendering := &Rendering{scene, *hpixels}
	img := rendering.Render(*parallelism)
	fmt.Println("done.")

	if *supersampling > 1 {
		fmt.Printf("Downsampling supersampled image...")
		img, err = Downsample(img, *supersampling)
		if err != nil {
			fatal("\nerror downsampling: %s\n", err)
		}
		fmt.Println("done")
	}

	fmt.Printf("Tone mapping image...")
	outImg := img.ToneMap()
	fmt.Println("done.")

	f, err := os.Create(*out)
	if err != nil {
		fatal("Cannot open output file %s: %s\n", *out, err)
	}
	if err := png.Encode(f, outImg); err != nil {
		fatal("Cannot write rendering as png: %s\n", err)
	}
	fmt.Printf("Image rendered to %s\n", *out)
}
