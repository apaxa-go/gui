// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import "top-view.h"

@implementation TopView

- (id)initWithFrame:(NSRect)frame { // TODO remove this constructor?
	return [super initWithFrame:frame];
}

- (id)initWithFrame:(NSRect)frame windowID:(int)windowID {
	self = [super initWithFrame:frame];
	if (self) { self.windowID = windowID; }
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

	drawCallback(self.windowID, context, frame);
}

- (void)dealloc {
	if (self.mouseTimer.valid) { [[self mouseTimer] invalidate]; }
	[super dealloc];
}

- (void)cursorUpdate:(NSEvent*)event {
	// Prevent default implementation from changing cursor to default.
}

- (NSPoint)mouseLocation {
	//NSPoint r = [NSEvent mouseLocation];
	//NSRect rect = NSMakeRect(r.x,r.y,0,0);
	//rect = [[self window] convertRectFromScreen: rect];
	//r=rect.origin;
	NSPoint r = [[self window] mouseLocationOutsideOfEventStream];
	r         = [self convertPoint:r fromView:nil];
	return r;
}

//
// Keyboard events
//

- (void)keyDown:(NSEvent*)event {
	// "ARepeat<<1" converts down event to 0 (first press) or 2 (repeat press).
	keyboardEventCallback(self.windowID, event.ARepeat << 1, event.keyCode, event.modifierFlags);
}

- (void)keyUp:(NSEvent*)event {
	keyboardEventCallback(self.windowID, 1, event.keyCode, event.modifierFlags);
}

- (void)flagsChanged:(NSEvent*)event {
	modifiersEventCallback(self.windowID, event.modifierFlags);
}

//
// Mouse button events
//
- (BOOL)acceptsFirstMouse:(NSEvent*)event {
	return true;
}

- (void)mouseDown:(NSEvent*)event {
	[self mouseButton:true Button:0 Event:event];
}

- (void)mouseUp:(NSEvent*)event {
	[self mouseButton:false Button:0 Event:event];
}

- (void)rightMouseDown:(NSEvent*)event {
	[self mouseButton:true Button:1 Event:event];
}

- (void)rightMouseUp:(NSEvent*)event {
	[self mouseButton:false Button:1 Event:event];
}

- (void)otherMouseDown:(NSEvent*)event {
	[self mouseButton:true Button:event.buttonNumber Event:event];
}

- (void)otherMouseUp:(NSEvent*)event {
	[self mouseButton:false Button:event.buttonNumber Event:event];
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
		pointerKeyEventCallback(self.windowID, self.mouseRepeatCount, self.mouseLastButton, self.mouseFirstPoint, self.mouseFirstModifiers);
		self.mouseRepeatCount = 0;
	}
}

- (void)mouseButton:(bool)down Button:(uint8)button Event:(NSEvent*)event {
	//
	// Related to drag
	//
	if (down) { // TODO multiple down
		NSPoint point      = [NSEvent mouseLocation];
		point.y            = -point.y;
		self.mouseDragBase = point;
		//NSLog(@"==========================%@", NSStringFromPoint(self.mouseDragBase));
		//self.mouseDragLast = NAN;
	}

	//
	// Related to clicks
	//
	NSPoint point = [self mouseLocation];
	if ([self mouseButtonDelayedCanPush:down Button:button Point:point]) {
		[self mouseButtonDelayedPush];
	} else {
		if (self.mouseRepeatCount > 0) { [self mouseButtonDelayedPop:nil]; }
		if (down) { [self mouseButtonDelayedInit:button Point:point Modifiers:event.modifierFlags]; }
	}

	//
	//
	//
	pointerKeyEventCallback(self.windowID, down ? 0 : 255, button, point, event.modifierFlags);
}

//
// Mouse move events
//

- (void)mouseDragged:(NSEvent*)event {
	NSPoint location = [NSEvent mouseLocation];
	//NSLog(@"%@", NSStringFromPoint(location));
	CGFloat x = location.x - self.mouseDragBase.x;
	CGFloat y = -location.y - self.mouseDragBase.y;
	pointerDragEventCallback(self.windowID, NSMakePoint(x, y));
}

- (void)mouseMoved:(NSEvent*)event {
	pointerMoveEventCallback(self.windowID, [self mouseLocation]);
}

int getTrackingAreaID(NSTrackingArea* area) {
	CFNumberRef idRef = (CFNumberRef)area.userInfo[@"id"]; // here we trust ...
	int         id;
	CFNumberGetValue(idRef, kCFNumberIntType, &id); // and here ...
	return id;
}

- (void)mouseEntered:(NSEvent*)event {
	if ([event trackingArea].userInfo[@"active"] != nil) { return; }
	//NSLog(@"Enter %d", getTrackingAreaID([event trackingArea]));
	[(NSMutableDictionary*)[event trackingArea].userInfo setObject:[NSNull null] forKey:@"active"];
	pointerEnterLeaveEventCallback(self.windowID, getTrackingAreaID([event trackingArea]), true);
}

- (void)mouseExited:(NSEvent*)event {
	if ([event trackingArea].userInfo[@"active"] == nil) { return; }
	//NSLog(@"Leave %d", getTrackingAreaID([event trackingArea]));
	[(NSMutableDictionary*)[event trackingArea].userInfo removeObjectForKey:@"active"];
	pointerEnterLeaveEventCallback(self.windowID, getTrackingAreaID([event trackingArea]), false);
}

//
// Scroll events
//

- (void)scrollWheel:(NSEvent*)event {
	NSPoint delta = NSMakePoint(event.deltaX, event.deltaY);
	scrollEventCallback(self.windowID, delta, [self mouseLocation], event.modifierFlags);
}

@end

NSView* CreateTopView(int windowID) {
	NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0) windowID:windowID];
	return view;
}

double distance(NSPoint p0, NSPoint p1) { return sqrt(pow(p1.x - p0.x, 2) + pow(p1.y - p0.y, 2)); }
