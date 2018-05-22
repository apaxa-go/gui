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
	}
}
