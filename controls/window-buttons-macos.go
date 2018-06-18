// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

const (
	windowButtonsMacOSNumButtons   = 3
	windowButtonsMacOSLeftPadding  = 8
	windowButtonsMacOSHorSpace     = 8
	windowButtonsMacOSRightPadding = 8
	windowButtonsMacOSWidth        = windowButtonsMacOSLeftPadding + windowButtonsMacOSNumButtons*windowButtonMacOSSize + (windowButtonsMacOSNumButtons-1)*windowButtonsMacOSHorSpace + windowButtonsMacOSRightPadding

	windowButtonsMacOSTopPadding    = 5
	windowButtonsMacOSBottomPadding = 5
	windowButtonsMacOSHeight        = windowButtonsMacOSTopPadding + windowButtonMacOSSize + windowButtonsMacOSBottomPadding
)

type windowButtonsMacOS struct {
	BaseControl
	closeButton    *windowButtonMacOS
	minimizeButton *windowButtonMacOS
	maximizeButton *windowButtonMacOS
	areaID         EnterLeaveAreaID
}

func (c *windowButtonsMacOS) Children() []Control {
	return []Control{c.closeButton, c.minimizeButton, c.maximizeButton}
}

func (c *windowButtonsMacOS) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return windowButtonsMacOSWidth, windowButtonsMacOSWidth, windowButtonsMacOSWidth
}

func (c *windowButtonsMacOS) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return windowButtonsMacOSHeight, windowButtonsMacOSHeight, windowButtonsMacOSHeight
}

func (c *windowButtonsMacOS) ComputeChildHorGeometry() (lefts, rights []float64) {
	const (
		left0 = windowButtonsMacOSLeftPadding
		left1 = left0 + windowButtonMacOSSize + windowButtonsMacOSHorSpace
		left2 = left1 + windowButtonMacOSSize + windowButtonsMacOSHorSpace

		right0 = left0 + windowButtonMacOSSize
		right1 = left1 + windowButtonMacOSSize
		right2 = left2 + windowButtonMacOSSize
	)
	left := c.Geometry().Left
	return []float64{left + left0, left + left1, left + left2}, []float64{left + right0, left + right1, left + right2}
}

func (c *windowButtonsMacOS) ComputeChildVerGeometry() (tops, bottoms []float64) {
	top := c.Geometry().Top + windowButtonsMacOSTopPadding
	bottom := top + windowButtonMacOSSize

	return []float64{top, top, top}, []float64{bottom, bottom, bottom}
}

func (c windowButtonsMacOS) Draw(canvas Canvas, region RectangleF64) {
	c.closeButton.Draw(canvas, region)
	c.minimizeButton.Draw(canvas, region)
	c.maximizeButton.Draw(canvas, region)
}

func (c *windowButtonsMacOS) AfterAttachToWindowEvent() {
	// Reserve TrackingAreaID.
	c.areaID = c.Window().AddEnterLeaveArea(c, RectangleF64{})
}

func (c *windowButtonsMacOS) BeforeDetachFromWindowEvent() {
	// Free TrackingArea.
	c.Window().RemoveEnterLeaveArea(c.areaID, false)
}

func (c *windowButtonsMacOS) OnGeometryChangeEvent() {
	// Update TrackingArea.
	origin := c.Geometry().LT()
	rect := RectangleF64{origin.X, origin.Y, origin.X + windowButtonsMacOSWidth, origin.Y + windowButtonsMacOSHeight}
	c.Window().ReplaceEnterLeaveArea(c.areaID, rect)
}

func (c *windowButtonsMacOS) OnPointerEnterLeaveEvent(event PointerEnterLeaveEvent) {
	c.closeButton.hover = event.Enter
	c.minimizeButton.hover = event.Enter
	c.maximizeButton.hover = event.Enter
	c.SetIR()
}

func (c *windowButtonsMacOS) OnWindowMainEvent(become bool) {
	c.SetIR()
}

func newWindowButtonsMacOS() *windowButtonsMacOS {
	r := &windowButtonsMacOS{}

	closeButton := newWindowButtonMacOSClose()
	minimizeButton := newWindowButtonMacOSMinimize()
	maximizeButton := newWindowButtonMacOSMaximize()

	r.BaseControl.SetParent(closeButton, r)
	r.BaseControl.SetParent(minimizeButton, r)
	r.BaseControl.SetParent(maximizeButton, r)

	r.closeButton = closeButton
	r.minimizeButton = minimizeButton
	r.maximizeButton = maximizeButton

	return r
}
