package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func fatal(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func main() {
	var (
		sceneFile = flag.String("scene", "scene.json", "The scene description json file")
		hpixels   = flag.Int("hpixels", 800, "Horizontal pixel size of the output image")
		out       = flag.String("out", "render.png", "Output png image")
		debug     = flag.Bool("debug", false, "Print verbose debugging information")
	)
	flag.Parse()

	f, err := os.Open(*sceneFile)
	if err != nil {
		fatal("Cannot open scene file %s: %s\n", *sceneFile, err)
	}

	fmt.Println("Loading scene...")
	scene := &Scene{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(scene); err != nil {
		fatal("Error loading scene: %s\n", err)
	}
	fmt.Println("done")

	if *debug {
		spew.Dump(scene)
	}

	fmt.Println("Rendering...")
	rendering := &Rendering{scene, *hpixels}
	img := rendering.Render()
	f, err = os.Create(*out)
	if err != nil {
		fatal("Cannot open output file %s: %s\n", *out, err)
	}
	if err := png.Encode(f, img); err != nil {
		fatal("Cannot write rendering as png: %s\n", err)
	}
	fmt.Printf("Done. Image rendered to %s\n", *out)
}