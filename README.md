<h1>Rectangle Coordinates Window</h1>

���{��:<br>
https://ja.wikipedia.org/wiki/%E7%9B%B4%E4%BA%A4%E5%BA%A7%E6%A8%99%E7%B3%BB<br>
English:<br>
https://en.wikipedia.org/wiki/Cartesian_coordinate_system<br>

<h2>get command/����R�}���h</h2>

<pre>
go get github.com/intelfike/rcwindow
</pre>

<h2>Usage/�g����</h2>

<pre>
rc := rcwindow.NewWindow(1, 1, 10)
//rc(struct) configration here //rc�\���̂������ŕҏW����
rc.Start()
rc.Dot(0.4, -0.3)
rc.Draw()
rc.Wait() //OR rc.End()
</pre>

<h2>Example/��</h2>

<img src="https://github.com/intelfike/images/blob/master/RCM2.png" width="400" height="400">
�_�����邭���]���܂��B
<pre>
package main

import (
	"math/cmplx"
	"github.com/intelfike/rcwindow"
	"time"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 10)
    rc.DotSize = 4
	rc.Start()
	c := cmplx.Pow(1i, 1.0/25.0)
	v := 1i
	rc.DrawTick(1 << 25)
	for n := 0; n <= 10000; n++{
		rc.Dot(real(v), imag(v))
		v *= c
		time.Sleep(1 << 22)
	}
	rc.Wait()
}
</pre>
<hr>
<img src="https://github.com/intelfike/images/blob/master/RCW.jpg" width="400" height="400">
<pre>
package main

import (
	"math/cmplx"
	"github.com/intelfike/rcwindow"
)

func main() {
	rc := rcwindow.NewWindow(1, 1, 1000)
    rc.DotSize = 4
	rc.Start()
	rc.DrawTick(1 << 25)
	c := cmplx.Pow(1i, 1.0/25.0)
	v := 1i
	for n := 0; n <= 400; n++{
		rc.Dot(float64(n)/200 - 1, imag(v))
		v *= c
	}
	rc.Wait()
}
</pre>

<h2>Argument/����</h2>
<b>func NewWindow(scaleX, scaleY float64, bufSize int) *rcConfig</b><br>
NewWindow(scale(max)X, scale(max)Y, bufSize=>Dots Array(ring buffer) Size)<br>
NewWindow(X���̍ő�l�AY���̍ő�l�A�o�b�t�@�T�C�Y=>�_�̔z��(�����O�o�b�t�@)�̑傫��)<br>
<br>
<b>func (rc *rcConfig) Dot(x, y float64)</b><br>
Dot(x, y)<br>
Dot(x���W�̈ʒu, y���W�̈ʒu)<br>
<br>
<b>func (rc *rcConfig) DrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) DrawTick((redraw loop interval))<br>
func (rc *rcConfig) DrawTick(�ĕ`��̊Ԋu���w��)<br>

<h2>Event/�C�x���g</h2>

click => Print x & y<br>
KeyPress Esc => close window<br>
KeyPress UpArrow => Scale / 1.2<br>
KeyPress DownArrow => Scale * 1.2<br>
KeyPress R => redraw(for debuging)<br>

�N���b�N => X��Y�̍��W���v�Z���ĕ\�����܂��B<br>
Esc�L�[ => �E�C���h�E����܂��B<br>
���L�[ => �X�P�[����1.2����1�ɏk�����܂��B<br>
���L�[ => �X�P�[����1.2�{�Ɋg�債�܂��B<br>
R�L�[ => �`����X�V���܂��B(�f�o�b�O�p)<br>
