// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

const (
	windowTitleMacOSHeight     = 16
	windowTitleMacOSVerPadding = 1
)

//var windowTitleMacOSBackground = ColorF64{}.MakeFromRGB8(230,230,230)
var windowTitleMacOSTitleColor = ColorF64{0, 0, 0, 1}

type windowTitleMacOS struct {
	BaseControl
	label *Label
	//initPointerPos PointF64
	//initWindowPos  PointF64
}

func (c *windowTitleMacOS) Children() []Control { return []Control{c.label} }

func (c *windowTitleMacOS) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return c.label.MinWidth(), c.label.MinWidth(), mathh.PositiveInfFloat64() // TODO implement WithOutBestWidth
}

func (c *windowTitleMacOS) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	height := mathh.Max2Float64(windowTitleMacOSHeight, c.label.MinHeight()+2*windowTitleMacOSVerPadding)
	return height, height, height
}

func (c *windowTitleMacOS) ComputeChildHorGeometry() (lefts, rights []float64) {
	left, right := AlignHor(0).MakeCenter().ApplyF64(c.Geometry().Left, c.Geometry().Right, c.label.MinWidth())
	return []float64{left}, []float64{right}
}

func (c *windowTitleMacOS) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top, bottom := AlignVer(0).MakeTop().ApplyF64(c.Geometry().Top, c.Geometry().Bottom, c.label.MinHeight())
	top += windowTitleMacOSVerPadding
	bottom += windowTitleMacOSVerPadding
	return []float64{top}, []float64{bottom}
}

func (c windowTitleMacOS) Draw(canvas Canvas, region RectangleF64) {
	// TODO do we need to draw background here???
	c.label.Draw(canvas, region)
}

func (c *windowTitleMacOS) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	return true
	/*if event.Kind.IsPress() && event.Button.IsLeft() {
			c.initPointerPos = event.Point
			c.initWindowPos = c.Window().Pos()
			processed=true
			{
				//log.Printf("SP: %v\tSW: %v\n",c.initPointerPos.String(),c.initWindowPos.String())
			}
	}
	return*/
}

func (c *windowTitleMacOS) OnPointerDragEvent(event PointerDragEvent) {
	//pos :=c.initWindowPos.Sub(c.initPointerPos).Add(event.Point)
	pos := c.Window().Pos().Add(event.Delta)
	c.Window().SetPos(pos)
	{
		//delta:=event.Point.Sub(c.initPointerPos)
		//log.Printf(" P: %v\t W: %v\tDelta: %v\tNow: %v\n",event.Point,pos.String(),delta,c.Window().Pos())
	}
}

func newWindowTitleMacOS() *windowTitleMacOS {
	r := &windowTitleMacOS{}
	label := NewLabel("", defaultFont, windowTitleMacOSTitleColor)
	r.BaseControl.SetParent(label, r)
	r.label = label
	return r
}
