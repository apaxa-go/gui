// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

// generateTrackingAreaID returns next free ID for TrackingArea.
// Window stores candidate for next ID, but it can be already assigned, so we never use it directly.
// This function in loop checks if ID is used starting from nextTrackingAreaID.
// Theoretically this loop may be infinite, but in real world looks like OOM has been happened earlier (MaxInt objects in memory).
// More over in most use case this loop will make only one step.
func (w *Window) generateTrackingAreaID() TrackingAreaID {
	id := w.nextTrackingAreaID
	for _, exists := w.trackingAreas[id]; exists; _, exists = w.trackingAreas[id] {
		id++
	}
	w.nextTrackingAreaID = id + 1
	return id
}

func (w *Window) AddTrackingArea(receiver Control, area TrackingArea) TrackingAreaID {
	id := w.generateTrackingAreaID()
	w.trackingAreas[id] = receiver
	w.driverWindow.AddTrackingArea(id, area)
	return id
}

func (w *Window) ReplaceTrackingArea(id TrackingAreaID, area TrackingArea) (ok bool) {
	_, ok = w.trackingAreas[id]
	if ok {
		w.driverWindow.ReplaceTrackingArea(id, area)
	}
	return
}

func (w *Window) RemoveTrackingArea(id TrackingAreaID) (ok bool) {
	_, ok = w.trackingAreas[id]
	if ok {
		w.driverWindow.RemoveTrackingArea(id)
	}
	return
}

func (w *Window) onPointerEnterLeave(e PointerEnterLeaveEvent) {
	receiver, ok := w.trackingAreas[e.Id]
	if !ok {
		return
	}
	receiver.OnPointerEnterLeaveEvent(e)
}
