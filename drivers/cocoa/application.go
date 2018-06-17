// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "application.h"
*/
import "C"

func init() {
	C.InitApplication()
}

func Run() (err error) {
	C.RunApplication()
	return nil
}

func Stop() {
	C.StopApplication()
}
