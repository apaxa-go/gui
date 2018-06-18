// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "github.com/apaxa-go/helper/mathh"

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
	zIndex         uint
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

func (c *BaseControl) ZIndex() uint { return c.zIndex }
func (c *BaseControl) setZIndex(index uint) {
	c.zIndex = index
}

func (c *BaseControl) GeometryHypervisorPause() {
	if c.window != nil {
		c.window.geometryHypervisorPause()
	}
}
func (c *BaseControl) GeometryHypervisorResume() {
	if c.window != nil {
		c.window.geometryHypervisorResume()
	}
}
func (c BaseControl) Parent() Control { return c.parent }

// Do not call this method directly - use SetParent function.
func (c *BaseControl) setParent(parent Control) {
	c.parent = parent
}

func (c BaseControl) MinSize() PointF64  { return c.minSize }
func (c BaseControl) BestSize() PointF64 { return c.bestSize }
func (c BaseControl) MaxSize() PointF64  { return c.maxSize }

func (c BaseControl) MinWidth() float64   { return c.minSize.X }
func (c BaseControl) BestWidth() float64  { return c.bestSize.X }
func (c BaseControl) MaxWidth() float64   { return c.maxSize.X }
func (c BaseControl) MinHeight() float64  { return c.minSize.Y }
func (c BaseControl) BestHeight() float64 { return c.bestSize.Y }
func (c BaseControl) MaxHeight() float64  { return c.maxSize.Y }

func (c *BaseControl) setPossibleHorGeometry(minWidth, bestWidth, maxWidth float64) (changed bool) {
	bestWidth = mathh.Max2Float64(minWidth, bestWidth)
	maxWidth = mathh.Max2Float64(bestWidth, maxWidth)
	changed = c.minSize.X != minWidth || c.bestSize.X != bestWidth || c.maxSize.X != maxWidth
	c.minSize.X = minWidth
	c.bestSize.X = bestWidth
	c.maxSize.X = maxWidth
	return
}

func (c *BaseControl) setPossibleVerGeometry(minHeight, bestHeight, maxHeight float64) (changed bool) {
	bestHeight = mathh.Max2Float64(minHeight, bestHeight)
	maxHeight = mathh.Max2Float64(bestHeight, maxHeight)
	changed = c.minSize.Y != minHeight || c.bestSize.X != bestHeight || c.maxSize.Y != maxHeight
	c.minSize.Y = minHeight
	c.bestSize.Y = bestHeight
	c.maxSize.Y = maxHeight
	return
}

func (c *BaseControl) setHorGeometry(left, right float64) (changed bool) {
	changed = c.geometry.Left != left || c.geometry.Right != right
	c.geometry.Left = left
	c.geometry.Right = right
	return
}

func (c *BaseControl) setVerGeometry(top, bottom float64) (changed bool) {
	changed = c.geometry.Top != top || c.geometry.Bottom != bottom
	c.geometry.Top = top
	c.geometry.Bottom = bottom
	return
}

func (c *BaseControl) Geometry() RectangleF64 {
	return c.geometry
}

//
//
// GeometryHypervisor related
//
//

//
// Update Possible Horizontal Geometry
//

func (c *BaseControl) getUPHG() bool          { return c.hypervisorData&HypervisorDataUPHG > 0 }
func (c *BaseControl) getUPHGRecursive() bool { return c.hypervisorData&HypervisorDataUPHGRecursive > 0 }
func (c *BaseControl) setUPHG(recursive bool) {
	c.hypervisorData |= HypervisorDataUPHG
	if recursive {
		c.hypervisorData |= HypervisorDataUPHGRecursive
	}
	for control := c.Parent(); control != nil && !control.getCUPHG(); control = control.Parent() {
		control.setCUPHG()
	}
}
func (c *BaseControl) SetUPHG(recursive bool) {
	if c.window == nil {
		return
	}
	c.setUPHG(recursive)
	c.window.geometryHypervisorRunIfActive()
}
func (c *BaseControl) unsetUPHG() {
	c.hypervisorData &= ^(HypervisorDataUPHG | HypervisorDataUPHGRecursive)
}

//
// Cache for Upgrade Possible Horizontal Geometry
//

func (c *BaseControl) getCUPHG() bool { return c.hypervisorData&HypervisorDataCUPHG > 0 }
func (c *BaseControl) setCUPHG()      { c.hypervisorData |= HypervisorDataCUPHG }
func (c *BaseControl) unsetCUPHG()    { c.hypervisorData &= ^HypervisorDataCUPHG }

//
//  Update Children's Horizontal Geometry
//

func (c *BaseControl) getUCHG() bool { return c.hypervisorData&HypervisorDataUCHG > 0 }
func (c *BaseControl) setUCHG() {
	c.hypervisorData |= HypervisorDataUCHG
	for control := c.Parent(); control != nil && !control.getCUCHG(); control = control.Parent() {
		control.setCUCHG()
	}
}
func (c *BaseControl) SetUCHG() {
	if c.window == nil {
		return
	}
	c.setUCHG()
	c.window.geometryHypervisorRunIfActive()
}
func (c *BaseControl) unsetUCHG() { c.hypervisorData |= HypervisorDataUCHG }

//
// Cache for Update Children's Horizontal Geometry
//

func (c *BaseControl) getCUCHG() bool { return c.hypervisorData&HypervisorDataCUCHG > 0 }
func (c *BaseControl) setCUCHG()      { c.hypervisorData |= HypervisorDataCUCHG }
func (c *BaseControl) unsetCUCHG()    { c.hypervisorData &= ^HypervisorDataCUCHG }

//
// Update Possible Vertical Geometry
//

func (c *BaseControl) getUPVG() bool          { return c.hypervisorData&HypervisorDataUPVG > 0 }
func (c *BaseControl) getUPVGRecursive() bool { return c.hypervisorData&HypervisorDataUPVGRecursive > 0 }
func (c *BaseControl) setUPVG(recursive bool) {
	c.hypervisorData |= HypervisorDataUPVG
	if recursive {
		c.hypervisorData |= HypervisorDataUPVGRecursive
	}
	for control := c.Parent(); control != nil && !control.getCUPVG(); control = control.Parent() {
		control.setCUPVG()
	}
}
func (c *BaseControl) SetUPVG(recursive bool) {
	if c.window == nil {
		return
	}
	c.setUPVG(recursive)
	c.window.geometryHypervisorRunIfActive()
}
func (c *BaseControl) unsetUPVG() {
	c.hypervisorData &= ^(HypervisorDataUPVG | HypervisorDataUPVGRecursive)
}

//
// Cache for Upgrade Possible Vertical Geometry
//

func (c *BaseControl) getCUPVG() bool { return c.hypervisorData&HypervisorDataCUPVG > 0 }
func (c *BaseControl) setCUPVG()      { c.hypervisorData |= HypervisorDataCUPVG }
func (c *BaseControl) unsetCUPVG()    { c.hypervisorData &= ^HypervisorDataCUPVG }

//
//  Update Children's Vertical Geometry
//

func (c *BaseControl) getUCVG() bool { return c.hypervisorData&HypervisorDataUCVG > 0 }
func (c *BaseControl) setUCVG() {
	c.hypervisorData |= HypervisorDataUCVG
	for control := c.Parent(); control != nil && !control.getCUCVG(); control = control.Parent() {
		control.setCUCVG()
	}
}
func (c *BaseControl) SetUCVG() {
	if c.window == nil {
		return
	}
	c.setUCVG()
	c.window.geometryHypervisorRunIfActive()
}
func (c *BaseControl) unsetUCVG() { c.hypervisorData |= HypervisorDataUCVG }

//
// Cache for Update Children's Vertical Geometry
//

func (c *BaseControl) getCUCVG() bool { return c.hypervisorData&HypervisorDataCUCVG > 0 }
func (c *BaseControl) setCUCVG()      { c.hypervisorData |= HypervisorDataCUCVG }
func (c *BaseControl) unsetCUCVG()    { c.hypervisorData &= ^HypervisorDataCUCVG }

//
// Invalidate Region
//

func (c *BaseControl) getIR() bool { return c.hypervisorData&HypervisorDataIR > 0 }
func (c *BaseControl) setIR() {
	c.hypervisorData |= HypervisorDataIR
	for control := c.Parent(); control != nil && !control.getCIR(); control = control.Parent() {
		control.setCIR()
	}
}
func (c *BaseControl) SetIR() {
	if c.window == nil {
		return
	}
	c.setIR()
	c.window.geometryHypervisorRunIfActive()
}
func (c *BaseControl) unsetIR() { c.hypervisorData &= ^HypervisorDataIR }

//
// Cache for Invalidate Region
//

func (c *BaseControl) getCIR() bool { return c.hypervisorData&HypervisorDataCIR > 0 }
func (c *BaseControl) setCIR()      { c.hypervisorData |= HypervisorDataCIR }
func (c *BaseControl) unsetCIR()    { c.hypervisorData &= ^HypervisorDataCIR }

//
// Invalidate Geometry
//

func (c *BaseControl) getIG() bool { return c.hypervisorData&HypervisorDataIG > 0 }
func (c *BaseControl) setIG() {
	c.hypervisorData |= HypervisorDataIG
	for control := c.Parent(); control != nil && !control.getCIG(); control = control.Parent() {
		control.setCIG()
	}
}
func (c *BaseControl) unsetIG() { c.hypervisorData &= ^HypervisorDataIG }

//
// Cache for Invalidate Geometry
//

func (c *BaseControl) getCIG() bool { return c.hypervisorData&HypervisorDataCIG > 0 }
func (c *BaseControl) setCIG()      { c.hypervisorData |= HypervisorDataCIG }
func (c *BaseControl) unsetCIG()    { c.hypervisorData &= ^HypervisorDataCIG }

//
// Shortcuts
//

// SetUPHG & SetUPVG
func (c *BaseControl) SetUPG(recursive bool) {
	c.GeometryHypervisorPause()
	c.setUPHG(recursive)
	c.setUPVG(recursive)
	c.GeometryHypervisorResume()
}

// SetUPHG, SetUPVG & SetIR
func (c *BaseControl) SetUPGIR(recursive bool) {
	c.GeometryHypervisorPause()
	c.setUPHG(recursive)
	c.setUPVG(recursive)
	c.setIR()
	c.GeometryHypervisorResume()
}

// SetUCHG() & SetUCVG()
func (c *BaseControl) SetUCG() {
	c.GeometryHypervisorPause()
	c.setUCHG()
	c.setUCVG()
	c.GeometryHypervisorResume()
}

// SetUCHG(), SetUCVG() & SetIR()
func (c *BaseControl) SetUCGIR() {
	c.GeometryHypervisorPause()
	c.setUCHG()
	c.setUCVG()
	c.setIR()
	c.GeometryHypervisorResume()
}

//
//
// Default implementations for method.
//
//

func (BaseControl) Children() []Control                                { return nil }
func (BaseControl) ComputeChildHorGeometry() (lefts, rights []float64) { return nil, nil }
func (BaseControl) ComputeChildVerGeometry() (tops, bottoms []float64) { return nil, nil }

//
// Default event handlers & related methods - controls may reimplement them.
//

func (BaseControl) AfterAttachToWindowEvent()                                  {}
func (BaseControl) BeforeDetachFromWindowEvent()                               {}
func (BaseControl) OnGeometryChangeEvent()                                     {}
func (BaseControl) OnKeyboardEvent(_ KeyboardEvent) (done bool)                { return false }
func (BaseControl) OnPointerButtonEvent(_ PointerButtonEvent) (processed bool) { return false }
func (BaseControl) OnPointerDragEvent(_ PointerDragEvent)                      {}
func (BaseControl) OnPointerMoveEvent(_ PointerMoveEvent) (processed bool)     { return false }
func (BaseControl) OnScrollEvent(_ ScrollEvent) (processed bool)               { return false }
func (BaseControl) OnPointerEnterLeaveEvent(_ PointerEnterLeaveEvent)          {}
func (BaseControl) OnFocus(_ FocusEvent)                                       {}
func (BaseControl) OnWindowMainEvent(_ bool)                                   {}

//
// Related to event handlers
//

// FocusCandidate is default implementation. It always returns nil - neither Control itself nor his children (if any) accepts focus.
func (BaseControl) FocusCandidate(reverse bool, current Control) Control { return nil }
