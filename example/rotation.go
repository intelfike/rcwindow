package main

import (
	"github.com/intelfike/rcwindow"
	"time"
	"math/cmplx"
)
func main() {
	rc := rcwindow.NewWindow(1, 1, 25)
	rc.SafeConfig(func(){rc.DotSize = 3})
	c := cmplx.Pow(1i, 1.0/100.0)
	v := 1i
	rc.RedrawTick(1 << 24)
	for n := 0; ; n++{
		rc.Dot(real(v), imag(v))
		v *= c
		time.Sleep(1)
	}
	rc.Wait()
}