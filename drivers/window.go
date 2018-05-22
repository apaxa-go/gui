package drivers

// Must be implemented by driver
type Window interface {
	//Run()
	Destroy()

	Geometry() RectangleI
	Pos() PointI
	Size() PointI

	SetGeometry(RectangleI)
	SetPos(PointI)
	SetSize(PointI)

	Title() string
	SetTitle(string)

	OfflineCanvas() OfflineCanvas
	InvalidateRegion(region RectangleI)
	Invalidate()

	RegisterDrawCallback(func(Canvas, RectangleI))
	RegisterEventCallback(func(Event) bool)
	RegisterResizeCallback(func())
	RegisterOfflineCanvasCallback(func())
}
