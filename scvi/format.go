// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package scvi

type SCVI struct {
	Size       PointF64
	KeepAspect bool
	Elements   []Primitive
}

type Primitive interface {
	Draw(Canvas, ColorF64)
	Equal(Primitive) bool
}

func (image SCVI) Equal(image2 SCVI) bool {
	r := image.Size == image2.Size && image.KeepAspect == image2.KeepAspect && len(image.Elements) == len(image2.Elements)
	if !r {
		return false
	}
	for i := range image.Elements {
		if !image.Elements[i].Equal(image2.Elements[i]) {
			return false
		}
	}
	return true
}

func (image SCVI) Draw(canvas Canvas, rect RectangleF64, color ColorF64) {
	canvas.SaveState()
	defer canvas.RestoreState()

	if !image.KeepAspect {
		canvas.Translate(rect.LT())
		canvas.ScaleXY(rect.Width()/image.Size.X, rect.Height()/image.Size.Y)
	} else {
		scale := rect.Width() / image.Size.X
		scale0 := rect.Height() / image.Size.Y

		var translate PointF64
		if scale <= scale0 {
			translate.X = rect.Left
			translate.Y = rect.Top + (rect.Height()-image.Size.Y*scale)/2
		} else {
			scale = scale0
			translate.X = rect.Left + (rect.Width()-image.Size.X*scale)/2
			translate.Y = rect.Top
		}
		canvas.Translate(translate)
		canvas.ScaleXY(scale, scale)
	}

	for _, e := range image.Elements {
		e.Draw(canvas, color)
	}
}

type Line struct {
	Point0 PointF64
	Point1 PointF64
	Width  float64
	Alpha  float64
}

func (p Line) Equal(primitive Primitive) bool { return p == primitive }

func (p Line) Draw(canvas Canvas, color ColorF64) {
	color.A *= p.Alpha
	canvas.DrawLine(p.Point0, p.Point1, color, p.Width)
}

func MakeLine(point0, point1 PointF64, width, alpha float64) Line {
	return Line{Point0: point0, Point1: point1, Width: width, Alpha: alpha}
}

type Lines struct {
	Points []PointF64
	Width  float64
	Alpha  float64
}

func (p Lines) Equal(primitive Primitive) bool { // nolint: dupl
	p2, ok := primitive.(Lines)
	ok = ok && len(p.Points) == len(p2.Points) && p.Width == p2.Width && p.Alpha == p2.Alpha
	if !ok {
		return false
	}
	for i := range p.Points {
		if p.Points[i] != p2.Points[i] {
			return false
		}
	}
	return true
}

func (p Lines) Draw(canvas Canvas, color ColorF64) {
	color.A *= p.Alpha
	canvas.DrawConnectedLines(p.Points, color, p.Width)
}

func MakeLines(points []PointF64, width, alpha float64) Lines {
	return Lines{Points: points, Width: width, Alpha: alpha}
}

type Rectangle struct {
	Rect  RectangleF64
	Width float64
	Alpha float64
}

func (p Rectangle) Equal(primitive Primitive) bool { return p == primitive }

func (p Rectangle) Draw(canvas Canvas, color ColorF64) {
	color.A *= p.Alpha
	canvas.DrawRectangle(p.Rect, color, p.Width)
}

type FilledContour struct {
	Points []PointF64
	Width  float64
	Alpha  float64
}

func (p FilledContour) Equal(primitive Primitive) bool { // nolint: dupl
	p2, ok := primitive.(FilledContour)
	ok = ok && len(p.Points) == len(p2.Points) && p.Width == p2.Width && p.Alpha == p2.Alpha
	if !ok {
		return false
	}
	for i := range p.Points {
		if p.Points[i] != p2.Points[i] {
			return false
		}
	}
	return true
}

func (p FilledContour) Draw(canvas Canvas, color ColorF64) {
	color.A *= p.Alpha
	canvas.FillStraightContour(p.Points, color)
}

func MakeFilledContour(points []PointF64, width, alpha float64) FilledContour {
	return FilledContour{Points: points, Width: width, Alpha: alpha}
}
