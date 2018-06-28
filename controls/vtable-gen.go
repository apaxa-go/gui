// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

type HTable struct {
	BaseControl
	children []Control
}

func (c *HTable) NumColumns() int { return len(c.children) }

func (c *HTable) Children() []Control { return c.children }

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

func (c *HTable) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	for _, child := range c.children {
		minWidth += child.MinWidth()
		bestWidth += child.BestWidth()
		maxWidth += child.MaxWidth()
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

// Compute child geometry if width >= maximum width.
// This condition also means that all children's max width is not inf.
func (c *HTable) computeChildHorGeometryMax() (lefts, rights []float64) {
	lefts = make([]float64, len(c.children))
	rights = make([]float64, len(c.children))

	left := c.Geometry().Left
	width := c.Geometry().Width()
	scale := c.MaxWidth()
	for i, child := range c.children {
		scalePart := child.MaxWidth()
		curWidth := width * scalePart / scale

		lefts[i] = left
		left += curWidth
		rights[i] = left

		width -= curWidth
		scale -= scalePart
	}
	return
}

// Compute child geometry if width >= best width and at least 1 child has max width = inf.
// In this case each child with non-inf max width will be sized to its best width,
// and estimated width will be split between children with max width = inf.
func (c *HTable) computeChildHorGeometryBestWithMaxInf(childrenWithMaxInf []int, childrenWithMaxNotInfSumBest float64) (lefts, rights []float64) {
	lefts = make([]float64, len(c.children))
	rights = make([]float64, len(c.children))

	left := c.Geometry().Left
	split := c.Geometry().Width() - childrenWithMaxNotInfSumBest
	scale := c.BestWidth() - childrenWithMaxNotInfSumBest
	for i, child := range c.children {
		var part float64
		if len(childrenWithMaxInf) > 0 && i == childrenWithMaxInf[0] {
			scalePart := child.BestWidth()
			if scale > 0 {
				part = split * scalePart / scale
			}
			split -= part
			scale -= scalePart
		} else {
			part = child.BestWidth()
		}
		lefts[i] = left
		left += part
		rights[i] = left
	}

	return
}

// Compute child geometry if width < maximum width and width >= best width, and all children max width != inf.
func (c *HTable) computeChildHorGeometryBest() (lefts, rights []float64) {
	lefts = make([]float64, len(c.children))
	rights = make([]float64, len(c.children))

	left := c.Geometry().Left
	split := c.Geometry().Width() - c.BestWidth()
	scale := c.MaxWidth() - c.BestWidth()
	for i, child := range c.children {
		scalePart := child.MaxWidth() - child.BestWidth()
		var part float64
		if scale > 0 {
			part = split * scalePart / scale
		}

		split -= part
		scale -= scalePart

		lefts[i] = left
		left += child.BestWidth() + part
		rights[i] = left
	}
	return
}

// Compute child geometry if width < best width.
func (c *HTable) computeChildHorGeometryMin() (lefts, rights []float64) {
	lefts = make([]float64, len(c.children))
	rights = make([]float64, len(c.children))

	left := c.Geometry().Left
	split := c.Geometry().Width() - c.MinWidth()
	scale := c.BestWidth() - c.MinWidth()
	for i, child := range c.children {
		scalePart := child.BestWidth() - child.MinWidth()
		var part float64
		if scale > 0 {
			part = split * scalePart / scale
		}

		split -= part
		scale -= scalePart

		lefts[i] = left
		left += child.MinWidth() + part
		rights[i] = left
	}
	return
}

func (c *HTable) ComputeChildHorGeometry() (lefts, rights []float64) {
	width := c.Geometry().Width()
	if width > c.MaxWidth() {
		return c.computeChildHorGeometryMax()
	}
	if width > c.BestWidth() {
		childrenWithMaxInf := make([]int, 0, len(c.children))
		var childrenWithMaxNotInfSumBest float64
		for i, child := range c.children {
			if child.MaxWidth() == mathh.PositiveInfFloat64() {
				childrenWithMaxInf = append(childrenWithMaxInf, i)
			} else {
				childrenWithMaxNotInfSumBest += child.BestWidth()
			}
		}
		if len(childrenWithMaxInf) > 0 {
			return c.computeChildHorGeometryBestWithMaxInf(childrenWithMaxInf, childrenWithMaxNotInfSumBest)
		}
		return c.computeChildHorGeometryBest()
	}
	return c.computeChildHorGeometryMin()
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

func (c *HTable) Insert(control Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	at = mathh.Max2Int(at, 0)
	at = mathh.Min2Int(at, c.NumColumns())
	c.BaseControl.SetParent(control, c)
	c.children = append(append(c.children[:at], control), c.children[at:]...)
	c.SetUPG(false) // TODO why not recursive?
}

func (c *HTable) Prepend(control Control) {
	c.Insert(control, 0)
}

func (c *HTable) Append(control Control) {
	c.Insert(control, len(c.children))
}

func (c *HTable) Remove(i int) Control {
	if i < 0 || i >= c.NumColumns() {
		return nil
	}
	control := c.children[i]
	c.BaseControl.SetParent(control, nil)
	c.children = append(c.children[:i], c.children[i+1:]...)
	c.SetUPG(false)
	return control
}

func NewHTable(children ...Control) *HTable {
	r := &HTable{
		children: children,
	}
	for _, child := range children {
		r.BaseControl.SetParent(child, r)
	}
	return r
}
