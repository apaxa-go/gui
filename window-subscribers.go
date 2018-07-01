// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

func (w *Window) SubscribeToModifiersEvent(s ModifiersEventSubscriber) {
	for _, subscriber := range w.modifiersEventSubscribers {
		if subscriber == s {
			return
		}
	}
	w.modifiersEventSubscribers = append(w.modifiersEventSubscribers, s)
}

func (w *Window) UnsubscribeFromModifiersEvent(s ModifiersEventSubscriber) {
	for i, subscriber := range w.modifiersEventSubscribers {
		if subscriber == s {
			w.modifiersEventSubscribers = append(w.modifiersEventSubscribers[:i], w.modifiersEventSubscribers[i+1:]...)
			return
		}
	}
}

func (w *Window) SubscribeToWindowDisplayStateEvent(s WindowDisplayStateEventSubscriber) {
	for _, subscriber := range w.windowDisplayStateEventSubscribers {
		if subscriber == s {
			return
		}
	}
	w.windowDisplayStateEventSubscribers = append(w.windowDisplayStateEventSubscribers, s)
}

func (w *Window) UnsubscribeFromWindowDisplayStateEvent(s WindowDisplayStateEventSubscriber) {
	for i, subscriber := range w.windowDisplayStateEventSubscribers {
		if subscriber == s {
			w.windowDisplayStateEventSubscribers = append(w.windowDisplayStateEventSubscribers[:i], w.windowDisplayStateEventSubscribers[i+1:]...)
			return
		}
	}
}

func (w *Window) SubscribeToWindowMainStateEvent(s WindowMainStateEventSubscriber) {
	for _, subscriber := range w.windowMainStateEventSubscribers {
		if subscriber == s {
			return
		}
	}
	w.windowMainStateEventSubscribers = append(w.windowMainStateEventSubscribers, s)
}

func (w *Window) UnsubscribeFromWindowMainStateEvent(s WindowMainStateEventSubscriber) {
	for i, subscriber := range w.windowMainStateEventSubscribers {
		if subscriber == s {
			w.windowMainStateEventSubscribers = append(w.windowMainStateEventSubscribers[:i], w.windowMainStateEventSubscribers[i+1:]...)
			return
		}
	}
}
