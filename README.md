<h1>Rectangle Coordinates Window</h1>

Japanese:<br>
https://ja.wikipedia.org/wiki/%E7%9B%B4%E4%BA%A4%E5%BA%A7%E6%A8%99%E7%B3%BB<br>
English:<br>
https://en.wikipedia.org/wiki/Cartesian_coordinate_system<br>


<h2>Usage/使い方</h2>

<pre>
rc := rcwindow.NewWindow(1, 1, 10)
//rc(struct) configration here //rc構造体をここで編集する
rc.Start()
rc.Dot(0.4, -0.3)
rc.Wait()
</pre>

<h2>Example/例</h2>

<pre>
package main

import (
	"math/cmplx"
	"github.com/intelfike/rcwindow"
	"time"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 10)
	rc.Start()
	c := cmplx.Pow(1i, 1.0/25.0)
	v := 1i
	for n := 0; n <= 10000; n++{
		rc.Dot(real(v), imag(v))
		v *= c
		time.Sleep(1 << 22)
	}
	rc.Wait()
}
</pre>

<h2>Argument/引数</h2>

NewWindow(scaleX, scaleY, bufSize=>Dots Array(ring buffer) Size)<br>
NewWindow(X軸の最大値、Y軸の最大値、バッファサイズ=>点の配列(リングバッファ)の大きさ)<br>
<br>
Dot(x, y)<br>
Dot(x座標の位置, y座標の位置)<br>

<h2>Event/イベント</h2>

click => Print x & y
KeyPress Esc => close window
KeyPress UpArrow => Scale / 1.2
KeyPress DownArrow => Scale * 1.2
KeyPress R => redraw(for debuging)

クリック => XとYの座標を計算して表示します。
Escキー => ウインドウを閉じます。
↑キー => スケールを1.2分の1に縮小します。
↓キー => スケールを1.2倍に拡大します。
Rキー => 描画を更新します。(デバッグ用)