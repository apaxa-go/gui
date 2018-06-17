// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type PointerEnterLeaveEvent struct {
	Id    TrackingAreaID
	Enter bool
}

func (e PointerEnterLeaveEvent) String() string {
	action := "leaves"
	if e.Enter {
		action = "enter"
	}
	return "pointer " + action + " area " + strconvh.FormatInt(int(e.Id))
}

func (e PointerEnterLeaveEvent) ShortString() string {
	action := "⍅"
	if e.Enter {
		action = "⍆"
	}
	return action + strconvh.FormatInt(int(e.Id))
}
