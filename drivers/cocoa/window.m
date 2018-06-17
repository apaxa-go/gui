// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include "window.h"
#include <stdlib.h>

@implementation PrimaryWindow

- (id)initWithStyleMask:(NSWindowStyleMask)styleMask windowP:(void*)window {
	self = [super initWithContentRect:NSMakeRect(0, 0, 0, 0) //
	                        styleMask:styleMask
	                          backing:NSBackingStoreBuffered
	                            defer:NO];
	if (self) { self.windowP = window; }
	return self;
}

- (BOOL)canBecomeKeyWindow {
	return TRUE;
}

@end

@implementation PrimaryWindowDelegate

- (void)windowDidBecomeKey:(NSNotification*)notification {
	PrimaryWindow* window = notification.object;
	windowMainEventCallback(window.windowP, true);
}
- (void)windowDidResignKey:(NSNotification*)notification {
	PrimaryWindow* window = notification.object;
	windowMainEventCallback(window.windowP, false);
}

@end

void* CreateWindow(void* goWindow) {
	// Attention: A lot of hacks here!
	// 1. We need to keep original window buttons or create minimize button to miniaturize method works, so we hide it manually.

	NSWindow* window                  = [[PrimaryWindow alloc]
        initWithStyleMask:NSWindowStyleMaskBorderless | NSWindowStyleMaskMiniaturizable // | NSWindowStyleMaskTitled | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskFullSizeContentView
                  windowP:goWindow];
	window.delegate                   = [PrimaryWindowDelegate alloc];
	window.titleVisibility            = NSWindowTitleHidden;
	window.titlebarAppearsTransparent = YES;
	NSView* topView                   = CreateTopView(goWindow); // TODO check for nil
	[window setContentView:topView];

	// Hide window buttons.

	NSButton* minimizeButton = [NSWindow standardWindowButton:NSWindowMiniaturizeButton
	                                             forStyleMask:NSWindowStyleMaskTitled | NSWindowStyleMaskMiniaturizable];
	[minimizeButton setHidden:YES];
	[window.contentView addSubview:minimizeButton];

	/*[[window standardWindowButton:NSWindowCloseButton] setHidden:YES];
    [[window standardWindowButton:NSWindowMiniaturizeButton] setHidden:YES];
    [[window standardWindowButton:NSWindowZoomButton] setHidden:YES];
    [[window standardWindowButton:NSWindowToolbarButton] setHidden:YES];
    [[window standardWindowButton:NSWindowDocumentIconButton] setHidden:YES];
    [[window standardWindowButton:NSWindowDocumentVersionsButton] setHidden:YES];*/

	/*for (id subview in window.contentView.superview.subviews) {
        if ([subview isKindOfClass:NSClassFromString(@"NSTitlebarContainerView")]) {
            NSView *titlebarView = [subview subviews][0];
            for (id button in titlebarView.subviews) {
                if ([button isKindOfClass:[NSButton class]]) {
                    [button setHidden:YES];
                }
            }
        }
    }*/

	[window makeKeyAndOrderFront:nil];

	return window;
}

void SetWindowAcceptMouseMoved(void* self, bool accept) {
	NSWindow* window = self;
	[window setAcceptsMouseMovedEvents:accept];
}

const char* GetWindowTitle(void* self) {
	NSWindow* window  = self;
	NSString* nsTitle = [window title];
	return [nsTitle UTF8String];
}

void SetWindowTitle(void* self, CFStringRef title) {
	NSWindow* window = self;
	[window setTitle:(__bridge NSString*)title];
	CFRelease(title);
}

NSRect GetWindowGeometry(void* self) {
	NSWindow* window = self;
	return [window frame];
}

void SetWindowPos(void* self, NSPoint pos) {
	NSWindow* window = self;
	[window setFrameTopLeftPoint:pos];
}

void SetWindowSize(void* self, CGSize size) {
	NSRect geometry;
	geometry.origin = GetWindowGeometry(self).origin;
	geometry.size   = size;

	NSWindow* window = self;
	[window setFrame:geometry display:YES];
}

void MinimizeWindow(void* self) {
	NSWindow* window = self;
	[window miniaturize:(id)nil];
}

void MaximizeWindow(void* self) {
	NSWindow* window = self;
	[window setFrame:[[NSScreen mainScreen] visibleFrame] display:YES];
}

void CloseWindow(void* self) {
	NSWindow* window = self;
	[window close];
}

CGContextRef GetWindowContext(void* self) {
	NSWindow* window = self;
	return [[window graphicsContext] CGContext];
}

CGFloat GetWindowScaleFactor(void* self) {
	NSWindow* window = self;
	return [window backingScaleFactor];
}

void InvalidateRegion(void* self, NSRect rect) {
	NSWindow* window = self;
	NSView*   view   = [window contentView];
	[view setNeedsDisplayInRect:rect];
}

void Invalidate(void* self) {
	NSWindow* window = self;
	NSView*   view   = [window contentView];
	[view setNeedsDisplay:YES];
}
