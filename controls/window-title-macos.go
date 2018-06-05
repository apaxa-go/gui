// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
)

const (
	windowTitleMacOSHeight     = 16
	windowTitleMacOSVerPadding = 1
)

//var windowTitleMacOSBackground = gui.ColorF64{}.MakeFromRGB8(230,230,230)
var windowTitleMacOSTitleColor = gui.ColorF64{0, 0, 0, 1}

type windowTitleMacOS struct {
	gui.BaseControl
	label *Label
}

func (c *windowTitleMacOS) Children() []gui.Control { return []gui.Control{c.label} }

func (c *windowTitleMacOS) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	return c.label.MinWidth(), gui.PosInfF64()
}

func (c *windowTitleMacOS) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	height := gui.Max2Float64(windowTitleMacOSHeight, c.label.MinHeight()+2*windowTitleMacOSVerPadding)
	return height, height
}

func (c *windowTitleMacOS) ComputeChildHorGeometry() (lefts, rights []float64) {
	left, right := basetypes.AlignHorCenter.ApplyF64(c.Geometry().Left, c.Geometry().Right, c.label.MinWidth())
	return []float64{left}, []float64{right}
}

func (c *windowTitleMacOS) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top, bottom := basetypes.AlignVerTop.ApplyF64(c.Geometry().Top, c.Geometry().Bottom, c.label.MinHeight())
	top += windowTitleMacOSVerPadding
	bottom += windowTitleMacOSVerPadding
	return []float64{top}, []float64{bottom}
}

func (c windowTitleMacOS) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	// TODO do we need to draw background here???
	c.label.Draw(canvas, region)
}

func newWindowTitleMacOS() *windowTitleMacOS {
	r := &windowTitleMacOS{}
	label := NewLabel("", defaultFont, windowTitleMacOSTitleColor)
	gui.SetParent(label, r)
	r.label = label
	return r
}
