// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import "github.com/apaxa-go/gui/drivers"

func NewFont(spec FontSpec) (Font, error) {
	return driverFontConstructor(spec)
}

func NewFontDefaultFont(size float64, monospace, italic, bold bool) Font {
	spec := drivers.MakeFontSpecDefaultFont(size, monospace, italic, bold)
	f, _ := NewFont(spec)
	return f
}

func NewFontByFamily(family string, size float64, monospace, italic, bold bool) Font {
	spec := drivers.MakeFontSpecByFamily(family, size, monospace, italic, bold)
	f, _ := NewFont(spec)
	return f
}

func NewFontByName(name string, size float64, monospace, italic, bold bool) Font {
	spec := drivers.MakeFontSpecByName(name, size, monospace, italic, bold)
	f, _ := NewFont(spec)
	return f
}

func NewFontByFile(file string, index FontIndex, size float64, monospace, italic, bold bool) (Font, error) {
	spec := drivers.MakeFontSpecByFile(file, index, size, monospace, italic, bold)
	return NewFont(spec)
}
