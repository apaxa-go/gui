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