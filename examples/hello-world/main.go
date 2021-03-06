// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/controls"
)

func main() {
	window := gui.NewWindow("Hello world")

	font := gui.NewFontDefaultFont(24, false, false, false)
	defer font.Release()

	color := basetypes.MakeColorF64RGB8(80, 80, 80)

	label := controls.NewLabel("Hello world!", font, color)
	window.SetChild(label)

	err := gui.Run()
	if err != nil {
		panic(err.Error())
	}
}
