// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

const (
	windowResizerWidth        = 3
	windowResizerHeight       = 3
	windowResizerCornerWidth  = 10
	windowResizerCornerHeight = 10
)

/*
type WindowResizerEnable uint8

const (
	// Sides
	WindowResizerLeft WindowResizerEnable = 1 << iota
	WindowResizerRight
	WindowResizerTop
	WindowResizerBottom

	// Corners
	WindowResizerLeftTop
	WindowResizerRightTop
	WindowResizerLeftBottom
	WindowResizerRightBottom
)

const WindowResizerAllDisabled WindowResizerEnable = 0

func (e WindowResizerEnable) AllDisabled() bool { return e == WindowResizerAllDisabled }
func (e WindowResizerEnable) Left() bool        { return e&WindowResizerLeft > 0 }
func (e WindowResizerEnable) Right() bool       { return e&WindowResizerRight > 0 }
func (e WindowResizerEnable) Top() bool         { return e&WindowResizerTop > 0 }
func (e WindowResizerEnable) Bottom() bool      { return e&WindowResizerBottom > 0 }
func (e WindowResizerEnable) LeftTop() bool     { return e&WindowResizerLeftTop > 0 }
func (e WindowResizerEnable) RightTop() bool    { return e&WindowResizerRightTop > 0 }
func (e WindowResizerEnable) LeftBottom() bool  { return e&WindowResizerLeftBottom > 0 }
func (e WindowResizerEnable) RightBottom() bool { return e&WindowResizerRightBottom > 0 }
*/

type resizeState int8

const (
	resizePositive resizeState = 1
	resizeNegative resizeState = -1
	resizeNone     resizeState = 0
)

//func (s resizeState) isActive() bool { return s == resizePositive || s == resizeNegative }

type WindowResizer struct {
	BaseControl

	leftAreaID   EnterLeaveAreaID
	rightAreaID  EnterLeaveAreaID
	topAreaID    EnterLeaveAreaID
	bottomAreaID EnterLeaveAreaID

	leftTopAreaID     EnterLeaveAreaID
	rightTopAreaID    EnterLeaveAreaID
	leftBottomAreaID  EnterLeaveAreaID
	rightBottomAreaID EnterLeaveAreaID

	ready bool // This is state.
	//lastAreaID EnterLeaveAreaID // This is state.
	resizeHor resizeState // This is state, not setting. -1 means resize left border, 1 means resize right border, no resize otherwise.
	resizeVer resizeState // This is state, not setting. -1 means resize top border, 1 means resize bottom border, no resize otherwise.

	//baseGeometry RectangleF64
	baseSize PointF64
}

func (c *WindowResizer) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	/*
		if c.enable.LeftTop() || c.enable.LeftBottom(){
			minWidth+=windowResizerCornerWidth
		}else if c.enable.Left(){ // here we assume that CornerWidth > Width
			minWidth+=windowResizerWidth
		}
		if c.enable.RightTop() || c.enable.RightBottom(){
			minWidth+=windowResizerCornerWidth
		}else if c.enable.Right(){ // here we assume that CornerWidth > Width
			minWidth+=windowResizerWidth
		}
		return minWidth,minWidth,mathh.PositiveInfFloat64() // TODO no best width
	*/
	return 2 * windowResizerCornerWidth, 2 * windowResizerCornerWidth, mathh.PositiveInfFloat64() // TODO no best width
}

func (c *WindowResizer) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	/*
		if c.enable.LeftTop() || c.enable.RightTop(){
			minHeight+=windowResizerCornerHeight
		}else if c.enable.Top(){ // here we assume that CornerHeight > Height
			minHeight+=windowResizerHeight
		}
		if c.enable.LeftBottom() || c.enable.RightBottom(){
			minHeight+=windowResizerCornerHeight
		}else if c.enable.Bottom(){ // here we assume that CornerHeight > Height
			minHeight+=windowResizerHeight
		}
		return minHeight,minHeight,mathh.PositiveInfFloat64() // TODO no best height
	*/
	return 2 * windowResizerCornerHeight, 2 * windowResizerCornerHeight, mathh.PositiveInfFloat64() // TODO no best height
}

func (c WindowResizer) Draw(canvas Canvas, region RectangleF64) {}

func (c *WindowResizer) AfterAttachToWindowEvent() {
	// Reserve EnterLeaveAreaIDs.
	c.leftAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.rightAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.topAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.bottomAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})

	c.leftTopAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.rightTopAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.leftBottomAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
	c.rightBottomAreaID = c.Window().AddEnterLeaveOverlappingArea(c, RectangleF64{})
}

func (c *WindowResizer) BeforeDetachFromWindowEvent() {
	// Free EnterLeaveAreas.
	c.Window().RemoveEnterLeaveArea(c.leftAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.rightAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.topAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.bottomAreaID, false)

	c.Window().RemoveEnterLeaveArea(c.leftTopAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.rightTopAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.leftBottomAreaID, false)
	c.Window().RemoveEnterLeaveArea(c.rightBottomAreaID, false)
}

func (c WindowResizer) areas() (l, r, t, b, lt, rt, lb, rb RectangleF64) {
	origin := c.Geometry()

	l.Left = origin.Left
	l.Right = origin.Left + windowResizerWidth
	l.Top = origin.Top + windowResizerCornerHeight
	l.Bottom = origin.Bottom - windowResizerCornerHeight

	r.Left = origin.Right - windowResizerWidth
	r.Right = origin.Right
	r.Top = origin.Top + windowResizerCornerHeight
	r.Bottom = origin.Bottom - windowResizerCornerHeight

	t.Left = origin.Left + windowResizerCornerWidth
	t.Right = origin.Right - windowResizerCornerWidth
	t.Top = origin.Top
	t.Bottom = origin.Top + windowResizerHeight

	b.Left = origin.Left + windowResizerCornerWidth
	b.Right = origin.Right - windowResizerCornerWidth
	b.Top = origin.Bottom - windowResizerHeight
	b.Bottom = origin.Bottom

	lt.Left = origin.Left
	lt.Right = origin.Left + windowResizerCornerWidth
	lt.Top = origin.Top
	lt.Bottom = origin.Top + windowResizerCornerHeight

	rt.Left = origin.Right - windowResizerCornerWidth
	rt.Right = origin.Right
	rt.Top = origin.Top
	rt.Bottom = origin.Top + windowResizerCornerHeight

	lb.Left = origin.Left
	lb.Right = origin.Left + windowResizerCornerWidth
	lb.Top = origin.Bottom - windowResizerCornerHeight
	lb.Bottom = origin.Bottom

	rb.Left = origin.Right - windowResizerCornerWidth
	rb.Right = origin.Right
	rb.Top = origin.Bottom - windowResizerCornerHeight
	rb.Bottom = origin.Bottom

	return
}

func (c *WindowResizer) OnGeometryChangeEvent() {
	// Update EnterLeaveAreas.
	l, r, t, b, lt, rt, lb, rb := c.areas()
	c.Window().ReplaceEnterLeaveArea(c.leftAreaID, l)
	c.Window().ReplaceEnterLeaveArea(c.rightAreaID, r)
	c.Window().ReplaceEnterLeaveArea(c.topAreaID, t)
	c.Window().ReplaceEnterLeaveArea(c.bottomAreaID, b)

	c.Window().ReplaceEnterLeaveArea(c.leftTopAreaID, lt)
	c.Window().ReplaceEnterLeaveArea(c.rightTopAreaID, rt)
	c.Window().ReplaceEnterLeaveArea(c.leftBottomAreaID, lb)
	c.Window().ReplaceEnterLeaveArea(c.rightBottomAreaID, rb)
}

func (c *WindowResizer) OnPointerEnterLeaveEvent(e PointerEnterLeaveEvent) {
	c.ready = e.Enter
	if !e.Enter {
		c.Window().SetCursor(Cursor(0).MakeDefault()) // TODO Parent Controls/areas must do this
		return
	}
	switch e.ID {
	case c.leftAreaID:
		c.resizeHor = resizeNegative
		c.resizeVer = resizeNone
		c.Window().SetCursor(Cursor(0).MakeResizeLeftRight())
	case c.rightAreaID:
		c.resizeHor = resizePositive
		c.resizeVer = resizeNone
		c.Window().SetCursor(Cursor(0).MakeResizeLeftRight())
	case c.topAreaID:
		c.resizeHor = resizeNone
		c.resizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeTopBottom())
	case c.bottomAreaID:
		c.resizeHor = resizeNone
		c.resizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeTopBottom())
	case c.leftTopAreaID:
		c.resizeHor = resizeNegative
		c.resizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeLeftTopRightBottom())
	case c.rightBottomAreaID:
		c.resizeHor = resizePositive
		c.resizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeLeftTopRightBottom())
	case c.leftBottomAreaID:
		c.resizeHor = resizeNegative
		c.resizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeLeftBottomRightTop())
	case c.rightTopAreaID:
		c.resizeHor = resizePositive
		c.resizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeLeftBottomRightTop())
	}
}

func (c *WindowResizer) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	processed = c.ready
	if !processed {
		return
	}
	if event.Kind.IsPress() {
		//c.baseGeometry = c.Window().WindowGeometry()
		c.baseSize = c.Window().WindowSize()
	}
	return
}

func (c *WindowResizer) OnPointerDragEvent(event PointerDragEvent) {
	size := c.baseSize

	switch c.resizeHor {
	case resizeNegative:
		size.X -= event.Delta.X
	case resizePositive:
		size.X += event.Delta.X
	}
	switch c.resizeVer {
	case resizeNegative:
		size.Y -= event.Delta.Y
	case resizePositive:
		size.Y += event.Delta.Y
	}

	c.Window().SetWindowSize(size, c.resizeHor == resizeNegative, c.resizeVer == resizeNegative)

	/*geometry := c.baseGeometry
	switch c.resizeHor {
	case resizeNegative:
		geometry.Left += event.Delta.X
	case resizePositive:
		geometry.Right += event.Delta.X
	}
	switch c.resizeVer {
	case resizeNegative:
		geometry.Top += event.Delta.Y
	case resizePositive:
		geometry.Bottom += event.Delta.Y
	}
	c.Window().SetWindowGeometry(geometry)*/
}

func NewWindowResizer() *WindowResizer {
	return &WindowResizer{}
}
