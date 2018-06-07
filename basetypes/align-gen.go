// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package basetypes

const (
	alignVerBits = 2
	alignVerMask = (1<<alignVerBits - 1) << alignVerShift
)

type AlignVer uint8

const (
	AlignVerCenter  AlignVer = 0 << alignVerShift
	AlignVerTop     AlignVer = 1 << alignVerShift
	AlignVerBottom  AlignVer = 2 << alignVerShift
	AlignVerStretch          = AlignVerTop | AlignVerBottom
)

func (_ AlignVer) MakeCenter() AlignVer  { return AlignVerCenter }
func (_ AlignVer) MakeTop() AlignVer     { return AlignVerTop }
func (_ AlignVer) MakeBottom() AlignVer  { return AlignVerBottom }
func (_ AlignVer) MakeStretch() AlignVer { return AlignVerStretch }

func (a AlignVer) IsCenter() bool  { return a == AlignVerCenter }
func (a AlignVer) IsTop() bool     { return a == AlignVerTop }
func (a AlignVer) IsBottom() bool  { return a == AlignVerBottom }
func (a AlignVer) IsStretch() bool { return a == AlignVerStretch }

func (a AlignVer) PinnedToTop() bool    { return a&AlignVerTop > 0 }
func (a AlignVer) PinnedToBottom() bool { return a&AlignVerBottom > 0 }

func (a AlignVer) KeepSize() AlignVer {
	if a.IsStretch() {
		return AlignVerCenter
	}
	return a
}

func (a AlignVer) AddHor(b AlignHor) Align { return Align(a) | Align(b) }

func (a AlignHor) ApplyF32(left, right, reqWidth float32) (Left, Right float32) {
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

func (a AlignHor) ApplyF32S(left, width, reqWidth float32) (Left, Width float32) {
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

func (a AlignVer) ApplyF64(top, bottom, reqWidth float64) (Top, Bottom float64) {
	switch a {
	case AlignVerTop:
		return top, top + reqWidth
	case AlignVerBottom:
		return bottom - reqWidth, bottom
	case AlignVerStretch:
		return top, bottom
	default:
		return (top + bottom - reqWidth) / 2, (top + bottom + reqWidth) / 2
	}
}

func (a AlignVer) ApplyF64S(top, width, reqWidth float64) (Top, Width float64) {
	switch a {
	case AlignVerTop:
		return top, reqWidth
	case AlignVerBottom:
		return top + width - reqWidth, reqWidth
	case AlignVerStretch:
		return top, width
	default:
		return top + (width-reqWidth)/2, reqWidth
	}
}

func (a AlignVer) ApplyF32(top, bottom, reqWidth float32) (Top, Bottom float32) {
	switch a {
	case AlignVerTop:
		return top, top + reqWidth
	case AlignVerBottom:
		return bottom - reqWidth, bottom
	case AlignVerStretch:
		return top, bottom
	default:
		return (top + bottom - reqWidth) / 2, (top + bottom + reqWidth) / 2
	}
}

func (a AlignVer) ApplyF32S(top, width, reqWidth float32) (Top, Width float32) {
	switch a {
	case AlignVerTop:
		return top, reqWidth
	case AlignVerBottom:
		return top + width - reqWidth, reqWidth
	case AlignVerStretch:
		return top, width
	default:
		return top + (width-reqWidth)/2, reqWidth
	}
}

func (a Align) ApplyF32(place RectangleF32, reqSize PointF32) RectangleF32 {
	place.Left, place.Right = a.Hor().ApplyF32(place.Left, place.Right, reqSize.X)
	place.Top, place.Bottom = a.Ver().ApplyF32(place.Top, place.Bottom, reqSize.Y)
	return place
}

func (a Align) ApplyF32S(place RectangleF32S, reqSize PointF32) RectangleF32S {
	place.Origin.X, place.Size.X = a.Hor().ApplyF32S(place.Origin.X, place.Size.X, reqSize.X)
	place.Origin.Y, place.Size.Y = a.Ver().ApplyF32S(place.Origin.Y, place.Size.Y, reqSize.Y)
	return place
}
