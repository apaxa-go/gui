package gui

// Implemented by BaseControl and can't be implemented by user.
// Any Control implementation must include BaseControl.
type BaseControlI interface {
	Window() *Window
	GeometryHypervisorPause()
	GeometryHypervisorResume()

	Parent() Control
	SetParent(Control)

	MinWidth() float64
	MinHeight() float64
	MaxWidth() float64
	MaxHeight() float64

	Geometry() RectangleF64

	setPossibleHorGeometry(minWidth, maxWidth float64) (changed bool)
	setPossibleVerGeometry(minHeight, maxHeight float64) (changed bool)

	setHorGeometry(left, right float64) (changed bool)
	setVerGeometry(top, bottom float64) (changed bool)

	//
	// GeometryHypervisor data
	//

	getUPHG() bool
	getUPHGRecursive() bool
	setUPHG(recursive bool)
	SetUPHG(recursive bool)
	unsetUPHG()

	getCUPHG() bool
	setCUPHG()
	unsetCUPHG()

	getUCHG() bool
	setUCHG()
	SetUCHG()
	unsetUCHG()

	getCUCHG() bool
	setCUCHG()
	unsetCUCHG()

	getUPVG() bool
	getUPVGRecursive() bool
	setUPVG(recursive bool)
	SetUPVG(recursive bool)
	unsetUPVG()

	getCUPVG() bool
	setCUPVG()
	unsetCUPVG()

	getUCVG() bool
	setUCVG()
	SetUCVG()
	unsetUCVG()

	getCUCVG() bool
	setCUCVG()
	unsetCUCVG()

	getIR() bool
	setIR()
	SetIR()
	unsetIR()

	getCIR() bool
	setCIR()
	unsetCIR()

	SetUPG(recursive bool)
	SetUPGIR(recursive bool)
}

// Must be implemented by user.
type Control interface {
	BaseControlI

	Children() []Control

	ComputePossibleHorGeometry() (minWidth, maxWidth float64)
	ComputePossibleVerGeometry() (minHeight, maxHeight float64)

	ComputeChildHorGeometry() (lefts, rights []float64) // index according Children()
	ComputeChildVerGeometry() (tops, bottoms []float64) // index according Children()

	Draw(canvas Canvas, region RectangleF64)
	ProcessEvent(Event) bool
}
