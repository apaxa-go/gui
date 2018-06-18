// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

// Implemented by BaseControl and can't be implemented by user.
// Any Control implementation must include BaseControl.
type BaseControlI interface {
	Window() *Window
	setWindow(*Window)

	Parent() Control
	setParent(Control)

	ZIndex() uint
	setZIndex(index uint)

	MinSize() PointF64
	BestSize() PointF64
	MaxSize() PointF64

	MinWidth() float64
	BestWidth() float64
	MaxWidth() float64
	MinHeight() float64
	BestHeight() float64
	MaxHeight() float64

	GeometryHypervisorPause()
	GeometryHypervisorResume()

	Geometry() RectangleF64

	setPossibleHorGeometry(minWidth, bestWidth, maxWidth float64) (changed bool)
	setPossibleVerGeometry(minHeight, bestHeight, maxHeight float64) (changed bool)

	setHorGeometry(left, right float64) (changed bool)
	setVerGeometry(top, bottom float64) (changed bool)

	//
	// GeometryHypervisor data
	//

	getUPHG() bool
	getUPHGRecursive() bool
	setUPHG(recursive bool)
	SetUPHG(recursive bool)
	unsetUPHG()

	getCUPHG() bool
	setCUPHG()
	unsetCUPHG()

	getUCHG() bool
	setUCHG()
	SetUCHG()
	unsetUCHG()

	getCUCHG() bool
	setCUCHG()
	unsetCUCHG()

	getUPVG() bool
	getUPVGRecursive() bool
	setUPVG(recursive bool)
	SetUPVG(recursive bool)
	unsetUPVG()

	getCUPVG() bool
	setCUPVG()
	unsetCUPVG()

	getUCVG() bool
	setUCVG()
	SetUCVG()
	unsetUCVG()

	getCUCVG() bool
	setCUCVG()
	unsetCUCVG()

	getIR() bool
	setIR()
	SetIR()
	unsetIR()

	getCIR() bool
	setCIR()
	unsetCIR()

	getIG() bool
	setIG()
	unsetIG()

	getCIG() bool
	setCIG()
	unsetCIG()

	SetUPG(recursive bool)
	SetUPGIR(recursive bool)
}

// Must be implemented by user.
type Control interface {
	BaseControlI

	Children() []Control // TODO may be use different methods for different purposes (mouse button press event; geometry hypervisor)?

	ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64)
	ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64)

	ComputeChildHorGeometry() (lefts, rights []float64) // index according Children()
	ComputeChildVerGeometry() (tops, bottoms []float64) // index according Children()

	Draw(canvas Canvas, region RectangleF64)

	AfterAttachToWindowEvent()
	BeforeDetachFromWindowEvent()
	OnGeometryChangeEvent()
	OnKeyboardEvent(event KeyboardEvent) (done bool)
	OnPointerButtonEvent(event PointerButtonEvent) (processed bool)
	OnPointerDragEvent(event PointerDragEvent)
	OnPointerMoveEvent(event PointerMoveEvent) (processed bool)
	OnPointerEnterLeaveEvent(event PointerEnterLeaveEvent)
	OnScrollEvent(event ScrollEvent) (processed bool)
	OnWindowMainEvent(become bool)

	// FocusCandidate returns candidate for keyboard event focus. This method is called by Window on Tab and Shift-Tab shortcuts.
	//
	// Process limited to the Controls itself and his direct children. Implementation usually do not work with child directly, but only returns them as candidates.
	// 3 different kind of results possible.
	// Returning nil means that there is no candidate for focus.
	// Returning receiver itself means that receiver decided become a focused Control (no other checks will be performed).
	// Returning child means that receiver suggest this child as Control for focus. In this case child's FocusCandidate method will be called.
	// In any case if x.FocusCandidate(*, *) returns "y" then x.FocusCandidate must accept "y" as "current" parameter.
	//
	// If reverse is true then looking backward (looking for previous before current, Shift-Tab), otherwise looking for forward (looking for next after current, Tab).
	// Passed current is the start point for search and must be direct child of receiver itself.
	// If current is nil then looking for first (if forward) or last (if backward) candidate.
	//
	// It is uncommon (but possible) to manipulate focus in some specific way there self.FocusCandidate(*, self)==self.
	// Usually this means that Control has children implemented not by this library (e.g. embedded browser).
	//
	// Common implementations:
	// 1. Control itself does not accept focus and it has no child:
	//	func (c *SomeControl)FocusCandidate(reverse bool, current Control)(Control){
	//		return nil
	//	}
	//
	// 2. Control itself accepts focus and has no child:
	//	func (c *SomeControl)FocusCandidate(reverse bool, current Control)(Control){
	//		if current==nil{ // First/last focus
	//			return c
	//		}
	//		return nil // Next/previous (usually after/before control itself)
	//	}
	//
	// 3. Control itself accepts focus and has single child (focus order: <Control itself> before <child>):
	//	func (c *SomeControl)FocusCandidate(reverse bool, current Control)(Control){
	//		switch {
	//		//
	//		// Forward
	//		//
	//		case !reverse && current==nil: // First focus
	//			return c
	//		case !reverse && current==c: // Next after control itself
	//			return c.child
	//		case !reverse && current==c.child: // Next after child
	//			return nil
	//		//
	//		// Backward
	//		//
	//		case reverse && current==nil: // Last focus
	//			return c.child
	//		case reverse && current==c.child: // Previous before child
	//			return c
	//		case reverse && current==c: // Previous before control itself
	//			return nil
	//		//
	//		// Fallback - unexpected case
	//		//
	//		default:
	//			return c
	//		}
	//	}
	//
	// 4. Control itself does not accept focus and has single child:
	//	func (c *SomeControl)FocusCandidate(reverse bool, current Control)(Control){
	//		if current==nil{ // First/last focus
	//			return c.child
	//		}
	//		return nil // Next/previous (usually after/before child)
	//	}
	FocusCandidate(reverse bool, current Control) Control

	OnFocus(event FocusEvent)
}
