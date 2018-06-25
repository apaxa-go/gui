// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "C"
import (
	"github.com/apaxa-go/gui/drivers"
	"github.com/apaxa-go/helper/mathh"
	"log"
)

// TODO implement best size logic for window.

type Window struct {
	driverWindow DriverWindow
	BaseControl
	child                     Control
	geometryHypervisorState   int // <0 means active, 0 means hypervisor is online (performs request immediately), otherwise it is paused geometryHypervisorState times.
	focusedControl            Control
	pointerPressControl       Control
	isMain                    bool
	enterLeaveAreas           map[EnterLeaveAreaID]enterLeaveArea // Lookup map to identify receiver by area id.
	overlappedEnterLeaveAreas []overlappedEnterLeaveArea
	nextEnterLeaveAreaID      EnterLeaveAreaID
	moveAreas                 map[MoveAreaID]Control // Lookup map to identify receiver by area id.
	nextMoveAreaID            MoveAreaID             // Candidate for next id.
}

//
// Unique methods
//

/*func (w *Window) Run() {
	w.driverWindow.Run()
}*/

func (w *Window) IsMain() bool {
	return w.isMain
}

func (w *Window) Title() string { return w.driverWindow.Title() }
func (w *Window) SetTitle(title string) {
	w.driverWindow.SetTitle(title)
}

func (w *Window) WindowPos() PointF64 { // TODO may be use BaseControl's geometry field for this? Not sure, but looks ugly.
	return w.driverWindow.Pos()
}

func (w *Window) WindowSize() PointF64 { // TODO why not use BaseControl's geometry field for this?
	return w.driverWindow.Size()
}

func (w *Window) WindowGeometry() RectangleF64 {
	return w.driverWindow.Geometry()
}

func (w *Window) SetWindowGeometry(geometry RectangleF64) {
	w.driverWindow.SetGeometry(geometry)
}

func (w *Window) SetWindowPos(pos PointF64) {
	w.driverWindow.SetPos(pos)
}
func (w *Window) SetWindowSize(size PointF64) {
	w.driverWindow.SetSize(size)
}

func (w *Window) Minimize() { w.driverWindow.Minimize() }
func (w *Window) Maximize() { w.driverWindow.Maximize() }

func (w *Window) Child() Control { return w.child }
func (w *Window) SetChild(child Control) {
	if w.child != nil {
		w.BaseControl.SetParent(w.child, nil)
	}
	w.child = child
	if w.child != nil {
		w.BaseControl.SetParent(w.child, w)
	}
	w.SetUPG(true)
}

func updateZIndex(c Control, zIndex uint) (maxZIndex uint) {
	c.setZIndex(zIndex)
	if c.SequentialZIndex() {
		for _, child := range c.Children() {
			zIndex = updateZIndex(child, zIndex+1)
		}
		return zIndex
	}
	for _, child := range c.Children() {
		tmp := updateZIndex(child, zIndex+1)
		maxZIndex = mathh.Max2Uint(maxZIndex, tmp)
	}
	return
}

func (w *Window) updateZIndex() { // TODO update ZIndex on each element adding is inefficient
	updateZIndex(w, 0)
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

func (w *Window) onResize(size PointF64) {
	w.GeometryHypervisorPause()
	if w.setHorGeometry(0, size.X) {
		w.setUCHG()
		w.setIR()
	}
	if w.setVerGeometry(0, size.Y) {
		w.setUCVG()
		w.setIR()
	}
	w.GeometryHypervisorResume()
}

func (w *Window) onOfflineCanvasChanged() {
	w.SetUPG(true)
}

func (w *Window) OfflineCanvas() OfflineCanvas { return w.driverWindow.OfflineCanvas() }

func (w *Window) SetCursor(cursor Cursor) {
	w.driverWindow.SetCursor(cursor)
}

//
// BaseControlI overrides
//

func (w *Window) setPossibleHorGeometry(minWidth, bestWidth, maxWidth float64) (changed bool) {
	changed = w.BaseControl.setPossibleHorGeometry(minWidth, bestWidth, maxWidth)
	if changed {
		w.driverWindow.SetPossibleSize(w.minSize, w.maxSize)
	}
	return
}

func (w *Window) setPossibleVerGeometry(minHeight, bestHeight, maxHeight float64) (changed bool) {
	changed = w.BaseControl.setPossibleVerGeometry(minHeight, bestHeight, maxHeight)
	if changed {
		w.driverWindow.SetPossibleSize(w.minSize, w.maxSize)
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

func (w *Window) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	if w.child == nil {
		return 100, 100, 100
	}
	return w.child.MinWidth(), w.child.BestWidth(), w.child.MaxWidth()
}

func (w *Window) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	if w.child == nil {
		return 100, 100, 100
	}
	return w.child.MinHeight(), w.child.BestHeight(), w.child.MaxHeight()
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

var count int

func (w *Window) Draw(canvas Canvas, region RectangleF64) {
	log.Println("X5.1")
	if w.child == nil {
		log.Println("X5.7")
		return
	}
	if !w.geometryHypervisorIsReady() {
		w.setIR()
		log.Println("X5.8")
		return
	}
	log.Println(count)
	count++
	w.child.Draw(canvas, region)
	log.Println("X5.9")
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
	//for _, candidate := range c.Children() {
	children := c.Children()
	for i := len(children) - 1; i >= 0; i-- { // we reverse order because of Layer Control
		child := children[i]
		processed = processPointerButtonEvent(child, e)
		if processed {
			return
		}
	}
	processed = c.OnPointerButtonEvent(e)
	if processed && e.Kind.IsPress() {
		// TODO not accurate - what if multiple buttons is pressed?
		c.Window().pointerPressControl = c
	}
	return
}

func (w *Window) onPointerKey(e PointerButtonEvent) {
	if e.Kind.IsRelease() {
		w.pointerPressControl.OnPointerButtonEvent(e)
		return
	}
	processed := processPointerButtonEvent(w, e)
	if !processed && e.Kind.IsPress() {
		// TODO not accurate - what if multiple buttons is pressed?
		w.pointerPressControl = w
	}
}

func processScrollEvent(c Control, e ScrollEvent) (processed bool) {
	if !c.Geometry().Contains(e.Point) {
		return false
	}
	//for _, candidate := range c.Children() {
	children := c.Children()
	for i := len(children) - 1; i >= 0; i-- { // we reverse order because of Layer Control
		child := children[i]
		processed = processScrollEvent(child, e)
		if processed {
			return
		}
	}
	return c.OnScrollEvent(e)
}

func (w *Window) onScroll(e ScrollEvent) {
	processScrollEvent(w, e)
}

func (w *Window) onPointerDrag(e PointerDragEvent) {
	w.pointerPressControl.OnPointerDragEvent(e)
}

func processWindowMainEvent(c Control, become bool) {
	c.OnWindowMainEvent(become)
	for _, child := range c.Children() {
		processWindowMainEvent(child, become)
	}
}

func (w *Window) onWindowMainEvent(become bool) {
	w.isMain = become
	processWindowMainEvent(w, become)
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
	w.enterLeaveAreas = make(map[EnterLeaveAreaID]enterLeaveArea)
	w.moveAreas = make(map[MoveAreaID]Control)
	w.focusedControl = w
	w.driverWindow.RegisterDrawCallback(w.Draw)
	w.driverWindow.RegisterResizeCallback(w.onResize)
	w.driverWindow.RegisterOfflineCanvasCallback(w.onOfflineCanvasChanged)
	w.driverWindow.RegisterKeyboardCallback(w.onKeyboardEvent)
	w.driverWindow.RegisterPointerKeyCallback(w.onPointerKey)
	w.driverWindow.RegisterPointerDragCallback(w.onPointerDrag)
	w.driverWindow.RegisterPointerMoveCallback(w.onPointerMove)
	w.driverWindow.RegisterPointerEnterLeaveCallback(w.onPointerEnterLeave)
	w.driverWindow.RegisterScrollCallback(w.onScroll)
	w.driverWindow.RegisterWindowMainCallback(w.onWindowMainEvent)
	w.BaseControl.window = w
	w.SetUPGIR(false)
}

func NewWindow(title string) *Window {
	var w Window
	w.driverWindow = driverWindowConstructor(title)
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

func (w *Window) Close() {
	w.GeometryHypervisorPause()
	w.SetChild(nil)
	w.driverWindow.Close()
	Stop() // TODO What if there are more than 1 window?
}
