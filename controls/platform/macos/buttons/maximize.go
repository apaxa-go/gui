// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package buttons

import (
	"github.com/apaxa-go/gui/drivers/cocoa"
	"github.com/apaxa-go/gui/scvi"
)

var maximizeBackgroundColor = ColorF64{}.MakeFromRGB8(68, 191, 66)

// "<\>"
var enterFullScreenImage = scvi.SCVI{
	Size:       PointF64{12, 12},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeFilledContour(
			[]PointF64{{3, 4.5}, {3, 9}, {7.5, 9}},
			1,
			0.5,
		),
		scvi.MakeFilledContour(
			[]PointF64{{4.5, 3}, {9, 3}, {9, 7.5}},
			1,
			0.5,
		),
	},
}

// ">\<"
var exitFullScreenImage = scvi.SCVI{
	Size:       PointF64{12, 12},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeFilledContour(
			[]PointF64{{3, 6}, {6, 6}, {6, 9}},
			1,
			1,
		),
		scvi.MakeFilledContour(
			[]PointF64{{6, 3}, {6, 6}, {9, 6}},
			1,
			1,
		),
	},
}

// "+"
var zoomImage = scvi.SCVI{
	Size:       PointF64{12, 12},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeLine(
			PointF64{6, 2},
			PointF64{6, 10},
			0.7,
			0.8,
		),
		scvi.MakeLine(
			PointF64{2, 6},
			PointF64{10, 6},
			0.7,
			0.8,
		),
	},
}

type MaximizeButton struct {
	BaseControl
	hover bool
	alt   bool
}

func (c *MaximizeButton) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return size, size, size
}

func (c *MaximizeButton) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return size, size, size
}

func (c MaximizeButton) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := Align(0).MakeCenter().ApplyF64(c.Geometry(), PointF64{size, size})
	circle := CircleF64{place.Center(), size / 2}.Inner(borderWidth)
	defer canvas.DrawCircle(circle, borderColor, borderWidth)

	if !c.Window().IsMain() && !c.hover {
		canvas.FillCircle(circle, inactiveBackgroundColor)
		return
	}

	switch c.Window().DisplayState() {
	case WindowDisplayState(0).MakeNormal():
		if c.Window().IsFullScreenAllowed() && !c.alt {
			canvas.FillCircle(circle, maximizeBackgroundColor)
			if c.hover {
				enterFullScreenImage.Draw(canvas, place, imageColor)
			}
			return
		} else { //nolint: golint
			canvas.FillCircle(circle, maximizeBackgroundColor)
			if c.hover {
				zoomImage.Draw(canvas, place, imageColor)
			}
			return
		}
	case WindowDisplayState(0).MakeFullScreen():
		if c.Window().IsNormalAllowed() {
			canvas.FillCircle(circle, maximizeBackgroundColor)
			if c.hover {
				exitFullScreenImage.Draw(canvas, place, imageColor)
			}
			return
		}
	}

	canvas.FillCircle(circle, inactiveBackgroundColor)
}

func (c *MaximizeButton) AfterAttachToWindowEvent() {
	if _, ok := c.Window().Driver().(*cocoa.Window); ok {
		c.Window().SubscribeToModifiersEvent(c)
	}
}

func (c *MaximizeButton) BeforeDetachFromWindowEvent() {
	c.Window().UnsubscribeFromModifiersEvent(c)
}

func (c *MaximizeButton) OnModifiersEvent(event ModifiersEvent) {
	alt := event.Modifiers.IsAltPressed()
	if c.alt != alt {
		c.alt = alt
		c.SetIR()
	}
}

func (c *MaximizeButton) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	processed = true
	if !event.Kind.IsPress() || !event.Button.IsLeft() {
		return
	}
	switch c.Window().DisplayState() {
	case WindowDisplayState(0).MakeNormal():
		if c.Window().IsMaximizeAllowed() && !c.alt {
			c.Window().EnterFullScreen()
		} else {
			c.Window().Driver().(*cocoa.Window).ToggleZoom() // Assume that driver can't be switched while running.
		}
	case WindowDisplayState(0).MakeFullScreen():
		if c.Window().IsNormalAllowed() {
			c.Window().ExitFullScreen()
		}
	}
	return
}

func NewMaximizeButton() *MaximizeButton { return &MaximizeButton{} }
