// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type MoveAreaID uint

func (id MoveAreaID) Base() uint { return uint(id) }

type PointerMoveEvent struct {
	ID    MoveAreaID
	Point PointF64
}

func (e PointerMoveEvent) String() string {
	return "in area " + strconvh.FormatUint(e.ID.Base()) + " move to " + e.Point.String()
}

func (e PointerMoveEvent) ShortString() string {
	return "[" + strconvh.FormatUint(e.ID.Base()) + "]↝(" + strconvh.FormatFloat64Prec(e.Point.X, 0) + ";" + strconvh.FormatFloat64Prec(e.Point.Y, 0) + ")"
}
