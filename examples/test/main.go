// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/controls"
	"log"
)

func main() {
	app, err := gui.InitApplication()
	if err != nil {
		log.Fatalln(err.Error())
	}

	window := gui.NewWindow()
	window.SetTitle("Hello world")

	b1 := controls.NewButton("Button 1")
	b2 := controls.NewButton("Button 2")
	cb1 := controls.NewCheckBox(true, controls.CheckBoxUnknown)
	cb2 := controls.NewCheckBox(false, controls.CheckBoxChecked)
	cb3 := controls.NewCheckBox(false, controls.CheckBoxUnchecked)
	cb4 := controls.NewCheckBox(false, controls.CheckBoxChecked)
	f := gui.NewFontDefaultFont(40, true, false, false)
	defer f.Release()
	l1 := controls.NewLabel("Label 1", f, gui.ColorF64{0, 1, 0, 1})
	l2 := controls.NewLabel("Label 2", f, gui.ColorF64{0, 0, 1, 1})
	vt1 := controls.NewVTable(b1, b2)
	vt2 := controls.NewVTable(cb1, cb2, cb3, cb4)
	ht := controls.NewHTable(vt1, vt2, l1, l2)

	wbs := controls.NewWindowButtons()
	l0 := controls.NewLayers(ht, wbs)

	window.SetChild(l0)

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
