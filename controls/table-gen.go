// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

//replacer:generated-file

package controls

import "github.com/apaxa-go/helper/mathh"

func (c *Table) ComputePossibleVerGeometry() (minHeight, bestHeight, maxHeight float64) {
	lColumn, lRow := c.ColumnsCount(), c.RowsCount()
	if lColumn == 0 || lRow == 0 {
		return 0, 0, 0
	}

	var prevMin, prevMax float64
	spannedGeometries := make(map[int]buttonSpannedGeometry)
	for iRow := 0; iRow < lRow; iRow++ { // row by row
		// values for current row
		min := float64(0)
		best := float64(0)
		max := mathh.PositiveInfFloat64()

		var effectiveCount float64 // Count of non-nil non-spanned cells in row
		for iColumn := 0; iColumn < lColumn; iColumn++ {
			span := c.span[iRow][iColumn].Ver
			if c.children[iRow][iColumn] == nil || span.IsSlave() {
				continue
			}
			cMinHeight := c.children[iRow][iColumn].MinHeight()
			cBestHeight := c.children[iRow][iColumn].MinHeight()
			cMaxHeight := c.children[iRow][iColumn].MaxHeight()
			if span.IsSingle() { // single row cell
				effectiveCount++
				min = mathh.Max2Float64(min, cMinHeight)
				best += cBestHeight
				max = mathh.Min2Float64(max, cMaxHeight)
			} else { // spanned cell
				im := iRow + span.AfterInt()
				if g, ok := spannedGeometries[im]; ok {
					g.min = mathh.Max2Float64(g.min, minHeight+min)
					g.max = mathh.Min2Float64(g.max, maxHeight+max)
				} else {
					spannedGeometries[im] = buttonSpannedGeometry{minHeight + min, maxHeight + max}
				}
			}
		}

		if spannedGeometry, spannedExists := spannedGeometries[iRow]; spannedExists { // Spanned geometry exists.
			// Convert saved spanned geometry from absolute to row-only values.
			spannedGeometry.min -= prevMin
			spannedGeometry.max -= prevMax

			if effectiveCount == 0 { // Only spanned geometry exists for row.
				min = mathh.Max2Float64(spannedGeometry.min, 0) // We need to be sure here what min is non negative (max will be fixed at normalization).
				max = spannedGeometry.max
			} else { // Both spanned geometry and row's geometry exists.
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

		c.rowsGeometry[iRow].min = min
		c.rowsGeometry[iRow].best = best
		c.rowsGeometry[iRow].max = max

		prevMin = minHeight
		prevMax = maxHeight

		minHeight += min
		bestHeight += best
		maxHeight += max
	}
	return
}

func (c *Table) ComputeChildVerGeometry() (tops, bottoms []float64) {
	lColumn, lRow := c.ColumnsCount(), c.RowsCount()
	if lColumn == 0 || lRow == 0 {
		return nil, nil
	}

	tops = make([]float64, lColumn*lRow)
	bottoms = make([]float64, lColumn*lRow)

	top := c.Geometry().Top
	height := c.Geometry().Height()
	switch {
	case height >= c.MaxHeight(): // scale according to MaxHeight
		scale := c.MaxHeight()
		for iRow, geometry := range c.rowsGeometry {
			scalePart := geometry.max
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			for iColumn := 0; iColumn < lColumn; iColumn++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				tops[i] = top
				bottoms[i] = bottom
			}

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	case height >= c.BestHeight(): // scale according to BestHeight
		scale := c.BestHeight()
		for iRow, geometry := range c.rowsGeometry {
			scalePart := geometry.best
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			for iColumn := 0; iColumn < lColumn; iColumn++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				tops[i] = top
				bottoms[i] = bottom
			}

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	default: // scale according to MinHeight
		scale := c.MinHeight()
		for iRow, geometry := range c.rowsGeometry {
			scalePart := geometry.min
			curHeight := height * scalePart / scale
			bottom := top + curHeight

			for iColumn := 0; iColumn < lColumn; iColumn++ {
				i := c.rowColumnToIndex(iRow, iColumn)
				tops[i] = top
				bottoms[i] = bottom
			}

			top = bottom
			height -= curHeight
			scale -= scalePart
		}
	}
	return
}

func (c *Table) fixNewColumnSpan(iColumn int, stretchSpan, bindSpanToRight bool) {
	lColumn, lRow := c.ColumnsCount(), c.RowsCount()
	if iColumn == 0 || iColumn == lColumn-1 {
		return
	}
	switch {
	case stretchSpan:
		for iRow := 0; iRow < lRow; iRow++ {
			span := c.span[iRow][iColumn+1]
			if span.Hor.IsMaster() {
				continue
			}
			c.span[iRow][iColumn] = span

			for ciColumn := iColumn - span.Hor.BeforeInt(); ciColumn <= iColumn; ciColumn++ {
				c.span[iRow][ciColumn].Hor.After++
			}
			for ciColumn := iColumn + 1; ciColumn <= iColumn+span.Hor.AfterInt()+1; ciColumn++ {
				c.span[iRow][ciColumn].Hor.Before++
			}
		}
	case bindSpanToRight:
		for iRow := 0; iRow < lRow; iRow++ {
			span := c.span[iRow][iColumn+1].Hor
			if span.IsMaster() {
				continue
			}

			for ciColumn := iColumn - span.BeforeInt(); ciColumn < iColumn; ciColumn++ {
				c.span[iRow][ciColumn+1] = c.span[iRow][ciColumn]
			}

			ciColumn := iColumn - span.BeforeInt()
			c.span[iRow][ciColumn] = TableSpanState{}
		}
	default: // bind span to top
		for iRow := 0; iRow < lRow; iRow++ {
			span := c.span[iRow][iColumn+1].Hor
			if span.IsMaster() {
				continue
			}

			for ciColumn := iColumn; ciColumn <= iColumn+span.AfterInt(); ciColumn++ {
				c.span[iRow][ciColumn] = c.span[iRow][ciColumn+1]
			}

			ciColumn := iColumn + span.AfterInt() + 1
			c.span[iRow][ciColumn] = TableSpanState{}
		}
	}
}

func (c *Table) InsertColumn(at int) {
	c.InsertColumnExtended(at, true, false)
}

func (c *Table) PrependColumn() {
	c.InsertColumn(0)
}

func (c *Table) AppendColumn() {
	c.InsertColumn(c.ColumnsCount())
}
