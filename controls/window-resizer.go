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

	// Block of states.
	ready         bool        // True if cursor over any of area.
	lastResizeHor resizeState // Horizontal resize settings of area under cursor or 0 if cursor outside.
	lastResizeVer resizeState // Vertical resize settings of area under cursor or 0 if cursor outside.
	resizeHor     resizeState // Same as lastResizeHor, but describe current resize and can't be change until resize will be done.
	resizeVer     resizeState // Same as lastResizeVer, but describe current resize and can't be change until resize will be done.

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
		c.Window().SetCursor(Cursor(0).MakeDefault(), false) // TODO Parent Controls/areas must do this
		return
	}
	switch e.ID {
	case c.leftAreaID:
		c.lastResizeHor = resizeNegative
		c.lastResizeVer = resizeNone
		c.Window().SetCursor(Cursor(0).MakeResizeLeftRight(), false)
	case c.rightAreaID:
		c.lastResizeHor = resizePositive
		c.lastResizeVer = resizeNone
		c.Window().SetCursor(Cursor(0).MakeResizeLeftRight(), false)
	case c.topAreaID:
		c.lastResizeHor = resizeNone
		c.lastResizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeTopBottom(), false)
	case c.bottomAreaID:
		c.lastResizeHor = resizeNone
		c.lastResizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeTopBottom(), false)
	case c.leftTopAreaID:
		c.lastResizeHor = resizeNegative
		c.lastResizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeLeftTopRightBottom(), false)
	case c.rightBottomAreaID:
		c.lastResizeHor = resizePositive
		c.lastResizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeLeftTopRightBottom(), false)
	case c.leftBottomAreaID:
		c.lastResizeHor = resizeNegative
		c.lastResizeVer = resizePositive
		c.Window().SetCursor(Cursor(0).MakeResizeLeftBottomRightTop(), false)
	case c.rightTopAreaID:
		c.lastResizeHor = resizePositive
		c.lastResizeVer = resizeNegative
		c.Window().SetCursor(Cursor(0).MakeResizeLeftBottomRightTop(), false)
	}
}

func (c *WindowResizer) OnPointerButtonEvent(event PointerButtonEvent) (processed bool) {
	if event.Kind.IsPress() {
		processed = c.ready && event.Button.IsLeft()
		if processed {
			c.baseSize = c.Window().WindowSize()
			c.resizeVer = c.lastResizeVer
			c.resizeHor = c.lastResizeHor
			c.Window().SetCursor(Cursor(0).MakeResizeLeftBottomRightTop(), true)
		}
	} else if event.Kind.IsRelease() {
		// As we accepts only left mouse button then here we may be sure that here button is left too.
		c.Window().StopCursorOverride()
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

	c.Window().SetWindowSize(size, c.lastResizeHor == resizeNegative, c.lastResizeVer == resizeNegative)
}

func NewWindowResizer() *WindowResizer {
	return &WindowResizer{}
}
