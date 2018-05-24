package basetypes

type AlignHor uint8

const (
	AlignHorCenter  AlignHor = 0
	AlignHorLeft    AlignHor = 1
	AlignHorRight   AlignHor = 2
	AlignHorStretch          = AlignHorLeft | AlignHorRight
)

const (
	alignHorBits = 2
	alignHorMask = 1<<alignHorBits - 1
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

type AlignVer uint8

const (
	AlignVerCenter  AlignVer = 0 << alignHorBits
	AlignVerTop     AlignVer = 1 << alignHorBits
	AlignVerBottom  AlignVer = 2 << alignHorBits
	AlignVerStretch          = AlignVerTop | AlignVerBottom
)

const (
	alignVerBits = 2
	alignVerMask = (1<<alignVerBits - 1) << alignHorBits
)

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
