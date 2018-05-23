package drivers

import "unsafe"

type Font interface {
	H() unsafe.Pointer
	IAmFont()
	Release()
}
