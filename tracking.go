// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

//
// Enter leave
//

func (w *Window) generateEnterLeaveAreaID() EnterLeaveAreaID {
	id := w.nextEnterLeaveAreaID
	for _, exists := w.enterLeaveAreas[id]; exists; _, exists = w.enterLeaveAreas[id] {
		id++
	}
	w.nextEnterLeaveAreaID = id + 1
	return id
}

func (w *Window) AddEnterLeaveArea(receiver Control, area RectangleF64) EnterLeaveAreaID {
	id := w.generateEnterLeaveAreaID()
	w.enterLeaveAreas[id] = receiver
	w.driverWindow.AddEnterLeaveArea(id, area)
	return id
}

func (w *Window) ReplaceEnterLeaveArea(id EnterLeaveAreaID, area RectangleF64) (ok bool) { // TODO may be implement replace in gui, not in driver?
	_, ok = w.enterLeaveAreas[id]
	if ok {
		w.driverWindow.ReplaceEnterLeaveArea(id, area)
	}
	return
}

func (w *Window) RemoveEnterLeaveArea(id EnterLeaveAreaID, keepID bool) (ok bool) {
	_, ok = w.enterLeaveAreas[id]
	if ok {
		w.driverWindow.RemoveEnterLeaveArea(id)
		if !keepID {
			delete(w.enterLeaveAreas, id)
		}
	}
	return
}

func (w *Window) onPointerEnterLeave(e PointerEnterLeaveEvent) {
	receiver, ok := w.enterLeaveAreas[e.ID]
	if !ok {
		return
	}
	receiver.OnPointerEnterLeaveEvent(e)
}

// generateMoveAreaID returns next free ID for TrackingArea.
// Window stores candidate for next ID, but it can be already assigned, so we never use it directly.
// This function in loop checks if ID is used starting from nextMoveAreaID.
// Theoretically this loop may be infinite, but in real world looks like OOM has been happened earlier (MaxInt objects in memory).
// More over in most use case this loop will make only one step.
func (w *Window) generateMoveAreaID() MoveAreaID {
	id := w.nextMoveAreaID
	for _, exists := w.moveAreas[id]; exists; _, exists = w.moveAreas[id] {
		id++
	}
	w.nextMoveAreaID = id + 1
	return id
}

func (w *Window) AddMoveArea(receiver Control, area RectangleF64) MoveAreaID {
	id := w.generateMoveAreaID()
	w.moveAreas[id] = receiver
	w.driverWindow.AddMoveArea(id, area)
	return id
}

func (w *Window) ReplaceMoveArea(id MoveAreaID, area RectangleF64) (ok bool) {
	_, ok = w.moveAreas[id]
	w.driverWindow.ReplaceMoveArea(id, area)
	return
}

func (w *Window) RemoveMoveArea(id MoveAreaID, keepID bool) (ok bool) {
	_, ok = w.moveAreas[id]
	if ok {
		w.driverWindow.RemoveMoveArea(id)
		if !keepID {
			delete(w.moveAreas, id)
		}
	}
	return
}

func (w *Window) onPointerMove(e PointerMoveEvent) {
	receiver, ok := w.moveAreas[e.ID]
	if !ok {
		return
	}
	receiver.OnPointerMoveEvent(e)
}
