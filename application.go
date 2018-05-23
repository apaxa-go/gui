package gui

func InitApplication() (Application, error) {
	return driverApplicationConstructor()
}
