// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

// Must be implemented by driver
type Window interface {
	//Run()
	Destroy()

	Geometry() RectangleF64
	Pos() PointF64
	Size() PointF64

	SetGeometry(RectangleF64)
	SetPos(PointF64)
	SetSize(PointF64)

	Title() string
	SetTitle(string)

	OfflineCanvas() OfflineCanvas
	InvalidateRegion(region RectangleF64)
	Invalidate()

	RegisterDrawCallback(func(Canvas, RectangleF64))
	RegisterEventCallback(func(Event) bool)
	RegisterResizeCallback(func())
	RegisterOfflineCanvasCallback(func())
}
