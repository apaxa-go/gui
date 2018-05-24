package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import "font.h"
*/
import "C"
import (
	"errors"
	"github.com/apaxa-go/gui/drivers"
	"unsafe"
)

type Font uintptr // CTFontRef

// Implements drivers.Font interface.
func (f Font) IAmFont() {}

// Handler
func (f Font) H() unsafe.Pointer { return unsafe.Pointer(f) }

func (f Font) Release() {
	C.releaseFont(C.CTFontRef(f.H()))
	// *f = 0 // TODO remove this or use *Font as receiver for all methods
}

func (f Font) Size() float64 {
	return float64(C.CTFontGetSize(C.CTFontRef(f.H())))
}

func NewFont(spec FontSpec) (f Font, err error) {
	spec = spec.Normalize()
	switch spec.Index {
	case drivers.FontIndexDefaultDriverFont:
		return newFontDefault(spec), nil
	case drivers.FontIndexFontFamily, drivers.FontIndexFontName:
		return newFont(spec)
	default:
		return newFontFromFile(spec)
	}
}

func newFontDefault(spec FontSpec) Font {
	// TODO current implementation totally ignore all requirements
	return Font(
		unsafe.Pointer(
			C.makeDefaultFont(
				C.CGFloat(spec.Size),

				C.bool(spec.Requirements.Monospace()),
				C.bool(spec.Monospace),

				C.bool(spec.Requirements.Italic()),
				C.CGFloat(spec.Italic),

				C.bool(spec.Requirements.Slant()),
				C.CGFloat(spec.Slant),

				C.bool(spec.Requirements.Width()),
				C.CGFloat(spec.Width),

				C.bool(spec.Requirements.Weight()),
				C.CGFloat(spec.Weight),
			),
		),
	)
}

func newFont(spec FontSpec) (f Font, err error) {
	nameLen := C.CFIndex(len(spec.Name))
	reqName := spec.Index.Name() && nameLen > 0
	reqFamily := spec.Index.Family() && nameLen > 0
	var name *C.UInt8
	if reqName || reqFamily {
		tmp := []byte(spec.Name)
		name = (*C.UInt8)(&tmp[0])
	}

	f = Font(unsafe.Pointer(C.makeFont(
		C.bool(reqName),
		name,
		nameLen,

		C.bool(reqFamily),
		name,
		nameLen,

		C.CGFloat(spec.Size),

		C.bool(spec.Requirements.Monospace()),
		C.bool(spec.Monospace),

		C.bool(spec.Requirements.Italic()),
		C.CGFloat(spec.Italic),

		C.bool(spec.Requirements.Slant()),
		C.CGFloat(spec.Slant),

		C.bool(spec.Requirements.Width()),
		C.CGFloat(spec.Width),

		C.bool(spec.Requirements.Weight()),
		C.CGFloat(spec.Weight),
	)))

	if f == 0 { // TODO is it possible
		err = errors.New("unable to create font")
	}

	return
}

func newFontFromFile(spec FontSpec) (f Font, err error) {
	nameLen := C.CFIndex(len(spec.Name))
	var name *C.UInt8
	if nameLen > 0 {
		tmp := []byte(spec.Name)
		name = (*C.UInt8)(&tmp[0])
	}

	f = Font(unsafe.Pointer(C.makeFontFromFile(
		name,
		nameLen,

		C.CGFloat(spec.Size),

		C.bool(spec.Requirements.Monospace()),
		C.bool(spec.Monospace),

		C.bool(spec.Requirements.Italic()),
		C.CGFloat(spec.Italic),

		C.bool(spec.Requirements.Slant()),
		C.CGFloat(spec.Slant),

		C.bool(spec.Requirements.Width()),
		C.CGFloat(spec.Width),

		C.bool(spec.Requirements.Weight()),
		C.CGFloat(spec.Weight),
	)))

	if f == 0 {
		err = errors.New("unable to create font from file\"" + spec.Name + "\"")
	}

	return
}
