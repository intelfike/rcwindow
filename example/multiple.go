package main
import (
	"github.com/intelfike/rcwindow"
	"math"
	"time"
	"image/color"
)

var c = []color.Color{
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
}
func main() {
	rc := rcwindow.NewWindow(10, 10, 10000)
	rc.SafeConfig(func(){rc.DotSize = 2})
	rc.DrawTick(1 << 24)
	
	count := 0
	rc.FillXc(
		func(x float64)(float64, color.Color){
			count++
			count %= 4
			return math.Pow(x, float64(count)), c[count]
		},
		func(){
			time.Sleep(1)
		},
	)
	rc.Wait()
}