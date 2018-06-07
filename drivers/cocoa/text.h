// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

CFAttributedStringRef makeAttrStringRef(CFStringRef str, CTFontRef font) {
	CFStringRef keys[]   = {kCTFontAttributeName};
	CFTypeRef   values[] = {font};

	CFDictionaryRef attributes = CFDictionaryCreate(
	    kCFAllocatorDefault, //
	    (const void**)&keys,
	    (const void**)&values,
	    sizeof(keys) / sizeof(keys[0]),
	    &kCFTypeDictionaryKeyCallBacks,
	    &kCFTypeDictionaryValueCallBacks);

	CFAttributedStringRef attrString = CFAttributedStringCreate(kCFAllocatorDefault, str, attributes);
	CFRelease(attributes);
	return attrString;
}

void DrawTextLine(CGContextRef context, CFStringRef str, CTFontRef font, CGPoint pos) {
	CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGContextSetTextPosition(context, pos.x, pos.y);
	CGContextSaveGState(context);
	CTLineDraw(line, context);
	CFRelease(line);
	CGContextRestoreGState(context);
}

CGRect GetTextLineGeometry(CGContextRef context, CFStringRef str, CTFontRef font) {
	CFAttributedStringRef attrStr = makeAttrStringRef(str, font);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGRect r = CTLineGetImageBounds(line, context);
	CFRelease(line);
	return r;
}