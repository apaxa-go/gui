// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

type Label struct {
	BaseControl
	text  string
	font  Font
	color ColorF64
	align Align
}

func (c *Label) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	width := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font).X
	return width, width, width
}

func (c *Label) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	height := c.Window().OfflineCanvas().TextLineGeometry(c.text, c.font).Y
	return height, height, height
}

func (c Label) Draw(canvas Canvas, _ RectangleF64) {
	// TODO use region
	place := c.align.ApplyF64(c.Geometry(), c.MinSize())
	canvas.DrawTextLine(c.text, c.font, place.LT(), c.color)
}

func (c *Label) GetText() string { return c.text }
func (c *Label) SetText(text string) {
	if c.text == text {
		return
	}
	c.text = text
	c.SetUPGIR(false)
}

func (c *Label) GetFont() Font { return c.font }
func (c *Label) SetFont(font Font) {
	if c.font == font {
		return
	}
	c.font = font
	c.SetUPGIR(false)
}

func (c *Label) GetColor() ColorF64 { return c.color }
func (c *Label) SetColor(color ColorF64) {
	if c.color == color {
		return
	}
	c.color = color
	c.SetIR()
}

func (c *Label) GetAlign() Align { return c.align }
func (c *Label) SetAlign(align Align) {
	align = align.KeepSize()
	if c.align == align {
		return
	}
	c.align = align
	c.SetIR()
}

func NewLabel(text string, font Font, color ColorF64) *Label {
	return &Label{
		text:  text,
		font:  font,
		color: color,
		align: Align(0).MakeCenter(),
	}
}
