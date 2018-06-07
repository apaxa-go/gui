// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import "font.h"

CFStringRef CFStringCreateFromGoString(_GoString_ str);

CTFontRef _CreateFont(_GoString_ name, _GoString_ family, struct FontSpec spec) {
	CFStringRef _name   = CFStringCreateFromGoString(name);
	CFStringRef _family = CFStringCreateFromGoString(family);
	CTFontRef r = CreateFont(_name, _family, spec);
	CFRelease(_family);
	CFRelease(_name);
	return r;
}

CTFontRef _CreateFontFromFile(_GoString_ path, struct FontSpec spec) {
	CFStringRef _path = CFStringCreateFromGoString(path);
	CTFontRef r = CreateFontFromFile(_path, spec);
	CFRelease(_path);
	return r;
}
*/
import "C"
import (
	"errors"
	"github.com/apaxa-go/gui/drivers"
)

func translateFontSpec(s FontSpec) (r C.struct_FontSpec) {
	r.size = C.CGFloat(s.Size)
	r.reqMonospace = C.bool(s.Requirements.Monospace())
	r.monospace = C.bool(s.Monospace)
	r.reqItalic = C.bool(s.Requirements.Italic())
	r.italic = C.CGFloat(s.Italic)
	r.reqSlant = C.bool(s.Requirements.Slant())
	r.slant = C.CGFloat(s.Slant)
	r.reqWidth = C.bool(s.Requirements.Width())
	r.width = C.CGFloat(s.Width)
	r.reqWeight = C.bool(s.Requirements.Weight())
	r.weight = C.CGFloat(s.Weight)
	return
}

type Font struct {
	pointer uintptr // CTFontRef
}

// Implements drivers.Font interface.
func (f Font) IAmFont() {}

func (f Font) Release() {
	C.CFRelease(C.CFTypeRef(f.pointer))
}

func (f Font) Size() float64 {
	return float64(C.CTFontGetSize(C.CTFontRef(f.pointer)))
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
	return Font{
		uintptr(
			C.CreateDefaultFont(translateFontSpec(spec)),
		),
	}
}

func newFont(spec FontSpec) (f Font, err error) {
	// TODO remove error from result if other drivers can return font in any cases (MacOS cocoa can).
	var name, family string
	if spec.Index.Name() && len(spec.Name) > 0 {
		name = spec.Name
	} else if spec.Index.Family() && len(spec.Name) > 0 {
		family = spec.Name
	}
	f = Font{
		uintptr(
			C._CreateFont(name, family, translateFontSpec(spec)),
		),
	}
	return
}

func newFontFromFile(spec FontSpec) (f Font, err error) {
	f = Font{
		uintptr(
			C._CreateFontFromFile(spec.Name, translateFontSpec(spec)),
		),
	}
	if f.pointer == 0 {
		err = errors.New("unable to create font from file\"" + spec.Name + "\"")
	}
	return
}
