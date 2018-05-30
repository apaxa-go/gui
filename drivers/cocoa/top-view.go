// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

import (
	"log"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "top-view.h"
*/
import "C"

type TopViewP unsafe.Pointer

func CreateTopView(window *Window) (view TopViewP, ok bool) {
	view = TopViewP(C.CreateTopView(unsafe.Pointer(window)))
	ok = view != nil
	return
}

//export drawCallback
func drawCallback(window, contextP unsafe.Pointer, rect C.CGRect) {
	if window == nil {
		panic("NIL window") // TODO
		//return
	}
	if contextP == nil {
		panic("Unable to retrieve context for drawing")
	}
	w := (*Window)(window)
	if w.drawCallback != nil {
		c := newContext(contextP)
		w.drawCallback(c, (*RectangleF64S)(unsafe.Pointer(&rect)).ToF64())
		//log.Println(c.GetTransform())
	}
}

//export keyboardEventCallback
func keyboardEventCallback(window unsafe.Pointer, event uint8, key uint16, modifiers uint64) {
	kEvent := KeyEvent(event)
	k := translateKey(key)
	kModifiers := translateKeyModifiers(modifiers)

	e := KeyboardEvent{kEvent, k, kModifiers}

	log.Println(e.ShortString())
}

//export pointerEventCallback
func pointerEventCallback(window unsafe.Pointer) {

}
