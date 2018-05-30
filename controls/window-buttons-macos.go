// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
)

const (
	windowButtonsMacOSNumButtons   = 3
	windowButtonsMacOSLeftPadding  = 7
	windowButtonsMacOSHorSpace     = 5
	windowButtonsMacOSRightPadding = 7
	windowButtonsMacOSWidth        = windowButtonsMacOSLeftPadding + windowButtonsMacOSNumButtons*windowButtonMacOSSize + (windowButtonsMacOSNumButtons-1)*windowButtonsMacOSHorSpace + windowButtonsMacOSRightPadding

	windowButtonsMacOSTopPadding    = 4
	windowButtonsMacOSBottomPadding = 4
	windowButtonsMacOSHeight        = windowButtonsMacOSTopPadding + windowButtonMacOSSize + windowButtonsMacOSBottomPadding
)

type windowButtonsMacOS struct {
	gui.BaseControl
	closeButton    *windowButtonMacOS
	hideButton     *windowButtonMacOS
	maximizeButton *windowButtonMacOS
}

func (c *windowButtonsMacOS) Children() []gui.Control {
	return []gui.Control{c.closeButton, c.hideButton, c.maximizeButton}
}

func (c *windowButtonsMacOS) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	return windowButtonsMacOSWidth, windowButtonsMacOSWidth
}

func (c *windowButtonsMacOS) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	return windowButtonsMacOSHeight, windowButtonsMacOSHeight
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

func (c windowButtonsMacOS) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	c.closeButton.Draw(canvas, region)
	c.hideButton.Draw(canvas, region)
	c.maximizeButton.Draw(canvas, region)
}

func (c windowButtonsMacOS) ProcessEvent(gui.Event) bool { return false } // TODO

func newWindowButtonsMacOS() *windowButtonsMacOS {
	r := &windowButtonsMacOS{}

	closeButton := newWindowButtonMacOS(windowButtonMacOSCloseImage, windowButtonMacOSImageColor, windowButtonMacOSCloseBackgroundColor)
	hideButton := newWindowButtonMacOS(windowButtonMacOSHideImage, windowButtonMacOSImageColor, windowButtonMacOSHideBackgroundColor)
	maximizeButton := newWindowButtonMacOS(windowButtonMacOSMaximizeImage, windowButtonMacOSImageColor, windowButtonMacOSMaximizeBackgroundColor)

	gui.SetParent(closeButton, r)
	gui.SetParent(hideButton, r)
	gui.SetParent(maximizeButton, r)

	r.closeButton = closeButton
	r.hideButton = hideButton
	r.maximizeButton = maximizeButton

	return r
}
