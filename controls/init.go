package controls

import "github.com/apaxa-go/gui"

func init() {
	ReInit()
}

func ReInit() {
	if defaultFont != nil {
		defaultFont.Release()
	}
	defaultFont = gui.NewFontDefaultFont(FontSize, false, false, false)
}
