// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
)

type VTable struct {
	gui.BaseControl
	children []gui.Control
}

func (c *VTable) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	if len(c.children) > 0 {
		maxWidth = gui.PosInfF64()
		for _, child := range c.children {
			minWidth = gui.Max2Float64(minWidth, child.MinWidth())
			maxWidth = gui.Min2Float64(maxWidth, child.MaxWidth())
		}
		maxWidth = gui.Max2Float64(minWidth, maxWidth)
	}
	return
}

func (c *VTable) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	for _, child := range c.children {
		minHeight += child.MinHeight()
		maxHeight += child.MaxHeight()
	}
	return
}

func (c *VTable) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	for _, child := range c.children {
		// TODO draw only required children
		child.Draw(canvas, region)
	}
}
func (c *VTable) FocusCandidate(reverse bool, current gui.Control) gui.Control {
	l := len(c.children)
	if l == 0 {
		return nil
	}
	switch {
	case current == nil && !reverse: // first
		return c.children[0]
	case current == nil && reverse: // last
		return c.children[l-1]
	default:
		i := 0
		for ; i < l && c.children[i] != current; i++ {
		}
		if i == l { // not found
			return c.children[0]
		}
		if reverse {
			i--
		} else {
			i++
		}
		if i < 0 || i >= l { // out of current control
			return nil
		}
		return c.children[i]
	}
}

func (c *VTable) PointerCandidates() []gui.Control {
	return c.children
}

func (c *VTable) ComputeChildHorGeometry() (lefts, rights []float64) {
	l := len(c.children)
	lefts = make([]float64, l)
	rights = make([]float64, l)
	left := c.Geometry().Left
	right := c.Geometry().Right
	for i := 0; i < l; i++ {
		lefts[i] = left
		rights[i] = right
	}
	return
}

func (c *VTable) ComputeChildVerGeometry() (tops, bottoms []float64) {
	l := len(c.children)
	tops = make([]float64, l)
	bottoms = make([]float64, l)

	top := c.Geometry().Top
	height := c.Geometry().Height()
	minHeight := c.MinHeight()
	for i, child := range c.children {
		childMinHeight := child.MinHeight()
		curHeight := height * childMinHeight / minHeight
		bottom := top + curHeight

		tops[i] = top
		bottoms[i] = bottom

		top = bottom
		height -= curHeight
		minHeight -= childMinHeight
	}
	return
}

func (c *VTable) Insert(control gui.Control, at int) {
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

func (c *VTable) Prepend(control gui.Control) {
	c.Insert(control, 0)
}

func (c *VTable) Append(control gui.Control) {
	c.Insert(control, len(c.children))
}

func (c *VTable) Remove(i int) gui.Control {
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

func (c *VTable) NumRows() int { return len(c.children) }

func (c *VTable) Children() []gui.Control { return c.children }

func NewVTable(children ...gui.Control) *VTable {
	r := &VTable{
		children: children,
	}
	for _, child := range children {
		gui.SetParent(child, r)
	}
	return r
}
