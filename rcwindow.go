package rcwindow
import (
	"fmt"
	"log"
	"image"
	"image/color"
	"sync"
	"os"
	"time"
	
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	
	"golang.org/x/image/draw"
)

type Dot struct{
	x, y float64
}
type rcConfig struct{
	ScaleX, ScaleY float64
	Width, Height int
	Dots []Dot
	DotSize int
	index int
	win screen.Window
	wg sync.WaitGroup
	mx sync.Mutex
	sx, sy float64
	state string
	change bool
	//Axis == Line
	DotColor ,AxisColor color.Color
}
func (rc *rcConfig) Start(){
	switch rc.state{
	case "":
		go xyWindow(rc)
		rc.wg.Add(1)
		rc.wg.Wait()
		rc.state = "running"
	default:
		fmt.Println("Start:Already started.")
		os.Exit(1)
	}
}
func (rc *rcConfig) Dot(x, y float64){
	switch rc.state{
	case "running":
		rc.Dots[rc.index % len(rc.Dots)] = Dot{x, y}
		rc.index++
		rc.index %= len(rc.Dots)
		rc.change = true
	default:
		fmt.Println("Dot:Not running.")
		os.Exit(1)
	}
}
func (rc *rcConfig) Draw(){
	switch rc.state{
	case "running":
		if rc.change{
			rc.win.Send(paint.Event{})
			rc.change = false
		}
	default:
		fmt.Println("Dot:Not running.")
		os.Exit(1)
	}
}
func (rc *rcConfig) DrawTick(tick time.Duration){
	go func(){
		for{
			rc.mx.Lock()
			if rc.state != "running"{
				rc.mx.Unlock()
				return
			}
			rc.Draw()
			rc.mx.Unlock()
			time.Sleep(tick)
		}
	}()
}
func (rc *rcConfig) Wait(){
	switch rc.state{
	case "running":
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
func (rc *rcConfig) End(){
	switch rc.state{
	case "running", "waiting":
		rc.state = "end"
		rc.win.Send(lifecycle.Event{To:lifecycle.StageDead})
	default:
		fmt.Println("End:Not running")
		os.Exit(1)
	}
	rc.wg.Add(1)
	rc.wg.Wait()
}
//Axis座標→window座標
func (rc *rcConfig)parse(fx, fy float64)(int, int){
	x := rc.Width / 2 + int(fx * rc.sx)
	y := rc.Height / 2 + int(fy * rc.sy)
	y = rc.Height - y
	return x, y
}
//window座標→Axis座標
func (rc *rcConfig)parseR(x, y float64)(float64, float64){
	fx := (x - float64(rc.Width/2)) / rc.sx
	fy := (y - float64(rc.Height/2)) / rc.sy * -1
	return fx, fy
}

func NewWindow(scaleX, scaleY float64, bufSize int) *rcConfig{
	//set default
	return &rcConfig{
		ScaleX: scaleX,
		ScaleY: scaleY,
		Width:700,
		Height:700,
		Dots: make([]Dot, bufSize),
		DotSize: 2,
		DotColor:	color.RGBA{0xff, 0xff, 0x00, 0xff},
		AxisColor:	color.White,
	}
	
}

func xyWindow(rc *rcConfig){
	driver.Main(func(s screen.Screen) {
		var op = screen.NewWindowOptions{
			Width : rc.Width,
			Height : rc.Height,
		}
		rc.win, _ = s.NewWindow(&op)
		if rc.win == nil {
			fmt.Println("err")
			return
		}
		defer rc.win.Release()
		var buf screen.Buffer
		defer func() {
			if buf != nil {
				buf.Release()
			}
		}()
		fmt.Println("draw.")
		rc.wg.Done()
		for {
			e := rc.win.NextEvent()
			//fmt.Printf("%#v\n", e)
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case mouse.Event:
				if e.Direction == mouse.DirPress{
					x, y := rc.parseR(float64(e.X), float64(e.Y))
					fmt.Printf("[x=%0.3f,y=%0.3f]", x, y)
				}
				
			case key.Event:
				if e.Direction == key.DirNone || e.Direction == key.DirPress{
					switch e.Code{
					case key.CodeEscape:
						return
					case key.CodeR:
	 					rc.win.Send(size.Event{WidthPx:rc.Width, HeightPx:rc.Height})
					case key.CodeUpArrow:
						rc.ScaleX /= 1.1
						rc.ScaleY /= 1.1
						rc.win.Send(paint.Event{})
					case key.CodeDownArrow:
						rc.ScaleX *= 1.1
						rc.ScaleY *= 1.1
						rc.win.Send(paint.Event{})
					}
				}
			case paint.Event:
 				rc.win.Send(size.Event{WidthPx:rc.Width, HeightPx:rc.Height})
			case size.Event:
				buf, _ = s.NewBuffer(e.Size())
				rc.Width = e.WidthPx
				rc.Height = e.HeightPx
				rc.sx = float64(rc.Width) / (rc.ScaleX * 2.0)
				rc.sy = float64(rc.Height) / (rc.ScaleY * 2.0)
				//vertical line
				r := image.Rect(rc.Width/2, 0, rc.Width/2 + 1, rc.Height)
				draw.Draw(buf.RGBA(), r, image.NewUniform(rc.AxisColor), image.ZP, draw.Src)
				//horizontal line
				r = image.Rect(0, rc.Height/2, rc.Width, rc.Height/2 + 1)
				draw.Draw(buf.RGBA(), r, image.NewUniform(rc.AxisColor), image.ZP, draw.Src)
				//point
				rc.mx.Lock()
				for _, v := range rc.Dots{
					if v.x != 0 && v.y != 0{
						x, y := rc.parse(v.x, v.y)
						r := image.Rect(x - rc.DotSize/2, y - rc.DotSize/2, x + int(float64(rc.DotSize)/2.0 + 0.5), y + int(float64(rc.DotSize)/2.0 + 0.5))
						draw.Draw(buf.RGBA(), r, image.NewUniform(rc.DotColor), image.ZP, draw.Src)
					}
				}
				rc.mx.Unlock()
				//更新
				rc.win.Upload(image.Point{0, 0}, buf, buf.Bounds())
				buf.Release()
				rc.win.Publish()
			case error:
				log.Print(e)
			}
		}
	})
	fmt.Println("\ndone.")
	if rc.state == "waiting"{
		rc.wg.Done()
	}
	if rc.state == "end"{
		rc.wg.Done()
	}
	rc.state = ""
	
}

