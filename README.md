<h1>Rectangle Coordinates Window</h1>

日本語:<br>
https://ja.wikipedia.org/wiki/%E7%9B%B4%E4%BA%A4%E5%BA%A7%E6%A8%99%E7%B3%BB<br>
これをGUI(shiny)で表示してくれます。<br>
<br>
English:<br>
https://en.wikipedia.org/wiki/Cartesian_coordinate_system<br>
this package display this with "shiny"<br>

<h2>get command/入手コマンド</h2>

<pre>
go get github.com/intelfike/rcwindow
</pre>

<h2>usage/使い方</h2>

<pre>
rc := rcwindow.NewWindow(1, 1, 10)	//Create window
rc.Dot(0.5, 0.5)	//Draw dot
rc.Wait()	//Wait or End
</pre>

<h2>method/メソッド</h2>
<pre>
type RCConfig struct {
	ScaleX, ScaleY float64	//set by the argument of NewWindow() function
	Dots    []*Dot 		//set by the argument of Dotc() and Dot() function
	DotSize int 	//Default Size
	DotColor, AxisColor, FrameColor color.Color 	//Default Color
	Move  float64
	Magni float64
}
</pre>

<b>func NewWindow(scaleX, scaleY float64, bufSize int) *RCConfig</b><br>
Create Window<br>
<br>
<b>func (rc *RCConfig) Clear()</b><br>
Clear Dots<br>
<br>
<b>func (rc *RCConfig) Dot(x, y float64)</b><br>
Add Dot (default color and don't display)<br>
<br>
<b>func (rc *RCConfig) Dotc(x, y float64, col color.Color)</b><br>
Add Dot with to setting color<br>
<br>
<b>func (rc *RCConfig) Dotcmplx(com complex128)</b><br>
x = real(com), y = imag(com)<br>
<br>
<b>func (rc *RCConfig) Dotcmplxc(com complex128, col color.Color)</b><br>
x = real(com), y = imag(com), setting color<br>
<br>
<b>func (rc *RCConfig) Draw()</b><br>
Additional display added Dot with Dot() function<br>
<br>
<b>func (rc *RCConfig) DrawTick(tick time.Duration)</b><br>
Automatically call the Draw()<br>
Recommend (1 << 23)　～　(1 << 25)<br>
<br>
<b>func (rc *RCConfig) End()</b><br>
Draw and End<br>
<br>
<b>func (rc *RCConfig) FillX(f func(float64) float64, delay func())</b><br>
Argument of func(float64) is value of x.<br>
<br>
<b>func (rc *RCConfig) FillXc(f func(float64) (float64, color.Color), delay func())</b><br>
FillX and setting color<br>
<br>
<b>func (rc *RCConfig) Len() int</b><br>
buffer size set by the argument of NewWindow() function<br>
<br>
<b>func (rc *RCConfig) Redraw()</b><br>
Display Dot in the ring buffer<br>
Must slower than the Dot() function<br>
<br>
<b>func (rc *RCConfig) RedrawTick(tick time.Duration)</b><br>
Automatically call the Redraw()<br>
Recommend (1 << 23)　～　(1 << 25)<br>
<br>
<b>func (rc *RCConfig) SafeConfig(f func())</b><br>
This function for Thread safe.<br>
<br>
<b>func (rc *RCConfig) Wait()</b><br>
Draw and Wait<br>
<br>


<h2>Argument/引数</h2>
<b>func NewWindow(scaleX, scaleY float64, bufSize int) *rcConfig</b><br>
NewWindow(scale(max)X, scale(max)Y, bufSize=>Dots Array(ring buffer) Size)<br>
NewWindow(X軸の最大値、Y軸の最大値、バッファサイズ=>点の配列(リングバッファ)の大きさ)<br>
<br>
<b>func (rc *rcConfig) Dot(x, y float64)</b><br>
Dot(x, y)<br>
Dot(x座標の位置, y座標の位置)<br>
<br>
<b>func (rc *rcConfig) DrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) DrawTick((draw loop interval))<br>
func (rc *rcConfig) DrawTick(描画の間隔を指定)<br>
<br>
<b>func (rc *rcConfig) RedrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) RedrawTick((redraw loop interval))<br>
func (rc *rcConfig) RedrawTick(再描画の間隔を指定)<br>
<br>
<b>func (rc *rcConfig) FillX(func(float64)(float64), func())</b><br>
func (rc *rcConfig) FillX((Argument:x Return:y), Delay)<br>
func (rc *rcConfig) FillX((引数:x 戻り値:y), 遅延処理)<br>


<h2>Event/イベント</h2>

click => Print x & y<br>
KeyPress Esc => close window<br>
KeyPress Arrow => Move<br>
KeyPress Z or X => Magnification<br>
KeyPress C => Undo the Move & Magnification<br>
KeyPress R => redraw(for debuging)<br>

クリック => XとYの座標を計算して表示します。<br>
Escキー => ウインドウを閉じます。<br>
矢印キー => 移動<br>
ZとXキー => 倍率変更<br>
Cキー => 移動と倍率をもとに戻す<br>
Rキー => 描画を更新します。(デバッグ用)<br>
