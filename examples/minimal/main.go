package main

import (
	_ "github.com/apaxa-go/gui/drivers/cocoa"

	"github.com/apaxa-go/gui"
	"log"
)

func main() {
	app, err := gui.InitApplication()
	if err != nil {
		log.Fatalln(err.Error())
	}
	_ = gui.NewWindow()
	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}
