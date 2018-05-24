package scvi

// drivers.Canvas implements this Canvas.
type Canvas interface {
	PushTransform()
	PopTransform()
	Translate(pos PointF64)
	ScaleXY(x, y float64)

	DrawLine(point0 PointF64, point1 PointF64, color ColorF64, width float64)
	DrawConnectedLines(points []PointF64, color ColorF64, width float64)
	DrawRectangle(rect RectangleF64, color ColorF64, width float64)
}
