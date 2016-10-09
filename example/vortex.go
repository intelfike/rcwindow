package main

import (
	"math/cmplx"
	"github.com/intelfike/rcwindow"
	"time"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 2000)
	rc.SafeConfig(func(){rc.DotSize = 3})
	c := cmplx.Pow(1i, 1.0/100.0)
	v := 1i
	rc.DrawTick(1 << 24)
	for n := 0; n <= rc.Len(); n++{
		rc.Dot(real(v), imag(v))
		v *= c / 1.001
		time.Sleep(1)
	}
	rc.Wait()
}