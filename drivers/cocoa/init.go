package cocoa

import "github.com/apaxa-go/gui"

func initApplication() (ApplicationI, error) {
	return InitApplication()
}

func mustCreateWindowI() WindowI {
	w, err := CreateWindow()
	if err != nil {
		panic(err.Error())
	}
	return w
}

func newFontI(spec FontSpec) (FontI, error) {
	return NewFont(spec)
}

func init() {
	gui.RegisterDriver(initApplication, mustCreateWindowI, newFontI)
}
