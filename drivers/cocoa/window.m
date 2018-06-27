// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include "window.h"
#include <stdlib.h>

@implementation PrimaryWindow

- (id)initWithStyleMask:(NSWindowStyleMask)styleMask windowID:(int)windowID {
	self = [super initWithContentRect:NSMakeRect(300, 300, 0, 0) //
	                        styleMask:styleMask
	                          backing:NSBackingStoreBuffered
	                            defer:NO];
	if (self) { self.windowID = windowID; }
	return self;
}

- (BOOL)canBecomeKeyWindow {
	return TRUE;
}

@end

@implementation PrimaryWindowDelegate

- (void)windowDidBecomeKey:(NSNotification*)notification {
	PrimaryWindow* window = notification.object;
	windowMainEventCallback(window.windowID, true);
}
- (void)windowDidResignKey:(NSNotification*)notification {
	PrimaryWindow* window = notification.object;
	windowMainEventCallback(window.windowID, false);
}

- (void)windowDidResize:(NSNotification*)notification {
	PrimaryWindow* window = notification.object;
	NSSize         size   = GetWindowGeometry(window).size;
	windowResizeCallback(window.windowID, size);
}

@end

void* CreateWindow(int windowID) {
	// Attention: A lot of hacks here!
	// 1. We need to keep original window buttons or create minimize button to miniaturize method works, so we hide it manually.

	NSWindow* window                  = [[PrimaryWindow alloc]
        initWithStyleMask:NSWindowStyleMaskBorderless | NSWindowStyleMaskMiniaturizable // | NSWindowStyleMaskTitled | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskFullSizeContentView
                 windowID:windowID];
	window.delegate                   = [PrimaryWindowDelegate alloc];
	window.titleVisibility            = NSWindowTitleHidden;
	window.titlebarAppearsTransparent = YES;
	NSView* topView                   = CreateTopView(windowID); // TODO check for nil
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

/*void SetWindowAcceptMouseMoved(void* self, bool accept) {
	NSWindow* window = self;
	[window setAcceptsMouseMovedEvents:accept];
}*/

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

void SetWindowPossibleSize(void* self, NSSize min, NSSize max) {
	NSWindow* window      = self;
	window.contentMinSize = min;
	window.contentMaxSize = max;
	//window.minSize = min;
	//window.maxSize = max;
	NSSize size   = GetWindowGeometry(self).size;
	bool   resize = false;

	if (size.width < min.width) {
		resize     = true;
		size.width = min.width;
	} else if (size.width > max.width) {
		resize     = true;
		size.width = max.width;
	}
	if (size.height < min.height) {
		resize      = true;
		size.height = min.height;
	} else if (size.height > max.height) {
		resize      = true;
		size.height = max.height;
	}

	if (resize) { SetWindowSize(self, size, false, false); }
}

// All {Get/Set}Window{Geometry/Pos/Left/Top} functions are related to LeftTop window corner with Y-axis increasing down (inverted Y).

NSRect GetWindowGeometry(void* self) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	geometry.origin.y  = -geometry.origin.y - geometry.size.height; // invert Y and move window base from LB to LT
	return geometry;
}

void SetWindowGeometry(void* self, NSRect geometry) {
	NSWindow* window  = self;
	geometry.origin.y = -geometry.origin.y - geometry.size.height;
	[window setFrame:geometry display:YES];
}

NSPoint GetWindowPos(void* self) { return GetWindowGeometry(self).origin; }

void SetWindowPos(void* self, NSPoint pos) {
	NSWindow* window = self;
	pos.y            = -pos.y;
	[window setFrameTopLeftPoint:pos];
}

CGFloat GetWindowLeft(void* self) {
	NSWindow* window = self;
	return [window frame].origin.x;
}

void SetWindowLeft(void* self, CGFloat left) {
	NSWindow* window = self;
	NSPoint   pos    = [window frame].origin;
	pos.x            = left;
	[window setFrameOrigin:pos];
}

CGFloat GetWindowRight(void* self) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	return geometry.origin.x + geometry.size.width;
}

void SetWindowRight(void* self, CGFloat right) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	geometry.origin.x  = right - geometry.size.width;
	[window setFrameOrigin:geometry.origin];
}

CGFloat GetWindowTop(void* self) { return GetWindowPos(self).y; }

void SetWindowTop(void* self, CGFloat top) {
	NSPoint pos = GetWindowPos(self);
	pos.y       = top;
	SetWindowPos(self, pos);
}

CGFloat GetWindowBottom(void* self) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	return -geometry.origin.y;
}

void SetWindowBottom(void* self, CGFloat bottom) {
	NSWindow* window = self;
	NSPoint   pos    = [window frame].origin;
	pos.y            = -bottom;
	[window setFrameOrigin:pos];
}

CGSize GetWindowSize(void* self) {
	NSWindow* window = self;
	return [window frame].size;
}

void SetWindowSize(void* self, CGSize size, bool fixedRight, bool fixedBottom) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	if (fixedRight) { geometry.origin.x += geometry.size.width - size.width; }
	if (!fixedBottom) { geometry.origin.y += geometry.size.height - size.height; }
	geometry.size = size;
	[window setFrame:geometry display:YES];
}

CGFloat GetWindowWidth(void* self) { return GetWindowSize(self).width; }

void SetWindowWidth(void* self, CGFloat width, bool fixedRight) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	if (fixedRight) { geometry.origin.x += geometry.size.width - width; }
	geometry.size.width = width;
	[window setFrame:geometry display:YES];
}

CGFloat GetWindowHeight(void* self) { return GetWindowSize(self).height; }

void SetWindowHeight(void* self, CGFloat height, bool fixedBottom) {
	NSWindow* window   = self;
	NSRect    geometry = [window frame];
	if (!fixedBottom) { geometry.origin.y += geometry.size.height - height; }
	geometry.size.height = height;
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

//
// Tracking area related
//

NSTrackingAreaOptions makeTrackingAreaOptions(bool move) {
	NSTrackingAreaOptions r = NSTrackingActiveAlways; // TODO allow to change Always to other values.
	r |= move ? NSTrackingMouseMoved : NSTrackingMouseEnteredAndExited;
	return r;
}

CFMutableDictionaryRef createTrackingAreaUserInfo(int id) {
	CFMutableDictionaryRef r = CFDictionaryCreateMutable(NULL, 0, NULL, NULL);

	CFNumberRef idRef = CFNumberCreate(NULL, kCFNumberSInt32Type, &id);
	CFDictionarySetValue(r, @"id", idRef);
	CFRelease(idRef);

	return r;
}

NSTrackingArea* getTrackingAreaByID(NSView* self, bool move, int id) {
	CFNumberRef idRef = CFNumberCreate(NULL, kCFNumberSInt32Type, &id);
	for (NSTrackingArea* area in [self trackingAreas]) {
		bool found = (CFNumberRef)area.userInfo[@"id"] == idRef &&
		             (area.options & (move ? NSTrackingMouseMoved : NSTrackingMouseEnteredAndExited));
		if (found) {
			CFRelease(idRef);
			return area;
		}
	}
	CFRelease(idRef);
	return nil;
}

void addTrackingArea(void* self, bool move, int id, NSRect rect) {
	NSWindow* window = self;
	NSView*   view   = [window contentView];

	CFMutableDictionaryRef userInfo = createTrackingAreaUserInfo(id);
	NSTrackingArea*        area     = [[NSTrackingArea alloc] initWithRect:rect //
                                                        options:makeTrackingAreaOptions(move)
                                                          owner:view
                                                       userInfo:(__bridge NSDictionary*)userInfo];
	CFRelease(userInfo);

	[view addTrackingArea:area];
	CFRelease(area);
}

void removeTrackingArea(void* self, bool move, int id) {
	NSWindow*       window = self;
	NSView*         view   = [window contentView];
	NSTrackingArea* area   = getTrackingAreaByID(view, move, id);
	if (area != nil) { [view removeTrackingArea:area]; }
}

void replaceTrackingArea(void* self, bool move, int id, NSRect rect) {
	removeTrackingArea(self, move, id);
	addTrackingArea(self, move, id, rect);
}

NSCursor* translateCursor(uint8 cursor) {
	switch (cursor) {
	case 1: return [[NSCursor class] performSelector:@selector(_windowResizeNorthSouthCursor)];
	case 2: return [[NSCursor class] performSelector:@selector(_windowResizeEastWestCursor)];
	case 3: return [[NSCursor class] performSelector:@selector(_windowResizeNorthWestSouthEastCursor)];
	case 4: return [[NSCursor class] performSelector:@selector(_windowResizeNorthEastSouthWestCursor)];
	default: return [NSCursor arrowCursor];
	}
}

void setCursor(uint8 cursor) {
	NSCursor* c = translateCursor(cursor);
	[c set];
}