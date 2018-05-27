// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
)

type Button struct {
	gui.BaseControl
	label *Label
	align basetypes.Align
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
	left, right := c.align.Hor().ApplyF64(c.Geometry().Left, c.Geometry().Right, c.MinSize().X)
	left += BorderWidth + HorPadding
	right -= BorderWidth + HorPadding
	return []float64{left}, []float64{right}
}

func (c *Button) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top, bottom := c.align.Ver().ApplyF64(c.Geometry().Top, c.Geometry().Bottom, c.MinSize().Y)
	top += BorderWidth + VerPadding
	bottom -= BorderWidth + VerPadding
	return []float64{top}, []float64{bottom}
}

func (c Button) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	rect := c.align.ApplyF64(c.Geometry(), c.MinSize()).Inner(BorderWidth).ToRounded(BorderRadius)
	canvas.FillRoundedRectangle(rect, backgroundColor)
	canvas.DrawRoundedRectangle(rect, borderColor, BorderWidth)
	c.label.Draw(canvas, region)
}

func (c Button) ProcessEvent(gui.Event) bool { return false }

func (c *Button) GetText() string { return c.label.text }
func (c *Button) SetText(text string) {
	c.label.SetText(text)
}

func (c *Button) GetAlign() basetypes.Align { return c.align }
func (c *Button) SetAlign(align basetypes.Align) {
	if c.align == align {
		return
	}
	c.align = align
	c.SetIR()
}

func NewButton(text string) *Button {
	r := &Button{
		label: NewLabel(text, defaultFont, labelColor),
		align: basetypes.AlignStretchCenter,
	}
	gui.SetParent(r.label, r)
	return r
}
