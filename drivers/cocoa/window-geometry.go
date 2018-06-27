// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

//#include "window.h"
import "C"
import "unsafe"

func (w *Window) Geometry() RectangleF64 {
	g := C.GetWindowGeometry(w.pointer)
	return (*(*RectangleF64S)(unsafe.Pointer(&g))).ToF64()
}

func (w *Window) SetGeometry(geometry RectangleF64) {
	g := geometry.ToF64S()
	C.SetWindowGeometry(w.pointer, *(*C.CGRect)(unsafe.Pointer(&g)))
}

func (w *Window) Pos() PointF64 {
	p := C.GetWindowPos(w.pointer)
	return *(*PointF64)(unsafe.Pointer(&p))
}

func (w *Window) SetPos(pos PointF64) {
	C.SetWindowPos(w.pointer, *(*C.CGPoint)(unsafe.Pointer(&pos)))
}

func (w *Window) Left() float64 {
	return float64(C.GetWindowLeft(w.pointer))
}

func (w *Window) SetLeft(left float64) {
	C.SetWindowLeft(w.pointer, C.CGFloat(left))
}

func (w *Window) Right() float64 {
	return float64(C.GetWindowRight(w.pointer))
}

func (w *Window) SetRight(right float64) {
	C.SetWindowRight(w.pointer, C.CGFloat(right))
}

func (w *Window) Top() float64 {
	return float64(C.GetWindowTop(w.pointer))
}

func (w *Window) SetTop(top float64) {
	C.SetWindowTop(w.pointer, C.CGFloat(top))
}

func (w *Window) Bottom() float64 {
	return float64(C.GetWindowBottom(w.pointer))
}

func (w *Window) SetBottom(bottom float64) {
	C.SetWindowBottom(w.pointer, C.CGFloat(bottom))
}

func (w *Window) Size() PointF64 {
	s := C.GetWindowSize(w.pointer)
	return *(*PointF64)(unsafe.Pointer(&s))
}

func (w *Window) SetSize(size PointF64, fixedRight, fixedBottom bool) {
	C.SetWindowSize(w.pointer, *(*C.CGSize)(unsafe.Pointer(&size)), C.bool(fixedRight), C.bool(fixedBottom))
}

func (w *Window) Width() float64 {
	return float64(C.GetWindowWidth(w.pointer))
}

func (w *Window) SetWidth(width float64, fixedRight bool) {
	C.SetWindowWidth(w.pointer, C.CGFloat(width), C.bool(fixedRight))
}

func (w *Window) Height() float64 {
	return float64(C.GetWindowHeight(w.pointer))
}

func (w *Window) SetHeight(height float64, fixedBottom bool) {
	C.SetWindowHeight(w.pointer, C.CGFloat(height), C.bool(fixedBottom))
}
