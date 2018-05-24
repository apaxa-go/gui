package gui

import "runtime"

func InitApplication() (Application, error) {
	runtime.LockOSThread()
	return driverApplicationConstructor()
}
