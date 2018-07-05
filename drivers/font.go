// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type Font interface {
	IAmFont()
	Release()
	Ascent() float64
	Descent() float64
	Leading() float64
}

type TextLineOrigin uint8

const (
	TextLineTop TextLineOrigin = iota
	TextLineBase
	TextLineBottom
)

func (TextLineOrigin) MakeTop() TextLineOrigin    { return TextLineTop }
func (TextLineOrigin) MakeBase() TextLineOrigin   { return TextLineBase }
func (TextLineOrigin) MakeBottom() TextLineOrigin { return TextLineBottom }

func (o TextLineOrigin) IsTop() bool    { return o == TextLineTop }
func (o TextLineOrigin) IsBase() bool   { return o == TextLineBase }
func (o TextLineOrigin) IsBottom() bool { return o == TextLineBottom }
