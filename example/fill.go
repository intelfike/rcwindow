package main

import (
	"github.com/intelfike/rcwindow"
)

func main() {
	scale := 1.0
	rc := rcwindow.NewWindow(scale, scale, 10000)
	rc.SafeConfig(func(){rc.DotSize = 3})
	pitch := 50
	for n := 0; n < pitch; n++{
		x := 1.0 / float64(pitch) * float64(n) * 2 - scale
		for m := 0; m < pitch; m++{
			y := 1.0 / float64(pitch) * float64(m) * 2 - scale
			rc.Dot(x, y)
		}
	}
	
	rc.Wait()
}