// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
)

func main() {
	_ = gui.NewWindow("Minimal")
	err := gui.Run()
	if err != nil {
		panic(err.Error())
	}
}
