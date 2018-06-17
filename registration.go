// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

var driverRunApplication func() error
var driverStopApplication func()
var driverWindowConstructor func(title string) DriverWindow
var driverFontConstructor func(FontSpec) (Font, error)

func RegisterDriver(
	runApplication func() error,
	stopApplication func(),
	windowConstructor func(title string) DriverWindow,
	fontConstructor func(FontSpec) (Font, error),
) {
	driverRunApplication = runApplication
	driverStopApplication = stopApplication
	driverWindowConstructor = windowConstructor
	driverFontConstructor = fontConstructor
}
