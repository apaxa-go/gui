// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type ScrollEvent struct {
	Delta     PointF64
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
	r := ""
	if e.Delta.X < 0 {
		r += "↞" + strconvh.FormatFloat64Prec(-e.Delta.X, -1)
	} else if e.Delta.X > 0 {
		r += "↠" + strconvh.FormatFloat64Prec(e.Delta.X, -1)
	}
	if len(r) != 0 && e.Delta.Y != 0 {
		r += ";"
	}
	if e.Delta.Y < 0 {
		r += "↡" + strconvh.FormatFloat64Prec(-e.Delta.Y, -1)
	} else if e.Delta.Y > 0 {
		r += "↟" + strconvh.FormatFloat64Prec(e.Delta.Y, -1)
	}
	return r
}
