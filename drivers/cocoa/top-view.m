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

    drawCallback(self.windowP,context,frame);
}
@end

void *CreateTopView(void *window) {
    NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 100, 100) windowP:window]; // TODO 100?
    return view;
}