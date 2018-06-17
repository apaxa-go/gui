// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type PointerDragEvent struct {
	Delta PointF64
}

func (e PointerDragEvent) String() string {
	return "drag " + e.Delta.String()
}

func (e PointerDragEvent) ShortString() string {
	return "⤞(" + strconvh.FormatFloat64Prec(e.Delta.X, 0) + ";" + strconvh.FormatFloat64Prec(e.Delta.Y, 0) + ")"
}
