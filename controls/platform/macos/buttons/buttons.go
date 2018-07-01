// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package buttons

import "github.com/apaxa-go/helper/mathh"

const (
	numButtons   = 3
	leftPadding  = 8
	horSpace     = 8
	rightPadding = 8
	width        = leftPadding + numButtons*size + (numButtons-1)*horSpace + rightPadding

	topPadding    = 5
	bottomPadding = 5
	height        = topPadding + size + bottomPadding
)

type Buttons struct {
	BaseControl
	closeButton    *CloseButton
	minimizeButton *MinimizeButton
	maximizeButton *MaximizeButton
	areaID         EnterLeaveAreaID
}

func (c *Buttons) Children() []Control {
	return []Control{c.closeButton, c.minimizeButton, c.maximizeButton}
}

func (c *Buttons) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return width, width, mathh.PositiveInfFloat64() //width
}

func (c *Buttons) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return height, height, height
}

func (c *Buttons) ComputeChildHorGeometry() (lefts, rights []float64) {
	const (
		left0 = leftPadding
		left1 = left0 + size + horSpace
		left2 = left1 + size + horSpace

		right0 = left0 + size
		right1 = left1 + size
		right2 = left2 + size
	)
	left := c.Geometry().Left
	return []float64{left + left0, left + left1, left + left2}, []float64{left + right0, left + right1, left + right2}
}

func (c *Buttons) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top := c.Geometry().Top + topPadding
	bottom := top + size

	return []float64{top, top, top}, []float64{bottom, bottom, bottom}
}

func (c Buttons) Draw(canvas Canvas, region RectangleF64) {
	c.closeButton.Draw(canvas, region)
	c.minimizeButton.Draw(canvas, region)
	c.maximizeButton.Draw(canvas, region)
}

func (c *Buttons) AfterAttachToWindowEvent() {
	c.areaID = c.Window().AddEnterLeaveArea(c, RectangleF64{}) // Reserve TrackingAreaID.
	c.Window().SubscribeToWindowMainStateEvent(c)
}

func (c *Buttons) BeforeDetachFromWindowEvent() {
	c.Window().UnsubscribeFromWindowMainStateEvent(c)
	c.Window().RemoveEnterLeaveArea(c.areaID, false) // Free TrackingArea.
}

func (c *Buttons) OnGeometryChangeEvent() {
	// Update TrackingArea.
	origin := c.Geometry().LT()
	rect := RectangleF64{origin.X, origin.Y, origin.X + width, origin.Y + height}
	c.Window().ReplaceEnterLeaveArea(c.areaID, rect)
}

func (c *Buttons) OnPointerEnterLeaveEvent(event PointerEnterLeaveEvent) {
	c.closeButton.hover = event.Enter
	c.minimizeButton.hover = event.Enter
	c.maximizeButton.hover = event.Enter
	c.SetIR()
}

func (c *Buttons) OnWindowMainStateEvent(event WindowMainStateEvent) {
	c.SetIR()
}

func NewButtons() *Buttons {
	r := &Buttons{}

	closeButton := NewCloseButton()
	minimizeButton := NewMinimizeButton()
	maximizeButton := NewMaximizeButton()

	r.BaseControl.SetParent(closeButton, r)
	r.BaseControl.SetParent(minimizeButton, r)
	r.BaseControl.SetParent(maximizeButton, r)

	r.closeButton = closeButton
	r.minimizeButton = minimizeButton
	r.maximizeButton = maximizeButton

	return r
}
