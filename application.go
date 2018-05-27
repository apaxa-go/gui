// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "runtime"

func InitApplication() (Application, error) {
	runtime.LockOSThread()
	return driverApplicationConstructor()
}
