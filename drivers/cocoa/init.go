// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

import "github.com/apaxa-go/gui"

func mustCreateWindowI(title string) WindowI {
	w, err := CreateWindow(title)
	if err != nil {
		panic(err.Error())
	}
	return w
}

func newFontI(spec FontSpec) (FontI, error) {
	return NewFont(spec)
}

func init() {
	gui.RegisterDriver(Run, Stop, mustCreateWindowI, newFontI)
}
