// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

import (
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/drivers"
)

type ColorF64 = basetypes.ColorF64

type PointF64 = basetypes.PointF64
type EllipseF64 = basetypes.EllipseF64
type CircleF64 = basetypes.CircleF64
type RectangleF64 = basetypes.RectangleF64
type RectangleF64S = basetypes.RectangleF64S
type RoundedRectangleF64 = basetypes.RoundedRectangleF64

type TransformF64 = basetypes.TransformF64

type WindowI = drivers.Window
type WindowDisplayState = drivers.WindowDisplayState
type CanvasI = drivers.Canvas
type FontI = drivers.Font
type OfflineCanvasI = drivers.OfflineCanvas
type FontSpec = drivers.FontSpec
type Cursor = drivers.Cursor

//
// Events
//

type KeyboardEvent = drivers.KeyboardEvent
type KeyEvent = drivers.KeyEvent
type Key = drivers.Key
type KeyModifiers = drivers.KeyModifiers

type PointerButtonEvent = drivers.PointerButtonEvent
type PointerButtonEventKind = drivers.PointerButtonEventKind
type PointerButton = drivers.PointerButton
type PointerDragEvent = drivers.PointerDragEvent
type PointerMoveEvent = drivers.PointerMoveEvent
type PointerEnterLeaveEvent = drivers.PointerEnterLeaveEvent

type ScrollEvent = drivers.ScrollEvent
type ModifiersEvent = drivers.ModifiersEvent
type WindowMainStateEvent = drivers.WindowMainStateEvent
type WindowDisplayStateEvent = drivers.WindowDisplayStateEvent

type EnterLeaveAreaID = drivers.EnterLeaveAreaID
type MoveAreaID = drivers.MoveAreaID
