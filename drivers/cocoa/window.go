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
	drawCallback          func(CanvasI, RectangleF64) // TODO Rectangle is too simple
	eventCallback         func(EventI) bool
	resizeCallback        func()
	offlineCanvasCallback func()
}

func CreateWindow() (window *Window, err error) {
	// We need initialize Window "in C memory" because we pass pointer to Window to top view (for long live).
	// If we do not do this Go GC can move Window to other location and applications crashes (at random moment).
	window = (*Window)(C.calloc(1, C.size_t(unsafe.Sizeof(Window{}))))

	window.pointer = C.CreateWindow(C.int(0), C.int(0), 0, 0)
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
	C.SetWindowTitle(w.pointer, C.CString(title)) // TODO do not use CString if possible in any place!!!
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
	return newContext(unsafe.Pointer(C.GetWindowContext(w.pointer)))
}

func (w *Window) ScaleFactor() float64 {
	return float64(C.GetWindowScaleFactor(w.pointer))
}

func (w *Window) RegisterDrawCallback(f func(CanvasI, RectangleF64)) { w.drawCallback = f }
func (w *Window) RegisterEventCallback(f func(EventI) bool)          { w.eventCallback = f }
func (w *Window) RegisterResizeCallback(f func())                    { w.resizeCallback = f }
func (w *Window) RegisterOfflineCanvasCallback(f func())             { w.offlineCanvasCallback = f }
