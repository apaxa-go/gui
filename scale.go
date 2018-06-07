// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

import (
	"github.com/apaxa-go/helper/mathh"
	"math"
)

var tan05 = math.Tan(0.5)

const (
	DefaultDPI      = 96 // dots per inch.
	DefaultDistance = 20 // inchces (~50cm).
)

var defaultDPR = DPI2DPR(DefaultDPI, DefaultDistance)

// Dots per radian. For approximate calculation may be used formula "DPR ~= DPI * distance * 1.1 ".
// Distance in inches.
func DPI2DPR(dpi, distance float64) float64 {
	return dpi * distance * 2 * tan05
}

// res2DPI computes DPI from resolution and diagonal.
// Diagonal in inches.
func res2DPI(resX, resY, diagonal float64) float64 {
	return math.Sqrt(resX*resX+resY*resY) / diagonal
}

// Res2DPR computes DPR from resolution, diagonal and distance.
// Distance & diagonal in inches.
func Res2DPR(resX, resY, diagonal, distance float64) float64 {
	return DPI2DPR(res2DPI(resX, resY, diagonal), distance)
}

type Scale struct {
	scale         float64
	lineWidth     float64
	borderRadius  float64
	fontHeight    float64
	controlHeight float64
	topPadding    float64
	bottomPadding float64
	horPadding    float64
}

func (s *Scale) update() {
	s.lineWidth = mathh.Max2Float64(1, mathh.RoundFloat64(s.scale))
	s.borderRadius = s.scale * 3
	s.fontHeight = mathh.Max2Float64(6, mathh.RoundFloat64(s.scale*12))
	s.controlHeight = mathh.Max2Float64(s.lineWidth*2+2+s.controlHeight, s.scale*20)
	s.topPadding = mathh.RoundFloat64((s.controlHeight - 2*s.lineWidth - s.fontHeight) / 2)
	s.bottomPadding = s.controlHeight - 2*s.lineWidth - s.fontHeight - s.topPadding
	s.horPadding = mathh.Max2Float64(s.topPadding, s.bottomPadding)
}

func MakeScale(dpi, distance float64) Scale {
	var s Scale
	s.scale = DPI2DPR(dpi, distance) / defaultDPR
	s.update()
	return s
}
