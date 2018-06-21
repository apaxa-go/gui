// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

//
// Enter leave
//

type enterLeaveArea struct {
	overlapping bool
	receiver    Control
}

type overlappedEnterLeaveArea struct {
	id       EnterLeaveAreaID
	receiver Control // This field helps to avoid multiple search in enterLeaveArea map.
}

func (w *Window) generateEnterLeaveAreaID() EnterLeaveAreaID {
	id := w.nextEnterLeaveAreaID
	for _, exists := w.enterLeaveAreas[id]; exists; _, exists = w.enterLeaveAreas[id] {
		id++
	}
	w.nextEnterLeaveAreaID = id + 1
	return id
}

func (w *Window) addEnterLeaveArea(receiver Control, area RectangleF64, overlapping bool) EnterLeaveAreaID {
	id := w.generateEnterLeaveAreaID()
	w.enterLeaveAreas[id] = enterLeaveArea{overlapping, receiver}
	w.driverWindow.AddEnterLeaveArea(id, area)
	return id
}

func (w *Window) AddEnterLeaveArea(receiver Control, area RectangleF64) EnterLeaveAreaID {
	return w.addEnterLeaveArea(receiver, area, false)
}

func (w *Window) AddEnterLeaveOverlappingArea(receiver Control, area RectangleF64) EnterLeaveAreaID {
	return w.addEnterLeaveArea(receiver, area, true)
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

// onPointerEnterLeave method processes PointerEnterLeaveEvent.
// If event's area marked as non-overlapping then receiver unconditionally receive event.
// Otherwise we consult overlappedEnterLeaveAreas.
func (w *Window) onPointerEnterLeave(e PointerEnterLeaveEvent) {
	area, ok := w.enterLeaveAreas[e.ID]
	if !ok {
		return
	}
	if !area.overlapping {
		area.receiver.OnPointerEnterLeaveEvent(e)
		return
	}
	if e.Enter {
		w.onPointerEnterOverlapped(e.ID, area.receiver)
	} else {
		w.onPointerLeaveOverlapped(e.ID, area.receiver)
	}
}

// If there is no yet overlapping element then send message and push self to overlappedEnterLeaveAreas.
// If new receiver has equal or higher zIndex than active entered area (the last one in the overlappedEnterLeaveAreas) then new Enter event overlaps active.
// Otherwise we add new event into overlappedEnterLeaveAreas (preserving zIndex ordering) as overlapped (no events will be sent to Controls).
func (w *Window) onPointerEnterOverlapped(id EnterLeaveAreaID, receiver Control) {
	l := len(w.overlappedEnterLeaveAreas)
	if l == 0 {
		receiver.OnPointerEnterLeaveEvent(PointerEnterLeaveEvent{id, true})
		w.overlappedEnterLeaveAreas = append(w.overlappedEnterLeaveAreas, overlappedEnterLeaveArea{id, receiver})
		return
	}
	zIndex := receiver.ZIndex()
	lastArea := w.overlappedEnterLeaveAreas[l-1]
	if zIndex >= lastArea.receiver.ZIndex() {
		lastArea.receiver.OnPointerEnterLeaveEvent(PointerEnterLeaveEvent{lastArea.id, false})
		receiver.OnPointerEnterLeaveEvent(PointerEnterLeaveEvent{id, true})
	} else {
		i := l - 1
		for i > 0 && zIndex < w.overlappedEnterLeaveAreas[i-1].receiver.ZIndex() {
			i--
		}
		w.overlappedEnterLeaveAreas = append(append(w.overlappedEnterLeaveAreas[:i], overlappedEnterLeaveArea{id, receiver}), w.overlappedEnterLeaveAreas[i:]...)
	}
}

// If Leave event related to active entered area then Leave current and Enter previous will be sent.
// In any case event will be removed from overlappedEnterLeaveAreas.
func (w *Window) onPointerLeaveOverlapped(id EnterLeaveAreaID, receiver Control) {
	l := len(w.overlappedEnterLeaveAreas)
	if l == 0 {
		return
	}
	if w.overlappedEnterLeaveAreas[l-1].id == id {
		receiver.OnPointerEnterLeaveEvent(PointerEnterLeaveEvent{id, false})
		l--
		w.overlappedEnterLeaveAreas = w.overlappedEnterLeaveAreas[:l]
		if l > 0 {
			lastArea := w.overlappedEnterLeaveAreas[l-1]
			lastArea.receiver.OnPointerEnterLeaveEvent(PointerEnterLeaveEvent{lastArea.id, true})
		}
	} else {
		for i := l - 2; i >= 0; i-- {
			if w.overlappedEnterLeaveAreas[i].id == id {
				w.overlappedEnterLeaveAreas = append(w.overlappedEnterLeaveAreas[:i], w.overlappedEnterLeaveAreas[i+1:]...)
				break
			}
		}
	}
}

// generateMoveAreaID returns next free ID for MoveArea.
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
