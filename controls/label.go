package controls

import "github.com/apaxa-go/gui"

type Label struct {
	gui.BaseControl
	text  string
	font  gui.Font
	color gui.ColorF64
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
	// TODO align text in geometry
	canvas.DrawTextLine(c.text, c.font, c.Geometry().LT(), c.color)
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

func (c *Label) SetColor(color gui.ColorF64) {
	if c.color == color {
		return
	}
	c.color = color
	c.SetIR()
}

func NewLabel(text string, font gui.Font, color gui.ColorF64) *Label {
	return &Label{
		text:  text,
		font:  font,
		color: color,
	}
}
