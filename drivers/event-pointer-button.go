// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type PointerButtonEvent struct {
	Kind      PointerButtonEventKind
	Button    PointerButton
	Point     PointF64
	Modifiers KeyModifiers
}

func (e PointerButtonEvent) String() string {
	return e.Button.String() + " is " + e.Kind.String() + " at " + e.Point.String() + " with modifiers: " + e.Modifiers.String()
}

func (e PointerButtonEvent) ShortString() string {
	mod := e.Modifiers.ShortString()
	if len(mod) == 0 {
		return e.Kind.ShortString() + e.Button.ShortString() + "@" + e.Point.String()
	}
	return e.Kind.ShortString() + e.Button.ShortString() + "@" + e.Point.String() + "+" + e.Modifiers.ShortString()
}
