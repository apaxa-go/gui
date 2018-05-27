// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

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
