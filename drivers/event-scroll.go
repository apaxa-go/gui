// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type ScrollEvent struct {
	// TODO if only one direction per event may changed then may be use single Delta and flag for coordinate (X or Y)?
	DeltaX float64
	DeltaY float64
}

func (e ScrollEvent) String() string {
	r := "scroll"
	if e.DeltaX != 0 {
		r += " horizontal " + strconvh.FormatFloat64(e.DeltaX)
	}
	if e.DeltaY != 0 {
		r += " vertical " + strconvh.FormatFloat64(e.DeltaY)
	}
	return r
}

func (e ScrollEvent) ShortString() string {
	// TODO
	return e.String()
}
