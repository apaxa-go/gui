// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

func (w *Window) DisplayState() WindowDisplayState { return w.driverWindow.DisplayState() }

func (w *Window) Minimize()        { w.driverWindow.Minimize() }
func (w *Window) Deminimize()      { w.driverWindow.Deminimize() }
func (w *Window) Maximize()        { w.driverWindow.Maximize() }
func (w *Window) Demaximize()      { w.driverWindow.Demaximize() }
func (w *Window) EnterFullScreen() { w.driverWindow.EnterFullScreen() }
func (w *Window) ExitFullScreen()  { w.driverWindow.ExitFullScreen() }

func (w *Window) IsNormalAllowed() bool     { return true }
func (w *Window) IsMinimizeAllowed() bool   { return w.minimizeAllowed }
func (w *Window) IsMaximizeAllowed() bool   { return w.maximizeAllowed }
func (w *Window) IsFullScreenAllowed() bool { return w.fullScreenAllowed }
func (w *Window) IsCloseAllowed() bool      { return w.closeAllowed }
