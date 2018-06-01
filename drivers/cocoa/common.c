// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include "common.h"

CFStringRef CreateStringRef(const UInt8* bytes, CFIndex numBytes) {
	// TODO check functions results for nil???
	return CFStringCreateWithBytes(NULL, bytes, numBytes, kCFStringEncodingUTF8, false);
}
