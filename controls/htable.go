// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

type HTable struct {
	BaseControl
	children []Control
}

func (c *HTable) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	for _, child := range c.children {
		minWidth += child.MinWidth()
		bestWidth += child.BestWidth()
		maxWidth += child.MaxWidth()
	}
	return
}

func (c *HTable) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	// There are multiple ways to calculate bestHeight.
	// Here we use average from children's bestHeights.
	// And in this case maxHeight has higher priority than bestHeight.
	if l := len(c.children); l > 0 {
		maxHeight = mathh.PositiveInfFloat64()
		for _, child := range c.children {
			minHeight = mathh.Max2Float64(minHeight, child.MinHeight())
			bestHeight += child.BestHeight()
			maxHeight = mathh.Min2Float64(maxHeight, child.MaxHeight())
		}
		maxHeight = mathh.Max2Float64(minHeight, maxHeight)
		bestHeight /= float64(l)
		bestHeight = mathh.Max2Float64(minHeight, bestHeight)
		bestHeight = mathh.Min2Float64(maxHeight, bestHeight)
	}
	return
}

func (c *HTable) Draw(canvas Canvas, region RectangleF64) {
	for _, child := range c.children {
		// TODO draw only required children
		child.Draw(canvas, region)
	}
}

func (c *HTable) FocusCandidate(reverse bool, current Control) Control {
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

func (c *HTable) ComputeChildHorGeometry() (lefts, rights []float64) {
	l := len(c.children)
	lefts = make([]float64, l)
	rights = make([]float64, l)

	left := c.Geometry().Left
	width := c.Geometry().Width()
	switch {
	case width >= c.MaxWidth(): // scale according to MaxWidth
		scale := c.MaxWidth()
		for i, child := range c.children {
			scalePart := child.MaxWidth()
			curWidth := width * scalePart / scale
			right := left + curWidth

			lefts[i] = left
			rights[i] = right

			left = right
			width -= curWidth
			scale -= scalePart
		}
	case width >= c.BestWidth(): // scale according to BestWidth
		scale := c.BestWidth()
		for i, child := range c.children {
			scalePart := child.BestWidth()
			curWidth := width * scalePart / scale
			right := left + curWidth

			lefts[i] = left
			rights[i] = right

			left = right
			width -= curWidth
			scale -= scalePart
		}
	default: // scale according to MinWidth
		scale := c.MinWidth()
		for i, child := range c.children {
			scalePart := child.MinWidth()
			curWidth := width * scalePart / scale
			right := left + curWidth

			lefts[i] = left
			rights[i] = right

			left = right
			width -= curWidth
			scale -= scalePart
		}
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

func (c *HTable) Insert(control Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	if at < 0 {
		at = 0
	} else if at > len(c.children) {
		at = len(c.children)
	}
	c.BaseControl.SetParent(control, c)
	c.children = append(append(c.children[:at], control), c.children[at:]...)
	c.SetUPG(false)
}

func (c *HTable) Prepend(control Control) {
	c.Insert(control, 0)
}

func (c *HTable) Append(control Control) {
	c.Insert(control, len(c.children))
}

func (c *HTable) Remove(i int) Control {
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

func (c *HTable) NumColumns() int { return len(c.children) }

func (c *HTable) Children() []Control { return c.children }

func NewHTable(children ...Control) *HTable {
	r := &HTable{
		children: children,
	}
	for _, child := range children {
		r.BaseControl.SetParent(child, r)
	}
	return r
}
