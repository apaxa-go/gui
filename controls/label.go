package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
)

type Label struct {
	gui.BaseControl
	text  string
	font  gui.Font
	color gui.ColorF64
	align basetypes.Align
}

//
// Empty implementations
//

func (*Label) Children() []gui.Control                            { return nil }
func (*Label) ComputeChildHorGeometry() (lefts, rights []float64) { return nil, nil }
func (*Label) ComputeChildVerGeometry() (tops, bottoms []float64) { return nil, nil }

//
//
//

func (c *Label) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	width := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font).X
	return width, width
}

func (c *Label) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	height := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font).Y
	return height, height
}

func (c Label) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	place := c.align.ApplyF64(c.Geometry(), c.MinSize())
	canvas.DrawTextLine(c.text, c.font, place.LT(), c.color)
}

func (c Label) ProcessEvent(gui.Event) bool { return false }

func (c *Label) GetText() string { return c.text }
func (c *Label) SetText(text string) {
	if c.text == text {
		return
	}
	c.text = text
	c.SetUPGIR(false)
}

func (c *Label) GetFont() gui.Font { return c.font }
func (c *Label) SetFont(font gui.Font) {
	if c.font == font {
		return
	}
	c.font = font
	c.SetUPGIR(false)
}

func (c *Label) GetColor() gui.ColorF64 { return c.color }
func (c *Label) SetColor(color gui.ColorF64) {
	if c.color == color {
		return
	}
	c.color = color
	c.SetIR()
}

func (c *Label) GetAlign() basetypes.Align { return c.align }
func (c *Label) SetAlign(align basetypes.Align) {
	align = align.KeepSize()
	if c.align == align {
		return
	}
	c.align = align
	c.SetIR()
}

func NewLabel(text string, font gui.Font, color gui.ColorF64) *Label {
	return &Label{
		text:  text,
		font:  font,
		color: color,
		align: basetypes.AlignCenter,
	}
}
