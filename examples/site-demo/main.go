// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/controls"
)

const defaultFontSize = 20
const largeFontSize = 40

var defaultFont gui.Font
var largeFont gui.Font

func init() {
	defaultFont = gui.NewFontDefaultFont(defaultFontSize, false, false, false)
	largeFont = gui.NewFontDefaultFont(largeFontSize, true, true, true)
}

func main() {
	window := gui.NewWindow("APAXA GUI demo")

	// Setup controls

	label0 := controls.NewLabel("Event:", defaultFont, gui.ColorF64{}.MakeFromRGB8(0, 0, 0))
	label1 := controls.NewLabel("", defaultFont, gui.ColorF64{}.MakeFromRGB8(255, 0, 0))
	label2 := controls.NewLabel("APAXA GUI site demo", largeFont, gui.ColorF64{}.MakeFromRGB8(74, 74, 74))

	button := controls.NewButton("Click me")
	checkbox := controls.NewCheckBox(false, controls.CheckBoxChecked)

	table := controls.NewTable(3, 2)
	table.Set(label0, 0, 0)
	table.Set(label1, 0, 1)
	table.Set(label2, 1, 0)
	table.AddSpan(1, 0, 1, 2)
	table.Set(button, 2, 0)
	table.Set(checkbox, 2, 1)

	winTitle := controls.NewWindowTitle()
	winButtons := controls.NewWindowButtons()
	verticalTable := controls.NewVTable(winTitle, table)
	layers := controls.NewLayers(verticalTable, winButtons)

	window.SetChild(layers)

	// Configure actions

	button.SetAction(func() {
		label1.SetText("Clicked the button, checkbox is " + checkbox.State().String())
	})

	// Run app

	err := gui.Run()
	if err != nil {
		panic(err.Error())
	}
}
