package main

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/controls"
	driver "github.com/apaxa-go/gui/drivers/cocoa"
	"log"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	app, err := driver.InitApplication()
	if err != nil {
		log.Fatalln(err.Error())
	}

	driverWindow, err := driver.CreateWindow(500, 500 /*, "someClass" /*, winProc,*/) // TODO remove size here?
	if err != nil {
		log.Fatalln(err.Error())
	}
	// defer window.Destroy() TODO call this from gui.Window

	window := gui.NewEmptyWindow(driverWindow)
	window.SetTitle("Hello world")

	font, ok := driver.NewFont("Times New Roman", 30)
	if !ok {
		panic("unable to create font")
	}
	label := controls.NewLabel("Hello world!", font, 30, basetypes.MakeColorF64RGB8(80, 80, 80))
	window.SetChild(label)

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}

}
