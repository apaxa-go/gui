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
    if (self) {
        self.windowP = window;
    }
    return self;
}

- (BOOL) isFlipped { return TRUE; }

- (BOOL) acceptsFirstResponder { return TRUE; }

- (void) drawRect:(NSRect)frame {
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
    CGContextSetFillColorSpace(context,colorSpace);
    CGContextSetStrokeColorSpace(context,colorSpace);

    // TODO do we need to set text matrix each drawRect?
    CGContextSetTextMatrix(context, (CGAffineTransform){1,0,0,-1,0,0});

    drawCallback(self.windowP, context, frame);
}

- (void) keyDown:(NSEvent *)event{
    // "ARepeat<<1" converts down event to 0 (first press) or 2 (repeat press).
    keyboardEventCallback(self.windowP, event.ARepeat<<1, event.keyCode, event.modifierFlags);
}

- (void) keyUp:(NSEvent *)event{
    keyboardEventCallback(self.windowP, 1, event.keyCode, event.modifierFlags);
}

@end

void *CreateTopView(void *window) {
    NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 100, 100) windowP:window]; // TODO 100?
    return view;
}