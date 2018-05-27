// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include <Cocoa/Cocoa.h>

void* CreateWindow(int x, int y, int width, int height);
void MakeWindowKeyAndOrderFront(void *self);
void SetWindowTopView(void *self, void *topView);
const char* GetWindowTitle(void *self);
void SetWindowTitle(void *self, char* title);

NSRect GetWindowGeometry(void *self);
void SetWindowPos(void *self, NSPoint pos);
void SetWindowSize(void *self, CGSize size);

CGContextRef GetWindowContext(void *self);
CGFloat GetWindowScaleFactor(void *self);

void InvalidateRegion(void *self, NSRect rect);
void Invalidate(void *self);