// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package gui

func (w *Window) DisplayState() WindowDisplayState { return w.driver.DisplayState() }

func (w *Window) Minimize()        { w.driver.Minimize() }
func (w *Window) Deminimize()      { w.driver.Deminimize() }
func (w *Window) Maximize()        { w.driver.Maximize() }
func (w *Window) Demaximize()      { w.driver.Demaximize() }
func (w *Window) EnterFullScreen() { w.driver.EnterFullScreen() }
func (w *Window) ExitFullScreen()  { w.driver.ExitFullScreen() }

func (w *Window) IsNormalAllowed() bool     { return true }
func (w *Window) IsMinimizeAllowed() bool   { return w.minimizeAllowed }
func (w *Window) IsMaximizeAllowed() bool   { return w.maximizeAllowed }
func (w *Window) IsFullScreenAllowed() bool { return w.fullScreenAllowed }
func (w *Window) IsCloseAllowed() bool      { return w.closeAllowed }
