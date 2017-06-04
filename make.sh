#!/bin/sh

set -eu -o pipefail

go build -o frosty
time ./frosty \
  -debug \
  -hpixels 1200 \
  -out out.png
