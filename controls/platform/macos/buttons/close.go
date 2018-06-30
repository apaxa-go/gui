// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package buttons

import "github.com/apaxa-go/gui/scvi"

var closeBackgroundColor = ColorF64{}.MakeFromRGB8(249, 100, 99)

// "x"
var closeImage = scvi.SCVI{
	Size:       PointF64{12, 12},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeLine(
			PointF64{3.25, 3.25},
			PointF64{8.75, 8.75},
			0.7,
			0.8,
		),
		scvi.MakeLine(
			PointF64{3.25, 8.75},
			PointF64{8.75, 3.25},
			0.7,
			0.7,
		),
	},
}

type CloseButton struct {
	BaseControl
	hover bool
}

func (c *CloseButton) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return size, size, size
}

func (c *CloseButton) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return size, size, size
}

func (c *CloseButton) actionAllowed() bool {
	return c.Window().IsCloseAllowed()
}

func (c CloseButton) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := Align(0).MakeCenter().ApplyF64(c.Geometry(), PointF64{size, size})
	circle := CircleF64{place.Center(), size / 2}.Inner(borderWidth)
	if (c.Window().IsMain() || c.hover) && c.actionAllowed() {
		canvas.FillCircle(circle, closeBackgroundColor)
		if c.hover {
			closeImage.Draw(canvas, place, imageColor)
		}
	} else {
		canvas.FillCircle(circle, inactiveBackgroundColor)
	}
	canvas.DrawCircle(circle, borderColor, borderWidth)
}

func (c *CloseButton) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	if event.Kind.IsPress() && event.Button.IsLeft() && c.actionAllowed() {
		c.Window().Close()
	}
	return true
}

func NewCloseButton() *CloseButton { return &CloseButton{} }
