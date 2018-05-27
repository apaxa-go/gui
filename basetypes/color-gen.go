// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package basetypes

type ColorF64 struct {
	R float64
	G float64
	B float64
	A float64
}

func MakeColorF64RGBA8(r, g, b, a uint8) ColorF64 {
	var res ColorF64
	res.R = float64(r) / max8
	res.G = float64(g) / max8
	res.B = float64(b) / max8
	res.A = float64(a) / max8
	return res
}

func MakeColorF64RGB8(r, g, b uint8) ColorF64 { return MakeColorF64RGBA8(r, g, b, max8) }

func (ColorF64) MakeFromRGBA8(r, g, b, a uint8) ColorF64 { return MakeColorF64RGBA8(r, g, b, a) }

func (ColorF64) MakeFromRGB8(r, g, b uint8) ColorF64 { return MakeColorF64RGB8(r, g, b) }
