// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#ifndef WINDOW_H
#define WINDOW_H

#include "top-view.h"

//static const CFStringRef kTrackingAreaID = CFSTR("id");

@interface      PrimaryWindow: NSWindow
@property void* windowP;
- (BOOL)canBecomeKeyWindow; // Allow create key window without title and resize bars.
@end

@interface PrimaryWindowDelegate: NSObject </*NSApplicationDelegate,*/ NSWindowDelegate>
- (void)windowDidBecomeKey:(NSNotification*)notification;
- (void)windowDidResignKey:(NSNotification*)notification;
- (void)windowDidResize:(NSNotification*)notification;
@end

void* CreateWindow(void* goWindow);
//void        SetWindowAcceptMouseMoved(void* self, bool accept);
const char* GetWindowTitle(void* self);
void        SetWindowTitle(void* self, CFStringRef title);

void SetWindowPossibleSize(void* self, NSSize min, NSSize max);

NSRect GetWindowGeometry(void* self);
void   SetWindowPos(void* self, NSPoint pos);
void   SetWindowSize(void* self, CGSize size);

void MinimizeWindow(void* self);
void MaximizeWindow(void* self);
void CloseWindow(void* self);

CGContextRef GetWindowContext(void* self);
CGFloat      GetWindowScaleFactor(void* self);

void InvalidateRegion(void* self, NSRect rect);
void Invalidate(void* self);

void windowMainEventCallback(void* window, unsigned char become);
void windowResizeCallback(void* window, NSSize size);

void addTrackingArea(void* self, bool move, int id, NSRect rect);
void replaceTrackingArea(void* self, bool move, int id, NSRect rect);
void removeTrackingArea(void* self, bool move, int id);

void setCursor(uint8 cursor);

#endif