package controls

import (
	"github.com/apaxa-go/gui"
	"github.com/apaxa-go/helper/mathh"
)

type HTable struct {
	gui.BaseControl
	children []gui.Control
}

func (c *HTable) ComputePossibleHorGeometry() (minWidth, maxWidth int) {
	for _, child := range c.children {
		minWidth += child.MinWidth()
		maxWidth += child.MaxWidth()
	}
	return
}

func (c *HTable) ComputePossibleVerGeometry() (minHeight, maxHeight int) {
	if len(c.children) > 0 {
		maxHeight = mathh.MaxInt
		for _, child := range c.children {
			minHeight = mathh.Max2Int(minHeight, child.MinWidth())
			maxHeight = mathh.Min2Int(maxHeight, child.MaxWidth())
		}
		maxHeight = mathh.Max2Int(minHeight, maxHeight)
	}
	return
}

func (c *HTable) Draw(canvas gui.Canvas, region gui.RectangleI) {
	for _, child := range c.children {
		// TODO draw only required children
		child.Draw(canvas, region)
	}
}

func (c *HTable) ProcessEvent(gui.Event) bool {
	// TODO
	return false
}

func (c *HTable) ComputeChildHorGeometry() (lefts, rights []int) {
	l := len(c.children)
	lefts = make([]int, l)
	rights = make([]int, l)

	left := c.Geometry().Left
	width := c.Geometry().Width()
	minWidth := c.MinWidth()
	for i, child := range c.children {
		cWidth := width * child.MinWidth() / minWidth
		right := left + cWidth

		lefts[i] = left
		rights[i] = right

		left = right
		width -= cWidth
		minWidth -= child.MinWidth()
	}
	return
}

func (c *HTable) ComputeChildVerGeometry() (tops, bottoms []int) {
	l := len(c.children)
	tops = make([]int, l)
	bottoms = make([]int, l)
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
	control.SetParent(c)
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
	control.SetParent(nil)
	c.children = append(c.children[:i], c.children[i+1:]...)
	c.SetUPG(false)
	return control
}

func (c *HTable) Children() []gui.Control { return c.children }

func NewHTable(children ...gui.Control) *HTable {
	return &HTable{
		children: children,
	}
}
