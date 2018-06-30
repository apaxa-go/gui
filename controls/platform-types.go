// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import "github.com/apaxa-go/gui/controls/platform"

type WindowButtons = platform.Buttons

func NewWindowButtons() *WindowButtons { // TODO nolint: golint
	return platform.NewButtons()
}
