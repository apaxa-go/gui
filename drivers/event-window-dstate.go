// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type WindowDisplayStateEvent struct {
	OldState WindowDisplayState
	NewState WindowDisplayState
}

func (e WindowDisplayStateEvent) String() string {
	return e.OldState.String() + " => " + e.NewState.String()
}
