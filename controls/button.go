// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

const ButtonBestWidth = 60

type Button struct {
	BaseControl
	label  *TextImage
	align  Align
	action func()
}

func (c *Button) Children() []Control { return []Control{c.label} }

func (c *Button) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	minWidth = c.label.BestWidth() + 2*(BorderWidth+HorPadding)
	bestWidth = mathh.Max2Float64(ButtonBestWidth, minWidth)
	maxWidth = mathh.PositiveInfFloat64()
	return
}

func (c *Button) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	height := c.label.BestHeight() + 2*(BorderWidth+VerPadding)
	return height, height, height
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

func (c *Button) Draw(canvas Canvas, region RectangleF64) {
	rect := c.align.ApplyF64(c.Geometry(), c.MinSize()).Inner(BorderWidth).ToRounded(BorderRadius)
	if c.Window().IfControlFocused(c) {
		canvas.FillRoundedRectangle(rect, ColorF64{}.MakeFromRGB8(255, 0, 0)) // TODO
	} else {
		canvas.FillRoundedRectangle(rect, backgroundColor)
	}
	canvas.DrawRoundedRectangle(rect, borderColor, BorderWidth)
	c.label.Draw(canvas, region)
}

func (c *Button) OnPointerButtonEvent(e PointerButtonEvent) (processed bool) {
	if c.action != nil {
		c.action()
	}
	return true
}

func (c *Button) FocusCandidate(reverse bool, current Control) Control {
	if current == nil {
		return c
	}
	return nil
}

func (c *Button) OnFocus(e FocusEvent) {
	c.SetIR()
}

func (c *Button) GetText() string { return c.label.text }
func (c *Button) SetText(text string) {
	c.label.SetText(text)
}

func (c *Button) GetAlign() Align { return c.align }
func (c *Button) SetAlign(align Align) {
	if c.align == align {
		return
	}
	c.align = align
	c.SetIR()
}

func (c *Button) SetAction(action func()) {
	c.action = action
}

func NewButton(text string) *Button {
	r := &Button{
		label: NewLabel(text, defaultFont, controlTextColor),
		align: Align(0).MakeStretchCenter(),
	}
	r.BaseControl.SetParent(r.label, r)
	return r
}
