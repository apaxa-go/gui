// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include "window.h"

#include <stdlib.h>

void *CreateWindow(int x, int y, int width, int height)
{
    NSWindow* window = [[NSWindow alloc] initWithContentRect:NSMakeRect(x, y, width, height)
                                              styleMask:NSWindowStyleMaskTitled  // TODO chack for valid deprecated NSTitledWindowMask replacement
                                              backing:NSBackingStoreBuffered defer:NO];
    [window setStyleMask:NSWindowStyleMaskBorderless]; // TODO chack for valid deprecated NSBorderlessWindowMask replacement
    return window;
}

void MakeWindowKeyAndOrderFront(void *self) {
    NSWindow *window = self;
    [window makeKeyAndOrderFront:nil];
}

void SetWindowTopView(void *self, void *topView){
    NSWindow *window = self;
    NSView *view = topView;
    [window setContentView:view];
}

const char* GetWindowTitle(void *self) {
    NSWindow *window = self;
    NSString *nsTitle = [window title];
    return [nsTitle UTF8String];
}

void SetWindowTitle(void *self, char *title) {
    NSWindow *window = self;
    NSString *nsTitle = [NSString stringWithUTF8String:title];
    [window setTitle:nsTitle];
    free(title);
}

NSRect GetWindowGeometry(void *self) {
    NSWindow *window = self;
    return [window frame];
}

void SetWindowPos(void *self, NSPoint pos) {
    NSWindow *window = self;
    [window setFrameTopLeftPoint:pos];
}

void SetWindowSize(void *self, CGSize size) {
    NSRect geometry;
    geometry.origin=GetWindowGeometry(self).origin;
    geometry.size=size;

    NSWindow *window = self;
    [window setFrame:geometry display:YES];
}

CGContextRef GetWindowContext(void *self){
    NSWindow *window = self;
    return [[window graphicsContext] CGContext];
}

CGFloat GetWindowScaleFactor(void *self){
    NSWindow *window = self;
    return [window backingScaleFactor];
}

void InvalidateRegion(void *self, NSRect rect){
    NSWindow *window = self;
    NSView *view = [window contentView];
    [view setNeedsDisplayInRect:rect];
}

void Invalidate(void *self){
    NSWindow *window = self;
    NSView *view = [window contentView];
    [view setNeedsDisplay:YES];
}