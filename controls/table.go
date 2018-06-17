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

const TableSpanMaxSize = mathh.MaxUint8 + 1

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

func (s TableSpanState) Master(iRow, iColumn int) (mRow, mColumn int) {
	return s.Ver.Master(iRow), s.Hor.Master(iColumn)
}
func (s TableSpanState) Size() (sizeRow, sizeColumn int) { return s.Ver.Size(), s.Hor.Size() }
func (s TableSpanState) Span(iRow, iColumn int) (mRow, mColumn, sizeRow, sizeColumn int) {
	mRow, mColumn = s.Master(iRow, iColumn)
	sizeRow, sizeColumn = s.Size()
	return
}

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

// ChildrenCount() returns number or children.
// Spanned cells N*M counts as 1 non-nil child and (N*M-1) nil children.
func (c *Table) ChildrenCount() int { return c.childrenCount }

func (c *Table) HasNilOrSpannedChild() bool {
	return c.ChildrenCount() != c.RowsCount()*c.ColumnsCount()
}

func (c *Table) Children() []Control {
	children := make([]Control, c.ChildrenCount())

	if !c.HasNilOrSpannedChild() { // fast path for all non-nil
		for _, row := range c.children {
			children = append(children, row...)
		}
		return children
	}

	i := 0
	for _, row := range c.children {
		for _, child := range row {
			if child != nil {
				children[i] = child
				i++
			}
		}
	}
	return children
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
//replacer:old Hor	left	Left	right	Right	width	Width	rowColumn	row		column	"(iRow, iColumn)"	Column	Row
//replacer:new Ver	top		Top		bottom	Bottom	height	Height	rowColumn	column	row		"(iRow, iColumn)"	Row		Column

// computeChildHorGeometry computes lefts and right for single row as if there are no spanned cells.
func (c *Table) computeChildHorGeometry() (lefts, rights []float64) {
	lColumn := c.ColumnsCount()

	lefts = make([]float64, lColumn)
	rights = make([]float64, lColumn)

	left := c.Geometry().Left
	width := c.Geometry().Width()
	switch {
	case width >= c.MaxWidth(): // scale according to MaxWidth
		scale := c.MaxWidth()
		for iColumn, geometry := range c.columnsGeometry {
			scalePart := geometry.max
			curWidth := width * scalePart / scale
			right := left + curWidth

			lefts[iColumn] = left
			rights[iColumn] = right

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

			lefts[iColumn] = left
			rights[iColumn] = right

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

			lefts[iColumn] = left
			rights[iColumn] = right

			left = right
			width -= curWidth
			scale -= scalePart
		}
	}
	return
}

//replacer:ignore

func (c *Table) ComputeChildHorGeometry() (lefts, rights []float64) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if lRow == 0 || lColumn == 0 {
		return nil, nil
	}

	lefts = make([]float64, c.ChildrenCount())
	rights = make([]float64, c.ChildrenCount())

	rowLefts, rowRights := c.computeChildHorGeometry()
	i := 0
	for iRow := 0; iRow < lRow; iRow++ {
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow][iColumn].Hor
			if c.children[iRow][iColumn] == nil || span.IsSlave() {
				continue
			}
			lefts[i] = rowLefts[iColumn]
			rights[i] = rowRights[iColumn+span.AfterInt()]
			i++
		}
	}
	return
}

func (c *Table) ComputeChildVerGeometry() (tops, bottoms []float64) {
	lColumn, lRow := c.ColumnsCount(), c.RowsCount()
	if lColumn == 0 || lRow == 0 {
		return nil, nil
	}

	tops = make([]float64, c.ChildrenCount())
	bottoms = make([]float64, c.ChildrenCount())

	columnTops, columnBottoms := c.computeChildVerGeometry()
	i := 0
	for iRow := 0; iRow < lRow; iRow++ {
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow][iColumn].Ver
			if c.children[iRow][iColumn] == nil || span.IsSlave() {
				continue
			}
			tops[i] = columnTops[iRow]
			bottoms[i] = columnBottoms[iRow+span.AfterInt()]
			i++
		}
	}
	return
}

func (c *Table) Draw(canvas Canvas, region RectangleF64) {
	// TODO draw only required children
	for _, row := range c.children {
		for _, child := range row {
			if child != nil {
				child.Draw(canvas, region)
			}
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

// For non nil control returns false if target cell is span's slave or if iRow or iColumn is invalid.
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

func (c *Table) detach(iRow, iColumn int, setNil bool) (control Control) {
	control = c.children[iRow][iColumn]
	if control == nil {
		return
	}
	c.BaseControl.SetParent(control, nil)
	if setNil {
		c.children[iRow][iColumn] = nil
	}
	c.childrenCount--
	c.SetUPG(false)
	return
}

// Returns false if target cell is span's slave or if iRow or iColumn is invalid.
// Returns (nil, true) if target cell is already nil.
func (c *Table) Detach(iRow, iColumn int) (control Control, ok bool) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	ok = iRow >= 0 && iRow < lRow && iColumn >= 0 && iColumn < lColumn
	ok = ok && c.span[iRow][iColumn].IsMaster()
	if !ok {
		return
	}
	control = c.detach(iRow, iColumn, true)
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
			if c.span[ciRow][iColumn].IsMaster() { // move master's child if required
				c.children[ciRow+1][iColumn] = c.children[ciRow][iColumn]
			}
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
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	at = mathh.Max2Int(at, 0)
	at = mathh.Min2Int(at, lRow)
	c.children = append(append(c.children[:at], make([]Control, lColumn)), c.children[at:]...)
	c.span = append(append(c.span[:at], make([]TableSpanState, lColumn)), c.span[at:]...)
	c.rowsGeometry = append(append(c.rowsGeometry[:at], tableBandGeometry{}), c.rowsGeometry[at:]...)
	c.fixNewRowSpan(at, stretchSpan, bindSpanToBottom)
	c.SetUPG(false) // Because inserted row may have intersection with spanned cell.
}

// bindSpanToRight is ignored id stretchSpan is true.
func (c *Table) InsertColumnExtended(at int, stretchSpan, bindSpanToRight bool) {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	at = mathh.Max2Int(at, 0)
	at = mathh.Min2Int(at, lColumn)
	for iRow := 0; iRow < lRow; iRow++ {
		c.children[iRow] = append(append(c.children[iRow][:at], nil), c.children[iRow][at:]...)
		c.span[iRow] = append(append(c.span[iRow][:at], TableSpanState{}), c.span[iRow][at:]...)
	}
	c.columnsGeometry = append(append(c.columnsGeometry[:at], tableBandGeometry{}), c.columnsGeometry[at:]...)
	c.fixNewColumnSpan(at, stretchSpan, bindSpanToRight)
	c.SetUPG(false) // Because inserted column may have intersection with spanned cell.
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

//replacer:replace
//replacer:old Ver	[iRow+1][iColumn]	[ciRow][iColumn]	[iRow	[iColumn	Row		Column
//replacer:new Hor	[iRow][iColumn+1]	[iRow][ciColumn]	[iRow	[iColumn	Column	Row

func (c *Table) fixRemoveRowSpan(iRow int, keepSpanMaster bool) {
	lColumn := c.ColumnsCount()

	for iColumn := 0; iColumn < lColumn; iColumn++ {
		span := c.span[iRow][iColumn].Ver
		if span.IsSingle() {
			continue
		}

		for ciRow := iRow - span.BeforeInt(); ciRow < iRow; ciRow++ {
			c.span[ciRow][iColumn].Ver.After--
		}
		for ciRow := iRow + 1; ciRow <= iRow+span.AfterInt(); ciRow++ {
			c.span[ciRow][iColumn].Ver.Before--
		}

		if keepSpanMaster && span.IsMaster() {
			c.children[iRow+1][iColumn] = c.children[iRow][iColumn]
		}
	}
}

//replacer:ignore

func (c *Table) RemoveRowExtended(iRow int, keepSpanMaster bool) []Control {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if iRow < 0 || iRow >= lRow {
		return nil
	}

	c.GeometryHypervisorPause()
	defer c.GeometryHypervisorResume()

	c.fixRemoveRowSpan(iRow, keepSpanMaster)
	r := c.children[iRow]
	for iColumn := 0; iColumn < lColumn; iColumn++ {
		c.detach(iRow, iColumn, false)
	}
	c.children = append(c.children[:iRow], c.children[iRow+1:]...)
	c.span = append(c.span[:iRow], c.span[iRow+1:]...)
	c.rowsGeometry = append(c.rowsGeometry[:iRow], c.rowsGeometry[iRow+1:]...)
	return r
}

func (c *Table) RemoveColumnExtended(iColumn int, keepSpanMaster bool) []Control {
	lRow, lColumn := c.RowsCount(), c.ColumnsCount()
	if iColumn < 0 || iColumn >= lColumn {
		return nil
	}

	c.GeometryHypervisorPause()
	defer c.GeometryHypervisorResume()

	c.fixRemoveColumnSpan(iColumn, keepSpanMaster)
	r := make([]Control, lRow)
	for iRow := 0; iRow < lRow; iRow++ {
		r[iRow] = c.children[iRow][iColumn]
		c.detach(iRow, iColumn, false)
		c.children[iRow] = append(c.children[iRow][:iColumn], c.children[iRow][iColumn+1:]...)
		c.span[iRow] = append(c.span[iRow][:iColumn], c.span[iRow][iColumn+1:]...)
	}
	c.columnsGeometry = append(c.columnsGeometry[:iColumn], c.columnsGeometry[iColumn+1:]...)
	return r
}

//replacer:replace
//replacer:old Row
//replacer:new Column

func (c *Table) RemoveRow(iRow int) []Control {
	return c.RemoveRowExtended(iRow, true)
}

func (c *Table) RemoveFirstRow() []Control {
	return c.RemoveRow(0)
}

func (c *Table) RemoveLastRow() []Control {
	return c.RemoveRow(c.RowsCount() - 1)
}

//replacer:ignore

func (c *Table) allSlavesAreSingleNil(iRow, iColumn, sizeRow, sizeColumn int) bool {
	// First row
	for ciColumn := iColumn + 1; ciColumn <= iColumn+sizeColumn; ciColumn++ {
		if !c.span[iRow][ciColumn].IsSingle() || c.children[iRow][ciColumn] != nil {
			return false
		}
	}
	// Other rows
	for ciRow := iRow + 1; ciRow < iRow+sizeRow; ciRow++ {
		for ciColumn := iColumn; ciColumn < iColumn+sizeColumn; ciColumn++ {
			if !c.span[ciRow][ciColumn].IsSingle() || c.children[ciRow][ciColumn] != nil {
				return false
			}
		}
	}
	return true
}

// iRow & iColumn are indexes of new span master.
// sizeRow & sizeColumn are size of new span (including master, so both must be >0).
func (c *Table) addNewSpan(iRow, iColumn, sizeRow, sizeColumn int) {
	top := uint8(0)
	bottom := uint8(sizeRow - 1)
	for ciRow := iRow; ciRow < iRow+sizeRow; ciRow++ {
		left := uint8(0)
		right := uint8(sizeColumn - 1)
		for ciColumn := iColumn; ciColumn < iColumn+sizeColumn; ciColumn++ {
			c.span[ciRow][ciColumn].Hor.Before = left
			c.span[ciRow][ciColumn].Hor.After = right
			c.span[ciRow][ciColumn].Ver.Before = top
			c.span[ciRow][ciColumn].Ver.After = bottom
			left++
			right--
		}
		top++
		bottom--
	}
}

// AddSpan create new span in table.
// iRow & iColumn are indexes of new span master (left-most top-most element).
// sizeRow & sizeColumn are size of new span in cells (including master).
// Requirements:
//  size* must be from [1;TableSpanMaxSize] (both sizes == 1 is valid and function performs nothing),
//  created span must fit in the table,
//  created span must not overflow existing one,
//  only child at master place may be non-nil (this child will be associated with span), all other children covered by created span must be nil.
// If not all requirements are met then function returns false.
func (c *Table) AddSpan(iRow, iColumn, sizeRow, sizeColumn int) (ok bool) {
	ok = sizeColumn > 0 && sizeColumn <= TableSpanMaxSize && sizeRow > 0 && sizeRow <= TableSpanMaxSize
	ok = ok && iColumn >= 0 && iColumn+sizeColumn <= c.ColumnsCount() && iRow >= 0 && iRow+sizeRow <= c.RowsCount()
	ok = ok && c.span[iRow][iColumn].IsSingle() && c.allSlavesAreSingleNil(iRow, iColumn, sizeRow, sizeColumn)
	if !ok {
		return
	}
	c.addNewSpan(iRow, iColumn, sizeRow, sizeColumn)
	return
}

// iRow & iColumn are indexes of span master.
func (c *Table) removeSpan(iRow, iColumn int) {
	lRow, lColumn := c.span[iRow][iColumn].Size()
	for ciRow := iRow; ciRow < iRow+lRow; ciRow++ {
		for ciColumn := iColumn; ciColumn < iColumn+lColumn; ciColumn++ {
			c.span[ciRow][ciColumn].Hor.Before = 0
			c.span[ciRow][ciColumn].Hor.After = 0
			c.span[ciRow][ciColumn].Ver.Before = 0
			c.span[ciRow][ciColumn].Ver.After = 0
		}
	}
}

// RemoveSpanExtended removes span from table.
// If indirect is false then iRow & iColumn must point to span's master, otherwise they may point to any place of span.
// iRow & iColumn must fit in table.
// It is valid to point to cells which is not under span (single cell), in this case function does nothing and returns true.
// If some requirements aren't met then function returns false.
// After span removing associated control (if any) will be at span's master cell.
func (c *Table) RemoveSpanExtended(iRow, iColumn int, indirect bool) (ok bool) {
	ok = iRow >= 0 && iRow < c.RowsCount() && iColumn >= 0 && iColumn < c.ColumnsCount()
	ok = ok && (indirect || c.span[iRow][iColumn].IsMaster())
	if !ok {
		return
	}
	if indirect {
		iRow, iColumn = c.span[iRow][iColumn].Master(iRow, iColumn)
	}
	c.removeSpan(iRow, iColumn)
	return
}

func (c *Table) RemoveSpan(iRow, iColumn int) (ok bool) {
	return c.RemoveSpanExtended(iRow, iColumn, true)
}

// TODO get span
// TODO get all spans

func NewTable(rowsCount, columnsCount int) *Table {
	r := &Table{
		children:        make([][]Control, rowsCount),
		span:            make([][]TableSpanState, rowsCount),
		rowsGeometry:    make([]tableBandGeometry, rowsCount),
		columnsGeometry: make([]tableBandGeometry, columnsCount),
	}
	for iRow := 0; iRow < rowsCount; iRow++ {
		r.children[iRow] = make([]Control, columnsCount)
		r.span[iRow] = make([]TableSpanState, columnsCount)
	}
	return r
}
