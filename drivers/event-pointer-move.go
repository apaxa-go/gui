// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type PointerMoveEvent struct {
	Point PointF64
}

func (e PointerMoveEvent) String() string {
	return "move to " + e.Point.String()
}

func (e PointerMoveEvent) ShortString() string {
	return "↝(" + strconvh.FormatFloat64Prec(e.Point.X, 0) + ";" + strconvh.FormatFloat64Prec(e.Point.Y, 0) + ")"
}
