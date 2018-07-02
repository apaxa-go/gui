// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "C"
import (
	"github.com/apaxa-go/gui/drivers"
	"github.com/apaxa-go/helper/mathh"
)

// TODO implement best size logic for window.

type geometryRequest struct {
	size            float64
	enabled         bool
	invertFixedSide bool
	//posShift float64
}

type ModifiersEventSubscriber interface {
	OnModifiersEvent(ModifiersEvent)
}

type WindowDisplayStateEventSubscriber interface {
	OnWindowDisplayStateEvent(WindowDisplayStateEvent)
}

type WindowMainStateEventSubscriber interface {
	OnWindowMainStateEvent(WindowMainStateEvent)
}

type Window struct {
	driver DriverWindow
	BaseControl
	child                           Control
	geometryHypervisorState         int // <0 means active, 0 means hypervisor is online (performs request immediately), otherwise it is paused geometryHypervisorState times.
	geometryHypervisorWidthRequest  geometryRequest
	geometryHypervisorHeightRequest geometryRequest

	focusedControl      Control
	pointerPressControl Control

	enterLeaveAreas           map[EnterLeaveAreaID]enterLeaveArea // Lookup map to identify receiver by area id.
	overlappedEnterLeaveAreas []overlappedEnterLeaveArea
	nextEnterLeaveAreaID      EnterLeaveAreaID
	moveAreas                 map[MoveAreaID]Control // Lookup map to identify receiver by area id.
	nextMoveAreaID            MoveAreaID             // Candidate for next id.

	modifiersEventSubscribers          []ModifiersEventSubscriber
	windowDisplayStateEventSubscribers []WindowDisplayStateEventSubscriber
	windowMainStateEventSubscribers    []WindowMainStateEventSubscriber

	cursor         Cursor // Current non-override cursor.
	cursorOverride bool   // Is cursor currently overrided.

	minimizeAllowed   bool
	maximizeAllowed   bool
	fullScreenAllowed bool
	closeAllowed      bool
}

//
// Unique methods
//

/*func (w *Window) Run() {
	w.driver.Run()
}*/

func (w *Window) Driver() DriverWindow { return w.driver }

func (w *Window) IsMain() bool {
	return w.driver.IsMain()
}

func (w *Window) Title() string { return w.driver.Title() }
func (w *Window) SetTitle(title string) {
	w.driver.SetTitle(title)
}

func (w *Window) SetWindowWidth(width float64, fixedRight bool) {
	w.geometryHypervisorWidthRequest.enabled = true
	w.geometryHypervisorWidthRequest.size = width
	w.geometryHypervisorWidthRequest.invertFixedSide = fixedRight
	w.geometryHypervisorRunIfReady()
}

func (w *Window) SetWindowHeight(height float64, fixedBottom bool) {
	w.geometryHypervisorHeightRequest.enabled = true
	w.geometryHypervisorHeightRequest.size = height
	w.geometryHypervisorHeightRequest.invertFixedSide = fixedBottom
	w.geometryHypervisorRunIfReady()
}

func (w *Window) SetWindowSize(size PointF64, fixedRight, fixedBottom bool) {
	w.GeometryHypervisorPause()
	w.SetWindowWidth(size.X, fixedRight)
	w.SetWindowHeight(size.Y, fixedBottom)
	w.GeometryHypervisorResume()
}

func (w *Window) WindowGeometry() RectangleF64 {
	return w.driver.Geometry() // TODO
}

/*func (w *Window) SetWindowGeometry(geometry RectangleF64) {
	w.driver.SetGeometry(geometry)
}
*/
func (w *Window) WindowPos() PointF64 {
	return w.driver.Pos()
}

func (w *Window) SetWindowPos(pos PointF64) {
	w.driver.SetPos(pos)
}

func (w *Window) WindowSize() PointF64 {
	return w.BaseControl.Geometry().GetSize()
}

/*func (w *Window) SetWindowSize(size PointF64, fixedRight, fixedBottom bool) {
	w.driver.SetSize(size, fixedRight, fixedBottom)
}*/

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

func (w *Window) invalidateRegion(region RectangleF64) {
	w.driver.InvalidateRegion(region)
}

func (w *Window) invalidate() {
	w.driver.Invalidate()
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

func (w *Window) OfflineCanvas() OfflineCanvas { return w.driver.OfflineCanvas() }

func (w *Window) SetCursor(cursor Cursor, override bool) {
	if override {
		w.driver.SetCursor(cursor)
		w.cursorOverride = true
		return
	}
	w.cursor = cursor
	if !w.cursorOverride {
		w.driver.SetCursor(cursor)
	}
}
func (w *Window) StopCursorOverride() {
	w.cursorOverride = false
	w.driver.SetCursor(w.cursor)
}

//
// BaseControlI overrides
//

func (w *Window) setPossibleHorGeometry(minWidth, bestWidth, maxWidth float64) (changed bool) {
	changed = w.BaseControl.setPossibleHorGeometry(minWidth, bestWidth, maxWidth)
	if changed {
		w.driver.SetPossibleSize(w.minSize, w.maxSize)
	}
	return
}

func (w *Window) setPossibleVerGeometry(minHeight, bestHeight, maxHeight float64) (changed bool) {
	changed = w.BaseControl.setPossibleVerGeometry(minHeight, bestHeight, maxHeight)
	if changed {
		w.driver.SetPossibleSize(w.minSize, w.maxSize)
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
	if w.child == nil {
		return
	}
	if !w.geometryHypervisorIsReady() {
		w.setIR()
		return
	}
	count++
	w.child.Draw(canvas, region)
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
	//fmt.Println(e)
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

func (w *Window) onModifiers(e ModifiersEvent) {
	for _, subscriber := range w.modifiersEventSubscribers {
		subscriber.OnModifiersEvent(e)
	}
}

func (w *Window) onPointerDrag(e PointerDragEvent) {
	w.pointerPressControl.OnPointerDragEvent(e)
}

func (w *Window) onWindowMainStateEvent(e WindowMainStateEvent) {
	for _, subscriber := range w.windowMainStateEventSubscribers {
		subscriber.OnWindowMainStateEvent(e)
	}
}

func (w *Window) onWindowDisplayStateEvent(e WindowDisplayStateEvent) {
	for _, subscriber := range w.windowDisplayStateEventSubscribers {
		subscriber.OnWindowDisplayStateEvent(e)
	}
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
	w.cursor = Cursor(0).MakeDefault()
	w.closeAllowed = true
	w.minimizeAllowed = true
	w.maximizeAllowed = true
	w.fullScreenAllowed = true
	w.enterLeaveAreas = make(map[EnterLeaveAreaID]enterLeaveArea)
	w.moveAreas = make(map[MoveAreaID]Control)
	w.focusedControl = w
	w.driver.RegisterDrawCallback(w.Draw)
	w.driver.RegisterResizeCallback(w.onResize)
	w.driver.RegisterOfflineCanvasCallback(w.onOfflineCanvasChanged)
	w.driver.RegisterKeyboardCallback(w.onKeyboardEvent)
	w.driver.RegisterPointerKeyCallback(w.onPointerKey)
	w.driver.RegisterPointerDragCallback(w.onPointerDrag)
	w.driver.RegisterPointerMoveCallback(w.onPointerMove)
	w.driver.RegisterPointerEnterLeaveCallback(w.onPointerEnterLeave)
	w.driver.RegisterScrollCallback(w.onScroll)
	w.driver.RegisterModifiersCallback(w.onModifiers)
	w.driver.RegisterWindowMainStateCallback(w.onWindowMainStateEvent)
	w.driver.RegisterWindowDisplayStateCallback(w.onWindowDisplayStateEvent)
	w.BaseControl.window = w
	w.SetUPGIR(false)
}

func NewWindow(title string) *Window {
	var w Window
	w.driver = driverWindowConstructor(title)
	w.baseInit()
	return &w
}

func NewWindowAdvanced(dw DriverWindow) *Window {
	var w Window
	w.driver = dw
	w.GeometryHypervisorPause()
	w.baseInit()
	w.GeometryHypervisorResume()
	return &w
}

func (w *Window) Close() {
	w.GeometryHypervisorPause()
	w.SetChild(nil)
	w.driver.Close()
	Stop() // TODO What if there are more than 1 window?
}
