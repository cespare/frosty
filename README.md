# Frosty

Frosty is an experimental raytracer, written in Go.

It was written for the experience of creating a raytracer, not as a serious project.

## Installation

    go install github.com/cespare/frosty@latest

## Usage

`frosty -h` tells you the options. Use the json scene file in the repo to get started.

## Status

Alpha. It makes pictures.

![status](/progress/5.png)

## Features

- Cubes
- Point lights
- Shadows
- Diffuse lighting
- Color clamping (naive tone mapping)
- Antialiasing (supersampling)
- Parallel ray computaiton
- Background (planes)

See open issues for other things I've thought about implementing.

## License

MIT licensed; see LICENSE.txt.
