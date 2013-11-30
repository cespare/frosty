package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var (
	scene = flag.String("scene", "scene.json", "The scene description json file")
	hpixels = flag.Int("hpixels", 800, "Horizontal pixel size of the output image")
)

func fatal(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

func main() {
	f, err := os.Open(*scene)
	if err != nil {
		fatal("Cannot open scene file %s: %s\n", *scene, err)
	}

	fmt.Println("Loading scene...")
	scene := &Scene{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(scene); err != nil {
		fatal("Error loading scene: %s\n", err)
	}
	fmt.Println("done")

	spew.Dump(scene)
}
