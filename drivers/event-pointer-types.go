// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import (
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/strconvh"
)

type PointerButton uint8

const (
	MouseButtonLeft PointerButton = iota
	MouseButtonRight
	MouseButtonMiddle
	MouseButton4 // 1-base, not 0-base
	MouseButton5
	MouseButton6
	MouseButton7
	MouseButton8
	MouseButton9
	MouseButton10
	MouseButton11
	MouseButton12
	MouseButton13
	MouseButton14
	MouseButton15
	MouseButton16
	MouseButton17
	MouseButton18
	MouseButton19
	MouseButton20
	// Other mouse buttons also may be used, but we do not define constants for them

	MouseButtonNamedCount = 3
)

var pointerButtonNames = [MouseButtonNamedCount]string{
	"left mouse button",
	"right mouse button",
	"middle mouse button",
}

var pointerButtonShortNames = [...]string{
	"❶",
	"❷",
	"❸",
	"❹",
	"❺",
	"❻",
	"❼",
	"❽",
	"❾",
	"❿",
	"⓫",
	"⓬",
	"⓭",
	"⓮",
	"⓯",
	"⓰",
	"⓱",
	"⓲",
	"⓳",
	"⓴",
}

func (b PointerButton) String() string {
	if b < MouseButtonNamedCount {
		return pointerButtonNames[b]
	}
	return "mouse button " + strconvh.FormatUint16(uint16(b)+1)
}

func (b PointerButton) ShortString() string {
	if b < PointerButton(len(pointerButtonShortNames)) {
		return pointerButtonShortNames[b]
	}
	return strconvh.FormatUint16(uint16(b) + 1)
}

type PointerButtonEventKind uint8

const (
	PointerButtonEventPress       PointerButtonEventKind = 0
	PointerButtonEventClick       PointerButtonEventKind = 1
	PointerButtonEventDoubleClick PointerButtonEventKind = 2
	PointerButtonEventTripleClick PointerButtonEventKind = 3
	PointerButtonEventRelease     PointerButtonEventKind = mathh.MaxUint8
)

func (k PointerButtonEventKind) IsBasic() bool {
	return k == PointerButtonEventPress || k == PointerButtonEventRelease
}
func (k PointerButtonEventKind) ClickCount() uint8 {
	if k.IsBasic() {
		return 0
	}
	return uint8(k)
}
func (k PointerButtonEventKind) String() string {
	switch k {
	case PointerButtonEventPress:
		return "pressed"
	case PointerButtonEventRelease:
		return "released"
	case PointerButtonEventClick:
		return "clicked"
	case PointerButtonEventDoubleClick:
		return "double-clicked"
	case PointerButtonEventTripleClick:
		return "triple-clicked"
	default:
		return strconvh.FormatUint8(uint8(k)) + "-clicked"
	}
}
func (k PointerButtonEventKind) ShortString() string {
	return k.String() // TODO
}
