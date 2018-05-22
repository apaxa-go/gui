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

    NSSize sizeInPoints = [self bounds].size;
    NSSize sizeInPixels = [self convertSizeToBacking:sizeInPoints];

    CGFloat scale = sizeInPixels.width/sizeInPoints.width; // TODO in window (for offline canvas) we get scale factor simpler

    if (scale<=0){ // TODO here we should sheck for integral scale factor, not for non-positive
        // Descale internal matrix
        CGContextScaleCTM (context,sizeInPoints.width/sizeInPixels.width,sizeInPoints.height/sizeInPixels.height);

        // Descale frame
        frame.size = [self convertSizeToBacking:frame.size]; // Do not convert frame fully because Y-axis in backing coordinate system is different (not flipped).
    }

    // Set color spaces to RGB(A)
    CGColorSpaceRef colorSpace = CGColorSpaceCreateWithName(kCGColorSpaceSRGB); // TODO may return NULL
    CGContextSetFillColorSpace(context,colorSpace);
    CGContextSetStrokeColorSpace(context,colorSpace);

    drawCallback(self.windowP,context,frame,scale);
}
@end

void *CreateTopView(void *window) {
    NSView* view = [[TopView alloc] initWithFrame:NSMakeRect(0, 0, 100, 100) windowP:window]; // TODO 100?
    return view;
}