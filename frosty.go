package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
)

func jsonError(raw []byte, err error) error {
	e, ok := err.(*json.SyntaxError)
	if !ok {
		return err
	}
	first := raw[:e.Offset]
	newlines := bytes.Count(first, []byte{'\n'})
	last := bytes.LastIndex(first, []byte{'\n'})
	return fmt.Errorf("JSON error at line %d, column %d: %s",
		newlines+1, len(first)-last-1, e)
}

var jsonCommentRegex = regexp.MustCompile(`(?m)^\s+(//|#).*$`)

// filterJSONComments is a quick and dirty way to filter out comments without
// disturbing offset/newline counts. Only whole-line comments are allowed.
func filterJSONComments(raw []byte) {
	for _, comment := range jsonCommentRegex.FindAllIndex(raw, -1) {
		for i := comment[0]; i < comment[1]; i++ {
			raw[i] = ' '
		}
	}
}

func main() {
	log.SetFlags(0)
	var (
		sceneFile     = flag.String("scene", "scene.json", "The scene description json file")
		hpixels       = flag.Int("hpixels", 800, "Horizontal pixel size of the output image")
		out           = flag.String("out", "render.png", "Output png image")
		debug         = flag.Bool("debug", false, "Print verbose debugging information")
		supersampling = flag.Int("supersampling", 1, "Supersampling (antialiasing) factor")
		// Default value of 4 * numcpu is based on some ad hoc testing.
		parallelism = flag.Int("parallelism", 4*runtime.NumCPU(), "Number of rays to compute in parallel")
		cpuProfile  = flag.Bool("cpuprofile", false, "Emit CPU profile")
	)
	flag.Parse()
	_ = *debug

	if *supersampling < 1 || *supersampling > 8 {
		log.Fatalf("Supersampling should be between 1 and 8; got %d", *supersampling)
	}
	*hpixels *= *supersampling

	if *parallelism < 1 {
		log.Fatalf("Bad value for parallelism (should be at least one): %d", *parallelism)
	}

	if *cpuProfile {
		const name = "cpu.pprof"
		f, err := os.Create(name)
		if err != nil {
			log.Fatalln("Error creating CPU profile file:", err)
		}
		pprof.StartCPUProfile(f)
		defer func() {
			pprof.StopCPUProfile()
			if err := f.Close(); err != nil {
				log.Fatalln("Error writing CPU profile", err)
			}
			log.Println("Stored CPU profile to", name)
		}()
	}

	raw, err := ioutil.ReadFile(*sceneFile)
	if err != nil {
		log.Fatalf("Cannot open scene file %s: %s", *sceneFile, err)
	}
	filterJSONComments(raw)

	fmt.Printf("Loading scene...")
	scene := &Scene{}
	if err := json.Unmarshal(raw, scene); err != nil {
		log.Fatalf("\nError loading scene: %s", jsonError(raw, err))
	}
	fmt.Println("done")

	fmt.Printf("Initializing primitives...")
	if err := scene.Initialize(); err != nil {
		fmt.Printf("\nError: %s\n", err)
	}
	fmt.Println("done")

	fmt.Printf("Rendering...")
	rendering := &Rendering{scene, *hpixels}
	img := rendering.Render(*parallelism)
	fmt.Println("done.")

	if *supersampling > 1 {
		fmt.Printf("Downsampling supersampled image...")
		img, err = Downsample(img, *supersampling)
		if err != nil {
			log.Fatalf("\nerror downsampling: %s", err)
		}
		fmt.Println("done")
	}

	fmt.Printf("Tone mapping image...")
	outImg := img.ToneMap()
	fmt.Println("done.")

	f, err := os.Create(*out)
	if err != nil {
		log.Fatalf("Cannot open output file %s: %s", *out, err)
	}
	if err := png.Encode(f, outImg); err != nil {
		log.Fatalf("Cannot write rendering as png: %s", err)
	}
	fmt.Printf("Image rendered to %s\n", *out)
}
