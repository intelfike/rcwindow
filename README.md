<h1>Rectangle Coordinates Window</h1>

���{��:<br>
https://ja.wikipedia.org/wiki/%E7%9B%B4%E4%BA%A4%E5%BA%A7%E6%A8%99%E7%B3%BB<br>
�����GUI(Shiny)�ŕ\�����Ă���܂��B<br>
<br>
English:<br>
https://en.wikipedia.org/wiki/Cartesian_coordinate_system<br>
Shiny display this<br>

<h2>get command/����R�}���h</h2>

<pre>
go get github.com/intelfike/rcwindow
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
func (rc *rcConfig) DrawTick((draw loop interval))<br>
func (rc *rcConfig) DrawTick(�`��̊Ԋu���w��)<br>
<br>
<b>func (rc *rcConfig) RedrawTick(tick time.Duration)</b><br>
func (rc *rcConfig) RedrawTick((redraw loop interval))<br>
func (rc *rcConfig) RedrawTick(�ĕ`��̊Ԋu���w��)<br>
<br>
<b>func (rc *rcConfig) FillX(func(float64)(float64), func())</b><br>
func (rc *rcConfig) FillX((Argument:x Return:y), Delay)<br>
func (rc *rcConfig) FillX((����:x �߂�l:y), �x������)<br>


<h2>Event/�C�x���g</h2>

click => Print x & y<br>
KeyPress Esc => close window<br>
KeyPress Arrow => Move<br>
KeyPress Z or X => Magnification<br>
KeyPress C => Undo the Move & Magnification<br>
KeyPress R => redraw(for debuging)<br>

�N���b�N => X��Y�̍��W���v�Z���ĕ\�����܂��B<br>
Esc�L�[ => �E�C���h�E����܂��B<br>
���L�[ => �ړ�<br>
Z��X�L�[ => �{���ύX<br>
C�L�[ => �ړ��Ɣ{�������Ƃɖ߂�<br>
R�L�[ => �`����X�V���܂��B(�f�o�b�O�p)<br>
