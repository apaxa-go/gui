// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type ModifiersEvent struct {
	Modifiers KeyModifiers
}

func (e ModifiersEvent) String() string      { return e.Modifiers.String() }
func (e ModifiersEvent) ShortString() string { return e.Modifiers.ShortString() }
