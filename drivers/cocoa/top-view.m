// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import "top-view.h"

@implementation TopView

- (id)initWithFrame:(NSRect)frame {
	return [super initWithFrame:frame];
}

- (id)initWithFrame:(NSRect)frame windowP:(void*)window {
	self = [super initWithFrame:frame];
	if (self) { self.windowP = window; }
	return self;
}

- (BOOL)isFlipped {
	return TRUE;
}

- (BOOL)acceptsFirstResponder {
	return TRUE;
}

- (void)drawRect:(NSRect)frame {
	//[[NSColor redColor] set];
	//[NSBezierPath fillRect:frame];

	CGContextRef context = [[NSGraphicsContext currentContext] graphicsPort]; // TODO

	/*
NSSize sizeInPoints = [self bounds].size;
NSSize sizeInPixels = [self convertSizeToBacking:sizeInPoints];

CGFloat scale = sizeInPixels.width/sizeInPoints.width;
*/

	// TODO do we need to set color space each drawRect?
	// Set color spaces to RGB(A)
	CGColorSpaceRef colorSpace = CGColorSpaceCreateWithName(kCGColorSpaceSRGB); // TODO may return NULL
	CGContextSetFillColorSpace(context, colorSpace);
	CGContextSetStrokeColorSpace(context, colorSpace);
	CFRelease(colorSpace);

	// TODO do we need to set text matrix each drawRect?
	CGContextSetTextMatrix(context, (CGAffineTransform){1, 0, 0, -1, 0, 0});

	drawCallback(self.windowP, context, frame);
}

- (void)dealloc {
	if (self.mouseTimer.valid) { [[self mouseTimer] invalidate]; }
	[super dealloc];
}

- (void)keyDown:(NSEvent*)event {
	// "ARepeat<<1" converts down event to 0 (first press) or 2 (repeat press).
	keyboardEventCallback(self.windowP, event.ARepeat << 1, event.keyCode, event.modifierFlags);
}

- (void)keyUp:(NSEvent*)event {
	keyboardEventCallback(self.windowP, 1, event.keyCode, event.modifierFlags);
}

//
// Mouse buttons
//

- (void)mouseDown:(NSEvent*)event {
	[self mouseButton:true //
	           Button:0
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)mouseUp:(NSEvent*)event {
	[self mouseButton:false //
	           Button:0
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)rightMouseDown:(NSEvent*)event {
	[self mouseButton:true //
	           Button:1
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)rightMouseUp:(NSEvent*)event {
	[self mouseButton:false //
	           Button:1
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)otherMouseDown:(NSEvent*)event {
	[self mouseButton:true //
	           Button:event.buttonNumber
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)otherMouseUp:(NSEvent*)event {
	[self mouseButton:false //
	           Button:event.buttonNumber
	            Point:[self currentMousePos]
	        Modifiers:event.modifierFlags];
}

- (void)mouseButtonTimerRestart {
	if (self.mouseTimer.valid) { [[self mouseTimer] invalidate]; }
	self.mouseTimer = [NSTimer scheduledTimerWithTimeInterval:0.15 // TODO move 0.15 to constant and tune
	                                                   target:self
	                                                 selector:@selector(mouseButtonDelayedPop:) // TODO does ":" required here?
	                                                 userInfo:nil
	                                                  repeats:NO];
}

- (void)mouseButtonDelayedInit:(uint8)button Point:(NSPoint)point Modifiers:(uint64)modifiers {
	self.mouseRepeatCount    = 1;
	self.mouseLastIsDown     = true;
	self.mouseLastButton     = button;
	self.mouseFirstPoint     = point;
	self.mouseFirstModifiers = modifiers;
	[self mouseButtonTimerRestart];
}

- (bool)mouseButtonDelayedCanPush:(bool)down Button:(uint8)button Point:(NSPoint)point {
	return self.mouseRepeatCount > 0 && down != self.mouseLastIsDown && button == self.mouseLastButton &&
	       distance(point, self.mouseFirstPoint) < 10; // TODO move to constant and tune value.
}

- (void)mouseButtonDelayedPush {
	self.mouseLastIsDown = !self.mouseLastIsDown;
	if (self.mouseLastIsDown) { self.mouseRepeatCount++; }
	[self mouseButtonTimerRestart];
}

- (void)mouseButtonDelayedPop:(NSTimer*)timer {
	if (self.mouseTimer.valid) { [[self mouseTimer] invalidate]; }
	if (self.mouseLastIsDown) { self.mouseRepeatCount--; }
	if (self.mouseRepeatCount > 0) {
		pointerKeyEventCallback(self.windowP, self.mouseRepeatCount, self.mouseLastButton, self.mouseFirstPoint, self.mouseFirstModifiers);
		self.mouseRepeatCount = 0;
	}
}

- (void)mouseButton:(bool)down Button:(uint8)button Point:(NSPoint)point Modifiers:(uint64)modifiers {
	if ([self mouseButtonDelayedCanPush:down Button:button Point:point]) {
		[self mouseButtonDelayedPush];
	} else {
		if (self.mouseRepeatCount > 0) { [self mouseButtonDelayedPop:nil]; }
		if (down) { [self mouseButtonDelayedInit:button Point:point Modifiers:modifiers]; }
	}
	pointerKeyEventCallback(self.windowP, down ? 0 : 255, button, point, modifiers);
}

//
// Mouse move
//

- (void)mouseMoved:(NSEvent*)event {
	pointerMoveEventCallback(self.windowP, [self currentMousePos]);
	// TODO There is no move events between press & release. May be send move event on release and/or
	// in dragging event?
}

- (NSPoint)currentMousePos {
	NSPoint r = [[self window] mouseLocationOutsideOfEventStream]; // NSEvent.mouseLocation;
	r         = [self convertPoint:r fromView:nil];
	return r;
}

//
// Scroll
//

- (void)scrollWheel:(NSEvent*)event {
	scrollEventCallback(self.windowP, event.deltaX, event.deltaY);
}

@end

void* CreateTopView(void* window) {
	NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 100, 100) windowP:window]; // TODO 100?
	return view;
}

double distance(NSPoint p0, NSPoint p1) { return sqrt(pow(p1.x - p0.x, 2) + pow(p1.y - p0.y, 2)); }
