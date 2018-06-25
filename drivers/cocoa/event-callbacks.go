// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

//#import <Cocoa/Cocoa.h>
import "C"
import (
	"unsafe"
)
import (
	"log"
)

//
// For window.h/m
//

//export windowMainEventCallback
func windowMainEventCallback(window unsafe.Pointer, become bool) {
	w := (*Window)(window)
	if w.windowMainEventCallback == nil {
		return
	}
	w.windowMainEventCallback(become)
}

//
// For top-view.h/m
//

//export drawCallback
func drawCallback(window unsafe.Pointer, context C.CGContextRef, rect C.CGRect) {
	log.Println("X1")
	if context == 0 {
		log.Println("X1.5")
		panic("Unable to retrieve context for drawing")
	}
	log.Println("X2")
	w := (*Window)(window)
	log.Println("X3")
	if w.drawCallback == nil {
		log.Println("X3.5")
		return
	}
	log.Println("X4")
	c := &Context{uintptr(context)}
	log.Println("X5", window, unsafe.Offsetof(w.drawCallback), w.drawCallback, unsafe.Pointer(context))
	w.drawCallback(c, RectangleF64{}) //(*RectangleF64S)(unsafe.Pointer(&rect)).ToF64())
	log.Println("X6")
}

//export keyboardEventCallback
func keyboardEventCallback(window unsafe.Pointer, event uint8, key uint16, modifiers uint64) {
	w := (*Window)(window)
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
func pointerKeyEventCallback(window unsafe.Pointer, event uint8, button uint8, point C.NSPoint, modifiers uint64) {
	w := (*Window)(window)
	if w.pointerKeyEventCallback == nil {
		return
	}
	tModifiers := translateKeyModifiers(modifiers)
	tPoint := *(*PointF64)(unsafe.Pointer(&point))
	e := PointerButtonEvent{PointerButtonEventKind(event), PointerButton(button), tModifiers, tPoint}
	w.pointerKeyEventCallback(e)
}

//export pointerDragEventCallback
func pointerDragEventCallback(window unsafe.Pointer, delta C.NSPoint) {
	w := (*Window)(window)
	if w.pointerDragEventCallback == nil {
		return
	}
	tDelta := *(*PointF64)(unsafe.Pointer(&delta))
	e := PointerDragEvent{tDelta}
	w.pointerDragEventCallback(e)
}

//export pointerMoveEventCallback
func pointerMoveEventCallback(window unsafe.Pointer, point C.NSPoint) {
	w := (*Window)(window)
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
func pointerEnterLeaveEventCallback(window unsafe.Pointer, trackingAreaID C.int, enter C.bool) {
	w := (*Window)(window)
	if w.pointerEnterLeaveEventCallback == nil {
		return
	}
	e := PointerEnterLeaveEvent{EnterLeaveAreaID(trackingAreaID), bool(enter)}
	w.pointerEnterLeaveEventCallback(e)
}

//export scrollEventCallback
func scrollEventCallback(window unsafe.Pointer, delta C.NSPoint, point C.NSPoint, modifiers uint64) {
	w := (*Window)(window)
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
func windowResizeCallback(window unsafe.Pointer, size C.NSSize) { // nolint: deadcode
	w := (*Window)(window)
	if w.resizeCallback == nil {
		return
	}
	w.resizeCallback(*(*PointF64)(unsafe.Pointer(&size)))
}
