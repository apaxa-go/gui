// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
)

type HTable struct {
	gui.BaseControl
	children []gui.Control
}

func (c *HTable) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	for _, child := range c.children {
		minWidth += child.MinWidth()
		maxWidth += child.MaxWidth()
	}
	return
}

func (c *HTable) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	if len(c.children) > 0 {
		maxHeight = gui.PosInfF64()
		for _, child := range c.children {
			minHeight = gui.Max2Float64(minHeight, child.MinHeight())
			maxHeight = gui.Min2Float64(maxHeight, child.MaxHeight())
		}
		maxHeight = gui.Max2Float64(minHeight, maxHeight)
	}
	return
}

func (c *HTable) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	for _, child := range c.children {
		// TODO draw only required children
		child.Draw(canvas, region)
	}
}

func (c *HTable) ProcessEvent(gui.Event) bool {
	// TODO
	return false
}

func (c *HTable) ComputeChildHorGeometry() (lefts, rights []float64) {
	l := len(c.children)
	lefts = make([]float64, l)
	rights = make([]float64, l)

	left := c.Geometry().Left
	width := c.Geometry().Width()
	minWidth := c.MinWidth()
	for i, child := range c.children {
		childMinWidth := child.MinWidth()
		curWidth := width * childMinWidth / minWidth
		right := left + curWidth

		lefts[i] = left
		rights[i] = right

		left = right
		width -= curWidth
		minWidth -= childMinWidth
	}
	return
}

func (c *HTable) ComputeChildVerGeometry() (tops, bottoms []float64) {
	l := len(c.children)
	tops = make([]float64, l)
	bottoms = make([]float64, l)
	top := c.Geometry().Top
	bottom := c.Geometry().Bottom
	for i := 0; i < l; i++ {
		tops[i] = top
		bottoms[i] = bottom
	}
	return
}

func (c *HTable) Insert(control gui.Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	if at < 0 {
		at = 0
	} else if at > len(c.children) {
		at = len(c.children)
	}
	gui.SetParent(control, c)
	c.children = append(append(c.children[:at], control), c.children[at:]...)
	c.SetUPG(false)
	{
		// TODO do smth with this:
		// c.window.Hypervisor().UpdatePossibleHorGeometry(c, false)
		// c.window.Hypervisor().UpdatePossibleVerGeometry(c, false)
		// c.window.Hypervisor().UpdateChildHorGeometry(c)
		// c.window.Hypervisor().UpdateChildVerGeometry(c)
		// c.window.Hypervisor().InvalidateRegion(control)
	}
}

func (c *HTable) Prepend(control gui.Control) {
	c.Insert(control, 0)
}

func (c *HTable) Append(control gui.Control) {
	c.Insert(control, len(c.children))
}

func (c *HTable) Remove(i int) gui.Control {
	if i < 0 {
		i = 0
	} else if i >= len(c.children) {
		i = len(c.children) - 1
	}
	control := c.children[i]
	gui.SetParent(control, nil)
	c.children = append(c.children[:i], c.children[i+1:]...)
	c.SetUPG(false)
	return control
}

func (c *HTable) NumColumns() int { return len(c.children) }

func (c *HTable) Children() []gui.Control { return c.children }

func NewHTable(children ...gui.Control) *HTable {
	r := &HTable{
		children: children,
	}
	for _, child := range children {
		gui.SetParent(child, r)
	}
	return r
}
