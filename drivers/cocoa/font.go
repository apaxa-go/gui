package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import "font.h"
*/
import "C"
import (
	"unsafe"
)

type Font uintptr // CTFontRef

// Implements drivers.Font interface.
func (f Font) IAmFont() {}

// Handler
func (f Font) H() unsafe.Pointer { return unsafe.Pointer(f) }

func NewFont(name string, size float64) (f Font, ok bool) {
	f = Font(
		unsafe.Pointer(
			C.makeFont(
				(*C.UInt8)(unsafe.Pointer(&[]byte(name)[0])),
				C.CFIndex(len(name)),
				C.CGFloat(size),
			),
		),
	)
	ok = f != 0
	return
}

func (f *Font) Release() {
	C.releaseFont(C.CTFontRef(f.H()))
	*f = 0
}

func (f Font) Size() float64 {
	return float64(C.CTFontGetSize(C.CTFontRef(f.H())))
}
