// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import "github.com/apaxa-go/helper/strconvh"

type EnterLeaveAreaID uint

func (id EnterLeaveAreaID) Base() uint { return uint(id) }

type PointerEnterLeaveEvent struct {
	ID    EnterLeaveAreaID
	Enter bool
}

func (e PointerEnterLeaveEvent) String() string {
	action := "leaves"
	if e.Enter {
		action = "enter"
	}
	return "pointer " + action + " area " + strconvh.FormatUint(e.ID.Base())
}

func (e PointerEnterLeaveEvent) ShortString() string {
	action := "⍅"
	if e.Enter {
		action = "⍆"
	}
	return action + strconvh.FormatUint(e.ID.Base())
}
