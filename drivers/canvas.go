package drivers

type OfflineCanvas interface {
	TextLineGeometry(text string, font Font, fontSize float64) PointF64
}

type Canvas interface {
	OfflineCanvas
	DrawLine(point1, point2 PointF64, color ColorF64, width float64)
	DrawConnectedLines(points []PointF64, color ColorF64, width float64)
	DrawRectangle(rectangle RectangleF64, color ColorF64, width float64)
	DrawRoundedRectangle(rectangle RoundedRectangleF64, color ColorF64, width float64)
	DrawRoundedRectangleExtended(rectangle RectangleF64, radiusLT, radiusRT, radiusRB, radiusLB PointF64, color ColorF64, width float64)
	DrawEllipse(ellipse EllipseF64, color ColorF64, width float64)
	DrawCircle(circle CircleF64, color ColorF64, width float64)

	FillRectangle(rectangle RectangleF64, color ColorF64)
	FillRoundedRectangle(rectangle RoundedRectangleF64, color ColorF64)
	FillRoundedRectangleExtended(rectangle RectangleF64, radiusLT, radiusRT, radiusRB, radiusLB PointF64, color ColorF64)
	FillEllipse(ellipse EllipseF64, color ColorF64)
	FillCircle(circle CircleF64, color ColorF64)

	DrawTextLine(text string, font Font, fontSize float64, pos PointF64, color ColorF64)
}
