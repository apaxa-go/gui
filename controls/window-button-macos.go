// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/scvi"
)

const windowButtonMacOSSize = SmallHeight

type windowButtonMacOS struct {
	gui.BaseControl
	image           scvi.SCVI
	imageColor      gui.ColorF64
	backgroundColor gui.ColorF64
}

func (c *windowButtonMacOS) Children() []gui.Control { return nil }

func (c *windowButtonMacOS) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	return windowButtonMacOSSize, windowButtonMacOSSize
}

func (c *windowButtonMacOS) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	return windowButtonMacOSSize, windowButtonMacOSSize
}

func (c *windowButtonMacOS) ComputeChildHorGeometry() (lefts, rights []float64) { return nil, nil }

func (c *windowButtonMacOS) ComputeChildVerGeometry() (tops, bottoms []float64) { return nil, nil }

func (c windowButtonMacOS) Draw(canvas gui.Canvas, _ gui.RectangleF64) {
	// TODO use region
	place := basetypes.AlignCenter.ApplyF64(c.Geometry(), c.MinSize())
	circle := gui.CircleF64{place.Center(), windowButtonMacOSSize / 2}.Inner(BorderWidth)
	canvas.FillCircle(circle, c.backgroundColor)
	canvas.DrawCircle(circle, borderColor, BorderWidth)
	c.image.Draw(canvas, place, c.imageColor)
}

func (c *windowButtonMacOS) Image() scvi.SCVI { return c.image }
func (c *windowButtonMacOS) SetImage(image scvi.SCVI) {
	if c.image.Equal(image) {
		return
	}
	c.image = image
	c.SetIR()
}

func (c *windowButtonMacOS) BackgroundColor() gui.ColorF64 { return c.backgroundColor }
func (c *windowButtonMacOS) SetBackgroundColor(backgroundColor gui.ColorF64) {
	if c.backgroundColor == backgroundColor {
		return
	}
	c.backgroundColor = backgroundColor
	c.SetIR()
}

func (c *windowButtonMacOS) ImageColor() gui.ColorF64 { return c.imageColor }
func (c *windowButtonMacOS) SetImageColor(imageColor gui.ColorF64) {
	if c.imageColor == imageColor {
		return
	}
	c.imageColor = imageColor
	c.SetIR()
}

func newWindowButtonMacOS(image scvi.SCVI, imageColor, backgroundColor gui.ColorF64) *windowButtonMacOS {
	return &windowButtonMacOS{
		image:           image,
		imageColor:      imageColor,
		backgroundColor: backgroundColor,
	}
}

// "x"
var windowButtonMacOSCloseImage = scvi.SCVI{
	Size:       gui.PointF64{13, 13},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeLine(
			gui.PointF64{3.5, 3.5},
			gui.PointF64{9.5, 9.5},
			0.7,
			1,
		),
		scvi.MakeLine(
			gui.PointF64{3.5, 9.5},
			gui.PointF64{9.5, 3.5},
			0.7,
			1,
		),
	},
}

// "-"
var windowButtonMacOSHideImage = scvi.SCVI{
	Size:       gui.PointF64{13, 13},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeLine(
			gui.PointF64{2, 6.5},
			gui.PointF64{11, 6.5},
			0.7,
			1,
		),
	},
}

// "<\>"
var windowButtonMacOSMaximizeImage = scvi.SCVI{
	Size:       gui.PointF64{13, 13},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeFilledContour(
			[]gui.PointF64{{3.5, 4.5}, {3.5, 9.5}, {8.5, 9.5}},
			1,
			1,
		),
		scvi.MakeFilledContour(
			[]gui.PointF64{{4.5, 3.5}, {9.5, 3.5}, {9.5, 8.5}},
			1,
			1,
		),
	},
}

// ">\<"
var windowButtonMacOSDemaximizeImage = scvi.SCVI{
	Size:       gui.PointF64{13, 13},
	KeepAspect: true,
	Elements: []scvi.Primitive{
		scvi.MakeFilledContour(
			[]gui.PointF64{{3.5, 6.5}, {6.5, 6.5}, {6.5, 9.5}},
			1,
			1,
		),
		scvi.MakeFilledContour(
			[]gui.PointF64{{6.5, 3.5}, {6.5, 6.5}, {9.5, 6.5}},
			1,
			1,
		),
	},
}

var windowButtonMacOSImageColor = gui.ColorF64{0, 0, 0, 1}

var windowButtonMacOSCloseBackgroundColor = gui.ColorF64{}.MakeFromRGB8(249, 100, 99)
var windowButtonMacOSHideBackgroundColor = gui.ColorF64{}.MakeFromRGB8(253, 185, 90)
var windowButtonMacOSMaximizeBackgroundColor = gui.ColorF64{}.MakeFromRGB8(68, 191, 66)
var windowButtonMacOSDisabledBackgroundColor = gui.ColorF64{}.MakeFromRGB8(210, 210, 210)
