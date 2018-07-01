// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#ifndef TOP_VIEW_H
#define TOP_VIEW_H

#import <Cocoa/Cocoa.h>

@interface    TopView: NSView
@property int windowID;
- (BOOL)isFlipped;             // Make coordinate left-top based.
- (BOOL)acceptsFirstResponder; // Allow view responds to keyboard events.

- (void)drawRect:(NSRect)frame;

- (void)keyDown:(NSEvent*)event;
- (void)keyUp:(NSEvent*)event;

- (NSPoint)mouseLocation;

@property (retain) NSTimer* mouseTimer;
@property uint8 mouseRepeatCount;
@property bool  mouseLastIsDown;
@property uint8 mouseLastButton;
@property NSPoint mouseFirstPoint;
@property uint64 mouseFirstModifiers;
- (BOOL)acceptsFirstMouse:(NSEvent*)event;
- (void)mouseDown:(NSEvent*)event;
- (void)mouseUp:(NSEvent*)event;
- (void)rightMouseDown:(NSEvent*)event;
- (void)rightMouseUp:(NSEvent*)event;
- (void)otherMouseDown:(NSEvent*)event;
- (void)otherMouseUp:(NSEvent*)event;
- (void)mouseMoved:(NSEvent*)event;
@property NSPoint mouseDragBase;
- (void)mouseDragged:(NSEvent*)event;

- (void)scrollWheel:(NSEvent*)event;

- (void)dealloc;
@end

double  distance(NSPoint p0, NSPoint p1);
NSView* CreateTopView(int windowID);

void drawCallback(int windowID, CGContextRef context, NSRect);
void keyboardEventCallback(int windowID, uint8 event, uint16_t key, uint64_t modifiers);
void pointerKeyEventCallback(int windowID, uint8 event, uint8 button, NSPoint point, uint64 modifiers);
void pointerDragEventCallback(int windowID, NSPoint delta);
void pointerMoveEventCallback(int windowID, NSPoint point);
void pointerEnterLeaveEventCallback(int windowID, int, bool);
void scrollEventCallback(int windowID, NSPoint delta, NSPoint point, uint64 modifiers);
void modifiersEventCallback(int windowID, uint64 modifiers);

#endif