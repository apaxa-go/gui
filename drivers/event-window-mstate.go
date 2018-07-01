// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type WindowMainStateEvent struct {
	Become bool
}

func (e WindowMainStateEvent) String() string {
	if e.Become {
		return "window become main"
	}
	return "window lost main state"
}
