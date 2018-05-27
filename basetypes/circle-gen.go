// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package basetypes

type CircleF32 struct {
	Point  PointF32
	Radius float32
}

func (c CircleF32) ToEllipse() EllipseF32 { return EllipseF32{c.Point, c.Radius, c.Radius} }
func (c CircleF32) OuterRectangle() RectangleF32 {
	return RectangleF32{c.Point.X - c.Radius, c.Point.Y - c.Radius, c.Point.X + c.Radius, c.Point.Y + c.Radius}
}

func (c CircleF32) Inset(delta float32) CircleF32 {
	c.Radius -= delta
	return c
}
func (c CircleF32) InsetXY(deltaX, deltaY float32) EllipseF32 {
	return EllipseF32{c.Point, c.Radius - deltaX, c.Radius - deltaY}
}
func (c CircleF32) Inner(lineWidth float32) CircleF32 { return c.Inset(lineWidth / 2) }
