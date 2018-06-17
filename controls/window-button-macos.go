// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui/scvi"
)

const windowButtonMacOSSize = SmallHeight
const windowButtonMacOSBorderWidth = 0.5

var windowButtonMacOSBorderColor = ColorF64{}.MakeFromRGBA8(0, 0, 0, 32)

type windowButtonMacOS struct {
	BaseControl
	image                   scvi.SCVI
	imageColor              ColorF64
	backgroundColor         ColorF64
	inactiveBackgroundColor ColorF64
	action                  WindowButtonAction
	hover                   bool
}

func (c *windowButtonMacOS) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return windowButtonMacOSSize, windowButtonMacOSSize, windowButtonMacOSSize
}

func (c *windowButtonMacOS) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return windowButtonMacOSSize, windowButtonMacOSSize, windowButtonMacOSSize
}

func (c windowButtonMacOS) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := Align(0).MakeCenter().ApplyF64(c.Geometry(), PointF64{windowButtonMacOSSize, windowButtonMacOSSize})
	circle := CircleF64{place.Center(), windowButtonMacOSSize / 2}.Inner(windowButtonMacOSBorderWidth)
	if c.Window().IsMain() || c.hover {
		canvas.FillCircle(circle, c.backgroundColor)
		if c.hover {
			c.image.Draw(canvas, place, c.imageColor)
		}
	} else {
		canvas.FillCircle(circle, c.inactiveBackgroundColor)
	}
	canvas.DrawCircle(circle, windowButtonMacOSBorderColor, windowButtonMacOSBorderWidth)
}

func (c *windowButtonMacOS) Image() scvi.SCVI { return c.image }
func (c *windowButtonMacOS) SetImage(image scvi.SCVI) {
	if c.image.Equal(image) {
		return
	}
	c.image = image
	c.SetIR()
}

func (c *windowButtonMacOS) BackgroundColor() ColorF64 { return c.backgroundColor }
func (c *windowButtonMacOS) SetBackgroundColor(backgroundColor ColorF64) {
	if c.backgroundColor == backgroundColor {
		return
	}
	c.backgroundColor = backgroundColor
	c.SetIR()
}

func (c *windowButtonMacOS) ImageColor() ColorF64 { return c.imageColor }
func (c *windowButtonMacOS) SetImageColor(imageColor ColorF64) {
	if c.imageColor == imageColor {
		return
	}
	c.imageColor = imageColor
	c.SetIR()
}

func (c *windowButtonMacOS) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	processed = true
	if !event.Kind.IsPress() {
		return
	}
	switch c.action {
	case WindowButtonActionClose:
		c.Window().Close()
	case WindowButtonActionMinimize:
		c.Window().Minimize()
	case WindowButtonActionMaximize:
		c.Window().Maximize()
	}
	return
}

func newWindowButtonMacOS(image scvi.SCVI, imageColor, backgroundColor, inactiveBackgroundColor ColorF64, action WindowButtonAction) *windowButtonMacOS {
	return &windowButtonMacOS{
		image:                   image,
		imageColor:              imageColor,
		backgroundColor:         backgroundColor,
		inactiveBackgroundColor: inactiveBackgroundColor,
		action:                  action,
	}
}

func newWindowButtonMacOSClose() *windowButtonMacOS {
	return newWindowButtonMacOS(windowButtonMacOSCloseImage, windowButtonMacOSImageColor, windowButtonMacOSCloseBackgroundColor, windowButtonMacOSInactiveBackgroundColor, WindowButtonActionClose)
}
func newWindowButtonMacOSMinimize() *windowButtonMacOS {
	return newWindowButtonMacOS(windowButtonMacOSMinimizeImage, windowButtonMacOSImageColor, windowButtonMacOSMinimizeBackgroundColor, windowButtonMacOSInactiveBackgroundColor, WindowButtonActionMinimize)
}
func newWindowButtonMacOSMaximize() *windowButtonMacOS {
	return newWindowButtonMacOS(windowButtonMacOSMaximizeImage, windowButtonMacOSImageColor, windowButtonMacOSMaximizeBackgroundColor, windowButtonMacOSInactiveBackgroundColor, WindowButtonActionMaximize)
}

// "x"
var windowButtonMacOSCloseImage = scvi.SCVI{
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

// "-"
var windowButtonMacOSMinimizeImage = scvi.SCVI{
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

// "<\>"
var windowButtonMacOSMaximizeImage = scvi.SCVI{
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
var windowButtonMacOSDemaximizeImage = scvi.SCVI{
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

var windowButtonMacOSImageColor = ColorF64{0, 0, 0, 1}

var windowButtonMacOSInactiveBackgroundColor = ColorF64{}.MakeFromRGB8(208, 208, 208)
var windowButtonMacOSCloseBackgroundColor = ColorF64{}.MakeFromRGB8(249, 100, 99)
var windowButtonMacOSMinimizeBackgroundColor = ColorF64{}.MakeFromRGB8(253, 185, 90)
var windowButtonMacOSMaximizeBackgroundColor = ColorF64{}.MakeFromRGB8(68, 191, 66)
var windowButtonMacOSDisabledBackgroundColor = ColorF64{}.MakeFromRGB8(210, 210, 210)
