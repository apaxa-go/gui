// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/gui"
)

type Layers struct {
	gui.BaseControl
	layers []gui.Control
}

func (c *Layers) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	if len(c.layers) > 0 {
		maxWidth = gui.PosInfF64()
		for _, child := range c.layers {
			minWidth = gui.Max2Float64(minWidth, child.MinWidth())
			maxWidth = gui.Min2Float64(maxWidth, child.MaxWidth())
		}
		maxWidth = gui.Max2Float64(minWidth, maxWidth)
	}
	return
}

func (c *Layers) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	if len(c.layers) > 0 {
		maxHeight = gui.PosInfF64()
		for _, child := range c.layers {
			minHeight = gui.Max2Float64(minHeight, child.MinHeight())
			maxHeight = gui.Min2Float64(maxHeight, child.MaxHeight())
		}
		maxHeight = gui.Max2Float64(minHeight, maxHeight)
	}
	return
}

func (c *Layers) Draw(canvas gui.Canvas, region gui.RectangleF64) {
	for _, child := range c.layers {
		child.Draw(canvas, region)
	}
}
func (c *Layers) FocusCandidate(reverse bool, current gui.Control) gui.Control {
	l := len(c.layers)
	if l == 0 {
		return nil
	}
	switch {
	case current == nil && !reverse: // first
		return c.layers[0]
	case current == nil && reverse: // last
		return c.layers[l-1]
	default:
		i := 0
		for ; i < l && c.layers[i] != current; i++ {
		}
		if i == l { // not found
			return c.layers[0]
		}
		if reverse {
			i--
		} else {
			i++
		}
		if i < 0 || i >= l { // out of current control
			return nil
		}
		return c.layers[i]
	}
}
func (c *Layers) ComputeChildHorGeometry() (lefts, rights []float64) {
	l := len(c.layers)
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

func (c *Layers) ComputeChildVerGeometry() (tops, bottoms []float64) {
	l := len(c.layers)
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

func (c *Layers) Insert(control gui.Control, at int) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	if at < 0 {
		at = 0
	} else if at > len(c.layers) {
		at = len(c.layers)
	}
	gui.SetParent(control, c)
	c.layers = append(append(c.layers[:at], control), c.layers[at:]...)
	c.SetUPGIR(false)
	{
		// TODO do smth with this:
		// c.window.Hypervisor().UpdatePossibleHorGeometry(c, false)
		// c.window.Hypervisor().UpdatePossibleVerGeometry(c, false)
		// c.window.Hypervisor().UpdateChildHorGeometry(c)
		// c.window.Hypervisor().UpdateChildVerGeometry(c)
		// c.window.Hypervisor().InvalidateRegion(control)
	}
}

func (c *Layers) Prepend(control gui.Control) {
	c.Insert(control, 0)
}

func (c *Layers) Append(control gui.Control) {
	c.Insert(control, len(c.layers))
}

func (c *Layers) Remove(i int) gui.Control {
	if i < 0 {
		i = 0
	} else if i >= len(c.layers) {
		i = len(c.layers) - 1
	}
	control := c.layers[i]
	gui.SetParent(control, nil)
	c.layers = append(c.layers[:i], c.layers[i+1:]...)
	c.SetUPGIR(false)
	return control
}

func (c *Layers) NumLayers() int { return len(c.layers) }

func (c *Layers) Children() []gui.Control { return c.layers }

func NewLayers(layers ...gui.Control) *Layers {
	r := &Layers{
		layers: layers,
	}
	for _, child := range layers {
		gui.SetParent(child, r)
	}
	return r
}
