package gui

import (
	"github.com/apaxa-go/gui/basetypes"
	"github.com/apaxa-go/gui/drivers"
)

//
// Base types
//

type ColorF64 = basetypes.ColorF64

type PointF64 = basetypes.PointF64
type RectangleF64 = basetypes.RectangleF64
type RoundedRectangleF64 = basetypes.RoundedRectangleF64
type EllipseF64 = basetypes.EllipseF64
type CircleF64 = basetypes.CircleF64

//
// Driver's types
//

type Application = drivers.Application
type DriverWindow = drivers.Window
type Canvas = drivers.Canvas
type OfflineCanvas = drivers.OfflineCanvas
type Event = drivers.Event

type Font = drivers.Font
type FontSpec = drivers.FontSpec
type FontIndex = drivers.FontIndex

//
// This is only to allow controls to not import basetypes
//

/*type Align = basetypes.Align
type AlignHor = basetypes.AlignHor
type AlignVer = basetypes.AlignVer*/
