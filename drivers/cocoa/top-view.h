#import <Cocoa/Cocoa.h>

@interface TopView : NSView
@property void *windowP;
//@property(getter=isFlipped, readonly) BOOL flipped;
- (BOOL) isFlipped;
- (void) drawRect:(NSRect)frame;
@end

void *CreateTopView(void *window);

void drawCallback(void*, void*, NSRect);