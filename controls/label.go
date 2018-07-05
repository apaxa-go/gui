// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import "github.com/apaxa-go/helper/mathh"

type TextImage struct {
	BaseControl
	text  string
	font  Font
	color ColorF64
	align Align
}

func (c *TextImage) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	// TODO compute geometry on text/font changes only
	width := c.Window().OfflineCanvas().TextImageGeometry(c.text, c.font).X
	return width, width, mathh.PositiveInfFloat64()
}

func (c *TextImage) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	height := c.Window().OfflineCanvas().TextImageGeometry(c.text, c.font).Y
	return height, height, height
}

func (c TextImage) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := c.align.ApplyF64(c.Geometry(), c.MinSize())
	//canvas.DrawRectangle(place,ColorF64{0,0,0,1},1)
	canvas.DrawTextImage(c.text, c.font, c.color, place.LT())
}

func (c *TextImage) GetText() string { return c.text }
func (c *TextImage) SetText(text string) {
	if c.text == text {
		return
	}
	c.text = text
	c.SetUPGIR(false)
}

func (c *TextImage) GetFont() Font { return c.font }
func (c *TextImage) SetFont(font Font) {
	if c.font == font {
		return
	}
	c.font = font
	c.SetUPGIR(false)
}

func (c *TextImage) GetColor() ColorF64 { return c.color }
func (c *TextImage) SetColor(color ColorF64) {
	if c.color == color {
		return
	}
	c.color = color
	c.SetIR()
}

func (c *TextImage) GetAlign() Align { return c.align }
func (c *TextImage) SetAlign(align Align) {
	align = align.KeepSize()
	if c.align == align {
		return
	}
	c.align = align
	c.SetIR()
}

func NewLabel(text string, font Font, color ColorF64) *TextImage {
	return &TextImage{
		text:  text,
		font:  font,
		color: color,
		align: Align(0).MakeCenter(),
	}
}
