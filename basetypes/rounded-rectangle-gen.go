// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package basetypes

type RoundedRectangleF32 struct {
	Rectangle RectangleF32
	RadiusX   float32
	RadiusY   float32
}

func (r RoundedRectangleF32) ToF64() RoundedRectangleF64 {
	return RoundedRectangleF64{r.Rectangle.ToF64(), float64(r.RadiusX), float64(r.RadiusY)}
}

func (r RoundedRectangleF32) Inset(delta float32) RoundedRectangleF32 { return r.InsetXY(delta, delta) }
func (r RoundedRectangleF32) InsetXY(deltaX, deltaY float32) RoundedRectangleF32 {
	r.Rectangle = r.Rectangle.InsetXY(deltaX, deltaY)
	r.RadiusX -= deltaX
	r.RadiusY -= deltaY
	return r
}
func (r RoundedRectangleF32) Inner(lineWidth float32) RoundedRectangleF32 {
	return r.Inset(lineWidth / 2)
}
