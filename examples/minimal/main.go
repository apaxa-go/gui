package main

import (
	"github.com/apaxa-go/gui"
	_ "github.com/apaxa-go/gui/drivers/cocoa"
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
