// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

//#import <Cocoa/Cocoa.h>
import "C"
import (
	"unsafe"
)

//
// For window.h/m
//

//export windowMainEventCallback
func windowMainEventCallback(windowID int, become bool) {
	w := windowByID(windowID)
	if w.windowMainEventCallback == nil {
		return
	}
	w.windowMainEventCallback(become)
}

//
// For top-view.h/m
//

//export drawCallback
func drawCallback(windowID int, context C.CGContextRef, rect C.CGRect) {
	if context == 0 {
		panic("Unable to retrieve context for drawing")
	}
	w := windowByID(windowID)
	if w.drawCallback == nil {
		return
	}
	c := &Context{uintptr(context)}
	w.drawCallback(c, RectangleF64{}) //(*RectangleF64S)(unsafe.Pointer(&rect)).ToF64())
}

//export keyboardEventCallback
func keyboardEventCallback(windowID int, event uint8, key uint16, modifiers uint64) {
	w := windowByID(windowID)
	if w.keyboardEventCallback == nil {
		return
	}
	tEvent := KeyEvent(event)
	tKey := translateKey(key)
	tModifiers := translateKeyModifiers(modifiers)
	e := KeyboardEvent{tEvent, tKey, tModifiers}
	w.keyboardEventCallback(e)
}

//export pointerKeyEventCallback
func pointerKeyEventCallback(windowID int, event uint8, button uint8, point C.NSPoint, modifiers uint64) {
	w := windowByID(windowID)
	if w.pointerKeyEventCallback == nil {
		return
	}
	tModifiers := translateKeyModifiers(modifiers)
	tPoint := *(*PointF64)(unsafe.Pointer(&point))
	e := PointerButtonEvent{PointerButtonEventKind(event), PointerButton(button), tModifiers, tPoint}
	w.pointerKeyEventCallback(e)
}

//export pointerDragEventCallback
func pointerDragEventCallback(windowID int, delta C.NSPoint) {
	w := windowByID(windowID)
	if w.pointerDragEventCallback == nil {
		return
	}
	tDelta := *(*PointF64)(unsafe.Pointer(&delta))
	e := PointerDragEvent{tDelta}
	w.pointerDragEventCallback(e)
}

//export pointerMoveEventCallback
func pointerMoveEventCallback(windowID int, point C.NSPoint) {
	w := windowByID(windowID)
	if w.pointerMoveEventCallback == nil {
		return
	}
	var e PointerMoveEvent
	e.Point = *(*PointF64)(unsafe.Pointer(&point))

	// Check all move areas for containing move point.
	// For all containing areas event will be sent.
	for _, area := range w.moveAreas {
		if area.Area.Contains(e.Point) {
			e.ID = area.ID
			w.pointerMoveEventCallback(e)
		}
	}
}

//export pointerEnterLeaveEventCallback
func pointerEnterLeaveEventCallback(windowID int, trackingAreaID C.int, enter C.bool) {
	w := windowByID(windowID)
	if w.pointerEnterLeaveEventCallback == nil {
		return
	}
	e := PointerEnterLeaveEvent{EnterLeaveAreaID(trackingAreaID), bool(enter)}
	w.pointerEnterLeaveEventCallback(e)
}

//export scrollEventCallback
func scrollEventCallback(windowID int, delta C.NSPoint, point C.NSPoint, modifiers uint64) {
	w := windowByID(windowID)
	if w.scrollEventCallback == nil {
		return
	}
	tDelta := *(*PointF64)(unsafe.Pointer(&delta))
	tModifiers := translateKeyModifiers(modifiers)
	tPoint := *(*PointF64)(unsafe.Pointer(&point))
	e := ScrollEvent{tDelta, tModifiers, tPoint}
	w.scrollEventCallback(e)
}

//export windowResizeCallback
func windowResizeCallback(windowID int, size C.NSSize) { // nolint: deadcode
	w := windowByID(windowID)
	if w.resizeCallback == nil {
		return
	}
	w.resizeCallback(*(*PointF64)(unsafe.Pointer(&size)))
}
