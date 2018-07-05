// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import "context.h"
#import "text.h"

CFStringRef CFStringCreateFromGoString(_GoString_ str);

CGSize _GetTextImageGeometry(CGContextRef context, _GoString_ str, CTFontRef font) {
	CFStringRef _str = CFStringCreateFromGoString(str);
	CGSize r = GetTextImageGeometry(context, _str, font);
	CFRelease(_str);
	return r;
}

void _DrawTextImage(CGContextRef context, _GoString_ str, CTFontRef font, CGFloat* color, CGPoint pos) {
	CFStringRef _str = CFStringCreateFromGoString(str);
	DrawTextImage(context, _str, font, color, pos);
	CFRelease(_str);
}

struct TextLineGeometry _GetTextLineGeometry(CGContextRef context, _GoString_ str, CTFontRef font) {
	CFStringRef _str = CFStringCreateFromGoString(str);
	struct TextLineGeometry r = GetTextLineGeometry(context, _str, font);
	CFRelease(_str);
	return r;
}

void _DrawTextLine(CGContextRef context, _GoString_ str, CTFontRef font, CGFloat* color, CGPoint pos, uint8_t origin) {
	CFStringRef _str = CFStringCreateFromGoString(str);
	DrawTextLine(context, _str, font, color, pos, origin);
	CFRelease(_str);
}
*/
import "C"
import "unsafe"

type Context struct {
	pointer uintptr // C.CGContextRef
}

// clip, transformation matrix, ???
func (c *Context) SaveState() {
	C.CGContextSaveGState(C.CGContextRef(c.pointer))
}
func (c *Context) RestoreState() {
	C.CGContextRestoreGState(C.CGContextRef(c.pointer))
}
func (c *Context) GetTransform() TransformF64 {
	t := C.CGContextGetCTM(C.CGContextRef(c.pointer))
	return *(*TransformF64)(unsafe.Pointer(&t))
}

func (c *Context) Rotate(angle float64) {
	C.CGContextRotateCTM(C.CGContextRef(c.pointer), C.CGFloat(-angle))
}
func (c *Context) Scale(x float64) {
	c.ScaleXY(x, x)
}
func (c *Context) ScaleXY(x, y float64) {
	C.CGContextScaleCTM(C.CGContextRef(c.pointer), C.CGFloat(x), C.CGFloat(y))
}
func (c *Context) Translate(pos PointF64) {
	C.CGContextTranslateCTM(C.CGContextRef(c.pointer), C.CGFloat(pos.X), C.CGFloat(pos.Y))
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

func (c *Context) ClipToRectangle(region RectangleF64) {
	rect := region.ToF64S()
	C.CGContextClipToRect(C.CGContextRef(c.pointer), *(*C.CGRect)(unsafe.Pointer(&rect)))
}

func (c *Context) setLineColor(color ColorF64) {
	C.CGContextSetStrokeColor(C.CGContextRef(c.pointer), (*C.CGFloat)(unsafe.Pointer(&color)))
}

func (c *Context) setLineWidth(width float64) {
	C.CGContextSetLineWidth(C.CGContextRef(c.pointer), C.CGFloat(width))
}

func (c *Context) setFillColor(color ColorF64) {
	C.CGContextSetFillColor(C.CGContextRef(c.pointer), (*C.CGFloat)(unsafe.Pointer(&color)))
}

func (c *Context) drawLine(point1, point2 PointF64) {
	points := [2]PointF64{point1, point2}
	C.CGContextStrokeLineSegments( // nolint: unparam
		C.CGContextRef(c.pointer),
		(*C.CGPoint)(unsafe.Pointer(&points)),
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
		C.CGContextRef(c.pointer),
		(*C.CGPoint)(unsafe.Pointer(&points[0])),
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
		C.CGContextRef(c.pointer),
		(*C.CGPoint)(unsafe.Pointer(&points[0])),
		C.size_t(len(points)),
	)
}

func (c *Context) FillStraightContour(points []PointF64, color ColorF64) {
	c.setFillColor(color)
	C.fillStraightContour(
		C.CGContextRef(c.pointer),
		(*C.CGPoint)(unsafe.Pointer(&points[0])),
		C.size_t(len(points)),
	)
}

func (c *Context) drawRectangleWithWidth(rect RectangleF64, width float64) {
	r := rect.ToF64S()
	C.CGContextStrokeRectWithWidth(
		C.CGContextRef(c.pointer),
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
		C.CGContextRef(c.pointer),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		C.CGFloat(rect.RadiusX),
		C.CGFloat(rect.RadiusY),
		(*C.CGFloat)(unsafe.Pointer(&color)),
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
		C.CGContextRef(c.pointer),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRB)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLB)),
		(*C.CGFloat)(unsafe.Pointer(&color)),
		C.CGFloat(width),
	)
}

func (c *Context) drawEllipse(ellipse EllipseF64) {
	r := ellipse.OuterRectangle().ToF64S()
	C.CGContextStrokeEllipseInRect(C.CGContextRef(c.pointer), *(*C.CGRect)(unsafe.Pointer(&r)))
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
	C.CGContextFillRect(C.CGContextRef(c.pointer), *(*C.CGRect)(unsafe.Pointer(&r)))
}

func (c *Context) FillRectangle(rect RectangleF64, color ColorF64) {
	c.setFillColor(color)
	c.fillRectangle(rect)
}

func (c *Context) FillRoundedRectangle(rect RoundedRectangleF64, color ColorF64) {
	r := rect.Rectangle.ToF64S()
	C.fillRoundedRectangle(
		C.CGContextRef(c.pointer),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		C.CGFloat(rect.RadiusX),
		C.CGFloat(rect.RadiusY),
		(*C.CGFloat)(unsafe.Pointer(&color)),
	)
}

func (c *Context) FillRoundedRectangleExtended(
	rect RectangleF64,
	radiusLT, radiusRT, radiusRB, radiusLB PointF64,
	color ColorF64,
) {
	r := rect.ToF64S()
	C.fillRoundedRectangleExtended(
		C.CGContextRef(c.pointer),
		*(*C.CGRect)(unsafe.Pointer(&r)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRT)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusRB)),
		*(*C.CGPoint)(unsafe.Pointer(&radiusLB)),
		(*C.CGFloat)(unsafe.Pointer(&color)),
	)
}

func (c *Context) fillEllipse(ellipse EllipseF64) {
	r := ellipse.OuterRectangle().ToF64S()
	C.CGContextFillEllipseInRect(C.CGContextRef(c.pointer), *(*C.CGRect)(unsafe.Pointer(&r)))
}

func (c *Context) FillEllipse(ellipse EllipseF64, color ColorF64) {
	c.setFillColor(color)
	c.fillEllipse(ellipse)
}

func (c *Context) FillCircle(circle CircleF64, color ColorF64) {
	c.FillEllipse(circle.ToEllipse(), color)
}

func (c *Context) TextImageGeometry(text string, font FontI) PointF64 {
	size := C._GetTextImageGeometry(
		C.CGContextRef(c.pointer),
		text,
		C.CTFontRef(font.(Font).pointer),
	)
	return *(*PointF64)(unsafe.Pointer(&size))
}

func (c *Context) DrawTextImage(text string, font FontI, color ColorF64, pos PointF64) {
	C._DrawTextImage(
		C.CGContextRef(c.pointer),
		text,
		C.CTFontRef(font.(Font).pointer),
		(*C.CGFloat)(unsafe.Pointer(&color)),
		*(*C.CGPoint)(unsafe.Pointer(&pos)),
	)
}

func (c *Context) TextLineGeometry(text string, font FontI) (width, ascent, descent, leading float64) {
	size := C._GetTextLineGeometry(
		C.CGContextRef(c.pointer),
		text,
		C.CTFontRef(font.(Font).pointer),
	)
	r := *(*struct{ width, ascent, descent, leading float64 })(unsafe.Pointer(&size))
	return r.width, r.ascent, r.descent, r.leading
}

func (c *Context) DrawTextLine(text string, font FontI, color ColorF64, pos PointF64, origin TextLineOrigin) {
	C._DrawTextLine(
		C.CGContextRef(c.pointer),
		text,
		C.CTFontRef(font.(Font).pointer),
		(*C.CGFloat)(unsafe.Pointer(&color)),
		*(*C.CGPoint)(unsafe.Pointer(&pos)),
		C.uint8_t(origin),
	)
}
