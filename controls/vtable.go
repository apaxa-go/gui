// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

type VTable struct {
	BaseControl
	children []Control
}

func (c *VTable) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	// There are multiple ways to calculate bestWidth.
	// Here we use average from children's bestWidths.
	// And in this case maxWidth has higher priority than bestWidth.
	if l := len(c.children); l > 0 {
		maxWidth = mathh.PositiveInfFloat64()
		for _, child := range c.children {
			minWidth = mathh.Max2Float64(minWidth, child.MinWidth())
			bestWidth += child.BestWidth()
			maxWidth = mathh.Min2Float64(maxWidth, child.MaxWidth())
		}
		maxWidth = mathh.Max2Float64(minWidth, maxWidth)
		bestWidth /= float64(l)
		bestWidth = mathh.Max2Float64(minWidth, bestWidth)
		bestWidth = mathh.Min2Float64(maxWidth, bestWidth)
	}
	return
}

func (c *VTable) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	for _, child := range c.children {
		minHeight += child.MinHeight()
		bestHeight += child.BestHeight()
		maxHeight += child.MaxHeight()
	}
	return
}

func (c *VTable) Draw(canvas Canvas, region RectangleF64) {
	for _, child := range c.children {
		// TODO draw only required children
		child.Draw(canvas, region)
	}
}
func (c *VTable) FocusCandidate(reverse bool, current Control) Control {
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
	switch {
	case height >= c.MaxHeight(): // scale according to MaxHeight
		scale := c.MaxHeight()
		for i, child := range c.children {
			scalePart := child.MaxHeight()
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			tops[i] = top
			bottoms[i] = bottom

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	case height >= c.BestHeight(): // scale according to BestHeight
		scale := c.BestHeight()
		for i, child := range c.children {
			scalePart := child.BestHeight()
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			tops[i] = top
			bottoms[i] = bottom

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	default: // scale according to MinWidth
		scale := c.MinHeight()
		for i, child := range c.children {
			scalePart := child.MinHeight()
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			tops[i] = top
			bottoms[i] = bottom

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	}
	return
}

func (c *VTable) Insert(control Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	if at < 0 {
		at = 0
	} else if at > len(c.children) {
		at = len(c.children)
	}
	c.BaseControl.SetParent(control, c)
	c.children = append(append(c.children[:at], control), c.children[at:]...)
	c.SetUPG(false) // TODO why not recursive?
}

func (c *VTable) Prepend(control Control) {
	c.Insert(control, 0)
}

func (c *VTable) Append(control Control) {
	c.Insert(control, len(c.children))
}

func (c *VTable) Remove(i int) Control {
	if i < 0 {
		i = 0
	} else if i >= len(c.children) {
		i = len(c.children) - 1
	}
	control := c.children[i]
	c.BaseControl.SetParent(control, nil)
	c.children = append(c.children[:i], c.children[i+1:]...)
	c.SetUPG(false)
	return control
}

func (c *VTable) NumRows() int { return len(c.children) }

func (c *VTable) Children() []Control { return c.children }

func NewVTable(children ...Control) *VTable {
	r := &VTable{
		children: children,
	}
	for _, child := range children {
		r.BaseControl.SetParent(child, r)
	}
	return r
}
