// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

// Must be implemented by driver
type Window interface {
	//Run()
	Close()

	SetPossibleSize(min, max PointF64)

	Geometry() RectangleF64
	SetGeometry(RectangleF64)
	Pos() PointF64
	SetPos(PointF64)
	Left() float64
	SetLeft(float64)
	Right() float64
	SetRight(float64)
	Top() float64
	SetTop(float64)
	Bottom() float64
	SetBottom(float64)
	Size() PointF64
	SetSize(size PointF64, fixedRight, fixedBottom bool)
	Width() float64
	SetWidth(width float64, fixedRight bool)
	Height() float64
	SetHeight(height float64, fixedBottom bool)

	Title() string
	SetTitle(string)

	DisplayState() WindowDisplayState
	IsMain() bool

	Minimize()
	Deminimize()
	Maximize()
	Demaximize()
	EnterFullScreen()
	ExitFullScreen()

	OfflineCanvas() OfflineCanvas
	InvalidateRegion(region RectangleF64)
	Invalidate()

	RegisterDrawCallback(func(Canvas, RectangleF64))
	RegisterResizeCallback(func(PointF64))
	RegisterOfflineCanvasCallback(func())

	RegisterKeyboardCallback(f func(KeyboardEvent))
	RegisterPointerKeyCallback(f func(PointerButtonEvent))
	RegisterPointerDragCallback(f func(PointerDragEvent))
	RegisterPointerMoveCallback(f func(PointerMoveEvent))
	RegisterPointerEnterLeaveCallback(f func(PointerEnterLeaveEvent))
	RegisterScrollCallback(f func(ScrollEvent))
	RegisterModifiersCallback(f func(ModifiersEvent))
	RegisterWindowMainStateCallback(f func(WindowMainStateEvent))
	RegisterWindowDisplayStateCallback(f func(WindowDisplayStateEvent))

	AddEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64)
	ReplaceEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64)
	RemoveEnterLeaveArea(id EnterLeaveAreaID)

	AddMoveArea(id MoveAreaID, area RectangleF64)
	ReplaceMoveArea(id MoveAreaID, area RectangleF64)
	RemoveMoveArea(id MoveAreaID)

	SetCursor(Cursor)
}

type WindowDisplayState uint8

const (
	NormalWindow                 WindowDisplayState = iota
	MinimizedWindow              WindowDisplayState = iota
	MaximizedWindow              WindowDisplayState = iota
	FullScreenWindow             WindowDisplayState = iota
	WindowDisplayStateUsedValues                    = iota
)

func (WindowDisplayState) MakeNormal() WindowDisplayState     { return NormalWindow }
func (WindowDisplayState) MakeMinimized() WindowDisplayState  { return MinimizedWindow }
func (WindowDisplayState) MakeMaximized() WindowDisplayState  { return MaximizedWindow }
func (WindowDisplayState) MakeFullScreen() WindowDisplayState { return FullScreenWindow }

func (s WindowDisplayState) IsNormal() bool     { return s == NormalWindow }
func (s WindowDisplayState) IsMinimized() bool  { return s == MinimizedWindow }
func (s WindowDisplayState) IsMaximized() bool  { return s == MaximizedWindow }
func (s WindowDisplayState) IsFullScreen() bool { return s == FullScreenWindow }

func (s WindowDisplayState) String() string {
	switch s {
	case NormalWindow:
		return "normal window"
	case MinimizedWindow:
		return "minimized window"
	case MaximizedWindow:
		return "maximized window"
	case FullScreenWindow:
		return "window in full screen mode"
	default:
		return "window in unknown state #" + strconvh.FormatUint8(uint8(s))
	}
}
