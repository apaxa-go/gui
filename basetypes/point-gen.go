// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package basetypes

type PointI struct {
	X int
	Y int
}

func (p PointI) Add(point PointI) PointI { return PointI{p.X + point.X, p.Y + point.Y} }
func (p PointI) Mul(k int) PointI        { return PointI{p.X * k, p.Y * k} }
func (p PointI) ToI() PointI             { return p }

type PointI32 struct {
	X int32
	Y int32
}

func (p PointI32) Add(point PointI32) PointI32 { return PointI32{p.X + point.X, p.Y + point.Y} }
func (p PointI32) Mul(k int32) PointI32        { return PointI32{p.X * k, p.Y * k} }
func (p PointI32) ToI32() PointI32             { return p }

type PointF32 struct {
	X float32
	Y float32
}

func (p PointF32) Add(point PointF32) PointF32 { return PointF32{p.X + point.X, p.Y + point.Y} }
func (p PointF32) Mul(k float32) PointF32      { return PointF32{p.X * k, p.Y * k} }
func (p PointF32) ToF32() PointF32             { return p }

type PointF64 struct {
	X float64
	Y float64
}

func (p PointF64) Add(point PointF64) PointF64 { return PointF64{p.X + point.X, p.Y + point.Y} }
func (p PointF64) Mul(k float64) PointF64      { return PointF64{p.X * k, p.Y * k} }
func (p PointF64) ToF64() PointF64             { return p }

func (p PointI64) ToF64() PointF64 { return PointF64{float64(p.X), float64(p.Y)} }

func (p PointI64) ToI() PointI { return PointI{int(p.X), int(p.Y)} }

func (p PointI64) ToI32() PointI32 { return PointI32{int32(p.X), int32(p.Y)} }

func (p PointI32) ToF32() PointF32 { return PointF32{float32(p.X), float32(p.Y)} }

func (p PointI32) ToF64() PointF64 { return PointF64{float64(p.X), float64(p.Y)} }

func (p PointI32) ToI() PointI { return PointI{int(p.X), int(p.Y)} }

func (p PointI32) ToI64() PointI64 { return PointI64{int64(p.X), int64(p.Y)} }

func (p PointI) ToF32() PointF32 { return PointF32{float32(p.X), float32(p.Y)} }

func (p PointI) ToF64() PointF64 { return PointF64{float64(p.X), float64(p.Y)} }

func (p PointI) ToI32() PointI32 { return PointI32{int32(p.X), int32(p.Y)} }

func (p PointI) ToI64() PointI64 { return PointI64{int64(p.X), int64(p.Y)} }

func (p PointF32) ToF64() PointF64 { return PointF64{float64(p.X), float64(p.Y)} }

func (p PointF32) ToI() PointI { return PointI{int(p.X), int(p.Y)} }

func (p PointF32) ToI32() PointI32 { return PointI32{int32(p.X), int32(p.Y)} }

func (p PointF32) ToI64() PointI64 { return PointI64{int64(p.X), int64(p.Y)} }

func (p PointF64) ToF32() PointF32 { return PointF32{float32(p.X), float32(p.Y)} }

func (p PointF64) ToI() PointI { return PointI{int(p.X), int(p.Y)} }

func (p PointF64) ToI32() PointI32 { return PointI32{int32(p.X), int32(p.Y)} }

func (p PointF64) ToI64() PointI64 { return PointI64{int64(p.X), int64(p.Y)} }
