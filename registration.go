package gui

var driverWindowConstructor func() DriverWindow
var driverFontConstructor func(FontSpec) (Font, error)
var driverApplicationConstructor func() (Application, error)

func RegisterDriver(
	applicationConstructor func() (Application, error),
	windowConstructor func() DriverWindow,
	fontConstructor func(FontSpec) (Font, error),
) {
	driverApplicationConstructor = applicationConstructor
	driverWindowConstructor = windowConstructor
	driverFontConstructor = fontConstructor
}
