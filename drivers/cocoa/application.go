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
import "unsafe"

type Application struct {
	pointer unsafe.Pointer // may be uintptr better here?
}

func InitApplication() (app *Application, err error) {
	return &Application{C.InitApplication()}, nil
}

func (a *Application) Run() (err error) {
	C.ApplicationRun(a.pointer)
	return nil
}
