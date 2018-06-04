// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import <Cocoa/Cocoa.h>

@interface      TopView: NSView
@property void* windowP;
- (BOOL)isFlipped;             // Make coordinate left-top based.
- (BOOL)acceptsFirstResponder; // Allow view responds to keyboard events.

- (void)drawRect:(NSRect)frame;

- (void)keyDown:(NSEvent*)event;
- (void)keyUp:(NSEvent*)event;

@property (retain) NSTimer* mouseTimer;
@property uint8 mouseRepeatCount;
@property bool  mouseLastIsDown;
@property uint8 mouseLastButton;
@property NSPoint mouseFirstPoint;
@property uint64 mouseFirstModifiers;
- (void)mouseDown:(NSEvent*)event;
- (void)mouseUp:(NSEvent*)event;
- (void)rightMouseDown:(NSEvent*)event;
- (void)rightMouseUp:(NSEvent*)event;
- (void)otherMouseDown:(NSEvent*)event;
- (void)otherMouseUp:(NSEvent*)event;
- (void)mouseMoved:(NSEvent*)event;

- (void)scrollWheel:(NSEvent*)event;

- (void)dealloc;
@end

double distance(NSPoint p0, NSPoint p1);
void*  CreateTopView(void* window);

void drawCallback(void* window, CGContextRef context, NSRect);
void keyboardEventCallback(void* window, uint8 event, uint16_t key, uint64_t modifiers);
void pointerKeyEventCallback(void* window, uint8 event, uint8 button, NSPoint point, uint64 modifiers);
void pointerMoveEventCallback(void* window, NSPoint point);
void scrollEventCallback(void* window, double deltaX, double deltaY);
