// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui/scvi"
)

type CheckBoxState uint8

const (
	CheckBoxUnchecked CheckBoxState = iota
	CheckBoxChecked   CheckBoxState = iota
	CheckBoxUnknown   CheckBoxState = iota
)

func (s CheckBoxState) IsChecked() bool   { return s == CheckBoxChecked }
func (s CheckBoxState) IsUnchecked() bool { return s == CheckBoxUnchecked }
func (s CheckBoxState) IsUnknown() bool   { return s >= CheckBoxUnknown }

var checkboxMark = scvi.SCVI{
	scvi.PointF64{14, 14},
	true,
	[]scvi.Primitive{
		scvi.MakeLines(
			[]scvi.PointF64{{3.5, 8}, {6, 10}, {10.5, 3}},
			1.3,
			1,
		),
	},
}

type CheckBox struct {
	BaseControl
	mayUnknown bool
	state      CheckBoxState
}

func (c *CheckBox) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	return SmallHeight, SmallHeight, SmallHeight
}

func (c *CheckBox) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	return SmallHeight, SmallHeight, SmallHeight
}

func (c CheckBox) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	space := Align(0).MakeCenter().ApplyF64(c.Geometry(), PointF64{SmallHeight, SmallHeight})
	rect := space.Inner(BorderWidth).ToRounded(BorderRadius)
	canvas.FillRoundedRectangle(rect, brightBackgroundColor)
	canvas.DrawRoundedRectangle(rect, brightBorderColor, BorderWidth)
	if c.state.IsChecked() { // TODO what if IsUnknown?
		checkboxMark.Draw(canvas, space, markColor)
	}
}

func (c *CheckBox) FocusCandidate(reverse bool, current Control) Control {
	if current == nil {
		return c
	}
	return nil
}

func (c CheckBox) correctState(state CheckBoxState) CheckBoxState {
	if !c.mayUnknown && state.IsUnknown() {
		return CheckBoxUnchecked
	}
	return state
}

func (c CheckBox) State() CheckBoxState {
	return c.state
}

func (c *CheckBox) SetState(state CheckBoxState) {
	state = c.correctState(state)
	if c.state == state {
		return
	}
	c.state = state
	c.SetIR()
}

func (c CheckBox) IsUnknown() bool   { return c.state.IsUnknown() }
func (c CheckBox) IsChecked() bool   { return c.state.IsChecked() }
func (c CheckBox) IsUnchecked() bool { return c.state.IsUnchecked() }

func (c *CheckBox) SetUnknown()   { c.SetState(CheckBoxUnknown) }
func (c *CheckBox) SetChecked()   { c.SetState(CheckBoxChecked) }
func (c *CheckBox) SetUnchecked() { c.SetState(CheckBoxUnchecked) }

func NewCheckBox(mayUnknown bool, state CheckBoxState) *CheckBox {
	r := &CheckBox{
		mayUnknown: mayUnknown,
	}
	r.state = r.correctState(state)
	return r
}
