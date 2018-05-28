// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

//#import "context.h"
//#import "text.h"
import "C"
import "unsafe"

//type ContextP unsafe.Pointer

// TODO C.Free each C.CString()

type Context struct {
	pointer C.CGContextRef // TODO fix changes in files
	//transforms []TransformF64 // Stack of transforms
}

func newContext(rawContext unsafe.Pointer) *Context {
	return &Context{C.CGContextRef(rawContext)}
}

/*func (c *Context) ResetTransform() {
	C.resetTransform(c.pointer)
}*/
func (c *Context) PushTransform() {
	C.CGContextSaveGState(C.CGContextRef(c.pointer))
	//c.transforms = append(c.transforms, c.GetTransform())
}
func (c *Context) PopTransform() {
	C.CGContextRestoreGState(C.CGContextRef(c.pointer)) // TODO this restore not just transform, but other thing too.
	/*
		l := len(c.transforms)
		if l <= 0 {
			return
		}
		c.SetTransform(c.transforms[l-1])
		c.transforms = c.transforms[:l-1]
	*/
}
func (c *Context) GetTransform() TransformF64 {
	t := C.CGContextGetCTM(c.pointer)
	return *(*TransformF64)(unsafe.Pointer(&t))
}

/*func (c *Context) SetTransform(transform TransformF64) {
	C.CGContextConcatCTM(c.pointer, *(*C.CGAffineTransform)(unsafe.Pointer(&transform)))
}*/
func (c *Context) Rotate(angle float64) {
	C.CGContextRotateCTM(c.pointer, C.CGFloat(-angle))
}
func (c *Context) Scale(x float64) {
	c.ScaleXY(x, x)
}
func (c *Context) ScaleXY(x, y float64) {
	C.CGContextScaleCTM(c.pointer, C.CGFloat(x), C.CGFloat(y))
}
func (c *Context) Translate(pos PointF64) {
	C.CGContextTranslateCTM(c.pointer, C.CGFloat(pos.X), C.CGFloat(pos.Y))
}
func (c *Context) Superpose(original, required RectangleF64) {
	scaleX := required.Width() / original.Width()
	scaleY := required.Height() / original.Height()
	c.ScaleXY(scaleX, scaleY)
	var translatePos PointF64
	translatePos.X = required.Left - scaleX*original.Left
	translatePos.Y = required.Top - scaleY*original.Top
	c.Translate(translatePos)
}

func (c *Context) setLineColor(color ColorF64) {
	C.CGContextSetStrokeColor(C.CGContextRef(unsafe.Pointer(c.pointer)), (*C.CGFloat)(unsafe.Pointer(&color)))
}

func (c *Context) setLineWidth(width float64) {
	C.CGContextSetLineWidth(C.CGContextRef(unsafe.Pointer(c.pointer)), C.CGFloat(width))
}

func (c *Context) setFillColor(color ColorF64) {
	C.CGContextSetFillColor(C.CGContextRef(unsafe.Pointer(c.pointer)), (*C.CGFloat)(unsafe.Pointer(&color)))
}

func (c *Context) drawLine(point1, point2 PointF64) {
	points := [2]PointF64{point1, point2}
	C.CGContextStrokeLineSegments(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		(*C.CGPoint)(unsafe.Pointer(&points)), // TODO is it safe (can Go GC remove points while this function works)?
		2,
	)
}

func (c *Context) DrawLine(point1, point2 PointF64, color ColorF64, width float64) {
	c.setLineColor(color)
	c.setLineWidth(width)
	c.drawLine(point1, point2)
}

func (c *Context) drawConnectedLines(points []PointF64) {
	C.drawConnectedLines(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		(*C.CGPoint)(unsafe.Pointer(&points[0])), // TODO is it safe (can Go GC remove points while this function works)?
		C.size_t(len(points)),
	)
}

func (c *Context) DrawConnectedLines(points []PointF64, color ColorF64, width float64) {
	c.setLineColor(color)
	c.setLineWidth(width)
	c.drawConnectedLines(points)
}

func (c *Context) DrawStraightContour(points []PointF64, color ColorF64, width float64) {
	c.setLineColor(color)
	c.setLineWidth(width)
	C.drawStraightContour(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		(*C.CGPoint)(unsafe.Pointer(&points[0])), // TODO is it safe (can Go GC remove points while this function works)?
		C.size_t(len(points)),
	)
}

func (c *Context) FillStraightContour(points []PointF64, color ColorF64) {
	c.setFillColor(color)
	C.fillStraightContour(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		(*C.CGPoint)(unsafe.Pointer(&points[0])), // TODO is it safe (can Go GC remove points while this function works)?
		C.size_t(len(points)),
	)
}

func (c *Context) drawRectangleWithWidth(rect RectangleF64, width float64) {
	r := rect.ToF64S()
	C.CGContextStrokeRectWithWidth(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		C.CGFloat(width),
	)
}

func (c *Context) DrawRectangle(rect RectangleF64, color ColorF64, width float64) {
	c.setLineColor(color)
	c.drawRectangleWithWidth(rect, width)
}

func (c *Context) DrawRoundedRectangle(rect RoundedRectangleF64, color ColorF64, width float64) {
	r := rect.Rectangle.ToF64S()
	C.drawRoundedRectangle(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		C.CGFloat(rect.RadiusX),
		C.CGFloat(rect.RadiusY),
		(*C.CGFloat)(unsafe.Pointer(&color)), // TODO is it safe to pass pointer here?
		C.CGFloat(width),
	)
}

func (c *Context) DrawRoundedRectangleExtended(
	rect RectangleF64,
	radiusLT, radiusRT, radiusRB, radiusLB PointF64,
	color ColorF64,
	width float64,
) {
	r := rect.ToF64S()
	C.drawRoundedRectangleExtended(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRB)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLB)),
		(*C.CGFloat)(unsafe.Pointer(&color)), // TODO is it safe to pass pointer here?
		C.CGFloat(width),
	)
}

func (c *Context) drawEllipse(ellipse EllipseF64) {
	r := ellipse.OuterRectangle().ToF64S()
	C.CGContextStrokeEllipseInRect(C.CGContextRef(unsafe.Pointer(c.pointer)), *(*C.CGRect)(unsafe.Pointer(&r)))
}

func (c *Context) DrawEllipse(ellipse EllipseF64, color ColorF64, width float64) {
	c.setLineColor(color)
	c.setLineWidth(width)
	c.drawEllipse(ellipse)
}

func (c *Context) DrawCircle(circle CircleF64, color ColorF64, width float64) {
	c.DrawEllipse(circle.ToEllipse(), color, width)
}

func (c *Context) fillRectangle(rect RectangleF64) {
	r := rect.ToF64S()
	C.CGContextFillRect(C.CGContextRef(unsafe.Pointer(c.pointer)), *(*C.CGRect)(unsafe.Pointer(&r)))
}

func (c *Context) FillRectangle(rect RectangleF64, color ColorF64) {
	c.setFillColor(color)
	c.fillRectangle(rect)
}

func (c *Context) FillRoundedRectangle(rect RoundedRectangleF64, color ColorF64) {
	r := rect.Rectangle.ToF64S()
	C.fillRoundedRectangle(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		C.CGFloat(rect.RadiusX),
		C.CGFloat(rect.RadiusY),
		(*C.CGFloat)(unsafe.Pointer(&color)), // TODO is it safe to pass pointer here?
	)
}

func (c *Context) FillRoundedRectangleExtended(
	rect RectangleF64,
	radiusLT, radiusRT, radiusRB, radiusLB PointF64,
	color ColorF64,
) {
	r := rect.ToF64S()
	C.fillRoundedRectangleExtended(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRB)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLB)),
		(*C.CGFloat)(unsafe.Pointer(&color)), // TODO is it safe to pass pointer here?
	)
}

func (c *Context) fillEllipse(ellipse EllipseF64) {
	r := ellipse.OuterRectangle().ToF64S()
	C.CGContextFillEllipseInRect(C.CGContextRef(unsafe.Pointer(c.pointer)), *(*C.CGRect)(unsafe.Pointer(&r)))
}

func (c *Context) FillEllipse(ellipse EllipseF64, color ColorF64) {
	c.setFillColor(color)
	c.fillEllipse(ellipse)
}

func (c *Context) FillCircle(circle CircleF64, color ColorF64) {
	c.FillEllipse(circle.ToEllipse(), color)
}

func (c *Context) DrawTextLine(text string, font FontI, pos PointF64, color ColorF64) {
	// TODO color
	// TODO fontSize
	buf := []byte(text)
	C.DrawTextLine(
		C.CGContextRef(unsafe.Pointer(c.pointer)),
		(*C.UInt8)(unsafe.Pointer(&buf[0])), // TODO is it safe to pass pointer here?
		C.CFIndex(len(buf)),
		C.CTFontRef(font.H()),
		*(*C.CGPoint)(unsafe.Pointer(&pos)),
	)
}

func (c *Context) TextLineGeometry(text string, font FontI) PointF64 {
	// TODO fontSize
	buf := []byte(text)
	rect := C.GetTextLineGeometry(
		c.pointer,
		(*C.UInt8)(unsafe.Pointer(&buf[0])), // TODO is it safe to pass pointer here?
		C.CFIndex(len(buf)),
		C.CTFontRef(font.H()),
	)
	return (*(*RectangleF64S)(unsafe.Pointer(&rect))).Size
}
