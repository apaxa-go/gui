// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <CoreFoundation/CoreFoundation.h>

CFStringRef CFStringCreateFromGoString(_GoString_ str) {
	return CFStringCreateWithBytes(NULL, (UInt8*)_GoStringPtr(str), (CFIndex)_GoStringLen(str), kCFStringEncodingUTF8, false);
}
*/
import "C"
