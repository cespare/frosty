#!/bin/sh

go build -o frosty && ./frosty -debug -hpixels 1200 -out out.png && open out.png
