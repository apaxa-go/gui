// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#include "common.h"

CFAttributedStringRef makeAttrStringRef(CFStringRef str, CTFontRef font){
    // TODO check functions results for nil???
    CFStringRef keys[] = { kCTFontAttributeName };
    CFTypeRef values[] = { font };

    CFDictionaryRef attributes =
        CFDictionaryCreate(kCFAllocatorDefault, (const void**)&keys,
            (const void**)&values, sizeof(keys) / sizeof(keys[0]),
            &kCFTypeDictionaryKeyCallBacks,
            &kCFTypeDictionaryValueCallBacks);

    CFAttributedStringRef attrString = CFAttributedStringCreate(kCFAllocatorDefault, str, attributes);
    CFRelease(attributes);
    return attrString;
}

/*
void drawTextLine(CGContextRef context, CFStringRef str, CTFontRef font, CGPoint pos){
    // TODO check functions results for nil???
    CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
    CTLineRef line = CTLineCreateWithAttributedString(attrStr);

    CGContextSetTextPosition(context, pos.x, pos.y);
    CTLineDraw(line, context);

    CFRelease(line);
    CFRelease(attrStr);
}
*/

void DrawTextLine(CGContextRef context, const UInt8 *strBytes, CFIndex strLen, CTFontRef font, CGPoint pos){
    // TODO check functions results for nil???
    CFStringRef str = makeStringRef(strBytes, strLen);
    CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
    CFRelease(str);
    CTLineRef line = CTLineCreateWithAttributedString(attrStr);
    CFRelease(attrStr);
    CGContextSetTextPosition(context, pos.x, pos.y);
    CTLineDraw(line, context);
    CFRelease(line);
}

/*
CGRect getTextLineGeometry(CGContextRef context, CFStringRef str, CTFontRef font){
    // TODO check functions results for nil???
    CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
    CTLineRef line = CTLineCreateWithAttributedString(attrStr);
    return CTLineGetImageBounds(line, context);
}
*/

CGRect GetTextLineGeometry(CGContextRef context, const UInt8 *strBytes, CFIndex strLen, CTFontRef font){
    // TODO check functions results for nil???
    CFStringRef str = makeStringRef(strBytes, strLen);
    CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
    CFRelease(str);
    CTLineRef line = CTLineCreateWithAttributedString(attrStr);
    CFRelease(attrStr);
    CGRect r = CTLineGetImageBounds(line, context);
    CFRelease(line);
    return r;
}