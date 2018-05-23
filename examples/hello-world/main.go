package main

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/controls"
	_ "github.com/apaxa-go/gui/drivers/cocoa"
	"log"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	app, err := gui.InitApplication()
	if err != nil {
		log.Fatalln(err.Error())
	}

	window := gui.NewWindow()
	window.SetTitle("Hello world")

	font := gui.NewFontDefaultFont(24, false, false, false)
	defer font.Release()

	color := basetypes.MakeColorF64RGB8(80, 80, 80)

	label := controls.NewLabel("Hello world!", font, color)
	window.SetChild(label)

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}

}
