// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package platform

import "github.com/apaxa-go/gui/controls/platform/macos/buttons"

type Buttons = buttons.Buttons

func NewButtons() *Buttons { // TODO nolint: golint
	return buttons.NewButtons()
}
