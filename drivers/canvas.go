// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type OfflineCanvas interface {
	TextLineGeometry(text string, font Font) PointF64
}

type Canvas interface {
	OfflineCanvas

	//ResetTransform()
	PushTransform()
	PopTransform()
	GetTransform() TransformF64
	//SetTransform(TransformF64)
	Rotate(angle float64) // In radians. Positive values rotate clockwise and negative values rotate counter-clockwise.
	Scale(x float64)
	ScaleXY(x, y float64)
	Translate(pos PointF64)                    // Move (0,0) to pos.
	Superpose(original, required RectangleF64) // Translate & scale canvas in the way that original rectangle becomes required rectangle.

	ClipToRectangle(region RectangleF64)

	DrawLine(point1, point2 PointF64, color ColorF64, width float64)
	DrawConnectedLines(points []PointF64, color ColorF64, width float64)
	DrawStraightContour(points []PointF64, color ColorF64, width float64)
	DrawRectangle(rectangle RectangleF64, color ColorF64, width float64)
	DrawRoundedRectangle(rectangle RoundedRectangleF64, color ColorF64, width float64)
	DrawRoundedRectangleExtended(rectangle RectangleF64, radiusLT, radiusRT, radiusRB, radiusLB PointF64, color ColorF64, width float64)
	DrawEllipse(ellipse EllipseF64, color ColorF64, width float64)
	DrawCircle(circle CircleF64, color ColorF64, width float64)

	FillStraightContour(points []PointF64, color ColorF64)
	FillRectangle(rectangle RectangleF64, color ColorF64)
	FillRoundedRectangle(rectangle RoundedRectangleF64, color ColorF64)
	FillRoundedRectangleExtended(rectangle RectangleF64, radiusLT, radiusRT, radiusRB, radiusLB PointF64, color ColorF64)
	FillEllipse(ellipse EllipseF64, color ColorF64)
	FillCircle(circle CircleF64, color ColorF64)

	DrawTextLine(text string, font Font, pos PointF64, color ColorF64)
}
