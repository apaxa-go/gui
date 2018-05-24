package controls

import (
	"github.com/apaxa-go/gui"
)

type Button struct {
	gui.BaseControl
	label *Label
}

func (c *Button) Children() []gui.Control { return []gui.Control{c.label} }

func (c *Button) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	minWidth = c.label.MinWidth() + 2*(BorderWidth+HorPadding)
	maxWidth = gui.PosInfF64()
	return
}

func (c *Button) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	height := c.label.MinHeight() + 2*(BorderWidth+VerPadding)
	return height, height
}

func (c *Button) ComputeChildHorGeometry() (lefts, rights []float64) {
	left := c.Geometry().Left + BorderWidth + HorPadding
	right := c.Geometry().Right - (BorderWidth + HorPadding)
	return []float64{left}, []float64{right}
}

func (c *Button) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top := c.Geometry().Top + BorderWidth + VerPadding
	bottom := c.Geometry().Bottom - (BorderWidth + VerPadding)
	return []float64{top}, []float64{bottom}
}

func (c Button) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	rect := c.Geometry().Inner(BorderWidth).ToRounded(BorderRadius)
	canvas.FillRoundedRectangle(rect, backgroundColor)
	canvas.DrawRoundedRectangle(rect, borderColor, BorderWidth)
	canvas.DrawLine(c.Geometry().LT(), c.Geometry().LT().Add(gui.PointF64{10, 10}), labelColor, BorderWidth*3) // TODO remove this debug
	c.label.Draw(canvas, region)
}

func (c Button) ProcessEvent(gui.Event) bool { return false }

func (c *Button) SetText(text string) {
	c.label.SetText(text)
}

func NewButton(text string) *Button {
	r := &Button{
		label: NewLabel(text, defaultFont, labelColor),
	}
	gui.SetParent(r.label, r)
	return r
}
