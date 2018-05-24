package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

const max8 = 255 // Maximum value for 8-bit color component

//replacer:replace
//replacer:old float32	F32
//replacer:new float64	F64

type ColorF32 struct {
	R float32
	G float32
	B float32
	A float32
}

func MakeColorF32RGBA8(r, g, b, a uint8) ColorF32 {
	var res ColorF32
	res.R = float32(r) / max8
	res.G = float32(g) / max8
	res.B = float32(b) / max8
	res.A = float32(a) / max8
	return res
}

func MakeColorF32RGB8(r, g, b uint8) ColorF32 { return MakeColorF32RGBA8(r, g, b, max8) }

func (ColorF32) MakeFromRGBA8(r, g, b, a uint8) ColorF32 { return MakeColorF32RGBA8(r, g, b, a) }

func (ColorF32) MakeFromRGB8(r, g, b uint8) ColorF32 { return MakeColorF32RGB8(r, g, b) }
