// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package scvi

var SelectMark = SCVI{
	PointF64{16, 20}, // nolint: vet
	true,
	[]Primitive{
		Lines{
			[]PointF64{{5, 8}, {8, 5}, {11, 8}}, // nolint: vet
			1.3,
			1,
		},
		Lines{
			[]PointF64{{5, 12}, {8, 15}, {11, 12}}, // nolint: vet
			1.3,
			1,
		},
	},
}
