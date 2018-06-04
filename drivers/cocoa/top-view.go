// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

import (
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "top-view.h"
*/
import "C"

func createTopView(window *Window) (view unsafe.Pointer, ok bool) {
	view = C.CreateTopView(unsafe.Pointer(window))
	ok = view != nil
	return
}

//export drawCallback
func drawCallback(window unsafe.Pointer, context C.CGContextRef, rect C.CGRect) {
	if window == nil {
		panic("NIL window") // TODO
	}
	if context == 0 {
		panic("Unable to retrieve context for drawing")
	}
	w := (*Window)(window)
	if w.drawCallback == nil {
		return
	}
	c := newContext(uintptr(context))
	w.drawCallback(c, (*RectangleF64S)(unsafe.Pointer(&rect)).ToF64())
}

//export keyboardEventCallback
func keyboardEventCallback(window unsafe.Pointer, event uint8, key uint16, modifiers uint64) {
	if window == nil {
		panic("NIL window") // TODO
	}
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
	if window == nil {
		panic("NIL window") // TODO
	}
	w := (*Window)(window)
	if w.pointerKeyEventCallback == nil {
		return
	}
	tModifiers := translateKeyModifiers(modifiers)
	tPoint := *(*PointF64)(unsafe.Pointer(&point))
	e := PointerButtonEvent{PointerButtonEventKind(event), PointerButton(button), tModifiers, tPoint}
	w.pointerKeyEventCallback(e)
}

//export pointerMoveEventCallback
func pointerMoveEventCallback(window unsafe.Pointer, point C.NSPoint) {
	if window == nil {
		panic("NIL window") // TODO
	}
	w := (*Window)(window)
	if w.pointerMoveEventCallback == nil {
		return
	}
	tPoint := *(*PointF64)(unsafe.Pointer(&point))
	e := PointerMoveEvent{tPoint}
	w.pointerMoveEventCallback(e)
}

//export scrollEventCallback
func scrollEventCallback(window unsafe.Pointer, deltaX float64, deltaY float64) {
	if window == nil {
		panic("NIL window") // TODO
	}
	w := (*Window)(window)
	if w.scrollEventCallback == nil {
		return
	}
	e := ScrollEvent{deltaX, deltaY}
	w.scrollEventCallback(e)
}
