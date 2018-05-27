// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

const (
	BorderWidth  = 1
	FontSize     = 12
	VerPadding   = 3
	HorPadding   = 3
	BorderRadius = 3
	Height       = 2*BorderWidth + 2*VerPadding + FontSize
	SmallHeight  = 14
)

/*
const (
	LineWidthShift      = -0.25
	MinLineWidth        = 1
	DefaultFontHeight   = 12
	MinFontHeight       = 8
	DefaultVerPadding   = 3
	MinVerPadding       = 1
	DefaultHorPadding   = 3
	MinHorPadding       = 1
	DefaultBorderRadius = 3
)

func LineWidth(scale float64) float64 {
	return gui.Max2Float64(MinLineWidth, gui.RoundF64(scale+LineWidthShift))
}

func LineWidthI(scale float64) int {
	return gui.RoundF64ToI(LineWidth(scale))
}

func ControlHeight(scale float64) float64 {
	return 2*(LineWidth(scale)+VerPadding(scale)) + FontHeight(scale)
}

func ControlHeightI(scale float64) int {
	return gui.RoundF64ToI(ControlHeight(scale))
}

func FontHeight(scale float64) float64 {
	return gui.Max2Float64(MinFontHeight, scale*DefaultFontHeight)
}

func FontHeightI(scale float64) int {
	return gui.RoundF64ToI(FontHeight(scale))
}

func VerPadding(scale float64) float64 {
	return gui.Max2Float64(MinVerPadding, scale*DefaultVerPadding)
}

func VerPaddingI(scale float64) int {
	return gui.RoundF64ToI(VerPadding(scale))
}

func HorPadding(scale float64) float64 {
	return gui.Max2Float64(MinHorPadding, scale*DefaultHorPadding)
}

func HorPaddingI(scale float64) int {
	return gui.RoundF64ToI(HorPadding(scale))
}

func BorderRadius(scale float64) float64 {
	return scale * DefaultBorderRadius
}
*/
// primaryHeight float64 // 20
// smallHeight   float64 // 14 (but should we still paint them in primaryHeight (in the middle, or + 1 additional offset from top)? )
