package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "window.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Window struct {
	pointer               unsafe.Pointer
	drawCallback          func(CanvasI, RectangleI) // TODO Rectangle is too simple
	eventCallback         func(EventI) bool
	resizeCallback        func()
	offlineCanvasCallback func()
}

func CreateWindow(width, height int) (window *Window, err error) {
	// We need initialize Window "in C memory" because we pass pointer to Window to top view (for long live).
	// If we do not do this Go GC can move Window to other location and applications crashes (at random moment).
	window = (*Window)(C.calloc(1, C.size_t(unsafe.Sizeof(Window{}))))

	window.pointer = C.CreateWindow(C.int(0), C.int(0), C.int(width), C.int(height)) // TODO 0 0
	if err != nil {
		return
	}

	view, ok := CreateTopView(window)
	if !ok {
		return nil, errors.New("Unable to create top NSView")
	}

	C.SetWindowTopView(unsafe.Pointer(window.pointer), unsafe.Pointer(view))
	C.MakeWindowKeyAndOrderFront(unsafe.Pointer(window.pointer))
	return
}

func (w *Window) Title() string {
	return C.GoString(C.GetWindowTitle(w.pointer))
}

func (w *Window) SetTitle(title string) {
	C.SetWindowTitle(w.pointer, C.CString(title))
}

func (w *Window) Destroy() {
	// TODO
}

func (w *Window) Geometry() RectangleI {
	r := C.GetWindowGeometry(w.pointer)
	return (*(*RectangleF64S)(unsafe.Pointer(&r))).ToI()
}

func (w *Window) Pos() PointI {
	return w.Geometry().LT()
}

func (w *Window) Size() PointI {
	return w.Geometry().GetSize()
}

func (w *Window) SetGeometry(geometry RectangleI) {
	w.SetPos(geometry.LT())
	w.SetSize(geometry.GetSize())
}

func (w *Window) SetPos(pos PointI) {
	posF := pos.ToF64()
	C.SetWindowPos(w.pointer, *(*C.CGPoint)(unsafe.Pointer(&posF)))
}
func (w *Window) SetSize(size PointI) {
	sizeF := size.ToF64()
	C.SetWindowSize(w.pointer, *(*C.CGSize)(unsafe.Pointer(&sizeF)))
}

func (w *Window) InvalidateRegion(region RectangleI) {
	regionS := region.ToF64S()
	C.InvalidateRegion(w.pointer, *(*C.NSRect)(unsafe.Pointer(&regionS)))
}

func (w *Window) Invalidate() {
	C.Invalidate(w.pointer)
}

func (w *Window) OfflineCanvas() OfflineCanvasI {
	return newContext(unsafe.Pointer(C.GetWindowContext(w.pointer)), float64(C.GetWindowScaleFactor(w.pointer)))
}

func (w *Window) RegisterDrawCallback(f func(CanvasI, RectangleI)) { w.drawCallback = f }
func (w *Window) RegisterEventCallback(f func(EventI) bool)        { w.eventCallback = f }
func (w *Window) RegisterResizeCallback(f func())                  { w.resizeCallback = f }
func (w *Window) RegisterOfflineCanvasCallback(f func())           { w.offlineCanvasCallback = f }
