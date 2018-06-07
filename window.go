// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "github.com/apaxa-go/gui/drivers"

type Window struct {
	driverWindow DriverWindow
	BaseControl
	child                   Control
	geometryHypervisorState uint // 0 means hypervisor is online (performs request immediately), otherwise it is paused geometryHypervisorState times.
	focusedControl          Control
}

//
// Unique methods
//

/*func (w *Window) Run() {
	w.driverWindow.Run()
}*/

func (w *Window) Title() string { return w.driverWindow.Title() }
func (w *Window) SetTitle(title string) {
	w.driverWindow.SetTitle(title)
}

func (w *Window) Child() Control { return w.child }
func (w *Window) SetChild(child Control) {
	if w.child != nil {
		w.BaseControl.SetParent(w.child, nil)
	}
	w.child = child
	w.BaseControl.SetParent(w.child, w)
	w.SetUPG(true)
}

func (w *Window) adjustSize() {
	reqSize := w.Geometry().GetSize()
	if w.driverWindow.Size() != reqSize {
		w.driverWindow.SetSize(reqSize)
	}
}

func (w *Window) invalidateRegion(region RectangleF64) {
	w.driverWindow.InvalidateRegion(region)
}

func (w *Window) invalidate() {
	w.driverWindow.Invalidate()
}

func (w *Window) onExternalResize() {
	w.SetUCGIR() // TODO we need some method to avoid invalid (according to Min/MaxSize) external resize.
}

func (w *Window) onOfflineCanvasChanged() {
	w.SetUPG(true)
}

func (w *Window) OfflineCanvas() OfflineCanvas { return w.driverWindow.OfflineCanvas() }

//
// BaseControlI overrides
//

func (w *Window) setPossibleHorGeometry(minWidth, maxWidth float64) (changed bool) {
	changed = w.BaseControl.setPossibleHorGeometry(minWidth, maxWidth)
	if !changed {
		return
	}
	if w.Geometry().Width() < w.MinWidth() {
		w.setHorGeometry(0, w.MinWidth())
		w.setUCHG()
	} else if w.Geometry().Width() > w.MaxWidth() {
		w.setHorGeometry(0, w.MaxWidth())
		w.setUCHG()
	}
	return
}

func (w *Window) setPossibleVerGeometry(minHeight, maxHeight float64) (changed bool) {
	changed = w.BaseControl.setPossibleVerGeometry(minHeight, maxHeight)
	if !changed {
		return
	}
	if w.Geometry().Height() < w.MinHeight() {
		w.setVerGeometry(0, w.MinHeight())
		w.setUCVG()
	} else if w.Geometry().Width() > w.MaxHeight() {
		w.setVerGeometry(0, w.MaxHeight())
		w.setUCVG()
	}
	return
}

// SetParent method does nothing for Window.
func (w *Window) SetParent(parent Control) {}

//
// Control interface implementation
//

func (w *Window) Children() []Control {
	if w.child == nil {
		return nil
	}
	return []Control{w.child}
}

func (w *Window) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	if w.child == nil {
		return 100, 100
	}
	return w.child.MinWidth(), w.child.MaxWidth()
}

func (w *Window) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	if w.child == nil {
		return 100, 100
	}
	return w.child.MinHeight(), w.child.MaxHeight()
}

func (w *Window) ComputeChildHorGeometry() (lefts, rights []float64) {
	if w.child == nil {
		return nil, nil
	}
	return []float64{0}, []float64{w.Geometry().Width()}
}
func (w *Window) ComputeChildVerGeometry() (tops, bottoms []float64) {
	if w.child == nil {
		return nil, nil
	}
	return []float64{0}, []float64{w.Geometry().Height()}
}

func (w *Window) Draw(canvas Canvas, region RectangleF64) {
	if w.child != nil {
		w.child.Draw(canvas, region)
	}
}

//
// Raw events handlers (process events directly from driver).
//

func (w *Window) processHotKeys(e KeyboardEvent) (processed bool) {
	switch {
	case e.Event.IsPressed() && e.Key == drivers.KeyTab && (e.Modifiers & ^drivers.KeyModifierShift == 0): // TODO use smth like e.Modifiers.OnlyShift()
		w.ShiftFocus(e.Modifiers.IsShiftPressed())
	default:
		return false
	}
	return true
}

func (w *Window) onKeyboardEvent(e KeyboardEvent) {
	if w.processHotKeys(e) {
		return
	}
	for c, done := w.focusedControl, false; !done && c != nil; c = c.Parent() {
		done = c.OnKeyboardEvent(e)
	}
}

func processPointerButtonEvent(c Control, e PointerButtonEvent) (processed bool) {
	if !c.Geometry().Contains(e.Point) {
		return false
	}
	for _, candidate := range c.Children() {
		processed = processPointerButtonEvent(candidate, e)
		if processed {
			return
		}
	}
	return c.OnPointerButtonEvent(e)
}

func (w *Window) onPointerKey(e PointerButtonEvent) {
	processPointerButtonEvent(w, e)
}

func processScrollEvent(c Control, e ScrollEvent) (processed bool) {
	if !c.Geometry().Contains(e.Point) {
		return false
	}
	for _, candidate := range c.Children() {
		processed = processScrollEvent(candidate, e)
		if processed {
			return
		}
	}
	return c.OnScrollEvent(e)
}

func (w *Window) onScroll(e ScrollEvent) {
	processScrollEvent(w, e)
}

//
// Events related.
//

func (w *Window) FocusedControl() Control         { return w.focusedControl }
func (w *Window) IfControlFocused(c Control) bool { return w.focusedControl == c }

//
// Constructors & destructor
//

func (w *Window) baseInit() {
	w.focusedControl = w
	w.driverWindow.RegisterDrawCallback(w.Draw)
	w.driverWindow.RegisterResizeCallback(w.onExternalResize)
	w.driverWindow.RegisterOfflineCanvasCallback(w.onOfflineCanvasChanged)
	w.driverWindow.RegisterKeyboardCallback(w.onKeyboardEvent)
	w.driverWindow.RegisterPointerKeyCallback(w.onPointerKey)
	w.driverWindow.RegisterScrollCallback(w.onScroll)
	w.BaseControl.window = w
	w.SetUPGIR(false)
}

func NewWindow() *Window {
	var w Window
	w.driverWindow = driverWindowConstructor()
	w.baseInit()
	return &w
}

func NewWindowAdvanced(dw DriverWindow) *Window {
	var w Window
	w.driverWindow = dw
	w.GeometryHypervisorPause()
	w.baseInit()
	w.GeometryHypervisorResume()
	return &w
}

func (w *Window) Destroy() {
	w.GeometryHypervisorPause()
	w.SetChild(nil)
	w.driverWindow.Destroy()
}
