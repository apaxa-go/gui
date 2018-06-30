// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#include "window.h"

CFStringRef CFStringCreateFromGoString(_GoString_ str);

void _SetWindowTitle(void* self, _GoString_ title){
	CFStringRef _title = CFStringCreateFromGoString(title);
	SetWindowTitle(self, _title);
	CFRelease(_title);
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

var windows []*Window
var windowsFree []int // slice of windows indexes with nil value (in any order)
var windowsMutex sync.RWMutex

func mapWindow(w *Window) (id int) {
	windowsMutex.Lock()
	defer windowsMutex.Unlock()

	l := len(windowsFree)
	if l > 0 {
		id = windowsFree[l-1]
		windowsFree = windowsFree[:l-1]
		windows[id] = w
		return
	}
	l = len(windows)
	id = l
	windows = append(windows, w)
	return
}

func unmapWindow(id int) {
	windowsMutex.Lock()
	defer windowsMutex.Unlock()

	windows[id] = nil
	windowsFree = append(windowsFree, id)
}

func windowByID(id int) *Window {
	windowsMutex.RLock()
	defer windowsMutex.RUnlock()

	return windows[id]
}

type moveArea struct {
	ID   MoveAreaID
	Area RectangleF64
}

type Window struct {
	pointer unsafe.Pointer

	displayState WindowDisplayState
	//restoreToMaximized bool // True if last of (normal, maximized) is maximized. Indicates to which mode restore from minimized & full screen.
	//lastNormalSize     PointF64

	drawCallback          func(CanvasI, RectangleF64) // TODO Rectangle is too simple
	resizeCallback        func(size PointF64)
	offlineCanvasCallback func()

	keyboardEventCallback          func(KeyboardEvent)
	pointerKeyEventCallback        func(PointerButtonEvent)
	pointerDragEventCallback       func(PointerDragEvent)
	pointerMoveEventCallback       func(PointerMoveEvent)
	pointerEnterLeaveEventCallback func(event PointerEnterLeaveEvent)
	scrollEventCallback            func(ScrollEvent)
	modifiersEventCallback         func(KeyModifiers)
	windowMainEventCallback        func(become bool)
	windowDisplayStateCallback     func(oldState, newState WindowDisplayState)

	moveAreas []moveArea // AppKit does not internally assign pointer move events to corresponding tracking areas, so we do it here.
}

func CreateWindow(title string) (window *Window, err error) {
	// TODO We need initialize Window "in C memory" because we pass pointer to Window to top view (for long live).
	// If we do not do this Go GC can move Window to other location and applications will crash (at random moment) - this is theory.
	// But if we use C.{c/m}alloc then applications will crash (at random moment) - this is real world (gdb shows what normal Go function may overwrite such memory).

	window = &Window{}
	id := mapWindow(window)
	window.pointer = C.CreateWindow(C.int(id))
	if err != nil {
		return
	}

	window.SetTitle(title)
	return
}

func CToGoString(cString unsafe.Pointer) string { // TODO move to other package
	r := C.GoString((*C.char)(cString)) // TODO is it possible to share C types between packages and pass *C.char to this function directly?
	C.free(cString)
	return r
}

func (w *Window) Title() string {
	return CToGoString(unsafe.Pointer(C.GetWindowTitle(w.pointer)))
}

func (w *Window) SetTitle(title string) {
	C._SetWindowTitle(w.pointer, title)
}

func (w *Window) Close() {
	C.CloseWindow(w.pointer)
}

func (w *Window) SetPossibleSize(min, max PointF64) {
	C.SetWindowPossibleSize(w.pointer, *(*C.CGSize)(unsafe.Pointer(&min)), *(*C.CGSize)(unsafe.Pointer(&max)))
}

/*func (*Window) PresentedDisplayStates() WindowDisplayState {
	return WindowDisplayState(0).MakeNormal() | WindowDisplayState(0).MakeMinimized() | WindowDisplayState(0).MakeMaximized() | WindowDisplayState(0).MakeFullScreen()
}
*/
func (w *Window) DisplayState() WindowDisplayState { return w.displayState }

/*func (*Window) possibleDisplayStates(current WindowDisplayState) WindowDisplayState {
	// On Mac OS it possible to go from any state to any with exception: deny direct switch between minimized and full screen mode.
	switch current {
	case WindowDisplayState(0).MakeNormal():
		return WindowDisplayState(0).MakeMinimized() | WindowDisplayState(0).MakeMaximized() | WindowDisplayState(0).MakeFullScreen()
	case WindowDisplayState(0).MakeMinimized():
		return WindowDisplayState(0).MakeNormal() | WindowDisplayState(0).MakeMaximized()
	case WindowDisplayState(0).MakeMaximized():
		return WindowDisplayState(0).MakeNormal() | WindowDisplayState(0).MakeMinimized() | WindowDisplayState(0).MakeFullScreen()
	case WindowDisplayState(0).MakeFullScreen():
		return WindowDisplayState(0).MakeNormal() | WindowDisplayState(0).MakeMaximized()
	default:
		return (*Window)(nil).PresentedDisplayStates()
	}
}

func (w *Window) PossibleDisplayStates() WindowDisplayState {
	return w.possibleDisplayStates(w.displayState)
}
*/

func (w *Window) Minimize() {
	C.MinimizeWindow(w.pointer)
}

func (w *Window) Deminimize() {
	C.DeminimizeWindow(w.pointer)
}

func (w *Window) Maximize() {
	/*if !w.displayState.IsNormal() {
		return
	}*/
	C.ZoomWindow(w.pointer)
	//w.setDisplayState(WindowDisplayState(0).MakeMaximized())
}

func (w *Window) Demaximize() {
	/*if !w.displayState.IsMaximized() {
		return
	}*/
	C.ZoomWindow(w.pointer)
	//w.setDisplayState(WindowDisplayState(0).MakeNormal())
}

func (w *Window) EnterFullScreen() {
	if !w.displayState.IsFullScreen() {
		C.ToggleFullScreen(w.pointer)
	}
}

func (w *Window) ExitFullScreen() {
	if w.displayState.IsFullScreen() {
		C.ToggleFullScreen(w.pointer)
	}
}

func (w *Window) setDisplayState(state WindowDisplayState) {
	if w.displayState == state {
		return
	}
	/*if state.IsNormal() {
		w.restoreToMaximized = false
	} else if state.IsMaximized() {
		w.restoreToMaximized = true
	}*/
	oldState := w.displayState
	w.displayState = state
	if w.windowDisplayStateCallback != nil {
		w.windowDisplayStateCallback(oldState, state)
	}
}

func (w *Window) InvalidateRegion(region RectangleF64) {
	regionS := region.ToF64S()
	C.InvalidateRegion(w.pointer, *(*C.NSRect)(unsafe.Pointer(&regionS)))
}

func (w *Window) Invalidate() {
	C.Invalidate(w.pointer)
}

func (w *Window) OfflineCanvas() OfflineCanvasI {
	return &Context{uintptr(C.GetWindowContext(w.pointer))}
}

func (w *Window) ScaleFactor() float64 {
	return float64(C.GetWindowScaleFactor(w.pointer))
}

func (w *Window) RegisterDrawCallback(f func(CanvasI, RectangleF64)) { w.drawCallback = f }
func (w *Window) RegisterResizeCallback(f func(PointF64))            { w.resizeCallback = f }
func (w *Window) RegisterOfflineCanvasCallback(f func())             { w.offlineCanvasCallback = f }

func (w *Window) RegisterKeyboardCallback(f func(KeyboardEvent)) { w.keyboardEventCallback = f }
func (w *Window) RegisterPointerKeyCallback(f func(PointerButtonEvent)) {
	w.pointerKeyEventCallback = f
}
func (w *Window) RegisterPointerDragCallback(f func(PointerDragEvent)) {
	w.pointerDragEventCallback = f
}
func (w *Window) RegisterPointerMoveCallback(f func(PointerMoveEvent)) {
	/*if w.pointerMoveEventCallback == nil && f != nil {
		C.SetWindowAcceptMouseMoved(w.pointer, true)
	} else if w.pointerMoveEventCallback != nil && f == nil {
		C.SetWindowAcceptMouseMoved(w.pointer, false)
	}*/
	w.pointerMoveEventCallback = f
}
func (w *Window) RegisterPointerEnterLeaveCallback(f func(PointerEnterLeaveEvent)) {
	w.pointerEnterLeaveEventCallback = f
}

func (w *Window) RegisterScrollCallback(f func(ScrollEvent))     { w.scrollEventCallback = f }
func (w *Window) RegisterModifiersCallback(f func(KeyModifiers)) { w.modifiersEventCallback = f }
func (w *Window) RegisterWindowMainCallback(f func(become bool)) { w.windowMainEventCallback = f }
func (w *Window) RegisterWindowDisplayStateCallback(f func(sizeState, possibleSizeState WindowDisplayState)) {
	w.windowDisplayStateCallback = f
}

func (w *Window) AddEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64) {
	rectS := area.ToF64S()
	rect := *(*C.NSRect)(unsafe.Pointer(&rectS))
	C.AddTrackingArea(w.pointer, false, C.int(id), rect)
}
func (w *Window) ReplaceEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64) {
	rectS := area.ToF64S()
	rect := *(*C.NSRect)(unsafe.Pointer(&rectS))
	C.ReplaceTrackingArea(w.pointer, false, C.int(id), rect)
}
func (w *Window) RemoveEnterLeaveArea(id EnterLeaveAreaID) {
	C.RemoveTrackingArea(w.pointer, false, C.int(id))
}

func (w *Window) AddMoveArea(id MoveAreaID, area RectangleF64) {
	w.moveAreas = append(w.moveAreas, moveArea{id, area})
	rectS := area.ToF64S()
	rect := *(*C.NSRect)(unsafe.Pointer(&rectS))
	C.AddTrackingArea(w.pointer, true, C.int(id), rect)
}
func (w *Window) moveAreaByID(id MoveAreaID) (index int) {
	for index = range w.moveAreas {
		if w.moveAreas[index].ID == id {
			return
		}
	}
	return -1
}
func (w *Window) ReplaceMoveArea(id MoveAreaID, area RectangleF64) {
	if index := w.moveAreaByID(id); index >= 0 {
		w.moveAreas[index].Area = area
	} else {
		w.moveAreas = append(w.moveAreas, moveArea{id, area})
	}
	rectS := area.ToF64S()
	rect := *(*C.NSRect)(unsafe.Pointer(&rectS))
	C.ReplaceTrackingArea(w.pointer, true, C.int(id), rect)
}
func (w *Window) RemoveMoveArea(id MoveAreaID) {
	C.RemoveTrackingArea(w.pointer, true, C.int(id))
	if index := w.moveAreaByID(id); index >= 0 {
		w.moveAreas = append(w.moveAreas[:index], w.moveAreas[index+1:]...)
	}
}

func (w *Window) SetCursor(cursor Cursor) {
	C.setCursor(C.uint8(cursor))
}
