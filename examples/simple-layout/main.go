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
	cb := controls.NewCheckBox(false, controls.CheckBoxChecked)
	ht := controls.NewHTable(b1, b2, cb)
	window.SetChild(ht)

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
