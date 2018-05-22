package gui

type Window struct {
	driverWindow DriverWindow
	BaseControl
	child                   Control
	geometryHypervisorState uint // 0 means hypervisor is online (performs request immediately), otherwise it is paused geometryHypervisorState times.
}

//
// Unique methods
//

/*func (w *Window) Run() {
	w.driverWindow.Run()
}*/

func (w *Window) Title() string         { return w.driverWindow.Title() }
func (w *Window) SetTitle(title string) { w.driverWindow.SetTitle(title) }
func (w *Window) Child() Control        { return w.child }
func (w *Window) SetChild(child Control) {
	if w.child != nil {
		w.child.SetParent(nil)
	}
	w.child = child
	w.child.SetParent(w)
	w.SetUPG(true)
}

func (w *Window) adjustSize() {
	reqSize := w.Geometry().GetSize()
	if w.driverWindow.Size() != reqSize {
		w.driverWindow.SetSize(reqSize)
	}
}

func (w *Window) invalidateRegion(region RectangleF64) {
	w.driverWindow.InvalidateRegion(region)
}

func (w *Window) invalidate() {
	w.driverWindow.Invalidate()
}

func (w *Window) onExternalResize() {
	w.SetUCGIR() // TODO we need some method to avoid invalid (according to Min/MaxSize) external resize.
}

func (w *Window) onOfflineCanvasChanged() {
	w.SetUPG(true)
}

func (w *Window) OfflineCanvas() OfflineCanvas { return w.driverWindow.OfflineCanvas() }

//
// BaseControlI overrides
//

func (w *Window) setPossibleHorGeometry(minWidth, maxWidth float64) (changed bool) {
	changed = w.BaseControl.setPossibleHorGeometry(minWidth, maxWidth)
	if !changed {
		return
	}
	if w.Geometry().Width() < w.MinWidth() {
		w.setHorGeometry(0, w.MinWidth())
		w.setUCHG()
	} else if w.Geometry().Width() > w.MaxWidth() {
		w.setHorGeometry(0, w.MaxWidth())
		w.setUCHG()
	}
	return
}

func (w *Window) setPossibleVerGeometry(minHeight, maxHeight float64) (changed bool) {
	changed = w.BaseControl.setPossibleVerGeometry(minHeight, maxHeight)
	if !changed {
		return
	}
	if w.Geometry().Height() < w.MinHeight() {
		w.setVerGeometry(0, w.MinHeight())
		w.setUCVG()
	} else if w.Geometry().Width() > w.MaxHeight() {
		w.setVerGeometry(0, w.MaxHeight())
		w.setUCVG()
	}
	return
}

func (w *Window) SetParent(parent Control) {
	panic("set parent for Window") // TODO may be less panic?
}

//
// Control interface implementation
//

func (w *Window) Children() []Control {
	if w.child == nil {
		return nil
	}
	return []Control{w.child}
}

func (w *Window) ComputePossibleHorGeometry() (minWidth, maxWidth float64) {
	if w.child == nil {
		return 100, 100
	}
	return w.child.MinWidth(), w.child.MaxWidth()
}

func (w *Window) ComputePossibleVerGeometry() (minHeight, maxHeight float64) {
	if w.child == nil {
		return 100, 100
	}
	return w.child.MinHeight(), w.child.MaxHeight()
}

func (w *Window) ComputeChildHorGeometry() (lefts, rights []float64) {
	if w.child == nil {
		return nil, nil
	}
	return []float64{0}, []float64{w.Geometry().Width()}
}
func (w *Window) ComputeChildVerGeometry() (tops, bottoms []float64) {
	if w.child == nil {
		return nil, nil
	}
	return []float64{0}, []float64{w.Geometry().Height()}
}

// TODO may remove this method?
func (w *Window) Draw(canvas Canvas, region RectangleF64) {
	if w.child != nil {
		w.child.Draw(canvas, region)
	}
}

// TODO may remove this method?
func (w *Window) ProcessEvent(event Event) bool {
	if w.child != nil {
		return w.child.ProcessEvent(event)
	}
	return false
}

//
// Constructors & destructor
//

func (w *Window) baseInit() {
	w.driverWindow.RegisterDrawCallback(w.Draw)
	w.driverWindow.RegisterEventCallback(w.ProcessEvent)
	w.driverWindow.RegisterResizeCallback(w.onExternalResize)
	w.driverWindow.RegisterOfflineCanvasCallback(w.onOfflineCanvasChanged)
	w.BaseControl.window = w
	w.SetUPGIR(false)
}

func NewEmptyWindow(dw DriverWindow) *Window {
	var w Window
	w.driverWindow = dw
	w.baseInit()
	return &w
}

func NewWindow(dw DriverWindow, title string, child Control) *Window {
	var w Window
	w.driverWindow = dw
	w.GeometryHypervisorPause()
	w.baseInit()
	w.SetTitle(title)
	w.SetChild(child)
	w.GeometryHypervisorResume()
	return &w
}

func (w *Window) Destroy() {
	w.GeometryHypervisorPause()
	w.SetChild(nil)
	w.driverWindow.Destroy()
}
