// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	I64
//replacer:new int		I
//replacer:new int32	I32
//replacer:new float32	F32
//replacer:new float64	F64

type RectangleI64 struct {
	Left   int64
	Top    int64
	Right  int64
	Bottom int64
}

type RectangleI64S struct {
	Origin PointI64
	Size   PointI64
}

func MakeRectangleI64(left, top, right, bottom int64) RectangleI64 {
	return RectangleI64{left, top, right, bottom}
}
func MakeRectangleI64S(left, top, right, bottom int64) RectangleI64S {
	return RectangleI64S{PointI64{left, top}, PointI64{right - left, bottom - top}}
}
func MakeSizedRectangleI64(origin PointI64, size PointI64) RectangleI64 {
	return RectangleI64{origin.X, origin.Y, origin.X + size.X, origin.Y + size.Y}
}
func MakeSizedRectangleI64S(origin PointI64, size PointI64) RectangleI64S {
	return RectangleI64S{origin, size}
}

func (r RectangleI64) String() string  { return "{" + r.LT().String() + "; " + r.RB().String() + "}" }
func (r RectangleI64S) String() string { return r.Origin.String() + " " + r.Size.String() }

func (r RectangleI64) ToI64S() RectangleI64S  { return MakeSizedRectangleI64S(r.LT(), r.GetSize()) }
func (r RectangleI64S) ToI64() RectangleI64   { return MakeSizedRectangleI64(r.Origin, r.Size) }
func (r RectangleI64) ToI64() RectangleI64    { return r }
func (r RectangleI64S) ToI64S() RectangleI64S { return r }

func (r RectangleI64) Width() int64       { return r.Right - r.Left }
func (r RectangleI64S) Width() int64      { return r.Size.X }
func (r RectangleI64) Height() int64      { return r.Bottom - r.Top }
func (r RectangleI64S) Height() int64     { return r.Size.Y }
func (r RectangleI64) GetSize() PointI64  { return PointI64{r.Width(), r.Height()} }
func (r RectangleI64S) GetSize() PointI64 { return r.Size }
func (r RectangleI64) GetLeft() int64     { return r.Left }
func (r RectangleI64S) GetLeft() int64    { return r.Origin.X }
func (r RectangleI64) GetTop() int64      { return r.Top }
func (r RectangleI64S) GetTop() int64     { return r.Origin.Y }
func (r RectangleI64) GetRight() int64    { return r.Right }
func (r RectangleI64S) GetRight() int64   { return r.Origin.X + r.Size.X }
func (r RectangleI64) GetBottom() int64   { return r.Bottom }
func (r RectangleI64S) GetBottom() int64  { return r.Origin.Y + r.Size.Y }
func (r RectangleI64) LT() PointI64       { return PointI64{r.Left, r.Top} }
func (r RectangleI64S) LT() PointI64      { return r.Origin }
func (r RectangleI64) RT() PointI64       { return PointI64{r.Right, r.Top} }
func (r RectangleI64S) RT() PointI64      { return PointI64{r.Origin.X + r.Size.X, r.Origin.Y} }
func (r RectangleI64) LB() PointI64       { return PointI64{r.Left, r.Bottom} }
func (r RectangleI64S) LB() PointI64      { return PointI64{r.Origin.X, r.Origin.Y + r.Size.Y} }
func (r RectangleI64) RB() PointI64       { return PointI64{r.Right, r.Bottom} }
func (r RectangleI64S) RB() PointI64      { return r.Origin.Add(r.Size) }

func (r RectangleI64) Shift(shift PointI64) RectangleI64 {
	return RectangleI64{r.Left + shift.X, r.Top + shift.Y, r.Right + shift.X, r.Bottom + shift.Y}
}
func (r RectangleI64S) Shift(shift PointI64) RectangleI64S {
	return RectangleI64S{r.Origin.Add(shift), r.Size.Add(shift)}
}

func (r RectangleI64) Contains(point PointI64) bool {
	return point.X >= r.Left &&
		point.X <= r.Right &&
		point.Y >= r.Top &&
		point.Y <= r.Bottom
}

func (r RectangleI64S) Contains(point PointI64) bool {
	return point.X >= r.Origin.X &&
		point.X <= r.GetRight() &&
		point.Y >= r.Origin.Y &&
		point.Y <= r.GetBottom()
}

//replacer:replace
//replacer:old I64	F32	float32
//replacer:new I64	F64	float64
//replacer:new I64	I	int
//replacer:new I64	I32	int32
//replacer:new I32	F32	float32
//replacer:new I32	F64	float64
//replacer:new I32	I	int
//replacer:new I32	I64	int64
//replacer:new I	F32	float32
//replacer:new I	F64	float64
//replacer:new I	I32	int32
//replacer:new I	I64	int64
//replacer:new F32	F64	float64
//replacer:new F32	I	int
//replacer:new F32	I32	int32
//replacer:new F32	I64	int64
//replacer:new F64	F32	float32
//replacer:new F64	I	int
//replacer:new F64	I32	int32
//replacer:new F64	I64	int64

func (r RectangleI64) ToF32() RectangleF32 {
	return RectangleF32{float32(r.Left), float32(r.Top), float32(r.Right), float32(r.Bottom)}
}
func (r RectangleI64S) ToF32S() RectangleF32S { return RectangleF32S{r.Origin.ToF32(), r.Size.ToF32()} }
func (r RectangleI64) ToF32S() RectangleF32S  { return r.ToI64S().ToF32S() }
func (r RectangleI64S) ToF32() RectangleF32   { return r.ToI64().ToF32() }

//replacer:replace
//replacer:old F64	float64
//replacer:new F32	float32

func (r RectangleF64) ToRounded(radius float64) RoundedRectangleF64 {
	return r.ToRoundedXY(radius, radius)
}
func (r RectangleF64S) ToRounded(radius float64) RoundedRectangleF64 {
	return r.ToRoundedXY(radius, radius)
}
func (r RectangleF64) ToRoundedXY(radiusX, radiusY float64) RoundedRectangleF64 {
	return RoundedRectangleF64{r, radiusX, radiusY}
}
func (r RectangleF64S) ToRoundedXY(radiusX, radiusY float64) RoundedRectangleF64 {
	return r.ToF64().ToRoundedXY(radiusX, radiusY)
}

func (r RectangleF64) Inset(delta float64) RectangleF64   { return r.InsetXY(delta, delta) }
func (r RectangleF64S) Inset(delta float64) RectangleF64S { return r.InsetXY(delta, delta) }
func (r RectangleF64) InsetXY(deltaX, deltaY float64) RectangleF64 {
	r.Left += deltaX
	r.Top += deltaY
	r.Right -= deltaX
	r.Bottom -= deltaX
	return r
}
func (r RectangleF64S) InsetXY(deltaX, deltaY float64) RectangleF64S {
	return RectangleF64S{r.Origin.Add(PointF64{deltaX, deltaY}), r.Size.Add(PointF64{deltaX, deltaY}.Mul(-2))}
}
func (r RectangleF64) Inner(lineWidth float64) RectangleF64   { return r.Inset(lineWidth / 2) }
func (r RectangleF64S) Inner(lineWidth float64) RectangleF64S { return r.Inset(lineWidth / 2) }

func (r RectangleF64) Center() PointF64 {
	return PointF64{(r.Left + r.Right) / 2, (r.Top + r.Bottom) / 2}
}
func (r RectangleF64S) Center() PointF64 {
	return PointF64{r.Origin.X + r.Size.X/2, r.Origin.Y + r.Size.Y/2}
}
