package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old float64	F64
//replacer:new float32	F32

type RoundedRectangleF64 struct {
	Rectangle RectangleF64
	RadiusX   float64
	RadiusY   float64
}

//replacer:replace
//replacer:old F32	F64 float32 float64
//replacer:new F64	F32 float64 float32

func (r RoundedRectangleF64) ToF32() RoundedRectangleF32 {
	return RoundedRectangleF32{r.Rectangle.ToF32(), float32(r.RadiusX), float32(r.RadiusY)}
}

func (r RoundedRectangleF64) Inset(delta float64) RoundedRectangleF64 { return r.InsetXY(delta, delta) }
func (r RoundedRectangleF64) InsetXY(deltaX, deltaY float64) RoundedRectangleF64 {
	r.Rectangle = r.Rectangle.InsetXY(deltaX, deltaY)
	r.RadiusX -= deltaX
	r.RadiusY -= deltaY
	return r
}
func (r RoundedRectangleF64) Inner(lineWidth float64) RoundedRectangleF64 {
	return r.Inset(lineWidth / 2)
}
