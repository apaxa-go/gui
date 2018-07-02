// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#ifndef WINDOW_H
#define WINDOW_H

#include "top-view.h"

//static const CFStringRef kTrackingAreaID = CFSTR("id");

@interface    PrimaryWindow: NSWindow
@property int windowID;
- (BOOL)canBecomeKeyWindow; // Allow create key window without title and resize bars.
@end

@interface PrimaryWindowDelegate: NSObject </*NSApplicationDelegate,*/ NSWindowDelegate>
- (void)windowDidBecomeKey:(NSNotification*)notification;
- (void)windowDidResignKey:(NSNotification*)notification;
- (void)windowDidResize:(NSNotification*)notification;
@end

void* CreateWindow(int windowID);
//void        SetWindowAcceptMouseMoved(void* self, bool accept);
const char* GetWindowTitle(void* self);
void        SetWindowTitle(void* self, CFStringRef title);

void SetWindowPossibleSize(void* self, NSSize min, NSSize max);

NSRect  GetWindowGeometry(void* self);
void    SetWindowGeometry(void* self, NSRect geometry);
NSPoint GetWindowPos(void* self);
void    SetWindowPos(void* self, NSPoint pos);
CGFloat GetWindowLeft(void* self);
void    SetWindowLeft(void* self, CGFloat left);
CGFloat GetWindowRight(void* self);
void    SetWindowRight(void* self, CGFloat right);
CGFloat GetWindowTop(void* self);
void    SetWindowTop(void* self, CGFloat top);
CGFloat GetWindowBottom(void* self);
void    SetWindowBottom(void* self, CGFloat bottom);
CGSize  GetWindowSize(void* self);
void    SetWindowSize(void* self, CGSize size, bool fixedRight, bool fixedBottom);
CGFloat GetWindowWidth(void* self);
void    SetWindowWidth(void* self, CGFloat width, bool fixedRight);
CGFloat GetWindowHeight(void* self);
void    SetWindowHeight(void* self, CGFloat height, bool fixedBottom);

void MinimizeWindow(void* self);
void DeminimizeWindow(void* self);
void MaximizeWindow(void* self);
void ZoomWindow(void* self);
void ToggleFullScreen(void* self);
void CloseWindow(void* self);

CGContextRef GetWindowContext(void* self);
CGFloat      GetWindowScaleFactor(void* self);

void InvalidateRegion(void* self, NSRect rect);
void Invalidate(void* self);

void windowMainStateEventCallback(int windowID, unsigned char become);
void windowResizeCallback(int windowID, NSSize size);
void windowMinimizeCallback(int windowID, unsigned char minimize);
void windowFullScreenCallback(int windowID, unsigned char enter);

void AddTrackingArea(void* self, bool move, int id, NSRect rect);
void ReplaceTrackingArea(void* self, bool move, int id, NSRect rect);
void RemoveTrackingArea(void* self, bool move, int id);

void setCursor(uint8 cursor);

#endif