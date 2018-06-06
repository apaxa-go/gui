// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type ScrollEvent struct {
	Delta     PointF64 // TODO if only one direction per event may changed then may be use float64 Delta and flag for coordinate (X or Y)?
	Modifiers KeyModifiers
	Point     PointF64
}

func (e ScrollEvent) String() string {
	r := "scroll"
	if e.Delta.X != 0 {
		r += " horizontal " + strconvh.FormatFloat64(e.Delta.X)
	}
	if e.Delta.Y != 0 {
		r += " vertical " + strconvh.FormatFloat64(e.Delta.Y)
	}
	return r
}

func (e ScrollEvent) ShortString() string {
	// TODO
	return e.String()
}
