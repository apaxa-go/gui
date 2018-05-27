// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package basetypes

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

const (
	alignHorShift = 0
	alignVerShift = alignHorShift + alignHorBits
)

//replacer:replace
//replacer:old Hor	Ver	Left	Right
//replacer:new Ver	Hor	Top		Bottom

const (
	alignHorBits = 2
	alignHorMask = (1<<alignHorBits - 1) << alignHorShift
)

type AlignHor uint8

const (
	AlignHorCenter  AlignHor = 0 << alignHorShift
	AlignHorLeft    AlignHor = 1 << alignHorShift
	AlignHorRight   AlignHor = 2 << alignHorShift
	AlignHorStretch          = AlignHorLeft | AlignHorRight
)

func (a AlignHor) IsCenter() bool  { return a == AlignHorCenter }
func (a AlignHor) IsLeft() bool    { return a == AlignHorLeft }
func (a AlignHor) IsRight() bool   { return a == AlignHorRight }
func (a AlignHor) IsStretch() bool { return a == AlignHorStretch }

func (a AlignHor) PinnedToLeft() bool  { return a&AlignHorLeft > 0 }
func (a AlignHor) PinnedToRight() bool { return a&AlignHorRight > 0 }

func (a AlignHor) KeepSize() AlignHor {
	if a.IsStretch() {
		return AlignHorCenter
	}
	return a
}

func (a AlignHor) AddVer(b AlignVer) Align { return Align(a) | Align(b) }

//replacer:replace
//replacer:old Hor	Left	left	Right	right	F64	float64
//replacer:new Hor	Left	left	Right	right	F32	float32
//replacer:new Ver	Top		top		Bottom	bottom	F64	float64
//replacer:new Ver	Top		top		Bottom	bottom	F32	float32

func (a AlignHor) ApplyF64(left, right, reqWidth float64) (Left, Right float64) {
	switch a {
	case AlignHorLeft:
		return left, left + reqWidth
	case AlignHorRight:
		return right - reqWidth, right
	case AlignHorStretch:
		return left, right
	default:
		return (left + right - reqWidth) / 2, (left + right + reqWidth) / 2
	}
}

func (a AlignHor) ApplyF64S(left, width, reqWidth float64) (Left, Width float64) {
	switch a {
	case AlignHorLeft:
		return left, reqWidth
	case AlignHorRight:
		return left + width - reqWidth, reqWidth
	case AlignHorStretch:
		return left, width
	default:
		return left + (width-reqWidth)/2, reqWidth
	}
}

//replacer:ignore

type Align uint8

const (
	AlignCenter  = Align(AlignHorCenter) | Align(AlignVerCenter)
	AlignStretch = Align(AlignHorStretch) | Align(AlignVerStretch)

	AlignCenterStretch = Align(AlignHorCenter) | Align(AlignVerStretch)
	AlignStretchCenter = Align(AlignHorStretch) | Align(AlignVerCenter)

	AlignLeftTop     = Align(AlignHorLeft) | Align(AlignVerTop)
	AlignRightTop    = Align(AlignHorRight) | Align(AlignVerTop)
	AlignLeftBottom  = Align(AlignHorLeft) | Align(AlignVerBottom)
	AlignRightBottom = Align(AlignHorRight) | Align(AlignVerBottom)

	AlignLeftCenter   = Align(AlignHorLeft) | Align(AlignVerCenter)
	AlignRightCenter  = Align(AlignHorRight) | Align(AlignVerCenter)
	AlignTopCenter    = Align(AlignHorCenter) | Align(AlignVerTop)
	AlignBottomCenter = Align(AlignHorCenter) | Align(AlignVerBottom)

	AlignLeftStretch   = Align(AlignHorLeft) | Align(AlignVerStretch)
	AlignRightStretch  = Align(AlignHorRight) | Align(AlignVerStretch)
	AlignTopStretch    = Align(AlignHorStretch) | Align(AlignVerTop)
	AlignBottomStretch = Align(AlignHorStretch) | Align(AlignVerBottom)
)

func (a Align) Hor() AlignHor   { return AlignHor(a & alignHorMask) }
func (a Align) Ver() AlignVer   { return AlignVer(a & alignVerMask) }
func (a Align) IsCenter() bool  { return a == AlignCenter }
func (a Align) IsStretch() bool { return a == AlignStretch }

func (a Align) PinnedToLeft() bool   { return a.Hor().PinnedToLeft() }
func (a Align) PinnedToRight() bool  { return a.Hor().PinnedToRight() }
func (a Align) PinnedToTop() bool    { return a.Ver().PinnedToTop() }
func (a Align) PinnedToBottom() bool { return a.Ver().PinnedToBottom() }

func (a Align) KeepSize() Align {
	if a.IsStretch() {
		return AlignCenter
	}
	return a
}

func (a Align) KeepHorSize() Align {
	if a.Hor().IsStretch() {
		return a.Ver().AddHor(AlignHorCenter)
	}
	return a
}

func (a Align) KeepVerSize() Align {
	if a.Ver().IsStretch() {
		return a.Hor().AddVer(AlignVerCenter)
	}
	return a
}

//replacer:replace
//replacer:old F64	float64
//replacer:new F32	float32

func (a Align) ApplyF64(place RectangleF64, reqSize PointF64) RectangleF64 {
	place.Left, place.Right = a.Hor().ApplyF64(place.Left, place.Right, reqSize.X)
	place.Top, place.Bottom = a.Ver().ApplyF64(place.Top, place.Bottom, reqSize.Y)
	return place
}

func (a Align) ApplyF64S(place RectangleF64S, reqSize PointF64) RectangleF64S {
	place.Origin.X, place.Size.X = a.Hor().ApplyF64S(place.Origin.X, place.Size.X, reqSize.X)
	place.Origin.Y, place.Size.Y = a.Ver().ApplyF64S(place.Origin.Y, place.Size.Y, reqSize.Y)
	return place
}
