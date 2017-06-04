#!/bin/sh

set -eu -o pipefail

go build -o frosty
time ./frosty \
  -debug \
  -cpuprofile \
  -hpixels 1200 \
  -out out.png
