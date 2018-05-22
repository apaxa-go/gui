package gui

// Implemented by BaseControl and can't be implemented by user.
// Any Control implementation must include BaseControl.
type BaseControlI interface {
	Window() *Window
	GeometryHypervisorPause()
	GeometryHypervisorResume()

	Parent() Control
	SetParent(Control)

	MinWidth() int
	MinHeight() int
	MaxWidth() int
	MaxHeight() int

	Geometry() RectangleI

	setPossibleHorGeometry(minWidth, maxWidth int) (changed bool)
	setPossibleVerGeometry(minHeight, maxHeight int) (changed bool)

	setHorGeometry(left, right int) (changed bool)
	setVerGeometry(top, bottom int) (changed bool)

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

	ComputePossibleHorGeometry() (minWidth, maxWidth int)
	ComputePossibleVerGeometry() (minHeight, maxHeight int)

	ComputeChildHorGeometry() (lefts, rights []int) // index according Children()
	ComputeChildVerGeometry() (tops, bottoms []int) // index according Children()

	Draw(canvas Canvas, region RectangleI)
	ProcessEvent(Event) bool
}
