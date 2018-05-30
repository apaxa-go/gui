// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

type WindowTitle = windowTitleMacOS

func NewWindowTitle() *WindowTitle {
	return newWindowTitleMacOS()
}
