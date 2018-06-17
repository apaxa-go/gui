// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import "top-view.h"

@implementation TopView

- (id)initWithFrame:(NSRect)frame { // TODO remove this constructor?
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
	CGContextRef context = [[NSGraphicsContext currentContext] graphicsPort];

	// TODO do we need to set color space each drawRect?
	// Set color spaces to RGB(A)
	CGColorSpaceRef colorSpace = CGColorSpaceCreateWithName(kCGColorSpaceSRGB);
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

- (NSPoint)mouseLocation {
	NSPoint r = [[self window] mouseLocationOutsideOfEventStream];
	r         = [self convertPoint:r fromView:nil];
	return r;
}

//
// Keyboard events
//

- (void)keyDown:(NSEvent*)event {
	// "ARepeat<<1" converts down event to 0 (first press) or 2 (repeat press).
	keyboardEventCallback(self.windowP, event.ARepeat << 1, event.keyCode, event.modifierFlags);
}

- (void)keyUp:(NSEvent*)event {
	keyboardEventCallback(self.windowP, 1, event.keyCode, event.modifierFlags);
}

//
// Mouse button events
//
- (BOOL)acceptsFirstMouse:(NSEvent*)event {
	return true;
}

- (void)mouseDown:(NSEvent*)event {
	//self.initialWindowLocation = [event locationInWindow];
	//self.windowDragging=true;
	[self mouseButton:true //
	           Button:0
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)mouseUp:(NSEvent*)event {
	//self.windowDragging = false;
	[self mouseButton:false //
	           Button:0
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)rightMouseDown:(NSEvent*)event {
	[self mouseButton:true //
	           Button:1
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)rightMouseUp:(NSEvent*)event {
	[self mouseButton:false //
	           Button:1
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)otherMouseDown:(NSEvent*)event {
	[self mouseButton:true //
	           Button:event.buttonNumber
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)otherMouseUp:(NSEvent*)event {
	[self mouseButton:false //
	           Button:event.buttonNumber
	            Point:[self mouseLocation]
	        Modifiers:event.modifierFlags];
}

- (void)mouseButtonTimerRestart {
	if (self.mouseTimer.valid) { [[self mouseTimer] invalidate]; }
	self.mouseTimer = [NSTimer scheduledTimerWithTimeInterval:0.15 // TODO move 0.15 to constant and tune
	                                                   target:self
	                                                 selector:@selector(mouseButtonDelayedPop:)
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
	//
	// Related to drag
	//
	if (down) {                     // TODO multiple down
		self.mouseDragBase = point; // TODO may we use mouseFirstPoint here?
		                            //self.mouseDragLast = NAN;
	}

	//
	// Related to clicks
	//
	if ([self mouseButtonDelayedCanPush:down Button:button Point:point]) {
		[self mouseButtonDelayedPush];
	} else {
		if (self.mouseRepeatCount > 0) { [self mouseButtonDelayedPop:nil]; }
		if (down) { [self mouseButtonDelayedInit:button Point:point Modifiers:modifiers]; }
	}

	//
	//
	//
	pointerKeyEventCallback(self.windowP, down ? 0 : 255, button, point, modifiers);
}

//
// Mouse move events
//

/*
- (void)startWindowDragging{
	self.windowDragging=true;
}
*/

- (void)mouseDragged:(NSEvent*)event {
	/*
	if (self.windowDragging){
		[self windowDragged:event];
		return;
	}
	*/
	NSPoint location = [self mouseLocation];
	CGFloat x        = location.x - self.mouseDragBase.x;
	CGFloat y        = location.y - self.mouseDragBase.y;
	pointerDragEventCallback(self.windowP, NSMakePoint(x, y));
}

/*
- (void)windowDragged:(NSEvent *)event {
    NSRect screenVisibleFrame = [[NSScreen mainScreen] visibleFrame];
    NSRect windowFrame = [self.window frame];
    NSPoint newOrigin = windowFrame.origin;

    // Get the mouse location in window coordinates.
    NSPoint currentLocation = [event locationInWindow];
    // Update the origin with the difference between the new mouse location and the old mouse location.
    newOrigin.x += (currentLocation.x - self.initialWindowLocation.x);
    newOrigin.y += (currentLocation.y - self.initialWindowLocation.y);

    // Don't let window get dragged up under the menu bar
    if ((newOrigin.y + windowFrame.size.height) > (screenVisibleFrame.origin.y + screenVisibleFrame.size.height)) {
        newOrigin.y = screenVisibleFrame.origin.y + (screenVisibleFrame.size.height - windowFrame.size.height);
    }

    // Move the window to the new location
    [self.window setFrameOrigin:newOrigin];
}
*/

- (void)mouseMoved:(NSEvent*)event {
	pointerMoveEventCallback(self.windowP, [self mouseLocation]);
}

//
// Scroll events
//

- (void)scrollWheel:(NSEvent*)event {
	NSPoint delta = NSMakePoint(event.deltaX, event.deltaY);
	scrollEventCallback(self.windowP, delta, [self mouseLocation], event.modifierFlags);
}

@end

NSView* CreateTopView(void* goWindow) {
	NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0) windowP:goWindow];
	return view;
}

double distance(NSPoint p0, NSPoint p1) { return sqrt(pow(p1.x - p0.x, 2) + pow(p1.y - p0.y, 2)); }
