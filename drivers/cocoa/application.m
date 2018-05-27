// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import "application.h"

void* InitApplication() {
	return [NSApplication sharedApplication];
}

void ApplicationRun(void *app) {
    @autoreleasepool {
        NSApplication* a = (NSApplication*)app;
        [a setActivationPolicy:NSApplicationActivationPolicyRegular];
        [a run];
	}
}