//replacer:generated-file

package basetypes

type RectangleI struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type RectangleIS struct {
	Origin PointI
	Size   PointI
}

func MakeRectangleI(left, top, right, bottom int) RectangleI {
	return RectangleI{left, top, right, bottom}
}
func MakeRectangleIS(left, top, right, bottom int) RectangleIS {
	return RectangleIS{PointI{left, top}, PointI{right - left, bottom - top}}
}
func MakeSizedRectangleI(origin PointI, size PointI) RectangleI {
	return RectangleI{origin.X, origin.Y, origin.X + size.X, origin.Y + size.Y}
}
func MakeSizedRectangleIS(origin PointI, size PointI) RectangleIS {
	return RectangleIS{origin, size}
}

func (r RectangleI) ToIS() RectangleIS  { return MakeSizedRectangleIS(r.LT(), r.GetSize()) }
func (r RectangleIS) ToI() RectangleI   { return MakeSizedRectangleI(r.Origin, r.Size) }
func (r RectangleI) ToI() RectangleI    { return r }
func (r RectangleIS) ToIS() RectangleIS { return r }

func (r RectangleI) Width() int       { return r.Right - r.Left }
func (r RectangleIS) Width() int      { return r.Size.X }
func (r RectangleI) Height() int      { return r.Bottom - r.Top }
func (r RectangleIS) Height() int     { return r.Size.Y }
func (r RectangleI) GetSize() PointI  { return PointI{r.Width(), r.Height()} }
func (r RectangleIS) GetSize() PointI { return r.Size }
func (r RectangleI) GetLeft() int     { return r.Left }
func (r RectangleIS) GetLeft() int    { return r.Origin.X }
func (r RectangleI) GetTop() int      { return r.Top }
func (r RectangleIS) GetTop() int     { return r.Origin.Y }
func (r RectangleI) GetRight() int    { return r.Right }
func (r RectangleIS) GetRight() int   { return r.Origin.X + r.Size.X }
func (r RectangleI) GetBottom() int   { return r.Bottom }
func (r RectangleIS) GetBottom() int  { return r.Origin.Y + r.Size.Y }
func (r RectangleI) LT() PointI       { return PointI{r.Left, r.Top} }
func (r RectangleIS) LT() PointI      { return r.Origin }
func (r RectangleI) RT() PointI       { return PointI{r.Right, r.Top} }
func (r RectangleIS) RT() PointI      { return PointI{r.Origin.X + r.Size.X, r.Origin.Y} }
func (r RectangleI) LB() PointI       { return PointI{r.Left, r.Bottom} }
func (r RectangleIS) LB() PointI      { return PointI{r.Origin.X, r.Origin.Y + r.Size.Y} }
func (r RectangleI) RB() PointI       { return PointI{r.Right, r.Bottom} }
func (r RectangleIS) RB() PointI      { return r.Origin.Add(r.Size) }

func (r RectangleI) Shift(shift PointI) RectangleI {
	return RectangleI{r.Left + shift.X, r.Top + shift.Y, r.Right + shift.X, r.Bottom + shift.Y}
}
func (r RectangleIS) Shift(shift PointI) RectangleIS {
	return RectangleIS{r.Origin.Add(shift), r.Size.Add(shift)}
}

type RectangleI32 struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type RectangleI32S struct {
	Origin PointI32
	Size   PointI32
}

func MakeRectangleI32(left, top, right, bottom int32) RectangleI32 {
	return RectangleI32{left, top, right, bottom}
}
func MakeRectangleI32S(left, top, right, bottom int32) RectangleI32S {
	return RectangleI32S{PointI32{left, top}, PointI32{right - left, bottom - top}}
}
func MakeSizedRectangleI32(origin PointI32, size PointI32) RectangleI32 {
	return RectangleI32{origin.X, origin.Y, origin.X + size.X, origin.Y + size.Y}
}
func MakeSizedRectangleI32S(origin PointI32, size PointI32) RectangleI32S {
	return RectangleI32S{origin, size}
}

func (r RectangleI32) ToI32S() RectangleI32S  { return MakeSizedRectangleI32S(r.LT(), r.GetSize()) }
func (r RectangleI32S) ToI32() RectangleI32   { return MakeSizedRectangleI32(r.Origin, r.Size) }
func (r RectangleI32) ToI32() RectangleI32    { return r }
func (r RectangleI32S) ToI32S() RectangleI32S { return r }

func (r RectangleI32) Width() int32       { return r.Right - r.Left }
func (r RectangleI32S) Width() int32      { return r.Size.X }
func (r RectangleI32) Height() int32      { return r.Bottom - r.Top }
func (r RectangleI32S) Height() int32     { return r.Size.Y }
func (r RectangleI32) GetSize() PointI32  { return PointI32{r.Width(), r.Height()} }
func (r RectangleI32S) GetSize() PointI32 { return r.Size }
func (r RectangleI32) GetLeft() int32     { return r.Left }
func (r RectangleI32S) GetLeft() int32    { return r.Origin.X }
func (r RectangleI32) GetTop() int32      { return r.Top }
func (r RectangleI32S) GetTop() int32     { return r.Origin.Y }
func (r RectangleI32) GetRight() int32    { return r.Right }
func (r RectangleI32S) GetRight() int32   { return r.Origin.X + r.Size.X }
func (r RectangleI32) GetBottom() int32   { return r.Bottom }
func (r RectangleI32S) GetBottom() int32  { return r.Origin.Y + r.Size.Y }
func (r RectangleI32) LT() PointI32       { return PointI32{r.Left, r.Top} }
func (r RectangleI32S) LT() PointI32      { return r.Origin }
func (r RectangleI32) RT() PointI32       { return PointI32{r.Right, r.Top} }
func (r RectangleI32S) RT() PointI32      { return PointI32{r.Origin.X + r.Size.X, r.Origin.Y} }
func (r RectangleI32) LB() PointI32       { return PointI32{r.Left, r.Bottom} }
func (r RectangleI32S) LB() PointI32      { return PointI32{r.Origin.X, r.Origin.Y + r.Size.Y} }
func (r RectangleI32) RB() PointI32       { return PointI32{r.Right, r.Bottom} }
func (r RectangleI32S) RB() PointI32      { return r.Origin.Add(r.Size) }

func (r RectangleI32) Shift(shift PointI32) RectangleI32 {
	return RectangleI32{r.Left + shift.X, r.Top + shift.Y, r.Right + shift.X, r.Bottom + shift.Y}
}
func (r RectangleI32S) Shift(shift PointI32) RectangleI32S {
	return RectangleI32S{r.Origin.Add(shift), r.Size.Add(shift)}
}

type RectangleF32 struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type RectangleF32S struct {
	Origin PointF32
	Size   PointF32
}

func MakeRectangleF32(left, top, right, bottom float32) RectangleF32 {
	return RectangleF32{left, top, right, bottom}
}
func MakeRectangleF32S(left, top, right, bottom float32) RectangleF32S {
	return RectangleF32S{PointF32{left, top}, PointF32{right - left, bottom - top}}
}
func MakeSizedRectangleF32(origin PointF32, size PointF32) RectangleF32 {
	return RectangleF32{origin.X, origin.Y, origin.X + size.X, origin.Y + size.Y}
}
func MakeSizedRectangleF32S(origin PointF32, size PointF32) RectangleF32S {
	return RectangleF32S{origin, size}
}

func (r RectangleF32) ToF32S() RectangleF32S  { return MakeSizedRectangleF32S(r.LT(), r.GetSize()) }
func (r RectangleF32S) ToF32() RectangleF32   { return MakeSizedRectangleF32(r.Origin, r.Size) }
func (r RectangleF32) ToF32() RectangleF32    { return r }
func (r RectangleF32S) ToF32S() RectangleF32S { return r }

func (r RectangleF32) Width() float32      { return r.Right - r.Left }
func (r RectangleF32S) Width() float32     { return r.Size.X }
func (r RectangleF32) Height() float32     { return r.Bottom - r.Top }
func (r RectangleF32S) Height() float32    { return r.Size.Y }
func (r RectangleF32) GetSize() PointF32   { return PointF32{r.Width(), r.Height()} }
func (r RectangleF32S) GetSize() PointF32  { return r.Size }
func (r RectangleF32) GetLeft() float32    { return r.Left }
func (r RectangleF32S) GetLeft() float32   { return r.Origin.X }
func (r RectangleF32) GetTop() float32     { return r.Top }
func (r RectangleF32S) GetTop() float32    { return r.Origin.Y }
func (r RectangleF32) GetRight() float32   { return r.Right }
func (r RectangleF32S) GetRight() float32  { return r.Origin.X + r.Size.X }
func (r RectangleF32) GetBottom() float32  { return r.Bottom }
func (r RectangleF32S) GetBottom() float32 { return r.Origin.Y + r.Size.Y }
func (r RectangleF32) LT() PointF32        { return PointF32{r.Left, r.Top} }
func (r RectangleF32S) LT() PointF32       { return r.Origin }
func (r RectangleF32) RT() PointF32        { return PointF32{r.Right, r.Top} }
func (r RectangleF32S) RT() PointF32       { return PointF32{r.Origin.X + r.Size.X, r.Origin.Y} }
func (r RectangleF32) LB() PointF32        { return PointF32{r.Left, r.Bottom} }
func (r RectangleF32S) LB() PointF32       { return PointF32{r.Origin.X, r.Origin.Y + r.Size.Y} }
func (r RectangleF32) RB() PointF32        { return PointF32{r.Right, r.Bottom} }
func (r RectangleF32S) RB() PointF32       { return r.Origin.Add(r.Size) }

func (r RectangleF32) Shift(shift PointF32) RectangleF32 {
	return RectangleF32{r.Left + shift.X, r.Top + shift.Y, r.Right + shift.X, r.Bottom + shift.Y}
}
func (r RectangleF32S) Shift(shift PointF32) RectangleF32S {
	return RectangleF32S{r.Origin.Add(shift), r.Size.Add(shift)}
}

type RectangleF64 struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

type RectangleF64S struct {
	Origin PointF64
	Size   PointF64
}

func MakeRectangleF64(left, top, right, bottom float64) RectangleF64 {
	return RectangleF64{left, top, right, bottom}
}
func MakeRectangleF64S(left, top, right, bottom float64) RectangleF64S {
	return RectangleF64S{PointF64{left, top}, PointF64{right - left, bottom - top}}
}
func MakeSizedRectangleF64(origin PointF64, size PointF64) RectangleF64 {
	return RectangleF64{origin.X, origin.Y, origin.X + size.X, origin.Y + size.Y}
}
func MakeSizedRectangleF64S(origin PointF64, size PointF64) RectangleF64S {
	return RectangleF64S{origin, size}
}

func (r RectangleF64) ToF64S() RectangleF64S  { return MakeSizedRectangleF64S(r.LT(), r.GetSize()) }
func (r RectangleF64S) ToF64() RectangleF64   { return MakeSizedRectangleF64(r.Origin, r.Size) }
func (r RectangleF64) ToF64() RectangleF64    { return r }
func (r RectangleF64S) ToF64S() RectangleF64S { return r }

func (r RectangleF64) Width() float64      { return r.Right - r.Left }
func (r RectangleF64S) Width() float64     { return r.Size.X }
func (r RectangleF64) Height() float64     { return r.Bottom - r.Top }
func (r RectangleF64S) Height() float64    { return r.Size.Y }
func (r RectangleF64) GetSize() PointF64   { return PointF64{r.Width(), r.Height()} }
func (r RectangleF64S) GetSize() PointF64  { return r.Size }
func (r RectangleF64) GetLeft() float64    { return r.Left }
func (r RectangleF64S) GetLeft() float64   { return r.Origin.X }
func (r RectangleF64) GetTop() float64     { return r.Top }
func (r RectangleF64S) GetTop() float64    { return r.Origin.Y }
func (r RectangleF64) GetRight() float64   { return r.Right }
func (r RectangleF64S) GetRight() float64  { return r.Origin.X + r.Size.X }
func (r RectangleF64) GetBottom() float64  { return r.Bottom }
func (r RectangleF64S) GetBottom() float64 { return r.Origin.Y + r.Size.Y }
func (r RectangleF64) LT() PointF64        { return PointF64{r.Left, r.Top} }
func (r RectangleF64S) LT() PointF64       { return r.Origin }
func (r RectangleF64) RT() PointF64        { return PointF64{r.Right, r.Top} }
func (r RectangleF64S) RT() PointF64       { return PointF64{r.Origin.X + r.Size.X, r.Origin.Y} }
func (r RectangleF64) LB() PointF64        { return PointF64{r.Left, r.Bottom} }
func (r RectangleF64S) LB() PointF64       { return PointF64{r.Origin.X, r.Origin.Y + r.Size.Y} }
func (r RectangleF64) RB() PointF64        { return PointF64{r.Right, r.Bottom} }
func (r RectangleF64S) RB() PointF64       { return r.Origin.Add(r.Size) }

func (r RectangleF64) Shift(shift PointF64) RectangleF64 {
	return RectangleF64{r.Left + shift.X, r.Top + shift.Y, r.Right + shift.X, r.Bottom + shift.Y}
}
func (r RectangleF64S) Shift(shift PointF64) RectangleF64S {
	return RectangleF64S{r.Origin.Add(shift), r.Size.Add(shift)}
}

func (r RectangleI64) ToF64() RectangleF64 {
	return RectangleF64{float64(r.Left), float64(r.Top), float64(r.Right), float64(r.Bottom)}
}
func (r RectangleI64S) ToF64S() RectangleF64S { return RectangleF64S{r.Origin.ToF64(), r.Size.ToF64()} }
func (r RectangleI64) ToF64S() RectangleF64S  { return r.ToI64S().ToF64S() }
func (r RectangleI64S) ToF64() RectangleF64   { return r.ToI64().ToF64() }

func (r RectangleI64) ToI() RectangleI {
	return RectangleI{int(r.Left), int(r.Top), int(r.Right), int(r.Bottom)}
}
func (r RectangleI64S) ToIS() RectangleIS { return RectangleIS{r.Origin.ToI(), r.Size.ToI()} }
func (r RectangleI64) ToIS() RectangleIS  { return r.ToI64S().ToIS() }
func (r RectangleI64S) ToI() RectangleI   { return r.ToI64().ToI() }

func (r RectangleI64) ToI32() RectangleI32 {
	return RectangleI32{int32(r.Left), int32(r.Top), int32(r.Right), int32(r.Bottom)}
}
func (r RectangleI64S) ToI32S() RectangleI32S { return RectangleI32S{r.Origin.ToI32(), r.Size.ToI32()} }
func (r RectangleI64) ToI32S() RectangleI32S  { return r.ToI64S().ToI32S() }
func (r RectangleI64S) ToI32() RectangleI32   { return r.ToI64().ToI32() }

func (r RectangleI32) ToF32() RectangleF32 {
	return RectangleF32{float32(r.Left), float32(r.Top), float32(r.Right), float32(r.Bottom)}
}
func (r RectangleI32S) ToF32S() RectangleF32S { return RectangleF32S{r.Origin.ToF32(), r.Size.ToF32()} }
func (r RectangleI32) ToF32S() RectangleF32S  { return r.ToI32S().ToF32S() }
func (r RectangleI32S) ToF32() RectangleF32   { return r.ToI32().ToF32() }

func (r RectangleI32) ToF64() RectangleF64 {
	return RectangleF64{float64(r.Left), float64(r.Top), float64(r.Right), float64(r.Bottom)}
}
func (r RectangleI32S) ToF64S() RectangleF64S { return RectangleF64S{r.Origin.ToF64(), r.Size.ToF64()} }
func (r RectangleI32) ToF64S() RectangleF64S  { return r.ToI32S().ToF64S() }
func (r RectangleI32S) ToF64() RectangleF64   { return r.ToI32().ToF64() }

func (r RectangleI32) ToI() RectangleI {
	return RectangleI{int(r.Left), int(r.Top), int(r.Right), int(r.Bottom)}
}
func (r RectangleI32S) ToIS() RectangleIS { return RectangleIS{r.Origin.ToI(), r.Size.ToI()} }
func (r RectangleI32) ToIS() RectangleIS  { return r.ToI32S().ToIS() }
func (r RectangleI32S) ToI() RectangleI   { return r.ToI32().ToI() }

func (r RectangleI32) ToI64() RectangleI64 {
	return RectangleI64{int64(r.Left), int64(r.Top), int64(r.Right), int64(r.Bottom)}
}
func (r RectangleI32S) ToI64S() RectangleI64S { return RectangleI64S{r.Origin.ToI64(), r.Size.ToI64()} }
func (r RectangleI32) ToI64S() RectangleI64S  { return r.ToI32S().ToI64S() }
func (r RectangleI32S) ToI64() RectangleI64   { return r.ToI32().ToI64() }

func (r RectangleI) ToF32() RectangleF32 {
	return RectangleF32{float32(r.Left), float32(r.Top), float32(r.Right), float32(r.Bottom)}
}
func (r RectangleIS) ToF32S() RectangleF32S { return RectangleF32S{r.Origin.ToF32(), r.Size.ToF32()} }
func (r RectangleI) ToF32S() RectangleF32S  { return r.ToIS().ToF32S() }
func (r RectangleIS) ToF32() RectangleF32   { return r.ToI().ToF32() }

func (r RectangleI) ToF64() RectangleF64 {
	return RectangleF64{float64(r.Left), float64(r.Top), float64(r.Right), float64(r.Bottom)}
}
func (r RectangleIS) ToF64S() RectangleF64S { return RectangleF64S{r.Origin.ToF64(), r.Size.ToF64()} }
func (r RectangleI) ToF64S() RectangleF64S  { return r.ToIS().ToF64S() }
func (r RectangleIS) ToF64() RectangleF64   { return r.ToI().ToF64() }

func (r RectangleI) ToI32() RectangleI32 {
	return RectangleI32{int32(r.Left), int32(r.Top), int32(r.Right), int32(r.Bottom)}
}
func (r RectangleIS) ToI32S() RectangleI32S { return RectangleI32S{r.Origin.ToI32(), r.Size.ToI32()} }
func (r RectangleI) ToI32S() RectangleI32S  { return r.ToIS().ToI32S() }
func (r RectangleIS) ToI32() RectangleI32   { return r.ToI().ToI32() }

func (r RectangleI) ToI64() RectangleI64 {
	return RectangleI64{int64(r.Left), int64(r.Top), int64(r.Right), int64(r.Bottom)}
}
func (r RectangleIS) ToI64S() RectangleI64S { return RectangleI64S{r.Origin.ToI64(), r.Size.ToI64()} }
func (r RectangleI) ToI64S() RectangleI64S  { return r.ToIS().ToI64S() }
func (r RectangleIS) ToI64() RectangleI64   { return r.ToI().ToI64() }

func (r RectangleF32) ToF64() RectangleF64 {
	return RectangleF64{float64(r.Left), float64(r.Top), float64(r.Right), float64(r.Bottom)}
}
func (r RectangleF32S) ToF64S() RectangleF64S { return RectangleF64S{r.Origin.ToF64(), r.Size.ToF64()} }
func (r RectangleF32) ToF64S() RectangleF64S  { return r.ToF32S().ToF64S() }
func (r RectangleF32S) ToF64() RectangleF64   { return r.ToF32().ToF64() }

func (r RectangleF32) ToI() RectangleI {
	return RectangleI{int(r.Left), int(r.Top), int(r.Right), int(r.Bottom)}
}
func (r RectangleF32S) ToIS() RectangleIS { return RectangleIS{r.Origin.ToI(), r.Size.ToI()} }
func (r RectangleF32) ToIS() RectangleIS  { return r.ToF32S().ToIS() }
func (r RectangleF32S) ToI() RectangleI   { return r.ToF32().ToI() }

func (r RectangleF32) ToI32() RectangleI32 {
	return RectangleI32{int32(r.Left), int32(r.Top), int32(r.Right), int32(r.Bottom)}
}
func (r RectangleF32S) ToI32S() RectangleI32S { return RectangleI32S{r.Origin.ToI32(), r.Size.ToI32()} }
func (r RectangleF32) ToI32S() RectangleI32S  { return r.ToF32S().ToI32S() }
func (r RectangleF32S) ToI32() RectangleI32   { return r.ToF32().ToI32() }

func (r RectangleF32) ToI64() RectangleI64 {
	return RectangleI64{int64(r.Left), int64(r.Top), int64(r.Right), int64(r.Bottom)}
}
func (r RectangleF32S) ToI64S() RectangleI64S { return RectangleI64S{r.Origin.ToI64(), r.Size.ToI64()} }
func (r RectangleF32) ToI64S() RectangleI64S  { return r.ToF32S().ToI64S() }
func (r RectangleF32S) ToI64() RectangleI64   { return r.ToF32().ToI64() }

func (r RectangleF64) ToF32() RectangleF32 {
	return RectangleF32{float32(r.Left), float32(r.Top), float32(r.Right), float32(r.Bottom)}
}
func (r RectangleF64S) ToF32S() RectangleF32S { return RectangleF32S{r.Origin.ToF32(), r.Size.ToF32()} }
func (r RectangleF64) ToF32S() RectangleF32S  { return r.ToF64S().ToF32S() }
func (r RectangleF64S) ToF32() RectangleF32   { return r.ToF64().ToF32() }

func (r RectangleF64) ToI() RectangleI {
	return RectangleI{int(r.Left), int(r.Top), int(r.Right), int(r.Bottom)}
}
func (r RectangleF64S) ToIS() RectangleIS { return RectangleIS{r.Origin.ToI(), r.Size.ToI()} }
func (r RectangleF64) ToIS() RectangleIS  { return r.ToF64S().ToIS() }
func (r RectangleF64S) ToI() RectangleI   { return r.ToF64().ToI() }

func (r RectangleF64) ToI32() RectangleI32 {
	return RectangleI32{int32(r.Left), int32(r.Top), int32(r.Right), int32(r.Bottom)}
}
func (r RectangleF64S) ToI32S() RectangleI32S { return RectangleI32S{r.Origin.ToI32(), r.Size.ToI32()} }
func (r RectangleF64) ToI32S() RectangleI32S  { return r.ToF64S().ToI32S() }
func (r RectangleF64S) ToI32() RectangleI32   { return r.ToF64().ToI32() }

func (r RectangleF64) ToI64() RectangleI64 {
	return RectangleI64{int64(r.Left), int64(r.Top), int64(r.Right), int64(r.Bottom)}
}
func (r RectangleF64S) ToI64S() RectangleI64S { return RectangleI64S{r.Origin.ToI64(), r.Size.ToI64()} }
func (r RectangleF64) ToI64S() RectangleI64S  { return r.ToF64S().ToI64S() }
func (r RectangleF64S) ToI64() RectangleI64   { return r.ToF64().ToI64() }

func (r RectangleF32) ToRounded(radius float32) RoundedRectangleF32 {
	return r.ToRoundedXY(radius, radius)
}
func (r RectangleF32S) ToRounded(radius float32) RoundedRectangleF32 {
	return r.ToRoundedXY(radius, radius)
}
func (r RectangleF32) ToRoundedXY(radiusX, radiusY float32) RoundedRectangleF32 {
	return RoundedRectangleF32{r, radiusX, radiusY}
}
func (r RectangleF32S) ToRoundedXY(radiusX, radiusY float32) RoundedRectangleF32 {
	return r.ToF32().ToRoundedXY(radiusX, radiusY)
}

func (r RectangleF32) Inset(delta float32) RectangleF32   { return r.InsetXY(delta, delta) }
func (r RectangleF32S) Inset(delta float32) RectangleF32S { return r.InsetXY(delta, delta) }
func (r RectangleF32) InsetXY(deltaX, deltaY float32) RectangleF32 {
	r.Left += deltaX
	r.Top += deltaY
	r.Right -= deltaX
	r.Bottom -= deltaX
	return r
}
func (r RectangleF32S) InsetXY(deltaX, deltaY float32) RectangleF32S {
	return RectangleF32S{r.Origin.Add(PointF32{deltaX, deltaY}), r.Size.Add(PointF32{deltaX, deltaY}.Mul(-2))}
}
func (r RectangleF32) Inner(lineWidth float32) RectangleF32   { return r.Inset(lineWidth / 2) }
func (r RectangleF32S) Inner(lineWidth float32) RectangleF32S { return r.Inset(lineWidth / 2) }
