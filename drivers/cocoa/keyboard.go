// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <AppKit/AppKit.h>
*/
import "C"
import (
	"github.com/apaxa-go/gui/drivers"
	"github.com/apaxa-go/helper/mathh"
)

const (
	keycodeAnsiA              = 0x00
	keycodeAnsiS              = 0x01
	keycodeAnsiD              = 0x02
	keycodeAnsiF              = 0x03
	keycodeAnsiH              = 0x04
	keycodeAnsiG              = 0x05
	keycodeAnsiZ              = 0x06
	keycodeAnsiX              = 0x07
	keycodeAnsiC              = 0x08
	keycodeAnsiV              = 0x09
	keycodeIsoSection         = 0x0A
	keycodeAnsiB              = 0x0B
	keycodeAnsiQ              = 0x0C
	keycodeAnsiW              = 0x0D
	keycodeAnsiE              = 0x0E
	keycodeAnsiR              = 0x0F
	keycodeAnsiY              = 0x10
	keycodeAnsiT              = 0x11
	keycodeAnsi1              = 0x12
	keycodeAnsi2              = 0x13
	keycodeAnsi3              = 0x14
	keycodeAnsi4              = 0x15
	keycodeAnsi6              = 0x16
	keycodeAnsi5              = 0x17
	keycodeAnsiEqual          = 0x18
	keycodeAnsi9              = 0x19
	keycodeAnsi7              = 0x1A
	keycodeAnsiMinus          = 0x1B
	keycodeAnsi8              = 0x1C
	keycodeAnsi0              = 0x1D
	keycodeAnsiRightBracket   = 0x1E
	keycodeAnsiO              = 0x1F
	keycodeAnsiU              = 0x20
	keycodeAnsiLeftBracket    = 0x21
	keycodeAnsiI              = 0x22
	keycodeAnsiP              = 0x23
	keycodeReturn             = 0x24
	keycodeAnsiL              = 0x25
	keycodeAnsiJ              = 0x26
	keycodeAnsiQuote          = 0x27
	keycodeAnsiK              = 0x28
	keycodeAnsiSemicolon      = 0x29
	keycodeAnsiBackslash      = 0x2A
	keycodeAnsiComma          = 0x2B
	keycodeAnsiSlash          = 0x2C
	keycodeAnsiN              = 0x2D
	keycodeAnsiM              = 0x2E
	keycodeAnsiPeriod         = 0x2F
	keycodeTab                = 0x30
	keycodeSpace              = 0x31
	keycodeAnsiGrave          = 0x32
	keycodeDelete             = 0x33 // 0x34 - skipped
	keycodeEscape             = 0x35
	keycodeRightCommand       = 0x36
	keycodeCommand            = 0x37
	keycodeShift              = 0x38
	keycodeCapsLock           = 0x39
	keycodeOption             = 0x3A
	keycodeControl            = 0x3B
	keycodeRightShift         = 0x3C
	keycodeRightOption        = 0x3D
	keycodeRightControl       = 0x3E
	keycodeFunction           = 0x3F
	keycodeF17                = 0x40
	keycodeAnsiKeypadDecimal  = 0x41 // 0x42 - skipped
	keycodeAnsiKeypadMultiply = 0x43 // 0x44 - skipped
	keycodeAnsiKeypadPlus     = 0x45 // 0x46 - skipped
	keycodeAnsiKeypadClear    = 0x47
	keycodeVolumeUp           = 0x48
	keycodeVolumeDown         = 0x49
	keycodeMute               = 0x4A
	keycodeAnsiKeypadDivide   = 0x4B
	keycodeAnsiKeypadEnter    = 0x4C // 0x4D - skipped
	keycodeAnsiKeypadMinus    = 0x4E
	keycodeF18                = 0x4F
	keycodeF19                = 0x50
	keycodeAnsiKeypadEquals   = 0x51
	keycodeAnsiKeypad0        = 0x52
	keycodeAnsiKeypad1        = 0x53
	keycodeAnsiKeypad2        = 0x54
	keycodeAnsiKeypad3        = 0x55
	keycodeAnsiKeypad4        = 0x56
	keycodeAnsiKeypad5        = 0x57
	keycodeAnsiKeypad6        = 0x58
	keycodeAnsiKeypad7        = 0x59
	keycodeF20                = 0x5A
	keycodeAnsiKeypad8        = 0x5B
	keycodeAnsiKeypad9        = 0x5C
	keycodeJisYen             = 0x5D
	keycodeJisUnderscore      = 0x5E
	keycodeJisKeypadComma     = 0x5F
	keycodeF5                 = 0x60
	keycodeF6                 = 0x61
	keycodeF7                 = 0x62
	keycodeF3                 = 0x63
	keycodeF8                 = 0x64
	keycodeF9                 = 0x65
	keycodeJisEisu            = 0x66
	keycodeF11                = 0x67
	keycodeJisKana            = 0x68
	keycodeF13                = 0x69
	keycodeF16                = 0x6A
	keycodeF14                = 0x6B // 0x6C - skipped
	keycodeF10                = 0x6D // 0x6E - skipped
	keycodeF12                = 0x6F // 0x70 - skipped
	keycodeF15                = 0x71
	keycodeHelp               = 0x72
	keycodeHome               = 0x73
	keycodePageUp             = 0x74
	keycodeForwardDelete      = 0x75
	keycodeF4                 = 0x76
	keycodeEnd                = 0x77
	keycodeF2                 = 0x78
	keycodePageDown           = 0x79
	keycodeF1                 = 0x7A
	keycodeLeftArrow          = 0x7B
	keycodeRightArrow         = 0x7C
	keycodeDownArrow          = 0x7D
	keycodeUpArrow            = 0x7E

	keycodeKnownBlockLen = 0x7F // maximum known keycode + 1
)

var translateKeyMap = [keycodeKnownBlockLen]Key{
	keycodeAnsiA:              drivers.KeyA,
	keycodeAnsiS:              drivers.KeyS,
	keycodeAnsiD:              drivers.KeyD,
	keycodeAnsiF:              drivers.KeyF,
	keycodeAnsiH:              drivers.KeyH,
	keycodeAnsiG:              drivers.KeyG,
	keycodeAnsiZ:              drivers.KeyZ,
	keycodeAnsiX:              drivers.KeyX,
	keycodeAnsiC:              drivers.KeyC,
	keycodeAnsiV:              drivers.KeyV,
	keycodeAnsiB:              drivers.KeyB,
	keycodeAnsiQ:              drivers.KeyQ,
	keycodeAnsiW:              drivers.KeyW,
	keycodeAnsiE:              drivers.KeyE,
	keycodeAnsiR:              drivers.KeyR,
	keycodeAnsiY:              drivers.KeyY,
	keycodeAnsiT:              drivers.KeyT,
	keycodeAnsi1:              drivers.Key1,
	keycodeAnsi2:              drivers.Key2,
	keycodeAnsi3:              drivers.Key3,
	keycodeAnsi4:              drivers.Key4,
	keycodeAnsi6:              drivers.Key6,
	keycodeAnsi5:              drivers.Key5,
	keycodeAnsiEqual:          drivers.KeyEqual,
	keycodeAnsi9:              drivers.Key9,
	keycodeAnsi7:              drivers.Key7,
	keycodeAnsiMinus:          drivers.KeyMinus,
	keycodeAnsi8:              drivers.Key8,
	keycodeAnsi0:              drivers.Key0,
	keycodeAnsiRightBracket:   drivers.KeyRightBracket,
	keycodeAnsiO:              drivers.KeyO,
	keycodeAnsiU:              drivers.KeyU,
	keycodeAnsiLeftBracket:    drivers.KeyLeftBracket,
	keycodeAnsiI:              drivers.KeyI,
	keycodeAnsiP:              drivers.KeyP,
	keycodeReturn:             drivers.KeyEnter,
	keycodeAnsiL:              drivers.KeyL,
	keycodeAnsiJ:              drivers.KeyJ,
	keycodeAnsiQuote:          drivers.KeyQuote,
	keycodeAnsiK:              drivers.KeyK,
	keycodeAnsiSemicolon:      drivers.KeySemicolon,
	keycodeAnsiBackslash:      drivers.KeyBackslash,
	keycodeAnsiComma:          drivers.KeyComma,
	keycodeAnsiSlash:          drivers.KeySlash,
	keycodeAnsiN:              drivers.KeyN,
	keycodeAnsiM:              drivers.KeyM,
	keycodeAnsiPeriod:         drivers.KeyPeriod,
	keycodeTab:                drivers.KeyTab,
	keycodeSpace:              drivers.KeySpace,
	keycodeAnsiGrave:          drivers.KeyGrave,
	keycodeDelete:             drivers.KeyBackspace,
	keycodeEscape:             drivers.KeyEscape,
	keycodeRightCommand:       drivers.KeyRightCommand,
	keycodeCommand:            drivers.KeyLeftCommand,
	keycodeShift:              drivers.KeyLeftShift,
	keycodeCapsLock:           drivers.KeyCapsLock,
	keycodeOption:             drivers.KeyLeftOption,
	keycodeControl:            drivers.KeyLeftControl,
	keycodeRightShift:         drivers.KeyRightShift,
	keycodeRightOption:        drivers.KeyRightOption,
	keycodeRightControl:       drivers.KeyRightControl,
	keycodeF17:                drivers.KeyF17,
	keycodeAnsiKeypadDecimal:  drivers.KeyNumpadDecimal,
	keycodeAnsiKeypadMultiply: drivers.KeyNumpadMultiply,
	keycodeAnsiKeypadPlus:     drivers.KeyNumpadPlus,
	keycodeVolumeUp:           drivers.KeyVolumeUp,
	keycodeVolumeDown:         drivers.KeyVolumeDown,
	keycodeMute:               drivers.KeyMute,
	keycodeAnsiKeypadDivide:   drivers.KeyNumpadDivide,
	keycodeAnsiKeypadEnter:    drivers.KeyNumpadEnter,
	keycodeAnsiKeypadMinus:    drivers.KeyNumpadMinus,
	keycodeF18:                drivers.KeyF18,
	keycodeF19:                drivers.KeyF19,
	keycodeAnsiKeypad0:        drivers.KeyNumpad0,
	keycodeAnsiKeypad1:        drivers.KeyNumpad1,
	keycodeAnsiKeypad2:        drivers.KeyNumpad2,
	keycodeAnsiKeypad3:        drivers.KeyNumpad3,
	keycodeAnsiKeypad4:        drivers.KeyNumpad4,
	keycodeAnsiKeypad5:        drivers.KeyNumpad5,
	keycodeAnsiKeypad6:        drivers.KeyNumpad6,
	keycodeAnsiKeypad7:        drivers.KeyNumpad7,
	keycodeF20:                drivers.KeyF20,
	keycodeAnsiKeypad8:        drivers.KeyNumpad8,
	keycodeAnsiKeypad9:        drivers.KeyNumpad9,
	keycodeF5:                 drivers.KeyF5,
	keycodeF6:                 drivers.KeyF6,
	keycodeF7:                 drivers.KeyF7,
	keycodeF3:                 drivers.KeyF3,
	keycodeF8:                 drivers.KeyF8,
	keycodeF9:                 drivers.KeyF9,
	keycodeF11:                drivers.KeyF11,
	keycodeF13:                drivers.KeyF13,
	keycodeF16:                drivers.KeyF16,
	keycodeF14:                drivers.KeyF14,
	keycodeF10:                drivers.KeyF10,
	keycodeF12:                drivers.KeyF12,
	keycodeF15:                drivers.KeyF15,
	keycodeHome:               drivers.KeyHome,
	keycodePageUp:             drivers.KeyPageUp,
	keycodeForwardDelete:      drivers.KeyDelete,
	keycodeF4:                 drivers.KeyF4,
	keycodeEnd:                drivers.KeyEnd,
	keycodeF2:                 drivers.KeyF2,
	keycodePageDown:           drivers.KeyPageDown,
	keycodeF1:                 drivers.KeyF1,
	keycodeLeftArrow:          drivers.KeyLeftArrow,
	keycodeRightArrow:         drivers.KeyRightArrow,
	keycodeDownArrow:          drivers.KeyDownArrow,
	keycodeUpArrow:            drivers.KeyUpArrow,

	//
	// Unknown
	//

	keycodeIsoSection: drivers.KeyFirstUnknown + 0,
	0x34:              drivers.KeyFirstUnknown + 1,
	keycodeFunction:   drivers.KeyFirstUnknown + 2,
	0x42:              drivers.KeyFirstUnknown + 3,
	0x44:              drivers.KeyFirstUnknown + 4,
	0x46:              drivers.KeyFirstUnknown + 5,
	keycodeAnsiKeypadClear: drivers.KeyFirstUnknown + 6,
	0x4D: drivers.KeyFirstUnknown + 7,
	keycodeAnsiKeypadEquals: drivers.KeyFirstUnknown + 8,
	keycodeJisYen:           drivers.KeyFirstUnknown + 9,
	keycodeJisUnderscore:    drivers.KeyFirstUnknown + 10,
	keycodeJisKeypadComma:   drivers.KeyFirstUnknown + 11,
	keycodeJisEisu:          drivers.KeyFirstUnknown + 12,
	keycodeJisKana:          drivers.KeyFirstUnknown + 13,
	0x6C:                    drivers.KeyFirstUnknown + 14,
	0x6E:                    drivers.KeyFirstUnknown + 15,
	0x70:                    drivers.KeyFirstUnknown + 16,
	keycodeHelp:             drivers.KeyFirstUnknown + 17,
}

const keycodeKnownBlockExceptions = 17 // number of unknown keycodes in known block

func translateKey(key uint16) Key {
	if key < keycodeKnownBlockLen {
		return translateKeyMap[key]
	}
	tmp := int(key) - keycodeKnownBlockLen + keycodeKnownBlockExceptions
	tmp = mathh.Min2Int(tmp, drivers.KeyLastUnknown)
	return Key(tmp)
}

const (
	keyModifierShift   = C.NSEventModifierFlagShift
	keyModifierCommand = C.NSEventModifierFlagCommand
	keyModifierControl = C.NSEventModifierFlagControl
	keyModifierOption  = C.NSEventModifierFlagOption

	keyModifierCapsLock   = C.NSEventModifierFlagCapsLock
	keyModifierNumericPad = C.NSEventModifierFlagNumericPad
)

func translateKeyModifiers(mod uint64) KeyModifiers {
	var r KeyModifiers
	if mod&keyModifierShift > 0 {
		r |= drivers.KeyModifierShift
	}
	if mod&keyModifierCommand > 0 {
		r |= drivers.KeyModifierMeta
	}
	if mod&keyModifierControl > 0 {
		r |= drivers.KeyModifierControl
	}
	if mod&keyModifierOption > 0 {
		r |= drivers.KeyModifierAlt
	}
	if mod&keyModifierCapsLock > 0 {
		r |= drivers.KeyModifierCapsLock
	}
	if mod&keyModifierNumericPad > 0 {
		r |= drivers.KeyModifierNumLock
	}
	return r
}
