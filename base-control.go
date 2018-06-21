// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

// HypervisorData is a bit flags for GeometryHypervisor.
// Example 1: it contains flag UPHG - "current control requests Update of self Possible Horizontal Geometry (minWidth & maxWidth)".
// Example 2: it contains flag CUPHG - "at least one Child of current control requests Update of self Possible Horizontal Geometry (minWidth & maxWidth)".
type HypervisorData uint16

const (
	HypervisorDataUPHG          HypervisorData = 1 << iota // Upgrade Possible Horizontal Geometry
	HypervisorDataUPHGRecursive HypervisorData = 1 << iota // Upgrade Possible Horizontal Geometry Recursively
	HypervisorDataCUPHG         HypervisorData = 1 << iota // Children (some of them) have Upgrade Possible Horizontal Geometry
	HypervisorDataUCHG          HypervisorData = 1 << iota // Upgrade Children Horizontal Geometry
	HypervisorDataCUCHG         HypervisorData = 1 << iota // Children (some of them) have Upgrade Children Horizontal Geometry
	HypervisorDataUPVG          HypervisorData = 1 << iota // Upgrade Possible Vertical Geometry
	HypervisorDataUPVGRecursive HypervisorData = 1 << iota // Upgrade Possible Vertical Geometry Recursively
	HypervisorDataCUPVG         HypervisorData = 1 << iota // Children (some of them) have Upgrade Possible Vertical Geometry
	HypervisorDataUCVG          HypervisorData = 1 << iota // Upgrade Children Vertical Geometry
	HypervisorDataCUCVG         HypervisorData = 1 << iota // Children (some of them) have Upgrade Children Vertical Geometry
	HypervisorDataIR            HypervisorData = 1 << iota // Invalidate Rectangle
	HypervisorDataCIR           HypervisorData = 1 << iota // Children (some of them) have Invalidate Rectangle
	HypervisorDataIG            HypervisorData = 1 << iota // Invalidate Geometry (must call OnGeometryChange)
	HypervisorDataCIG           HypervisorData = 1 << iota // Children (some of them) have Invalidate Geometry
)
const HypervisorDataNil HypervisorData = 0

type BaseControl struct {
	window         *Window
	parent         Control
	zIndex         uint // TODO implement this (currently always 0)
	minSize        PointF64
	bestSize       PointF64
	maxSize        PointF64
	geometry       RectangleF64
	hypervisorData HypervisorData
}

// SetParent is a static method.
func (BaseControl) SetParent(child, parent Control) { setParent(child, parent) }

func setParent(control, parent Control) {
	oldWindow := control.Window()
	oldParent := control.Parent()

	var newWindow *Window
	if parent != nil {
		newWindow = parent.Window()
	}

	control.setParent(parent)

	if newWindow != oldWindow {
		if newWindow != nil {
			newWindow.updateZIndex()
		}

		setWindow(control, oldWindow, newWindow)

		// Validate focus
		if oldWindow != nil && oldWindow.focusedControl.Window() != oldWindow {
			// Assume that "control" which has been moved to other parent itself has focus or contains such child.
			// Make old parent focused (or window itself if something goes wrong).
			if !oldWindow.SetFocus(oldParent) { // TODO this may causes unfocusable Control receive focus.
				oldWindow.SetFocus(oldWindow)
			}
		}
	}
}

func setWindow(control Control, oldWindow, newWindow *Window) {
	if oldWindow != nil {
		control.BeforeDetachFromWindowEvent()
	}
	control.setWindow(newWindow)
	if newWindow != nil {
		control.AfterAttachToWindowEvent()
	}
	for _, child := range control.Children() {
		setWindow(child, oldWindow, newWindow)
	}
}

func (c *BaseControl) Window() *Window { return c.window }

// Do not call this method directly - use SetParent function.
func (c *BaseControl) setWindow(window *Window) {
	c.window = window
}

func (c BaseControl) Parent() Control { return c.parent }

// Do not call this method directly - use SetParent function.
func (c *BaseControl) setParent(parent Control) {
	c.parent = parent
}

func (c *BaseControl) ZIndex() uint { return c.zIndex }
func (c *BaseControl) setZIndex(index uint) {
	c.zIndex = index
}
