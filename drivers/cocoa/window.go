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
	"errors"
	"unsafe"
)

type Window struct {
	pointer unsafe.Pointer

	drawCallback          func(CanvasI, RectangleF64) // TODO Rectangle is too simple
	resizeCallback        func()
	offlineCanvasCallback func()

	keyboardEventCallback    func(KeyboardEvent)
	pointerKeyEventCallback  func(PointerButtonEvent)
	pointerMoveEventCallback func(PointerMoveEvent)
	scrollEventCallback      func(ScrollEvent)
}

func CreateWindow() (window *Window, err error) {
	// We need initialize Window "in C memory" because we pass pointer to Window to top view (for long live).
	// If we do not do this Go GC can move Window to other location and applications crashes (at random moment).
	window = (*Window)(C.calloc(1, C.size_t(unsafe.Sizeof(Window{}))))

	window.pointer = C.CreateWindow(C.int(0), C.int(0), 0, 0)
	if err != nil {
		return
	}

	view, ok := createTopView(window)
	if !ok {
		return nil, errors.New("Unable to create top NSView")
	}

	C.SetWindowTopView(window.pointer, view)
	C.MakeWindowKeyAndOrderFront(window.pointer)
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

func (w *Window) Destroy() {
	// TODO
}

func (w *Window) Geometry() RectangleF64 {
	r := C.GetWindowGeometry(w.pointer)
	return (*(*RectangleF64S)(unsafe.Pointer(&r))).ToF64()
}

func (w *Window) Pos() PointF64 {
	return w.Geometry().LT()
}

func (w *Window) Size() PointF64 {
	return w.Geometry().GetSize()
}

func (w *Window) SetGeometry(geometry RectangleF64) {
	w.SetPos(geometry.LT())
	w.SetSize(geometry.GetSize())
}

func (w *Window) SetPos(pos PointF64) {
	C.SetWindowPos(w.pointer, *(*C.CGPoint)(unsafe.Pointer(&pos)))
}
func (w *Window) SetSize(size PointF64) {
	C.SetWindowSize(w.pointer, *(*C.CGSize)(unsafe.Pointer(&size)))
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
func (w *Window) RegisterResizeCallback(f func())                    { w.resizeCallback = f }
func (w *Window) RegisterOfflineCanvasCallback(f func())             { w.offlineCanvasCallback = f }

func (w *Window) RegisterKeyboardCallback(f func(KeyboardEvent)) { w.keyboardEventCallback = f }
func (w *Window) RegisterPointerKeyCallback(f func(PointerButtonEvent)) {
	w.pointerKeyEventCallback = f
}
func (w *Window) RegisterPointerMoveCallback(f func(PointerMoveEvent)) {
	if w.pointerMoveEventCallback == nil && f != nil {
		C.SetWindowAcceptMouseMoved(w.pointer, true)
	} else if w.pointerMoveEventCallback != nil && f == nil {
		C.SetWindowAcceptMouseMoved(w.pointer, false)
	}
	w.pointerMoveEventCallback = f
}
func (w *Window) RegisterScrollCallback(f func(ScrollEvent)) { w.scrollEventCallback = f }
