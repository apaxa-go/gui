// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import "application.h"

void InitApplication() { [NSApplication sharedApplication]; }

void RunApplication() {
	@autoreleasepool {
		[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
		[NSApp run];
	}
}

void StopApplication() { [NSApp terminate:nil]; }