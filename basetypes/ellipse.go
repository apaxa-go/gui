// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old float64	F64
//replacer:new float32	F32

type EllipseF64 struct {
	Point   PointF64
	RadiusX float64
	RadiusY float64
}

func (e EllipseF64) OuterRectangle() RectangleF64 {
	return RectangleF64{e.Point.X - e.RadiusX, e.Point.Y - e.RadiusY, e.Point.X + e.RadiusX, e.Point.Y + e.RadiusY}
}
