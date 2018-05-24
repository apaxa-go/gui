package gui

func (w *Window) geometryHypervisorPause() { w.geometryHypervisorState++ }
func (w *Window) geometryHypervisorResume() {
	if w.geometryHypervisorState == 0 {
		return
	}
	w.geometryHypervisorState--
	w.geometryHypervisorRunIfActive()
}
func (w *Window) geometryHypervisorIsActive() bool { return w.geometryHypervisorState == 0 }

func (w *Window) geometryHypervisorRunIfActive() {
	if w.geometryHypervisorIsActive() {
		w.geometryHypervisorDo()
	}
}

func (w *Window) geometryHypervisorDo() {
	w.geometryHypervisorDoUPHG(w)
	w.geometryHypervisorDoUCHG(w)
	w.geometryHypervisorDoUPVG(w)
	w.geometryHypervisorDoUCVG(w)
	w.adjustSize()
	w.geometryHypervisorDoIR(w)
	w.invalidate()
}

//
// Update Possible Horizontal Geometry (from down to up).
//

func (w *Window) geometryHypervisorDoUPHG(control Control) {
	if control.getCUPHG() && (!control.getUPHG() || !control.getUPHGRecursive()) {
		for _, child := range control.Children() {
			w.geometryHypervisorDoUPHG(child)
		}
	}
	if control.getUPHG() {
		var changed bool
		if control.getUPHGRecursive() {
			changed = w.geometryHypervisorDoUPHGRecursive(control)
		} else {
			changed = control.setPossibleHorGeometry(control.ComputePossibleHorGeometry())
		}
		if changed && control.Parent() != nil {
			control.Parent().setUPHG(false)
			control.Parent().setUCHG()
		}
	}
	control.unsetCUPHG()
	control.unsetUPHG()
}

func (w *Window) geometryHypervisorDoUPHGRecursive(control Control) (changed bool) {
	for _, child := range control.Children() {
		w.geometryHypervisorDoUPHGRecursive(child)
	}
	changed = control.setPossibleHorGeometry(control.ComputePossibleHorGeometry())
	if changed && control.Parent() != nil {
		control.Parent().setUCHG()
	}
	control.unsetCUPHG()
	control.unsetUPHG()
	return
}

//
// Update Child Horizontal Geometry (from up to down).
//

func (w *Window) geometryHypervisorDoUCHG(control Control) {
	if control.getUCHG() {
		children := control.Children()
		lefts, rights := control.ComputeChildHorGeometry()
		for i, child := range children {
			changed := child.setHorGeometry(lefts[i], rights[i])
			if changed {
				child.setUCHG()
				child.setUPVG(false)
				child.setIR()
			}
		}
		control.unsetUCHG()
	}
	if control.getCUCHG() {
		children := control.Children()
		for _, child := range children {
			w.geometryHypervisorDoUCHG(child)
		}
		control.unsetCUCHG()
	}
}

//
// Update Possible Vertical Geometry (from down to up).
//

func (w *Window) geometryHypervisorDoUPVG(control Control) {
	if control.getCUPVG() && (!control.getUPVG() || !control.getUPVGRecursive()) {
		for _, child := range control.Children() {
			w.geometryHypervisorDoUPVG(child)
		}
	}
	if control.getUPVG() {
		var changed bool
		if control.getUPVGRecursive() {
			changed = w.geometryHypervisorDoUPVGRecursive(control)
		} else {
			changed = control.setPossibleVerGeometry(control.ComputePossibleVerGeometry())
		}
		if changed && control.Parent() != nil {
			control.Parent().setUPVG(false)
			control.Parent().setUCVG()
		}
	}
	control.unsetCUPVG()
	control.unsetUPVG()
}

func (w *Window) geometryHypervisorDoUPVGRecursive(control Control) (changed bool) {
	for _, child := range control.Children() {
		w.geometryHypervisorDoUPVGRecursive(child)
	}
	changed = control.setPossibleVerGeometry(control.ComputePossibleVerGeometry())
	if changed && control.Parent() != nil {
		control.Parent().setUCVG()
	}
	control.unsetCUPVG()
	control.unsetUPVG()
	return
}

//
// Update Child Vertical Geometry (from up to down).
//

func (w *Window) geometryHypervisorDoUCVG(control Control) {
	if control.getUCVG() {
		children := control.Children()
		tops, bottoms := control.ComputeChildVerGeometry()
		for i, child := range children {
			changed := child.setVerGeometry(tops[i], bottoms[i])
			if changed {
				child.setUCVG()
				child.setIR()
			}
		}
		control.unsetUCVG()
	}
	if control.getCUCVG() {
		children := control.Children()
		for _, child := range children {
			w.geometryHypervisorDoUCVG(child)
		}
		control.unsetCUCVG()
	}
}

//
// Invalidate Region (from up to down).
//

func (w *Window) geometryHypervisorDoIR(control Control) {
	if control.getIR() {
		w.invalidateRegion(control.Geometry())
		control.unsetIR()
		if control.getCIR() {
			w.geometryHypervisorDoIRUnsetRecursive(control)
		}
	} else if control.getCIR() {
		children := control.Children()
		for _, child := range children {
			w.geometryHypervisorDoIR(child)
		}
		control.unsetCIR()
	}
}

func (w *Window) geometryHypervisorDoIRUnsetRecursive(control Control) {
	if control.getCIR() {
		children := control.Children()
		for _, child := range children {
			w.geometryHypervisorDoIRUnsetRecursive(child)
		}
		control.unsetCIR()
	}
	control.unsetIR()
}
