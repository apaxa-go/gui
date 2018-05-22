package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old int64	I64
//replacer:new int		I
//replacer:new int32	I32
//replacer:new float32	F32
//replacer:new float64	F64

type PointI64 struct {
	X int64
	Y int64
}

func (p PointI64) Add(point PointI64) PointI64 { return PointI64{p.X + point.X, p.Y + point.Y} }
func (p PointI64) Mul(k int64) PointI64        { return PointI64{p.X * k, p.Y * k} }
func (p PointI64) ToI64() PointI64             { return p }

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

func (p PointI64) ToF32() PointF32 { return PointF32{float32(p.X), float32(p.Y)} }
