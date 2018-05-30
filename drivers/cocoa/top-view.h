// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import <Cocoa/Cocoa.h>

@interface TopView : NSView
    @property void *windowP;
    - (BOOL) isFlipped; // Make coordinate left-top based.
    - (BOOL) acceptsFirstResponder; // Allow view responds to keyboard events.
    - (void) drawRect:(NSRect)frame;
    - (void) keyDown:(NSEvent *)event;
    - (void) keyUp:(NSEvent *)event;
@end

void *CreateTopView(void *window);

void drawCallback(void*, void*, NSRect);

void keyboardEventCallback(void*, uint8 event, uint16_t key, uint64_t modifiers);