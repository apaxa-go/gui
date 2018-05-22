package controls

import "github.com/apaxa-go/gui"

type Label struct {
	gui.BaseControl
	text       string
	font       gui.Font
	fontHeight float64
	color      gui.ColorF64
}

//
// Empty implementations
//

func (*Label) Children() []gui.Control                        { return nil }
func (*Label) ComputeChildHorGeometry() (lefts, rights []int) { return nil, nil }
func (*Label) ComputeChildVerGeometry() (tops, bottoms []int) { return nil, nil }

//
//
//

func (c *Label) ComputePossibleHorGeometry() (minWidth, maxWidth int) {
	width := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font, c.fontHeight).ToI().X // TODO ToI() may cut size
	return width, width
}

func (c *Label) ComputePossibleVerGeometry() (minHeight, maxHeight int) {
	height := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font, c.fontHeight).ToI().Y
	return height, height
}

func (c Label) Draw(canvas gui.Canvas, region gui.RectangleI) {
	canvas.DrawTextLine(c.text, c.font, c.fontHeight, c.Geometry().LT().ToF64(), c.color)
}

func (c Label) ProcessEvent(gui.Event) bool { return false }

func (c *Label) SetText(text string) {
	if c.text == text {
		return
	}
	c.text = text
	c.SetUPGIR(false)
}

func (c *Label) SetFont(font gui.Font) {
	if c.font == font {
		return
	}
	c.font = font
	c.SetUPGIR(false)
}

func (c *Label) SetFontHeight(fontHeight float64) {
	if c.fontHeight == fontHeight {
		return
	}
	c.fontHeight = fontHeight
	c.SetUPGIR(false)
}

func (c *Label) SetColor(color gui.ColorF64) {
	if c.color == color {
		return
	}
	c.color = color
	c.SetIR()
}

func NewLabel(text string, font gui.Font, fontHeight float64, color gui.ColorF64) *Label {
	return &Label{
		text:       text,
		font:       font,
		fontHeight: fontHeight,
		color:      color,
	}
}
