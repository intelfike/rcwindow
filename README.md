<h1>Rectangle Coordinates Window</h1>

日本語:<br>
https://ja.wikipedia.org/wiki/%E7%9B%B4%E4%BA%A4%E5%BA%A7%E6%A8%99%E7%B3%BB<br>
これをGUI(shiny)で表示してくれます。<br>
<br>

<h2>getコマンド</h2>

<pre>
go get github.com/intelfike/rcwindow
</pre>

<h2>使い方</h2>

<pre>
rc := rcwindow.NewWindow(1, 1, 10)	//ウィンドウを表示
rc.Dot(0.5, 0.5)	//ドットをセット
rc.Wait()	//処理を待機(待機しないとメインスレッドと同時にウィンドウの描画
終了してしまう)
</pre>

<h2>構造体とメソッド</h2>
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
ウィンドウを作成して表示します。<br>
<br>
<b>func (rc *RCConfig) Clear()</b><br>
セットされたドットをすべて削除して再描画します。<br>
<br>
<b>func (rc *RCConfig) Dot(x, y float64)</b><br>
ドットを追加します(表示されません)<br>
<br>
<b>func (rc *RCConfig) Dotc(x, y float64, col color.Color)</b><br>
ドットを追加します(表示されません)。色を指定できます。<br>
<br>
<b>func (rc *RCConfig) Dotcmplx(com complex128)</b><br>
x = real(com), y = imag(com)<br>
<br>
<b>func (rc *RCConfig) Dotcmplxc(com complex128, col color.Color)</b><br>
x = real(com), y = imag(com), 色を指定できます。<br>
<br>
<b>func (rc *RCConfig) Draw()</b><br>
Dot()やDotc()によって追加された最新のドットを表示します。<br>
最新のドットを追加表示するだけなので、バッファサイズが足りなくなった場合にも過去のドットは消去されません。<br>
<br>
<b>func (rc *RCConfig) DrawTick(tick time.Duration)</b><br>
指定された時間ごとにDraw()関数を呼び出します。<br>
推奨値 (1 << 23)　～　(1 << 25)<br>
<br>
<b>func (rc *RCConfig) End()</b><br>
Draw()を呼び出して終了します。ウィンドウはメインスレッドが終了するまで消えません。<br>
<br>
<b>func (rc *RCConfig) FillX(f func(float64) float64, delay func())</b><br>
func(float64)の引数はx軸の値で、戻り値はy軸の値です。<br>
NewWindow()で指定されたバッファサイズを埋めるようにxの増加量が調整されます。<br>
<br>
<b>func (rc *RCConfig) FillXc(f func(float64) (float64, color.Color), delay func())</b><br>
色を指定できるFillX()です。<br>
<br>
<b>func (rc *RCConfig) Len() int</b><br>
NewWindow()で指定したバッファサイズです。<br>
<br>
<b>func (rc *RCConfig) Redraw()</b><br>
バッファ内のドットを表示します。<br>
基本的にDraw()と比べ低速ですが、バッファサイズが小さいとき、アニメーションとして動かしたいときに使います。<br>
<br>
<b>func (rc *RCConfig) RedrawTick(tick time.Duration)</b><br>
指定された時間ごとにRedraw()関数を呼び出します。<br>
推奨値(1 << 23)　～　(1 << 25)<br>
<br>
<b>func (rc *RCConfig) SafeConfig(f func())</b><br>
この関数はスレッドセーフに値を設定できます。RCConfigの値を設定したいときに使います。<br>
<br>
<b>func (rc *RCConfig) Wait()</b><br>
Draw()を呼び出してウィンドウの終了を待機します。<br>
<br>


<h2>引数</h2>
<b>func NewWindow(scaleX, scaleY float64, bufSize int) *rcConfig</b><br>
NewWindow(X軸の最大値、Y軸の最大値、バッファサイズ=>点の配列(リングバッファ)の大きさ)<br>
<br>
<b>func (rc *rcConfig) Dot(x, y float64)</b><br>
func (rc *rcConfig) Dot(x座標の位置, y座標の位置)<br>
<br>
<b>func (rc *rcConfig) DrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) DrawTick(描画の間隔を指定)<br>
<br>
<b>func (rc *rcConfig) RedrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) RedrawTick(再描画の間隔を指定)<br>
<br>
<b>func (rc *rcConfig) FillX(func(float64)(float64), func())</b><br>
func (rc *rcConfig) FillX((引数:x 戻り値:y), 遅延処理)<br>


<h2>イベント</h2>

クリック => XとYの座標を計算して表示します。<br>
Escキー => ウインドウを閉じます。<br>
矢印キー => 移動<br>
ZとXキー => 倍率変更<br>
Cキー => 移動と倍率をもとに戻す<br>
Rキー => 描画を更新します。<br>
