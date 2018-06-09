// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package controls

import (
	"github.com/apaxa-go/helper/mathh"
)

//replacer:ignore
//go:generate go run $GOPATH/src/github.com/apaxa-go/generator/replacer/main.go -- $GOFILE

// BUG(Anton Bekker): Table geometry computation is ugly if spanned cells are presented.
// Spanned cell may cause adding space to columns/rows which it covers.
// This space is not distributed between columns/rows - all space is assigned to the last covered column/row.
// Also best size of spanned cells is completely ignored.

type tableBandGeometry struct {
	min  float64
	best float64
	max  float64
}

type TableSpanDirectionState struct {
	Before uint8
	After  uint8
}

func (s TableSpanDirectionState) IsMaster() bool  { return s.Before == 0 }
func (s TableSpanDirectionState) IsSlave() bool   { return !s.IsMaster() }
func (s TableSpanDirectionState) IsSingle() bool  { return s.Before == 0 && s.After == 0 }
func (s TableSpanDirectionState) IsSpanned() bool { return !s.IsSingle() }

func (s TableSpanDirectionState) BeforeInt() int { return int(s.Before) }
func (s TableSpanDirectionState) AfterInt() int  { return int(s.After) }

func (s TableSpanDirectionState) Master(origin int) int { return origin - s.BeforeInt() }
func (s TableSpanDirectionState) Size() int             { return s.BeforeInt() + 1 + s.AfterInt() }
func (s TableSpanDirectionState) Span(origin int) (master, size int) {
	return s.Master(origin), s.Size()
}

type TableSpanState struct {
	Hor TableSpanDirectionState
	Ver TableSpanDirectionState
}

func (s TableSpanState) IsMaster() bool  { return s.Hor.IsMaster() && s.Ver.IsMaster() }
func (s TableSpanState) IsSlave() bool   { return !s.IsMaster() }
func (s TableSpanState) IsSingle() bool  { return s.Hor.IsSingle() && s.Ver.IsSingle() }
func (s TableSpanState) IsSpanned() bool { return !s.IsSingle() }

func (s TableSpanState) Master(origin PointI) PointI {
	return PointI{s.Hor.Master(origin.X), s.Ver.Master(origin.Y)}
}
func (s TableSpanState) Size() PointI                             { return PointI{s.Hor.Size(), s.Ver.Size()} }
func (s TableSpanState) Span(origin PointI) (master, size PointI) { return s.Master(origin), s.Size() }

type Table struct {
	BaseControl
	children        [][]Control
	span            [][]TableSpanState
	rowsGeometry    []tableBandGeometry
	columnsGeometry []tableBandGeometry
	childrenCount   int // real (non-nil) children count
}

func (c *Table) ColumnsCount() int { return len(c.columnsGeometry) }
func (c *Table) RowsCount() int    { return len(c.rowsGeometry) }

// Spanned cells N*M counts as 1 non-nil child and (N*M-1) nil children.
func (c *Table) NonNilChildrenCount() int { return c.childrenCount }

// ChildrenCount() returns number or children including nil children.
// For table of size N*M results is exactly (N*M) even if there are some spanned children.
func (c *Table) ChildrenCount() int { return c.RowsCount() * c.ColumnsCount() }

func (c *Table) HasNilOrSpannedChild() bool { return c.NonNilChildrenCount() != c.ChildrenCount() }

func (c *Table) Children() []Control {
	children := make([]Control, 0, c.ChildrenCount())
	for _, row := range c.children {
		children = append(children, row...)
	}
	return children
	/*
		if !c.HasNilOrSpannedChild() { // fast path for all non-nil
			for _, row := range c.children {
				children = append(children, row...)
			}
			return children
		}

		for _, row := range c.children {
			for _, child := range row {
				if child != nil {
					children = append(children, child)
				}
			}
		}
		return children
	*/
}

func (c *Table) horSpan(y, x int) int {
	// TODO
	return 1
}

// is spanned and is not span base
func (c *Table) isHorSpanned(y, x int) bool {
	return false
}

func (c *Table) verSpan(y, x int) int {
	// TODO
	return 1
}

// is spanned and is not span base
func (c *Table) isVerSpanned(y, x int) bool {
	return false
}

type buttonSpannedGeometry struct {
	min float64
	max float64
}

//replacer:replace
//replacer:old Hor	Width	column	"[iRow][iColumn]"	"(iRow, iColumn)"	Column	Row
//replacer:new Ver	Height	row		"[iRow][iColumn]"	"(iRow, iColumn)"	Row		Column

func (c *Table) ComputePossibleHorGeometry() (minWidth, bestWidth, maxWidth float64) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if lRow == 0 || lColumn == 0 {
		return 0, 0, 0
	}

	var prevMin, prevMax float64
	spannedGeometries := make(map[int]buttonSpannedGeometry)
	for iColumn := 0; iColumn < lColumn; iColumn++ { // column by column
		// values for current column
		min := float64(0)
		best := float64(0)
		max := mathh.PositiveInfFloat64()

		var effectiveCount float64 // Count of non-nil non-spanned cells in column
		for iRow := 0; iRow < lRow; iRow++ {
			span := c.span[iRow][iColumn].Hor
			if c.children[iRow][iColumn] == nil || span.IsSlave() {
				continue
			}
			cMinWidth := c.children[iRow][iColumn].MinWidth()
			cBestWidth := c.children[iRow][iColumn].MinWidth()
			cMaxWidth := c.children[iRow][iColumn].MaxWidth()
			if span.IsSingle() { // single column cell
				effectiveCount++
				min = mathh.Max2Float64(min, cMinWidth)
				best += cBestWidth
				max = mathh.Min2Float64(max, cMaxWidth)
			} else { // spanned cell
				im := iColumn + span.AfterInt()
				if g, ok := spannedGeometries[im]; ok {
					g.min = mathh.Max2Float64(g.min, minWidth+min)
					g.max = mathh.Min2Float64(g.max, maxWidth+max)
				} else {
					spannedGeometries[im] = buttonSpannedGeometry{minWidth + min, maxWidth + max}
				}
			}
		}

		if spannedGeometry, spannedExists := spannedGeometries[iColumn]; spannedExists { // Spanned geometry exists.
			// Convert saved spanned geometry from absolute to column-only values.
			spannedGeometry.min -= prevMin
			spannedGeometry.max -= prevMax

			if effectiveCount == 0 { // Only spanned geometry exists for column.
				min = mathh.Max2Float64(spannedGeometry.min, 0) // We need to be sure here what min is non negative (max will be fixed at normalization).
				max = spannedGeometry.max
			} else { // Both spanned geometry and column's geometry exists.
				best /= effectiveCount
				min = mathh.Max2Float64(min, spannedGeometry.min)
				max = mathh.Min2Float64(max, spannedGeometry.max)
			}
		} else { // No spanned geometry.
			if effectiveCount == 0 {
				max = 0
			} else {
				best /= effectiveCount
			}
		}

		// Normalization
		max = mathh.Max2Float64(max, min)
		best = mathh.Max2Float64(min, best)
		best = mathh.Min2Float64(max, best)

		c.columnsGeometry[iColumn].min = min
		c.columnsGeometry[iColumn].best = best
		c.columnsGeometry[iColumn].max = max

		prevMin = minWidth
		prevMax = maxWidth

		minWidth += min
		bestWidth += best
		maxWidth += max
	}
	return
}

//replacer:ignore

func (c *Table) rowColumnToIndex(rowIndex, columnIndex int) (childIndex int) {
	return rowIndex*c.ColumnsCount() + columnIndex
}

//replacer:replace
//replacer:old Hor	left	Left	right	Right	width	Width	rowColumn	column	"(iRow, iColumn)"	Column	Row
//replacer:new Ver	top		Top		bottom	Bottom	height	Height	rowColumn	row		"(iRow, iColumn)"	Row		Column

func (c *Table) ComputeChildHorGeometry() (lefts, rights []float64) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if lRow == 0 || lColumn == 0 {
		return nil, nil
	}

	lefts = make([]float64, lRow*lColumn)
	rights = make([]float64, lRow*lColumn)

	left := c.Geometry().Left
	width := c.Geometry().Width()
	switch {
	case width >= c.MaxWidth(): // scale according to MaxWidth
		scale := c.MaxWidth()
		for iColumn, geometry := range c.columnsGeometry {
			scalePart := geometry.max
			curWidth := width * scalePart / scale
			right := left + curWidth

			for iRow := 0; iRow < lRow; iRow++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				lefts[i] = left
				rights[i] = right
			}

			left = right
			width -= curWidth
			scale -= scalePart
		}
	case width >= c.BestWidth(): // scale according to BestWidth
		scale := c.BestWidth()
		for iColumn, geometry := range c.columnsGeometry {
			scalePart := geometry.best
			curWidth := width * scalePart / scale
			right := left + curWidth

			for iRow := 0; iRow < lRow; iRow++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				lefts[i] = left
				rights[i] = right
			}

			left = right
			width -= curWidth
			scale -= scalePart
		}
	default: // scale according to MinWidth
		scale := c.MinWidth()
		for iColumn, geometry := range c.columnsGeometry {
			scalePart := geometry.min
			curWidth := width * scalePart / scale
			right := left + curWidth

			for iRow := 0; iRow < lRow; iRow++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				lefts[i] = left
				rights[i] = right
			}

			left = right
			width -= curWidth
			scale -= scalePart
		}
	}
	return
}

//replacer:ignore

func (c *Table) Draw(canvas Canvas, region RectangleF64) {
	// TODO draw only required children
	for _, row := range c.children {
		for _, child := range row {
			child.Draw(canvas, region)
		}
	}
}

func (c *Table) FocusCandidate(reverse bool, current Control) Control {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if lRow == 0 || lColumn == 0 {
		return nil
	}
	switch {
	case current == nil && !reverse: // first
		for _, row := range c.children {
			for _, child := range row {
				if child != nil {
					return child
				}
			}
		}
		return nil
	case current == nil && reverse: // last
		for iRow := lRow - 1; iRow >= 0; iRow-- {
			for iColumn := lColumn - 1; iColumn >= 0; iColumn-- {
				if child := c.children[iRow][iColumn]; child != nil {
					return child
				}
			}
		}
		return nil
	case !reverse: // next
		found := false
		for _, row := range c.children {
			for _, child := range row {
				if found && child != nil {
					return child
				} else if !found {
					found = child == current
				}
			}
		}
		return nil
	default: // previous
		found := false
		for iRow := lRow - 1; iRow >= 0; iRow-- {
			for iColumn := lColumn - 1; iColumn >= 0; iColumn-- {
				child := c.children[iRow][iColumn]
				if found && child != nil {
					return child
				} else if !found {
					found = child == current
				}
			}
		}
		return nil
	}
}

// For non nil control returns false if target cell is span slave or if iRow or iColumn is invalid.
// If control is nil then Detach is performed and its "ok" result will be returned.
func (c *Table) Set(control Control, iRow, iColumn int) (ok bool) {
	// TODO what if control already assigned to some other/the same parent ?
	// TODO control.geometry must be ={0,0,-1,-1} & min/maxSize must be ={0,-1} (for simplify Hypervisor calling)
	if control == nil {
		_, ok = c.Detach(iRow, iColumn)
		return
	}
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	ok = iRow >= 0 && iRow < lRow && iColumn >= 0 && iColumn < lColumn
	ok = ok && c.span[iRow][iColumn].IsMaster()
	if !ok {
		return
	}

	if c.children[iRow][iColumn] != nil {
		c.BaseControl.SetParent(c.children[iRow][iColumn], nil)
	} else {
		c.childrenCount++
	}

	c.BaseControl.SetParent(control, c)
	c.children[iRow][iColumn] = control
	c.SetUPG(false) // TODO why not recursive?
	return
}

// Returns "ok" == false if target cell is span slave or if iRow or iColumn is invalid.
// Returns (nil, true) if target cell is already nil.
func (c *Table) Detach(iRow, iColumn int) (control Control, ok bool) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	ok = iRow >= 0 && iRow < lRow && iColumn >= 0 && iColumn < lColumn
	ok = ok && c.span[iRow][iColumn].IsMaster()
	if !ok {
		return
	}
	control = c.children[iRow][iColumn]
	if control == nil {
		return
	}
	c.BaseControl.SetParent(control, nil)
	c.children[iRow][iColumn] = nil
	c.childrenCount--
	c.SetUPG(false)
	return
}

//replacer:replace
//replacer:old Ver	[iRow+1][iColumn]	[ciRow+1][iColumn]	[ciRow][iColumn]	[iRow	[iColumn	[ciRow	[ciColumn	row		column	Row		Column	Bottom
//replacer:new Hor	[iRow][iColumn+1]	[iRow][ciColumn+1]	[iRow][ciColumn]	[iRow	[iColumn	[ciRow	[ciColumn	column	row		Column	Row		Right

func (c *Table) fixNewRowSpan(iRow int, stretchSpan, bindSpanToBottom bool) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if iRow == 0 || iRow == lRow-1 {
		return
	}
	switch {
	case stretchSpan:
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow+1][iColumn]
			if span.Ver.IsMaster() {
				continue
			}
			c.span[iRow][iColumn] = span

			for ciRow := iRow - span.Ver.BeforeInt(); ciRow <= iRow; ciRow++ {
				c.span[ciRow][iColumn].Ver.After++
			}
			for ciRow := iRow + 1; ciRow <= iRow+span.Ver.AfterInt()+1; ciRow++ {
				c.span[ciRow][iColumn].Ver.Before++
			}
		}
	case bindSpanToBottom:
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow+1][iColumn].Ver
			if span.IsMaster() {
				continue
			}

			for ciRow := iRow - span.BeforeInt(); ciRow < iRow; ciRow++ {
				c.span[ciRow+1][iColumn] = c.span[ciRow][iColumn]
			}

			ciRow := iRow - span.BeforeInt()
			c.span[ciRow][iColumn] = TableSpanState{}
		}
	default: // bind span to top
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow+1][iColumn].Ver
			if span.IsMaster() {
				continue
			}

			for ciRow := iRow; ciRow <= iRow+span.AfterInt(); ciRow++ {
				c.span[ciRow][iColumn] = c.span[ciRow+1][iColumn]
			}

			ciRow := iRow + span.AfterInt() + 1
			c.span[ciRow][iColumn] = TableSpanState{}
		}
	}
}

//replacer:ignore

// bindSpanToBottom is ignored id stretchSpan is true.
func (c *Table) InsertRowExtended(at int, stretchSpan, bindSpanToBottom bool) {
	lColumn := c.ColumnsCount()
	at = mathh.Max2Int(at, 0)
	at = mathh.Min2Int(at, c.RowsCount())
	c.children = append(append(c.children[:at], make([]Control, lColumn)), c.children[at:]...)
	c.span = append(append(c.span[:at], make([]TableSpanState, lColumn)), c.span[at:]...)
	c.fixNewRowSpan(at, stretchSpan, bindSpanToBottom)
	c.SetUPG(false) // Because inserted row may have intersection with spanned cell.
}

//replacer:replace
//replacer:old Row
//replacer:new Column

func (c *Table) InsertRow(at int) {
	c.InsertRowExtended(at, true, false)
}

func (c *Table) PrependRow() {
	c.InsertRow(0)
}

func (c *Table) AppendRow() {
	c.InsertRow(c.RowsCount())
}

//replacer:ignore

func (c *Table) RemoveRow(iRow int) []Control {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if iRow < 0 || iRow >= lRow {
		return nil
	}
	r := c.children[iRow]
	// TODO
	return r
}

func NewTable(children ...Control) *Table { // TODO
	r := &Table{
		children: children,
	}
	for _, child := range children {
		r.BaseControl.SetParent(child, r)
	}
	return r
}
