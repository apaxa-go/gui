// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

// Returns false if "c" does not belong to "w".
func (w *Window) SetFocus(newFocus Control) (ok bool) {
	ok = newFocus.Window() == w
	if ok {
		prevFocus := w.focusedControl
		prevFocus.OnFocus(FocusEvent{false, newFocus})
		w.focusedControl = newFocus
		newFocus.OnFocus(FocusEvent{true, prevFocus})
	}
	return
}

func (w *Window) ShiftFocus(reverse bool) {
	current := w.focusedControl
	where := current
	for {
		candidate := where.FocusCandidate(reverse, current)
		switch candidate {
		case where:
			// Control has decided to become a focus Control.
			w.SetFocus(candidate)
			return
		case nil:
			// There is no candidate in "where" and his children.
			// Go to upper level and look for next or previous candidates.
			current = where
			where = where.Parent()
			if where == nil { // TODO remove this case?
				w.SetFocus(w)
				return
			}
		default:
			// Some child is a candidate.
			// Try to find first or last candidate in it.
			current = nil
			where = candidate
		}
	}
}

func (w *Window) FocusCandidate(_ bool, current Control) Control {
	if current == w {
		return w.child
	}
	return w
}
