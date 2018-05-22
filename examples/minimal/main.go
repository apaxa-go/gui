package main

import (
	"github.com/apaxa-go/gui"
	driver "github.com/apaxa-go/gui/drivers/cocoa"
	"log"
)

func main() {
	app, err := driver.InitApplication()
	if err != nil {
		log.Fatalln(err.Error())
	}

	driverWindow, err := driver.CreateWindow(500, 500 /*, "someClass" /*, winProc,*/) // TODO remove size here?
	if err != nil {
		log.Fatalln(err.Error())
	}
	// defer window.Destroy() TODO call this from gui.Window

	_ = gui.NewEmptyWindow(driverWindow)

	err = app.Run()
	if err != nil {
		panic(err.Error())
	}

}
