// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

func (BaseControl) Children() []Control                                { return nil }
func (BaseControl) SequentialZIndex() bool                             { return false }
func (BaseControl) ComputeChildHorGeometry() (lefts, rights []float64) { return nil, nil }
func (BaseControl) ComputeChildVerGeometry() (tops, bottoms []float64) { return nil, nil }

func (BaseControl) AfterAttachToWindowEvent()                                  {}
func (BaseControl) BeforeDetachFromWindowEvent()                               {}
func (BaseControl) OnGeometryChangeEvent()                                     {}
func (BaseControl) OnKeyboardEvent(_ KeyboardEvent) (done bool)                { return false }
func (BaseControl) OnPointerButtonEvent(_ PointerButtonEvent) (processed bool) { return false }
func (BaseControl) OnPointerDragEvent(_ PointerDragEvent)                      {}
func (BaseControl) OnPointerMoveEvent(_ PointerMoveEvent) (processed bool)     { return false }
func (BaseControl) OnScrollEvent(_ ScrollEvent) (processed bool)               { return false }
func (BaseControl) OnPointerEnterLeaveEvent(_ PointerEnterLeaveEvent)          {}
func (BaseControl) OnFocus(_ FocusEvent)                                       {}
func (BaseControl) OnWindowMainEvent(_ bool)                                   {}

// FocusCandidate is default implementation. It always returns nil - neither Control itself nor his children (if any) accepts focus.
func (BaseControl) FocusCandidate(reverse bool, current Control) Control { return nil }
