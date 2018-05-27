// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old float64	F64
//replacer:new float32	F32

type CircleF64 struct {
	Point  PointF64
	Radius float64
}

func (c CircleF64) ToEllipse() EllipseF64 { return EllipseF64{c.Point, c.Radius, c.Radius} }
func (c CircleF64) OuterRectangle() RectangleF64 {
	return RectangleF64{c.Point.X - c.Radius, c.Point.Y - c.Radius, c.Point.X + c.Radius, c.Point.Y + c.Radius}
}

func (c CircleF64) Inset(delta float64) CircleF64 {
	c.Radius -= delta
	return c
}
func (c CircleF64) InsetXY(deltaX, deltaY float64) EllipseF64 {
	return EllipseF64{c.Point, c.Radius - deltaX, c.Radius - deltaY}
}
func (c CircleF64) Inner(lineWidth float64) CircleF64 { return c.Inset(lineWidth / 2) }
