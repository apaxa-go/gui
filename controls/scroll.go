// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

type ScrollState struct {
	enabled     bool
	minimumSize float64
	shift       float64 // Shift from beginning of child to beginning of visible area. Non negative. 0 means that beginning of view is visible.
}

func (s *ScrollState) validateShift(containerSize, childSize float64) {
	if s.shift < 0 || childSize < containerSize {
		s.shift = 0
	} else if s.shift > childSize-containerSize {
		s.shift = childSize - containerSize
	}
}

type Scroll struct {
	BaseControl
	hor   ScrollState
	ver   ScrollState
	child Control
}

func (c *Scroll) Children() []Control {
	if c.child == nil {
		return nil
	}
	return []Control{c.child}
}

func (c *Scroll) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	if c.child == nil {
		return 0, 0, 0
	}
	minWidth = c.child.MinWidth()
	if c.hor.enabled && c.hor.minimumSize < minWidth {
		minWidth = c.hor.minimumSize
	}
	bestWidth = c.child.BestWidth()
	maxWidth = c.child.MaxWidth()
	return
}

func (c *Scroll) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	if c.child == nil {
		return 0, 0, 0
	}
	minHeight = c.child.MinHeight()
	if c.ver.enabled && c.ver.minimumSize < minHeight {
		minHeight = c.ver.minimumSize
	}
	bestHeight = c.child.BestHeight()
	maxHeight = c.child.MaxHeight()
	return
}

func (c *Scroll) ComputeChildHorGeometry() (lefts, rights []float64) {
	if c.child == nil {
		return nil, nil
	}
	if !c.hor.enabled { // disabled by settings
		return []float64{c.Geometry().Left}, []float64{c.Geometry().Right}
	}

	containerSize := c.Geometry().Width()
	childSize := c.child.BestWidth()
	if containerSize >= childSize { // disabled by sizes
		return []float64{c.Geometry().Left}, []float64{c.Geometry().Right}
	}

	c.hor.validateShift(containerSize, childSize)
	left := c.Geometry().Left - c.hor.shift
	right := left + childSize
	return []float64{left}, []float64{right}
}

func (c *Scroll) ComputeChildVerGeometry() (tops, bottoms []float64) {
	if c.child == nil {
		return nil, nil
	}
	if !c.ver.enabled { // disabled by settings
		return []float64{c.Geometry().Top}, []float64{c.Geometry().Bottom}
	}

	containerSize := c.Geometry().Height()
	childSize := c.child.BestHeight()
	if containerSize >= childSize { // disabled by sizes
		return []float64{c.Geometry().Top}, []float64{c.Geometry().Bottom}
	}

	c.ver.validateShift(containerSize, childSize)
	top := c.Geometry().Top - c.ver.shift
	bottom := top + childSize
	return []float64{top}, []float64{bottom}
}

func (c *Scroll) Draw(canvas Canvas, region RectangleF64) {
	// TODO draw scrolls itself
	canvas.SaveState()
	canvas.ClipToRectangle(c.Geometry())
	if c.child != nil { // TODO call child draw from window method directly ?
		c.child.Draw(canvas, region)
	}
	canvas.RestoreState()
}

func (c *Scroll) FocusCandidate(reverse bool, current Control) Control {
	switch {
	//
	// Forward
	//
	case !reverse && current == nil: // First focus
		return c
	case !reverse && current == c: // Next after control itself
		if c.child != nil {
			return c.child
		}
		fallthrough
	case !reverse && current == c.child: // Next after child
		return nil
	//
	// Backward
	//
	case reverse && current == nil: // Last focus
		if c.child != nil {
			return c.child
		}
		fallthrough
	case reverse && current == c.child: // Previous before child
		return c
	case reverse && current == c: // Previous before control itself
		return nil
	//
	// Fallback - unexpected case
	//
	default:
		return c
	}
}

func (c *Scroll) HorState() ScrollState { return c.hor }
func (c *Scroll) VerState() ScrollState { return c.ver }

func (c *Scroll) enableHor(enable bool) {
	if c.hor.enabled == enable {
		return
	}
	c.hor.enabled = enable
	{ // TODO not sure
		c.GeometryHypervisorPause()
		c.SetUPHG(false)
		c.SetUCHG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
}
func (c *Scroll) EnableHor()  { c.enableHor(true) }
func (c *Scroll) DisableHor() { c.enableHor(false) }

func (c *Scroll) enableVer(enable bool) {
	if c.ver.enabled == enable {
		return
	}
	c.ver.enabled = enable
	{ // TODO not sure
		c.GeometryHypervisorPause()
		c.SetUPVG(false)
		c.SetUCVG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
}
func (c *Scroll) EnableVer()  { c.enableVer(true) }
func (c *Scroll) DisableVer() { c.enableVer(false) }

func (c *Scroll) SetMinimumVisibleWidth(width float64) {
	width = mathh.Max2Float64(0, width)
	if c.hor.minimumSize == width {
		return
	}
	c.hor.minimumSize = width
	{ // TODO not sure
		c.GeometryHypervisorPause()
		c.SetUPHG(false)
		c.SetUCHG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
}

func (c *Scroll) SetMinimumVisibleHeight(height float64) {
	height = mathh.Max2Float64(0, height)
	if c.ver.minimumSize == height {
		return
	}
	c.ver.minimumSize = height
	{ // TODO not sure
		c.GeometryHypervisorPause()
		c.SetUPVG(false)
		c.SetUCVG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
}

func (c *Scroll) ScrollHorTo(pos float64) (hasEffect bool) {
	if !c.hor.enabled || c.child == nil {
		return false
	}
	oldShift := c.hor.shift
	c.hor.shift = pos
	c.hor.validateShift(c.Geometry().Width(), c.child.MinWidth()) // TODO use recommended instead of minimum size
	hasEffect = oldShift != c.hor.shift
	if hasEffect {
		c.GeometryHypervisorPause()
		c.SetUCHG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
	return
}

func (c *Scroll) ScrollVerTo(pos float64) (hasEffect bool) {
	if !c.ver.enabled || c.child == nil {
		return false
	}
	oldShift := c.ver.shift
	c.ver.shift = pos
	c.ver.validateShift(c.Geometry().Height(), c.child.MinHeight()) // TODO use recommended instead of minimum size
	hasEffect = oldShift != c.ver.shift
	if hasEffect {
		c.GeometryHypervisorPause()
		c.SetUCVG()
		c.SetIR()
		c.GeometryHypervisorResume()
	}
	return
}

func (c *Scroll) ScrollHor(delta float64) (hasEffect bool) {
	return c.ScrollHorTo(c.hor.shift + delta)
}
func (c *Scroll) ScrollVer(delta float64) (hasEffect bool) {
	return c.ScrollVerTo(c.ver.shift + delta)
}

func (c *Scroll) Child() Control { return c.child }
func (c *Scroll) SetChild(child Control) {
	if c.child != nil {
		c.BaseControl.SetParent(c.child, nil) // TODO move SetParent to some class as static method.
	}
	c.BaseControl.SetParent(child, c)
	c.child = child
	c.SetUPGIR(true)
}

func (c *Scroll) OnScrollEvent(event ScrollEvent) (processed bool) {
	return c.ScrollHor(event.Delta.X) || c.ScrollVer(event.Delta.Y)
}

func NewScroll(child Control, minimumVisibleWidth, minimumVisibleHeight float64) *Scroll {
	var hor, ver ScrollState
	hor.enabled = minimumVisibleWidth >= 0
	if hor.enabled {
		hor.minimumSize = minimumVisibleWidth
	}
	ver.enabled = minimumVisibleHeight >= 0
	if ver.enabled {
		ver.minimumSize = minimumVisibleHeight
	}
	r := &Scroll{
		hor: hor,
		ver: ver,
	}
	r.SetChild(child)
	return r
}
