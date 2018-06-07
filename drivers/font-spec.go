// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

import (
	"github.com/apaxa-go/helper/mathh"
	"math"
)

type FontSpec struct {
	Index FontIndex
	Name  string  // Name of font file of name of font family (depending on "index" field).
	Size  float64 // Size of font in device independent points.

	Requirements FontRequirements
	Monospace    bool
	Italic       FontVariationItalic
	Slant        FontVariationSlant
	Width        FontVariationWidth
	Weight       FontVariationWeight
}

func (s FontSpec) Normalize() FontSpec {
	s.Italic = s.Italic.Normalize()
	s.Slant = s.Slant.Normalize()
	s.Width = s.Width.Normalize()
	s.Weight = s.Weight.Normalize()
	return s
}

type FontIndex uint64

// Special values for font index.
const (
	FontIndexFontFamily        FontIndex = mathh.MaxUint64 - iota
	FontIndexFontName          FontIndex = mathh.MaxUint64 - iota
	FontIndexDefaultDriverFont FontIndex = mathh.MaxUint64 - iota
	FontIndexNotCollection     FontIndex = 0
	FontIndexFirst             FontIndex = 0
)

func (i FontIndex) Family() bool        { return i == FontIndexFontFamily }
func (i FontIndex) Name() bool          { return i == FontIndexFontName }
func (i FontIndex) DefaultDriver() bool { return i == FontIndexDefaultDriverFont }

type FontRequirements uint8

const (
	FontRequirementMonospace FontRequirements = 1 << iota
	FontRequirementItalic    FontRequirements = 1 << iota
	FontRequirementSlant     FontRequirements = 1 << iota
	FontRequirementWidth     FontRequirements = 1 << iota
	FontRequirementWeight    FontRequirements = 1 << iota
	fontRequirementCount     FontRequirements = iota
)

const (
	FontRequirementNone      FontRequirements = 0
	FontRequirementValidMask FontRequirements = 2<<fontRequirementCount - 1
)

func (r FontRequirements) Normalize() FontRequirements {
	return r & FontRequirementValidMask
}

func (r FontRequirements) Monospace() bool { return r&FontRequirementMonospace > 0 }
func (r FontRequirements) Italic() bool    { return r&FontRequirementItalic > 0 }
func (r FontRequirements) Slant() bool     { return r&FontRequirementSlant > 0 }
func (r FontRequirements) Width() bool     { return r&FontRequirementWidth > 0 }
func (r FontRequirements) Weight() bool    { return r&FontRequirementWeight > 0 }

// Value must be in [0;1] where 0 means non-italic and 1 means italic.
type FontVariationItalic float64

// Common values for italic variation.
const (
	FontRoman     FontVariationItalic = 0
	FontNonItalic                     = FontRoman
	FontItalic    FontVariationItalic = 1
)

func (v FontVariationItalic) Normalize() FontVariationItalic {
	if v < FontRoman {
		return FontRoman
	}
	if v > FontItalic {
		return FontItalic
	}
	return v
}

// Slant (the angle, in clockwise degrees) of text. Value must be in (-90;+90) where 0 means upright text.
type FontVariationSlant float64

// Common values for slant variation.
const (
	// 89.9999999999999857891452847979962825775146484375 = math.Nextafter(90,0)

	FontSlantRight FontVariationSlant = 89.9999999999999857891452847979962825775146484375
	FontUpright    FontVariationSlant = 0
	FontSlantLeft                     = -FontSlantRight
)

func (v FontVariationSlant) Normalize() FontVariationSlant {
	if v < FontSlantRight {
		return FontSlantRight
	}
	if v > FontSlantLeft {
		return FontSlantLeft
	}
	return v
}

// Width of text. Values must be > 0. Value is percent of "normal" (as designer decided) width. Value 100 means "normal".
type FontVariationWidth float64

// Common values for width variation.
const (
	FontNarrowest      FontVariationWidth = math.SmallestNonzeroFloat64
	FontUltraCondensed FontVariationWidth = 50
	FontExtraCondensed FontVariationWidth = 62.5
	FontCondensed      FontVariationWidth = 75
	FontSemiCondensed  FontVariationWidth = 87.5
	FontMediumWidth    FontVariationWidth = 100
	FontNormalWidth                       = FontMediumWidth
	FontSemiExpanded   FontVariationWidth = 112.5
	FontExpanded       FontVariationWidth = 125
	FontExtraExpanded  FontVariationWidth = 150
	FontUltraExpanded  FontVariationWidth = 200
)

func (v FontVariationWidth) Normalize() FontVariationWidth {
	if v < FontNarrowest {
		return FontNarrowest
	}
	return v
}

// Value must be in [1;1000] where 1 means the lightest text, 400 is "normal" and 1000 is the blackest.
type FontVariationWeight float64

// Common values for weight variation.
const (
	FontLightest     FontVariationWeight = 1
	FontThin         FontVariationWeight = 100
	FontExtraLight   FontVariationWeight = 200
	FontLight        FontVariationWeight = 300
	FontNormalWeight FontVariationWeight = 400
	FontMediumWeight FontVariationWeight = 500
	FontSemiBold     FontVariationWeight = 600
	FontBold         FontVariationWeight = 700
	FontExtraBold    FontVariationWeight = 800
	FontBlack        FontVariationWeight = 900
	FontBlackest     FontVariationWeight = 1000
)

func (v FontVariationWeight) Normalize() FontVariationWeight {
	if v < FontLightest {
		return FontLightest
	}
	if v > FontBlackest {
		return FontBlackest
	}
	return v
}

func initFontSpecCommonFields(size float64, monospace, italic, bold bool) FontSpec {
	r := FontSpec{
		Size:         size,
		Requirements: FontRequirementMonospace | FontRequirementMonospace | FontRequirementWeight,
		Monospace:    monospace,
	}
	if italic {
		r.Italic = FontItalic
	}
	if bold {
		r.Weight = FontBold
	} else {
		r.Weight = FontNormalWeight
	}
	return r
}

func MakeFontSpecDefaultFont(size float64, monospace, italic, bold bool) FontSpec {
	r := initFontSpecCommonFields(size, monospace, italic, bold)
	r.Index = FontIndexDefaultDriverFont
	return r
}

func MakeFontSpecByFamily(family string, size float64, monospace, italic, bold bool) FontSpec {
	r := initFontSpecCommonFields(size, monospace, italic, bold)
	r.Index = FontIndexFontFamily
	r.Name = family
	return r
}

func MakeFontSpecByName(name string, size float64, monospace, italic, bold bool) FontSpec {
	r := initFontSpecCommonFields(size, monospace, italic, bold)
	r.Index = FontIndexFontName
	r.Name = name
	return r
}

func MakeFontSpecByFile(file string, index FontIndex, size float64, monospace, italic, bold bool) FontSpec {
	r := initFontSpecCommonFields(size, monospace, italic, bold)
	r.Index = index
	r.Name = file
	return r
}
