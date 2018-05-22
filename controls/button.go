package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/helper/mathh"
)

type Button struct {
	gui.BaseControl
	label *Label
}

func (c *Button) Children() []gui.Control { return []gui.Control{c.label} }

func (c *Button) ComputePossibleHorGeometry() (minWidth, maxWidth int) {
	scale := c.Window().OfflineCanvas().ScaleFactor()
	minWidth = c.label.MinWidth() + 2*(LineWidthI(scale)+HorPaddingI(scale))
	maxWidth = mathh.MaxInt
	return
}

func (c *Button) ComputePossibleVerGeometry() (minHeight, maxHeight int) {
	scale := c.Window().OfflineCanvas().ScaleFactor()
	height := c.label.MinHeight() + 2*(LineWidthI(scale)+VerPaddingI(scale))
	return height, height
}

func (c *Button) ComputeChildHorGeometry() (lefts, rights []int) {
	scale := c.Window().OfflineCanvas().ScaleFactor()
	left := c.Geometry().Left + LineWidthI(scale) + HorPaddingI(scale)
	right := c.Geometry().Right - LineWidthI(scale) - HorPaddingI(scale)
	return []int{left}, []int{right}
}

func (c *Button) ComputeChildVerGeometry() (tops, bottoms []int) {
	scale := c.Window().OfflineCanvas().ScaleFactor()
	top := c.Geometry().Top + LineWidthI(scale) + VerPaddingI(scale)
	bottom := c.Geometry().Bottom - LineWidthI(scale) - VerPaddingI(scale)
	return []int{top}, []int{bottom}
}

func (c Button) Draw(canvas gui.Canvas, region gui.RectangleI) {
	scale := canvas.ScaleFactor()
	lw := LineWidth(scale)
	br := BorderRadius(scale)
	rect := c.Geometry().ToF64().Inner(lw).ToRounded(br)
	canvas.FillRoundedRectangle(rect, backgroundColor)
	canvas.DrawRoundedRectangle(rect, borderColor, lw)
	c.label.Draw(canvas, region)
}

func (c Button) ProcessEvent(gui.Event) bool { return false }

func (c *Button) SetText(text string) {
	c.label.SetText(text)
}

func NewButton(text string) *Button {
	label := NewLabel(text, font, DefaultFontHeight, labelColor)
	return &Button{
		label: label,
	}
}
