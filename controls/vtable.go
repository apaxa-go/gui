// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE
//replacer:replace
//replacer:old VTable	Row		Hor	Ver	Width	width	Height	height	left	Left	top		Top		right	Right	bottom	Bottom
//replacer:new HTable	Column	Ver	Hor	Height	height	Width	width	top		Top		left	Left	bottom	Bottom	right	Right

type VTable struct {
	BaseControl
	children []Control
}

func (c *VTable) NumRows() int { return len(c.children) }

func (c *VTable) Children() []Control { return c.children }

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

// Compute child geometry if height >= maximum height.
// This condition also means that all children's max height is not inf.
func (c *VTable) computeChildVerGeometryMax() (tops, bottoms []float64) {
	tops = make([]float64, len(c.children))
	bottoms = make([]float64, len(c.children))

	top := c.Geometry().Top
	height := c.Geometry().Height()
	scale := c.MaxHeight()
	for i, child := range c.children {
		scalePart := child.MaxHeight()
		curHeight := height * scalePart / scale

		tops[i] = top
		top += curHeight
		bottoms[i] = top

		height -= curHeight
		scale -= scalePart
	}
	return
}

// Compute child geometry if height >= best height and at least 1 child has max height = inf.
// In this case each child with non-inf max height will be sized to its best height,
// and estimated height will be split between children with max height = inf.
func (c *VTable) computeChildVerGeometryBestWithMaxInf(childrenWithMaxInf []int, childrenWithMaxNotInfSumBest float64) (tops, bottoms []float64) {
	tops = make([]float64, len(c.children))
	bottoms = make([]float64, len(c.children))

	top := c.Geometry().Top
	split := c.Geometry().Height() - childrenWithMaxNotInfSumBest
	scale := c.BestHeight() - childrenWithMaxNotInfSumBest
	for i, child := range c.children {
		var part float64
		if len(childrenWithMaxInf) > 0 && i == childrenWithMaxInf[0] {
			scalePart := child.BestHeight()
			if scale > 0 {
				part = split * scalePart / scale
			}
			split -= part
			scale -= scalePart
		} else {
			part = child.BestHeight()
		}
		tops[i] = top
		top += part
		bottoms[i] = top
	}

	return
}

// Compute child geometry if height < maximum height and height >= best height, and all children max height != inf.
func (c *VTable) computeChildVerGeometryBest() (tops, bottoms []float64) {
	tops = make([]float64, len(c.children))
	bottoms = make([]float64, len(c.children))

	top := c.Geometry().Top
	split := c.Geometry().Height() - c.BestHeight()
	scale := c.MaxHeight() - c.BestHeight()
	for i, child := range c.children {
		scalePart := child.MaxHeight() - child.BestHeight()
		var part float64
		if scale > 0 {
			part = split * scalePart / scale
		}

		split -= part
		scale -= scalePart

		tops[i] = top
		top += child.BestHeight() + part
		bottoms[i] = top
	}
	return
}

// Compute child geometry if height < best height.
func (c *VTable) computeChildVerGeometryMin() (tops, bottoms []float64) {
	tops = make([]float64, len(c.children))
	bottoms = make([]float64, len(c.children))

	top := c.Geometry().Top
	split := c.Geometry().Height() - c.MinHeight()
	scale := c.BestHeight() - c.MinHeight()
	for i, child := range c.children {
		scalePart := child.BestHeight() - child.MinHeight()
		var part float64
		if scale > 0 {
			part = split * scalePart / scale
		}

		split -= part
		scale -= scalePart

		tops[i] = top
		top += child.MinHeight() + part
		bottoms[i] = top
	}
	return
}

func (c *VTable) ComputeChildVerGeometry() (tops, bottoms []float64) {
	height := c.Geometry().Height()
	if height > c.MaxHeight() {
		return c.computeChildVerGeometryMax()
	}
	if height > c.BestHeight() {
		childrenWithMaxInf := make([]int, 0, len(c.children))
		var childrenWithMaxNotInfSumBest float64
		for i, child := range c.children {
			if child.MaxHeight() == mathh.PositiveInfFloat64() {
				childrenWithMaxInf = append(childrenWithMaxInf, i)
			} else {
				childrenWithMaxNotInfSumBest += child.BestHeight()
			}
		}
		if len(childrenWithMaxInf) > 0 {
			return c.computeChildVerGeometryBestWithMaxInf(childrenWithMaxInf, childrenWithMaxNotInfSumBest)
		}
		return c.computeChildVerGeometryBest()
	}
	return c.computeChildVerGeometryMin()
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

func (c *VTable) Insert(control Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	at = mathh.Max2Int(at, 0)
	at = mathh.Min2Int(at, c.NumRows())
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
	if i < 0 || i >= c.NumRows() {
		return nil
	}
	control := c.children[i]
	c.BaseControl.SetParent(control, nil)
	c.children = append(c.children[:i], c.children[i+1:]...)
	c.SetUPG(false)
	return control
}

func NewVTable(children ...Control) *VTable {
	r := &VTable{
		children: children,
	}
	for _, child := range children {
		r.BaseControl.SetParent(child, r)
	}
	return r
}
