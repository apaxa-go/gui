// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

type FocusEvent struct {
	Receive bool    // True on receiving focus and false on lost.
	Another Control // Another focus Control. If Received then this field is a Control which previously been focused, otherwise it is a Control which now receive focus.
}
