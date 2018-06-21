// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package drivers

type Cursor uint8

const (
	DefaultCursor                  Cursor = iota // default cursor, usually arrow
	ResizeTopBottomCursor                        // ↕
	ResizeLeftRightCursor                        // ↔
	ResizeLeftTopRightBottomCursor               // ⤡
	ResizeLeftBottomRightTopCursor               // ⤢
)

func (Cursor) MakeDefault() Cursor                  { return DefaultCursor }
func (Cursor) MakeResizeTopBottom() Cursor          { return ResizeTopBottomCursor }
func (Cursor) MakeResizeLeftRight() Cursor          { return ResizeLeftRightCursor }
func (Cursor) MakeResizeLeftTopRightBottom() Cursor { return ResizeLeftTopRightBottomCursor }
func (Cursor) MakeResizeLeftBottomRightTop() Cursor { return ResizeLeftBottomRightTopCursor }
