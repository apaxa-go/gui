// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/controls"
)

func main() {
	window := gui.NewWindow("Simple layout")

	b1 := controls.NewButton("Button 1")
	b2 := controls.NewButton("Button 2")
	cb := controls.NewCheckBox(false, controls.CheckBoxChecked)
	ht := controls.NewHTable(b1, b2, cb)
	window.SetChild(ht)

	err := gui.Run()
	if err != nil {
		panic(err.Error())
	}
}
