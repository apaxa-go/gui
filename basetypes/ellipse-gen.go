//replacer:generated-file

package basetypes

type EllipseF32 struct {
	Point   PointF32
	RadiusX float32
	RadiusY float32
}

func (e EllipseF32) OuterRectangle() RectangleF32 {
	return RectangleF32{e.Point.X - e.RadiusX, e.Point.Y - e.RadiusY, e.Point.X + e.RadiusX, e.Point.Y + e.RadiusY}
}
