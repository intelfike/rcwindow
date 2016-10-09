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
			switch count{
			case 0:
				return math.Tan(x), c[count]
			case 1:
				return x + 1, c[count]
			case 2:
				return x + 3, c[count]
			case 3:
				return math.Cos(math.Tan(x)), c[count]
			}
			return 0, c[0]
		},
		func(){
			time.Sleep(1)
		},
	)
	rc.Wait()
}