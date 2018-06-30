// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package buttons

import "github.com/apaxa-go/gui/scvi"

var minimizeBackgroundColor = ColorF64{}.MakeFromRGB8(253, 185, 90)

// "-"
var minimizeImage = scvi.SCVI{
	Size:       PointF64{12, 12},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeLine(
			PointF64{2.25, 6},
			PointF64{9.75, 6},
			0.7,
			0.7,
		),
	},
}

type MinimizeButton struct {
	BaseControl
	hover bool
}

func (c *MinimizeButton) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return size, size, size
}

func (c *MinimizeButton) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return size, size, size
}

func (c *MinimizeButton) actionAllowed() bool {
	return !c.Window().DisplayState().IsFullScreen() && c.Window().IsMinimizeAllowed()
}

func (c MinimizeButton) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := Align(0).MakeCenter().ApplyF64(c.Geometry(), PointF64{size, size})
	circle := CircleF64{place.Center(), size / 2}.Inner(borderWidth)
	if (c.Window().IsMain() || c.hover) && c.actionAllowed() {
		canvas.FillCircle(circle, minimizeBackgroundColor)
		if c.hover {
			minimizeImage.Draw(canvas, place, imageColor)
		}
	} else {
		canvas.FillCircle(circle, inactiveBackgroundColor)
	}
	canvas.DrawCircle(circle, borderColor, borderWidth)
}

func (c *MinimizeButton) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	if event.Kind.IsPress() && event.Button.IsLeft() && c.actionAllowed() {
		c.Window().Minimize()
	}
	return true
}

func NewMinimizeButton() *MinimizeButton { return &MinimizeButton{} }
