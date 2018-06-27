// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

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

	Minimize()
	Maximize()

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
	RegisterPointerEnterLeaveCallback(f func(event PointerEnterLeaveEvent))
	RegisterScrollCallback(f func(ScrollEvent))
	RegisterWindowMainCallback(f func(become bool))

	AddEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64)
	ReplaceEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64)
	RemoveEnterLeaveArea(id EnterLeaveAreaID)

	AddMoveArea(id MoveAreaID, area RectangleF64)
	ReplaceMoveArea(id MoveAreaID, area RectangleF64)
	RemoveMoveArea(id MoveAreaID)

	SetCursor(Cursor)
}
