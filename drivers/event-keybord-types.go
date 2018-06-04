// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import (
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/strconvh"
)

// TODO PrintScreen ⎙

type Key uint8

// com
const (
	//
	// Function keys
	//

	KeyF1  Key = iota
	KeyF2  Key = iota
	KeyF3  Key = iota
	KeyF4  Key = iota
	KeyF5  Key = iota
	KeyF6  Key = iota
	KeyF7  Key = iota
	KeyF8  Key = iota
	KeyF9  Key = iota
	KeyF10 Key = iota
	KeyF11 Key = iota
	KeyF12 Key = iota
	KeyF13 Key = iota
	KeyF14 Key = iota
	KeyF15 Key = iota
	KeyF16 Key = iota
	KeyF17 Key = iota
	KeyF18 Key = iota
	KeyF19 Key = iota
	KeyF20 Key = iota
	KeyF21 Key = iota
	KeyF22 Key = iota
	KeyF23 Key = iota
	KeyF24 Key = iota

	//
	// Other (control) keys
	//

	KeyEscape    Key = iota
	KeyTab       Key = iota
	KeyEnter     Key = iota
	KeyBackspace Key = iota
	KeyDelete    Key = iota

	KeyLeftShift    Key = iota
	KeyRightShift   Key = iota
	KeyLeftControl  Key = iota
	KeyRightControl Key = iota
	KeyLeftMeta     Key = iota
	KeyRightMeta    Key = iota
	KeyLeftAlt      Key = iota
	KeyRightAlt     Key = iota

	KeyHome     Key = iota
	KeyEnd      Key = iota
	KeyPageUp   Key = iota
	KeyPageDown Key = iota

	KeyLeftArrow  Key = iota
	KeyRightArrow Key = iota
	KeyUpArrow    Key = iota
	KeyDownArrow  Key = iota

	//KeyHelp Key = iota
	//KeyFunction     Key = iota

	//
	// Lock keys
	//

	KeyCapsLock   Key = iota
	KeyNumLock    Key = iota
	KeyScrollLock Key = iota

	//
	// Multimedia keys
	//

	KeyVolumeUp   Key = iota
	KeyVolumeDown Key = iota
	KeyMute       Key = iota

	//
	// Character keys
	//

	// row 1

	KeyGrave Key = iota // `~
	Key0     Key = iota
	Key1     Key = iota
	Key2     Key = iota
	Key3     Key = iota
	Key4     Key = iota
	Key5     Key = iota
	Key6     Key = iota
	Key7     Key = iota
	Key8     Key = iota
	Key9     Key = iota
	KeyMinus Key = iota // -_
	KeyEqual Key = iota // =+

	// row 2

	KeyQ            Key = iota
	KeyW            Key = iota
	KeyE            Key = iota
	KeyR            Key = iota
	KeyT            Key = iota
	KeyY            Key = iota
	KeyU            Key = iota
	KeyI            Key = iota
	KeyO            Key = iota
	KeyP            Key = iota
	KeyLeftBracket  Key = iota // [{
	KeyRightBracket Key = iota // ]}
	KeyBackslash    Key = iota // \|

	// row 3

	KeyA         Key = iota
	KeyS         Key = iota
	KeyD         Key = iota
	KeyF         Key = iota
	KeyG         Key = iota
	KeyH         Key = iota
	KeyJ         Key = iota
	KeyK         Key = iota
	KeyL         Key = iota
	KeySemicolon Key = iota // ;:
	KeyQuote     Key = iota // '"
	KeyZ         Key = iota
	KeyX         Key = iota
	KeyC         Key = iota
	KeyV         Key = iota
	KeyB         Key = iota
	KeyN         Key = iota
	KeyM         Key = iota
	KeyComma     Key = iota // ,<
	KeyPeriod    Key = iota // .>
	KeySlash     Key = iota // /?

	// row 4

	KeySpace Key = iota

	//
	// Keypad keys
	//

	KeyNumpad0 Key = iota
	KeyNumpad1 Key = iota
	KeyNumpad2 Key = iota
	KeyNumpad3 Key = iota
	KeyNumpad4 Key = iota
	KeyNumpad5 Key = iota
	KeyNumpad6 Key = iota
	KeyNumpad7 Key = iota
	KeyNumpad8 Key = iota
	KeyNumpad9 Key = iota

	KeyNumpadDecimal  Key = iota // .
	KeyNumpadDivide   Key = iota // /
	KeyNumpadMultiply Key = iota // *
	KeyNumpadMinus    Key = iota // -
	KeyNumpadPlus     Key = iota // +
	KeyNumpadEnter    Key = iota // <Enter>
	//KeyNumpadClear    Key = iota
	//KeyNumpadEquals  Key = iota

	//
	// Unknown keys
	//

	KeyKnownCount  = iota
	KeyLastUnknown = mathh.MaxUint8
)

const KeyFirstUnknown Key = KeyKnownCount

// MacOS alternative names.
// TODO this may be wrong
const (
	KeyLeftCommand  = KeyLeftMeta
	KeyRightCommand = KeyRightMeta
	KeyLeftOption   = KeyLeftAlt
	KeyRightOption  = KeyRightAlt
)

var keyNames = [KeyKnownCount]string{
	KeyF1:             "F1",
	KeyF2:             "F2",
	KeyF3:             "F3",
	KeyF4:             "F4",
	KeyF5:             "F5",
	KeyF6:             "F6",
	KeyF7:             "F7",
	KeyF8:             "F8",
	KeyF9:             "F9",
	KeyF10:            "F10",
	KeyF11:            "F11",
	KeyF12:            "F12",
	KeyF13:            "F13",
	KeyF14:            "F14",
	KeyF15:            "F15",
	KeyF16:            "F16",
	KeyF17:            "F17",
	KeyF18:            "F18",
	KeyF19:            "F19",
	KeyF20:            "F20",
	KeyF21:            "F21",
	KeyF22:            "F22",
	KeyF23:            "F23",
	KeyF24:            "F24",
	KeyEscape:         "Escape",
	KeyTab:            "Tab",
	KeyEnter:          "Enter",
	KeyBackspace:      "Backspace",
	KeyDelete:         "Delete",
	KeyLeftShift:      "Left Shift",
	KeyRightShift:     "Right Shift",
	KeyLeftControl:    "Left Control",
	KeyRightControl:   "Right Control",
	KeyLeftMeta:       "Left Meta",
	KeyRightMeta:      "Right Meta",
	KeyLeftAlt:        "Left Alt",
	KeyRightAlt:       "Right Alt",
	KeyHome:           "Home",
	KeyEnd:            "End",
	KeyPageUp:         "PageUp",
	KeyPageDown:       "PageDown",
	KeyLeftArrow:      "Left Arrow",
	KeyRightArrow:     "Right Arrow",
	KeyUpArrow:        "Up Arrow",
	KeyDownArrow:      "Down Arrow",
	KeyCapsLock:       "CapsLock",
	KeyNumLock:        "NumLock",
	KeyScrollLock:     "ScrollLock",
	KeyVolumeUp:       "Volume Up",
	KeyVolumeDown:     "Volume Down",
	KeyMute:           "Mute",
	KeyGrave:          "Grave",
	Key0:              "0",
	Key1:              "1",
	Key2:              "2",
	Key3:              "3",
	Key4:              "4",
	Key5:              "5",
	Key6:              "6",
	Key7:              "7",
	Key8:              "8",
	Key9:              "9",
	KeyMinus:          "Minus",
	KeyEqual:          "Equal",
	KeyQ:              "Q",
	KeyW:              "W",
	KeyE:              "E",
	KeyR:              "R",
	KeyT:              "T",
	KeyY:              "Y",
	KeyU:              "U",
	KeyI:              "I",
	KeyO:              "O",
	KeyP:              "P",
	KeyLeftBracket:    "Left Bracket",
	KeyRightBracket:   "Right Bracket",
	KeyBackslash:      "Backslash",
	KeyA:              "A",
	KeyS:              "S",
	KeyD:              "D",
	KeyF:              "F",
	KeyG:              "G",
	KeyH:              "H",
	KeyJ:              "J",
	KeyK:              "K",
	KeyL:              "L",
	KeySemicolon:      "Semicolon",
	KeyQuote:          "Quote",
	KeyZ:              "Z",
	KeyX:              "X",
	KeyC:              "C",
	KeyV:              "V",
	KeyB:              "B",
	KeyN:              "N",
	KeyM:              "M",
	KeyComma:          "Comma",
	KeyPeriod:         "Period",
	KeySlash:          "Slash",
	KeySpace:          "Space",
	KeyNumpad0:        "Numpad 0",
	KeyNumpad1:        "Numpad 1",
	KeyNumpad2:        "Numpad 2",
	KeyNumpad3:        "Numpad 3",
	KeyNumpad4:        "Numpad 4",
	KeyNumpad5:        "Numpad 5",
	KeyNumpad6:        "Numpad 6",
	KeyNumpad7:        "Numpad 7",
	KeyNumpad8:        "Numpad 8",
	KeyNumpad9:        "Numpad 9",
	KeyNumpadDecimal:  "Numpad Decimal",
	KeyNumpadDivide:   "Numpad Divide",
	KeyNumpadMultiply: "Numpad Multiply",
	KeyNumpadMinus:    "Numpad Minus",
	KeyNumpadPlus:     "Numpad Plus",
	KeyNumpadEnter:    "Numpad Enter",
}

var keyShortNames = [KeyKnownCount]string{
	KeyF1:             "①",
	KeyF2:             "②",
	KeyF3:             "③",
	KeyF4:             "④",
	KeyF5:             "⑤",
	KeyF6:             "⑥",
	KeyF7:             "⑦",
	KeyF8:             "⑧",
	KeyF9:             "⑨",
	KeyF10:            "⑩",
	KeyF11:            "⑪",
	KeyF12:            "⑫",
	KeyF13:            "⑬",
	KeyF14:            "⑭",
	KeyF15:            "⑮",
	KeyF16:            "⑯",
	KeyF17:            "⑰",
	KeyF18:            "⑱",
	KeyF19:            "⑲",
	KeyF20:            "⑳",
	KeyF21:            "㉑",
	KeyF22:            "㉒",
	KeyF23:            "㉓",
	KeyF24:            "㉔",
	KeyEscape:         "⎋", // alternatives: "␛","␘"
	KeyTab:            "⇥", // alternatives: "↹","␉"
	KeyEnter:          "⏎", // alternatives: "↵","␤","␍"
	KeyBackspace:      "⌫", // alternatives: "␈"
	KeyDelete:         "⌦", // alternatives: "␡"
	KeyLeftShift:      "⇧", // alternatives: "␏"
	KeyRightShift:     "⇧.",
	KeyLeftControl:    "⌃",
	KeyRightControl:   "⌃.",
	KeyLeftMeta:       "◆",
	KeyRightMeta:      "◆.",
	KeyLeftAlt:        "⎇",
	KeyRightAlt:       "⎇.",
	KeyHome:           "⇱",
	KeyEnd:            "⇲",
	KeyPageUp:         "⇞",
	KeyPageDown:       "⇟",
	KeyLeftArrow:      "←", // alternatives: "⇠"
	KeyRightArrow:     "→", // alternatives: "⇢"
	KeyUpArrow:        "↑", // alternatives: "⇡"
	KeyDownArrow:      "↓", // alternatives: "⇣"
	KeyCapsLock:       "⇪",
	KeyNumLock:        "⇭",
	KeyScrollLock:     "⤓", // alternatives: "⇳"
	KeyVolumeUp:       "🔊",
	KeyVolumeDown:     "🔈",
	KeyMute:           "🔇",
	KeyGrave:          "`",
	Key0:              "0",
	Key1:              "1",
	Key2:              "2",
	Key3:              "3",
	Key4:              "4",
	Key5:              "5",
	Key6:              "6",
	Key7:              "7",
	Key8:              "8",
	Key9:              "9",
	KeyMinus:          "-",
	KeyEqual:          "=",
	KeyQ:              "Q",
	KeyW:              "W",
	KeyE:              "E",
	KeyR:              "R",
	KeyT:              "T",
	KeyY:              "Y",
	KeyU:              "U",
	KeyI:              "I",
	KeyO:              "O",
	KeyP:              "P",
	KeyLeftBracket:    "[",
	KeyRightBracket:   "]",
	KeyBackslash:      "\\",
	KeyA:              "A",
	KeyS:              "S",
	KeyD:              "D",
	KeyF:              "F",
	KeyG:              "G",
	KeyH:              "H",
	KeyJ:              "J",
	KeyK:              "K",
	KeyL:              "L",
	KeySemicolon:      ";",
	KeyQuote:          "\"",
	KeyZ:              "Z",
	KeyX:              "X",
	KeyC:              "C",
	KeyV:              "V",
	KeyB:              "B",
	KeyN:              "N",
	KeyM:              "M",
	KeyComma:          ",",
	KeyPeriod:         ".",
	KeySlash:          "/",
	KeySpace:          "␣", // alternatives: "␠"
	KeyNumpad0:        "🄀",
	KeyNumpad1:        "⒈",
	KeyNumpad2:        "⒉",
	KeyNumpad3:        "⒊",
	KeyNumpad4:        "⒋",
	KeyNumpad5:        "⒌",
	KeyNumpad6:        "⒍",
	KeyNumpad7:        "⒎",
	KeyNumpad8:        "⒏",
	KeyNumpad9:        "⒐",
	KeyNumpadDecimal:  "‥",
	KeyNumpadDivide:   "/.",
	KeyNumpadMultiply: "*.",
	KeyNumpadMinus:    "-.",
	KeyNumpadPlus:     "+.",
	KeyNumpadEnter:    "⏎.",
}

func (k Key) String() string {
	if k < KeyKnownCount {
		return keyNames[k]
	}
	return "unknown key " + strconvh.FormatUint8(uint8(k-KeyKnownCount))
}

func (k Key) ShortString() string {
	if k < KeyKnownCount {
		return keyShortNames[k]
	}
	return strconvh.FormatUint8(uint8(k-KeyKnownCount)) + "?"
}

type KeyModifiers uint8

const KeyNoModifiers KeyModifiers = 0

const (
	KeyModifierShift      KeyModifiers = 1 << iota
	KeyModifierControl    KeyModifiers = 1 << iota
	KeyModifierMeta       KeyModifiers = 1 << iota
	KeyModifierAlt        KeyModifiers = 1 << iota
	KeyModifierCapsLock   KeyModifiers = 1 << iota
	KeyModifierNumLock    KeyModifiers = 1 << iota
	KeyModifierScrollLock KeyModifiers = 1 << iota
	KeyModifierKnownCount              = iota // TODO remove this (single use)?
)

const keyModifierFirstUnknown KeyModifiers = 1 << KeyModifierKnownCount

// MacOS altarnative names.
// TODO this may be wrong
const (
	KeyModifierCommand = KeyModifierMeta
	KeyModifierOption  = KeyModifierAlt
)

var keyModifierNames = [KeyModifierKnownCount]string{
	keyNames[KeyLeftShift],
	keyNames[KeyLeftControl],
	keyNames[KeyLeftMeta],
	keyNames[KeyLeftAlt],
	keyNames[KeyCapsLock],
	keyNames[KeyNumLock],
	keyNames[KeyScrollLock],
}

var keyModifierShortNames = [KeyModifierKnownCount]string{
	keyShortNames[KeyLeftShift],
	keyShortNames[KeyLeftControl],
	keyShortNames[KeyLeftMeta],
	keyShortNames[KeyLeftAlt],
	keyShortNames[KeyCapsLock],
	keyShortNames[KeyNumLock],
	keyShortNames[KeyScrollLock],
}

func (m KeyModifiers) IsNothingPressed() bool    { return m == KeyNoModifiers }
func (m KeyModifiers) IsShiftPressed() bool      { return m&KeyModifierShift > 0 }
func (m KeyModifiers) IsControlPressed() bool    { return m&KeyModifierControl > 0 }
func (m KeyModifiers) IsMetaPressed() bool       { return m&KeyModifierMeta > 0 }
func (m KeyModifiers) IsAltPressed() bool        { return m&KeyModifierAlt > 0 }
func (m KeyModifiers) IsAnyUnknownPressed() bool { return m >= keyModifierFirstUnknown }

func (m KeyModifiers) String() string {
	if m.IsNothingPressed() {
		return "no modifiers"
	}
	r := ""
	for i := uint(0); i < KeyModifierKnownCount; i++ {
		if m&(KeyModifiers(1)<<i) > 0 {
			r += keyModifierNames[i] + ", "
		}
	}
	if m.IsAnyUnknownPressed() {
		for i := keyModifierFirstUnknown; i >= keyModifierFirstUnknown; i = i << 1 {
			if m&i > 0 {
				r += "unknown modifier " + strconvh.FormatUint8(uint8(m)) + ", "
			}
		}
	}
	return r[:len(r)-2]
}

func (m KeyModifiers) ShortString() string {
	if m.IsNothingPressed() {
		return ""
	}
	r := ""
	for i := uint(0); i < KeyModifierKnownCount; i++ {
		if m&(KeyModifiers(1)<<i) > 0 {
			r += keyModifierShortNames[i]
		}
	}
	if m.IsAnyUnknownPressed() {
		for i := keyModifierFirstUnknown; i >= keyModifierFirstUnknown; i = i << 1 {
			if m&i > 0 {
				r += strconvh.FormatUint8(uint8(m)-KeyModifierKnownCount) + "?"
			}
		}
	}
	return r
}

type KeyEvent uint8

const (
	KeyEventPress       KeyEvent = 0
	KeyEventRelease     KeyEvent = 1
	KeyEventRepeatPress          = KeyEventPress | 1<<1
)

// Any: first or repeat.
func (e KeyEvent) IsPressed() bool { return e&KeyEventRelease == KeyEventPress }

func (e KeyEvent) IsFirstPressed() bool  { return e == KeyEventPress }
func (e KeyEvent) IsRepeatPressed() bool { return e == KeyEventRepeatPress }
func (e KeyEvent) IsReleased() bool      { return e == KeyEventRelease }

func (e KeyEvent) String() string {
	switch e {
	case KeyEventPress:
		return "pressed"
	case KeyEventRelease:
		return "released"
	case KeyEventRepeatPress:
		return "repeat pressed"
	default:
		return "unknown event " + strconvh.FormatUint8(uint8(e))
	}
}

func (e KeyEvent) ShortString() string {
	switch e {
	case KeyEventPress:
		return "↓"
	case KeyEventRelease:
		return "↑"
	case KeyEventRepeatPress:
		return "⇊"
	default:
		return "?"
	}
}
