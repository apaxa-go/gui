// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type KeyboardEvent struct {
	Event     KeyEvent
	Key       Key
	Modifiers KeyModifiers
}

func (e KeyboardEvent) String() string {
	return "key " + e.Key.String() + " is " + e.Event.String() + " with following modifiers: " + e.Modifiers.String()
}

func (e KeyboardEvent) ShortString() string {
	mod := e.Modifiers.ShortString()
	if len(mod) == 0 {
		return e.Event.ShortString() + e.Key.ShortString()
	}
	return e.Event.ShortString() + e.Key.ShortString() + "+" + e.Modifiers.ShortString()
}
