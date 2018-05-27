// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import "github.com/apaxa-go/gui"

func init() {
	ReInit()
}

func ReInit() {
	if defaultFont != nil {
		defaultFont.Release()
	}
	defaultFont = gui.NewFontDefaultFont(FontSize, false, false, false)
}
