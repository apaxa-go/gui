// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type TrackingAreaID int

type TrackingArea struct {
	Area       RectangleF64
	EnterLeave bool
	Move       bool
}

func (a TrackingArea) Valid() bool {
	return a.EnterLeave || a.Move
}
