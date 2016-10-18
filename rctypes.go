package rcwindow

import (
	"fmt"
	"image/color"
	"os"
	"sync"
	"time"

	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Base struct {
	right, top, left, bottom int
}

type Dot struct {
	x, y float64
	col  color.Color
}
type RCConfig struct {
	ScaleX, ScaleY float64
	width, height  int
	Dots           []*Dot
	DotSize        int
	index          int

	wg     sync.WaitGroup
	mx     sync.Mutex
	sx, sy float64
	state  string
	change bool
	//Axis == Line
	DotColor, AxisColor, FrameColor color.Color

	Move  float64
	Magni float64
}

func (rc *RCConfig) SafeConfig(f func()) {
	rc.mx.Lock()
	defer rc.mx.Unlock()
	f()
}
func (rc *RCConfig) Dot(x, y float64) {
	rc.Dotc(x, y, rc.DotColor)
}
func (rc *RCConfig) Dotc(x, y float64, col color.Color) {
	if col == nil{
		return
	}
	rc.mx.Lock()
	defer rc.mx.Unlock()
	switch rc.state {
	case "running":
		//fmt.Println(rc.index, x, y)
		rc.Dots[rc.index%len(rc.Dots)] = &Dot{x, y, col}
		rc.index++
		icount++
		rc.index %= len(rc.Dots)
		rc.change = true
	default:
		fmt.Println("Dot:Not running.")
		os.Exit(1)
	}
}
func (rc *RCConfig) Dotcmplx(com complex128) {
	rc.Dot(real(com), imag(com))
}
func (rc *RCConfig) Dotcmplxc(com complex128, col color.Color) {
	rc.Dotc(real(com), imag(com), col)
}
func (rc *RCConfig) FillX(f func(float64) float64, delay func()) {
	xv := fscaleX * 2 / float64(rc.Len())
	for n := 0; ; n++ {
		if rc.state != "running" {
			return
		}
		x := xv * float64(n)
		x -= fscaleX
		if x > fscaleX {
			return
		}
		y := f(x)
		rc.Dot(x, y)
		if fscaleY*-1 < y && y < fscaleY {
			delay()
		}
	}
}
func (rc *RCConfig) FillXc(f func(float64) (float64, color.Color), delay func()) {
	xv := fscaleX * 2 / float64(rc.Len())
	for n := 0; ; n++ {
		if rc.state != "running" {
			return
		}
		x := xv * float64(n)
		x -= fscaleX
		if x > fscaleX {
			return
		}
		y, c := f(x)
		if c == nil{
			continue
		}
		rc.Dotc(x, y, c)
		if fscaleY*-1 < y && y < fscaleY {
			delay()
		}
	}
}
func (rc *RCConfig) Draw() {
	switch rc.state {
	case "running":
		if rc.change {
			win.Send(paint.Event{})
			rc.change = false
		}
	}
}
func (rc *RCConfig) Redraw() {
	switch rc.state {
	case "running":
		if rc.change {
			rc.mx.Lock()
			win.Send(size.Event{WidthPx: rc.width, HeightPx: rc.height})
			rc.mx.Unlock()
			rc.change = false
		}
	}
}
func (rc *RCConfig) DrawTick(tick time.Duration) {
	go func() {
		for {
			if rc.state != "running" {
				break
			}
			rc.Draw()
			time.Sleep(tick)
		}
	}()
}
func (rc *RCConfig) RedrawTick(tick time.Duration) {
	go func() {
		for {
			if rc.state != "running" {
				break
			}
			rc.Redraw()
			time.Sleep(tick)
		}
	}()
}
func (rc *RCConfig) Clear() {
	rc.mx.Lock()
	defer rc.mx.Unlock()
	switch rc.state {
	case "running", "waiting":
		for n, _ := range rc.Dots {
			rc.Dots[n] = nil
		}
		rc.index = 0
		rc.Draw()
	}
}
func (rc *RCConfig) Wait() {
	switch rc.state {
	case "running":
		rc.Draw()
		fmt.Println("wait")
		rc.state = "waiting"
		rc.wg.Add(1)
		rc.wg.Wait()
		rc.state = "running"
	case "waiting":
		fmt.Println("Wait:Already waiting.")
		os.Exit(1)
	default:
		fmt.Println("Wait:Not running.")
		os.Exit(1)
	}
}
func (rc *RCConfig) End() {
	switch rc.state {
	case "running", "waiting":
		rc.state = "end"
		win.Send(lifecycle.Event{To: lifecycle.StageDead})
	default:
		fmt.Println("End:Not running")
		os.Exit(1)
	}
	rc.wg.Add(1)
	rc.wg.Wait()
}
func (rc *RCConfig) Len() int {
	return len(rc.Dots)
}

//Axis座標→window座標
func (rc *RCConfig) parse(fx, fy float64) (int, int) {
	fx += relX
	fy += relY
	x := rc.width/2 + int(fx*rc.sx)
	y := rc.height/2 + int(fy*rc.sy)
	y = rc.height - y
	return x - 1, y - 1
}

//window座標→Axis座標
func (rc *RCConfig) parseR(x, y float64) (float64, float64) {
	x++
	y++
	fx := (x - float64(rc.width/2)) / rc.sx
	fy := (y - float64(rc.height/2)) / rc.sy * -1
	return fx - relX, fy - relY
}
