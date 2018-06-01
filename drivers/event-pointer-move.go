// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type PointerMoveEvent struct {
	Point PointF64
}

func (e PointerMoveEvent) String() string {
	return "move to " + e.Point.String()
}

func (e PointerMoveEvent) ShortString() string {
	//TODO
	return e.String()
}
