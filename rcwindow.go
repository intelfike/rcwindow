package rcwindow

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"

	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"

	"golang.org/x/image/draw"
)

var (
	atomic           = true
	win              screen.Window
	buf              screen.Buffer
	fscaleX, fscaleY float64 //fix scale
	relX, relY       float64
	icount           int
)

func NewWindow(scaleX, scaleY float64, bufSize int) *RCConfig {
	if !atomic {
		fmt.Println("NewWindow:already make window")
		os.Exit(1)
	}
	atomic = false
	//set default
	rc := &RCConfig{
		ScaleX:     scaleX,
		ScaleY:     scaleY,
		width:      700,
		height:     700,
		Dots:       make([]*Dot, bufSize),
		DotSize:    2,
		DotColor:   color.RGBA{0xff, 0xff, 0x00, 0xff},
		AxisColor:  color.White,
		FrameColor: color.White,
		Move:       0.1,
		Magni:      1.1,
	}
	rc.sx = float64(rc.width) / (rc.ScaleX * 2.0)
	rc.sy = float64(rc.height) / (rc.ScaleY * 2.0)
	rc.state = "running"
	fscaleX = scaleX
	if fscaleX < 0 {
		fscaleX *= -1.0
	}
	fscaleY = scaleY
	if fscaleY < 0 {
		fscaleY *= -1.0
	}
	go xyWindow(rc)
	rc.wg.Add(1)
	rc.wg.Wait()
	return rc
}

func xyWindow(rc *RCConfig) {
	driver.Main(func(s screen.Screen) {
		var op = screen.NewWindowOptions{
			Width:  rc.width,
			Height: rc.height,
		}
		win, _ = s.NewWindow(&op)
		if win == nil {
			fmt.Println("err")
			return
		}
		defer win.Release()
		defer func() {
			if buf != nil {
				buf.Release()
			}
		}()
		cur := 0
		fmt.Println("draw.")
		rc.wg.Done()
		for {
			e := win.NextEvent()
			//fmt.Printf("%#v\n", e)
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case mouse.Event:
				if e.Direction == mouse.DirPress {
					x, y := rc.parseR(float64(e.X), float64(e.Y))
					fmt.Printf("[x=%0.10f,y=%0.10f]", x, y)
				}

			case key.Event:
				if e.Direction == key.DirNone || e.Direction == key.DirPress {
					rc.mx.Lock()
					switch e.Code {
					case key.CodeEscape:
						rc.mx.Unlock()
						return
					case key.CodeR:

					case key.CodeC:
						relX = 0.0
						relY = 0.0
						rc.ScaleX = fscaleX
						rc.ScaleY = fscaleY
					case key.CodeX:
						rc.ScaleX /= rc.Magni
						rc.ScaleY /= rc.Magni
					case key.CodeZ:
						rc.ScaleX *= rc.Magni
						rc.ScaleY *= rc.Magni
					case key.CodeUpArrow:
						relY -= rc.ScaleY * rc.Move
					case key.CodeDownArrow:
						relY += rc.ScaleY * rc.Move
					case key.CodeLeftArrow:
						relX += rc.ScaleX * rc.Move
					case key.CodeRightArrow:
						relX -= rc.ScaleX * rc.Move
					default:
						goto none
					}
					win.Send(size.Event{WidthPx: rc.width, HeightPx: rc.height})
				none:
					rc.mx.Unlock()
				}
			case paint.Event:
				//win.Send(size.Event{WidthPx:rc.width, HeightPx:rc.height})
				rc.mx.Lock()
				for n := cur; n != icount; n++ {
					m := n % len(rc.Dots)
					if rc.Dots[m] == nil {
						continue
					}
					v := *rc.Dots[m]
					x, y := rc.parse(v.x, v.y)
					rectDraw(x-rc.DotSize/2, y-rc.DotSize/2, x+int(float64(rc.DotSize)/2.0+0.5), y+int(float64(rc.DotSize)/2.0+0.5), v.col)
				}
				cur = icount

				rc.mx.Unlock()
				//更新
				win.Upload(image.Point{0, 0}, buf, buf.Bounds())
				win.Publish()
			case size.Event:
				if buf != nil {
					buf.Release()
				}
				var err error
				buf, err = s.NewBuffer(e.Size())
				if err != nil {
					continue
				}
				rc.width = e.WidthPx
				rc.height = e.HeightPx

				rc.mx.Lock()
				rc.sx = float64(rc.width) / (rc.ScaleX * 2.0)
				rc.sy = float64(rc.height) / (rc.ScaleY * 2.0)
				cenx, ceny := rc.parse(0, 0)

				var dy Base
				dy.right, dy.top = rc.parse(rc.ScaleX-relX, rc.ScaleY-relY)
				dy.left, dy.bottom = rc.parse(-rc.ScaleX-relX, -rc.ScaleY-relY)
				//vertical line
				rectDraw(cenx, dy.top, cenx+1, dy.bottom, rc.AxisColor)
				//horizontal line
				rectDraw(dy.right, ceny, dy.left, ceny+1, rc.AxisColor)
				//frame
				var st Base
				st.right, st.top = rc.parse(fscaleX, fscaleY)
				st.left, st.bottom = rc.parse(-fscaleX, -fscaleY)

				rectDraw(st.left, st.top, st.right, st.top+1, rc.FrameColor)
				rectDraw(st.left, st.bottom, st.right, st.bottom+1, rc.FrameColor)
				rectDraw(st.left, st.top, st.left+1, st.bottom, rc.FrameColor)
				rectDraw(st.right, st.top, st.right+1, st.bottom, rc.FrameColor)
				//point
				for _, p := range rc.Dots {
					if p != nil {
						v := *p
						x, y := rc.parse(v.x, v.y)
						rectDraw(x-rc.DotSize/2, y-rc.DotSize/2, x+int(float64(rc.DotSize)/2.0+0.5), y+int(float64(rc.DotSize)/2.0+0.5), v.col)
					} else {
						break
					}
				}
				rc.mx.Unlock()
				//更新
				win.Upload(image.Point{0, 0}, buf, buf.Bounds())
				win.Publish()
			case error:
				log.Print(e)
			}
		}
	})
	fmt.Println("\ndone.")
	if rc.state == "waiting" {
		rc.wg.Done()
	}
	if rc.state == "end" {
		rc.wg.Done()
	}
	rc.state = ""

}
func rectDraw(x, y, x2, y2 int, c color.Color) {
	r := image.Rect(x, y, x2, y2)
	draw.Draw(buf.RGBA(), r, image.NewUniform(c), image.ZP, draw.Src)
}
