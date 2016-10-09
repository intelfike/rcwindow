package main

import (
	"github.com/intelfike/rcwindow"
)
func main() {
	rc := rcwindow.NewWindow(1, 1, 100000)
	rc.SafeConfig(func(){rc.DotSize = 2})
	rc.DrawTick(1 << 24)
	rc.FillX(
		func(x float64)float64{
			return x * 2
		},
		func(){
			for n := 0; n < 100000; n++{}
		},
	)
	rc.Wait()
}